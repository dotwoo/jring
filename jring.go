package jring

import (
	"fmt"
	"hash/fnv"
	"sort"

	"sync"

	"errors"

	jump "github.com/dgryski/go-jump"
)

var hasher = fnv.New64a()

// Node interface
type Node interface {
	GetAddr() string
	GetWeight() int
	GetHash() uint64
}

// JRing interface
type JRing interface {
	Add(addr string, weight int) error
	Remove(addr string)
	Set(map[string]int)
	Get(key string) Node
	GetTwo(key string) (Node, Node)
	AllNodes() []Node
	Hash(string) uint64
}

type node struct {
	addr   string
	weight int
	hash   uint64
}

func (n node) GetAddr() string {
	return n.addr
}

func (n node) GetWeight() int {
	return n.weight
}

func (n node) GetHash() uint64 {
	return n.hash
}

func (n node) Eq(naddr string) bool {
	return n.addr == naddr
}

// NewNode create a node
func NewNode(addr string, weight int) Node {
	return node{addr: addr, weight: weight}
}

type nodeList []node

type hashJRing struct {
	nodes   nodeList
	members map[string]Node
	sync.RWMutex
}

func (h nodeList) Len() int {
	return len(h)
}

func (h nodeList) Less(i, j int) bool {
	return h[i].hash < h[j].hash
}

func (h nodeList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h nodeList) sort() {
	sort.Sort(h)
}

func (h *hashJRing) Add(addr string, weight int) error {
	h.Lock()
	defer h.Unlock()
	_, ok := h.members[addr]

	if ok {
		return errors.New("The node is exist")
	}
	hlen := len(h.nodes)
	cap := hlen + weight

	// resize node list
	nodes := make(nodeList, cap)
	copy(nodes, h.nodes)

	for i := hlen; i < cap; i++ {
		// hash: 0:localhost:7000:0
		// adding the index at the start and end seemed to give better distribution
		hasher.Write([]byte(fmt.Sprint(i, ":", addr, ":", i)))

		// hash value
		value := hasher.Sum64()

		// create node
		n := node{hash: value, addr: addr, weight: weight}

		// insert node
		nodes[i] = n

		// reset hash
		hasher.Reset()
	}
	nodes.sort()
	h.nodes = nodes
	h.members[addr] = NewNode(addr, weight)
	return nil
}

func (h *hashJRing) Remove(addr string) {
	h.Lock()
	defer h.Unlock()
	hlen := len(h.nodes)
	_, ok := h.members[addr]
	if !ok {
		return
	}

	nodes := make(nodeList, 0, hlen)
	delete(h.members, addr)
	isremoved := false
	for _, n := range h.nodes {
		if n.Eq(addr) {
			isremoved = true
			continue
		}
		nodes = append(nodes, n)
	}
	if isremoved {
		h.nodes = nodes
	}

}

func (h *hashJRing) Set(addrs map[string]int) {
	for addr, weight := range addrs {
		h.Add(addr, weight)
	}
}

func (h *hashJRing) Hash(key string) uint64 {
	hasher.Write([]byte(key))
	defer hasher.Reset()
	return hasher.Sum64()
}

func (h *hashJRing) Get(key string) Node {
	ukey := h.Hash(key)
	h.RLock()
	hlen := len(h.nodes)
	if hlen == 0 {
		h.RUnlock()
		return nil
	}
	index := int(jump.Hash(ukey, hlen))
	nd := h.nodes[index]
	h.RUnlock()
	return nd
}
func (h *hashJRing) GetTwo(key string) (Node, Node) {
	ukey := h.Hash(key)
	h.RLock()
	defer h.RUnlock()
	hlen := len(h.nodes)
	nlen := len(h.members)
	if hlen == 0 {
		return nil, nil
	}
	start := int(jump.Hash(ukey, hlen))
	nd := h.nodes[start]
	var nd2 node
	if nlen <= 1 {
		return nd, nil
	}
	for i := start + 1; i != start; i++ {
		if i >= hlen {
			i = 0
		}
		nd2 = h.nodes[i]
		if !nd.Eq(nd2.addr) {
			break
		}
	}
	return nd, nd2
}

func (h *hashJRing) AllNodes() []Node {
	out := make([]Node, 0)
	h.RLock()
	defer h.RUnlock()
	for _, n := range h.members {
		out = append(out, n)
	}
	return out
}

// NewHashJRing creates a new hash ring.
func NewHashJRing() JRing {
	return &hashJRing{
		nodes:   make([]node, 0, 16),
		members: make(map[string]Node),
	}
}
