package jring

import (
	"reflect"
	"sync"
	"testing"
)

func Test_node_GetAddr(t *testing.T) {
	type fields struct {
		addr   string
		weight int
		hash   uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := node{
				addr:   tt.fields.addr,
				weight: tt.fields.weight,
				hash:   tt.fields.hash,
			}
			if got := n.GetAddr(); got != tt.want {
				t.Errorf("node.GetAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_GetWeight(t *testing.T) {
	type fields struct {
		addr   string
		weight int
		hash   uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := node{
				addr:   tt.fields.addr,
				weight: tt.fields.weight,
				hash:   tt.fields.hash,
			}
			if got := n.GetWeight(); got != tt.want {
				t.Errorf("node.GetWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_GetHash(t *testing.T) {
	type fields struct {
		addr   string
		weight int
		hash   uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := node{
				addr:   tt.fields.addr,
				weight: tt.fields.weight,
				hash:   tt.fields.hash,
			}
			if got := n.GetHash(); got != tt.want {
				t.Errorf("node.GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Eq(t *testing.T) {
	type fields struct {
		addr   string
		weight int
		hash   uint64
	}
	type args struct {
		naddr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := node{
				addr:   tt.fields.addr,
				weight: tt.fields.weight,
				hash:   tt.fields.hash,
			}
			if got := n.Eq(tt.args.naddr); got != tt.want {
				t.Errorf("node.Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	type args struct {
		addr   string
		weight int
	}
	tests := []struct {
		name string
		args args
		want Node
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.addr, tt.args.weight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodeList_Len(t *testing.T) {
	tests := []struct {
		name string
		h    nodeList
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("nodeList.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodeList_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    nodeList
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("nodeList.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodeList_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    nodeList
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_nodeList_sort(t *testing.T) {
	tests := []struct {
		name string
		h    nodeList
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.sort()
		})
	}
}

func Test_hashJRing_Add(t *testing.T) {
	type fields struct {
		nodes   nodeList
		members map[string]Node
		RWMutex sync.RWMutex
	}
	type args struct {
		addr   string
		weight int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashJRing{
				nodes:   tt.fields.nodes,
				members: tt.fields.members,
				RWMutex: tt.fields.RWMutex,
			}
			h.Add(tt.args.addr, tt.args.weight)
		})
	}
}

func Test_hashJRing_Remove(t *testing.T) {
	type fields struct {
		nodes   nodeList
		members map[string]Node
		RWMutex sync.RWMutex
	}
	type args struct {
		addr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashJRing{
				nodes:   tt.fields.nodes,
				members: tt.fields.members,
				RWMutex: tt.fields.RWMutex,
			}
			h.Remove(tt.args.addr)
		})
	}
}

func Test_hashJRing_Set(t *testing.T) {
	type fields struct {
		nodes   nodeList
		members map[string]Node
		RWMutex sync.RWMutex
	}
	type args struct {
		addrs map[string]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashJRing{
				nodes:   tt.fields.nodes,
				members: tt.fields.members,
				RWMutex: tt.fields.RWMutex,
			}
			h.Set(tt.args.addrs)
		})
	}
}

func Test_hashJRing_Hash(t *testing.T) {
	type fields struct {
		nodes   nodeList
		members map[string]Node
		RWMutex sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashJRing{
				nodes:   tt.fields.nodes,
				members: tt.fields.members,
				RWMutex: tt.fields.RWMutex,
			}
			if got := h.Hash(tt.args.key); got != tt.want {
				t.Errorf("hashJRing.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashJRing_Get(t *testing.T) {
	type fields struct {
		nodes   nodeList
		members map[string]Node
		RWMutex sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Node
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashJRing{
				nodes:   tt.fields.nodes,
				members: tt.fields.members,
				RWMutex: tt.fields.RWMutex,
			}
			if got := h.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashJRing.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashJRing_GetTwo(t *testing.T) {
	type fields struct {
		nodes   nodeList
		members map[string]Node
		RWMutex sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Node
		want1  Node
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashJRing{
				nodes:   tt.fields.nodes,
				members: tt.fields.members,
				RWMutex: tt.fields.RWMutex,
			}
			got, got1 := h.GetTwo(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashJRing.GetTwo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("hashJRing.GetTwo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_hashJRing_AllNodes(t *testing.T) {
	type fields struct {
		nodes   nodeList
		members map[string]Node
		RWMutex sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   []Node
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashJRing{
				nodes:   tt.fields.nodes,
				members: tt.fields.members,
				RWMutex: tt.fields.RWMutex,
			}
			if got := h.AllNodes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashJRing.AllNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashJRing(t *testing.T) {
	tests := []struct {
		name string
		want JRing
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashJRing(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashJRing() = %v, want %v", got, tt.want)
			}
		})
	}
}
