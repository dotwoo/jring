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
		{
			name: "testGetHost",
			fields: fields{
				addr:   "localhost:7788",
				weight: 7,
			},
			want: "localhost:7788",
		},
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
		{
			name: "testGetWeight",
			fields: fields{
				addr:   "localhost:7788",
				weight: 7,
			},
			want: 7,
		},
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
		{
			name: "testGetHash",
			fields: fields{
				addr:   "localhost:7788",
				weight: 7,
				hash:   2342343243,
			},
			want: 2342343243,
		},
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
		{
			name: "testNodeEqFalse",
			fields: fields{
				addr:   "localhost:7788",
				weight: 7,
				hash:   2342343243,
			},
			args: args{
				naddr: "localhost:7788 ",
			},
			want: false,
		},
		{
			name: "testNodeEq",
			fields: fields{
				addr:   "localhost:7788",
				weight: 7,
				hash:   2342343243,
			},
			args: args{
				naddr: "localhost:7788",
			},
			want: true,
		},
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
		{
			name: "testNewNode",
			args: args{
				addr:   "localhost:7788",
				weight: 7,
			},
			want: node{
				addr:   "localhost:7788",
				weight: 7,
			},
		},
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
		{
			name: "nodeListLen0",
			h:    []node{},
			want: 0,
		},
		{
			name: "nodeListLen1",
			h: []node{node{
				addr:   "test.com",
				weight: 7,
			}},
			want: 1,
		},
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
		{
			name: "testNodeListLessTrue",
			h: nodeList{
				node{
					addr: "node1",
					hash: 1,
				},
				node{
					addr: "node2",
					hash: 2,
				},
			},
			args: args{
				i: 0,
				j: 1,
			},
			want: true,
		},
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
		name   string
		h      nodeList
		args   args
		before bool
		after  bool
	}{
		// TODO: Add test cases.
		{
			name: "testNodeListSwap",
			h: nodeList{
				node{
					addr: "node1",
					hash: 1,
				},
				node{
					addr: "node2",
					hash: 2,
				},
			},
			args: args{
				i: 0,
				j: 1,
			},
			before: true,
			after:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.before {
				t.Errorf("before swap nodeList.Less() = %v, want %v", got, tt.before)
			}
			tt.h.Swap(tt.args.i, tt.args.j)
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.after {
				t.Errorf("after swap nodeList.Less() = %v, want %v", got, tt.after)
			}
		})
	}
}

func Test_nodeList_sort(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		h      nodeList
		args   args
		before bool
		after  bool
	}{
		// TODO: Add test cases.
		{
			name: "testNodeListsort",
			h: nodeList{
				node{
					addr: "node2",
					hash: 2,
				},
				node{
					addr: "node1",
					hash: 1,
				},
			},
			args: args{
				i: 0,
				j: 1,
			},
			before: false,
			after:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.before {
				t.Errorf("befor sort nodeList.Less() = %v, want %v", got, tt.before)
			}
			tt.h.sort()
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.after {
				t.Errorf("after sort nodeList.Less() = %v, want %v", got, tt.after)
			}
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
		name        string
		fields      fields
		args        args
		nodeswant   int
		memberswant int
	}{
		// TODO: Add test cases.
		{
			name: "zeroadd",
			fields: fields{
				nodes:   nodeList{},
				members: make(map[string]Node, 0),
			},
			args: args{
				addr:   "test1",
				weight: 2,
			},
			nodeswant:   2,
			memberswant: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashJRing{
				nodes:   tt.fields.nodes,
				members: tt.fields.members,
				RWMutex: tt.fields.RWMutex,
			}
			h.Add(tt.args.addr, tt.args.weight)
			if got := len(h.nodes); got != tt.nodeswant {
				t.Errorf("after add len(h.nodes) = %v, want %v", got, tt.nodeswant)
			}
			if got := len(h.members); got != tt.memberswant {
				t.Errorf("after add len(h.members) = %v, want %v", got, tt.memberswant)
			}
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
		name        string
		fields      fields
		args        args
		nodeswant   int
		memberswant int
	}{
		// TODO: Add test cases.
		{
			name: "zeroremove",
			fields: fields{
				nodes:   nodeList{},
				members: make(map[string]Node, 0),
			},
			args: args{
				addr: "test1",
			},
			nodeswant:   0,
			memberswant: 0,
		},
		{
			name: "one remove",
			fields: fields{
				nodes:   nodeList{node{addr: "test1", hash: 1, weight: 2}, node{addr: "test1", hash: 2, weight: 2}},
				members: map[string]Node{"test1": node{addr: "test1"}},
			},
			args: args{
				addr: "test1",
			},
			nodeswant:   0,
			memberswant: 0,
		},
		{
			name: "not match remove",
			fields: fields{
				nodes:   nodeList{node{addr: "test1", hash: 1, weight: 2}, node{addr: "test1", hash: 2, weight: 2}},
				members: map[string]Node{"test1": node{addr: "test1", weight: 2}},
			},
			args: args{
				addr: "test2",
			},
			nodeswant:   2,
			memberswant: 1,
		},
		{
			name: "two remove",
			fields: fields{
				nodes:   nodeList{node{addr: "test1", hash: 1, weight: 2}, node{addr: "test1", hash: 2, weight: 2}, node{addr: "test2", weight: 1}},
				members: map[string]Node{"test1": node{addr: "test1", weight: 2}, "test2": node{addr: "test2", weight: 1}},
			},
			args: args{
				addr: "test2",
			},
			nodeswant:   2,
			memberswant: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashJRing{
				nodes:   tt.fields.nodes,
				members: tt.fields.members,
				RWMutex: tt.fields.RWMutex,
			}
			h.Remove(tt.args.addr)
			if got := len(h.nodes); got != tt.nodeswant {
				t.Errorf("after remove len(h.nodes) = %v, want %v", got, tt.nodeswant)
			}
			if got := len(h.members); got != tt.memberswant {
				t.Errorf("after remove len(h.members) = %v, want %v", got, tt.memberswant)
			}
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
		name        string
		fields      fields
		args        args
		nodeswant   int
		memberswant int
	}{
		// TODO: Add test cases.
		{
			name: "zeroadd",
			fields: fields{
				nodes:   nodeList{},
				members: make(map[string]Node, 0),
			},
			args: args{
				addrs: map[string]int{"test1": 2, "test2": 3},
			},
			nodeswant:   5,
			memberswant: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashJRing{
				nodes:   tt.fields.nodes,
				members: tt.fields.members,
				RWMutex: tt.fields.RWMutex,
			}
			h.Set(tt.args.addrs)
			if got := len(h.nodes); got != tt.nodeswant {
				t.Errorf("after set len(h.nodes) = %v, want %v", got, tt.nodeswant)
			}
			if got := len(h.members); got != tt.memberswant {
				t.Errorf("after set len(h.members) = %v, want %v", got, tt.memberswant)
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
		{
			name: "newHashJRing",
			want: &hashJRing{
				nodes:   nodeList{},
				members: make(map[string]Node, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashJRing(); !reflect.DeepEqual(got.AllNodes(), tt.want.AllNodes()) {
				t.Errorf("NewHashJRing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashJRing_Hash(t *testing.T) {

	tests := []struct {
		name  string
		addrs map[string]int
		key   string
	}{
		// TODO: Add test cases.
		{
			name:  "Get Hash",
			addrs: map[string]int{"test1": 3, "test2": 3, "test3": 1, "test4": 4},
			key:   "keysone",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashJRing()
			h.Set(tt.addrs)
			data := []byte(tt.key)

			// Hash input data
			hasher.Write(data)

			// calculate hash
			want := hasher.Sum64()

			// reset hasher
			hasher.Reset()
			if got := h.Hash(tt.key); got != want {
				t.Errorf("hashJRing.Hash() = %v, want %v", got, want)
			}
		})
	}
}

func Test_hashJRing_AllNodes(t *testing.T) {

	tests := []struct {
		name  string
		addrs map[string]int
		want  int
	}{
		// TODO: Add test cases.
		{
			name:  "Get AllNodes",
			addrs: map[string]int{"test1": 3, "test2": 3, "test3": 1, "test4": 4},
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashJRing()
			h.Set(tt.addrs)
			if got := h.AllNodes(); len(got) != tt.want {
				t.Errorf("hashJRing.AllNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashJRing_Get(t *testing.T) {
	tests := []struct {
		name  string
		addrs map[string]int
		key   string
		want  Node
	}{
		// TODO: Add test cases.
		{
			name:  "GetNode",
			addrs: map[string]int{"test1": 3, "test2": 3, "test3": 1},
			key:   "testkey",
		},
		{
			name:  "GetNode nil",
			addrs: map[string]int{},
			key:   "testkey",
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashJRing()
			h.Set(tt.addrs)
			got2 := h.Get(tt.key)
			if got := h.Get(tt.key); !reflect.DeepEqual(got, got2) {
				t.Errorf("hashJRing.Get() = %v, want %v", got, got2)
			}
		})
	}
}

func Test_hashJRing_GetTwo(t *testing.T) {
	tests := []struct {
		name  string
		addrs map[string]int
		key   string
	}{
		// TODO: Add test cases.
		{
			name:  "Get2 threenodes",
			addrs: map[string]int{"test1": 3, "test2": 3, "test3": 1},
			key:   "dfsdfsfsd",
		},
		{
			name:  "Get2 twonodes",
			addrs: map[string]int{"test1": 1, "test3": 1},
			key:   "dfsdfsfsd",
		},
		{
			name:  "Get onenode",
			addrs: map[string]int{"test1": 1},
			key:   "dfsdfsfsd",
		},
		{
			name:  "Get zeronode",
			addrs: map[string]int{},
			key:   "dfsdfsfsd",
		},
		{
			name:  "Get circle",
			addrs: map[string]int{"test1": 1, "test2": 1, "test3": 1},
			key:   "testkey",
		},
		{
			name:  "Get circle2",
			addrs: map[string]int{"test1": 1, "test2": 1, "test3": 1},
			key:   "testkey2",
		},
		{
			name:  "Get circle3",
			addrs: map[string]int{"test1": 1, "test2": 1, "test3": 1},
			key:   "testkey3",
		},
		{
			name:  "Get circle4",
			addrs: map[string]int{"test1": 1, "test2": 1, "test3": 1},
			key:   "testkey4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashJRing()
			h.Set(tt.addrs)
			want := h.Get(tt.key)
			got, got1 := h.GetTwo(tt.key)
			if !reflect.DeepEqual(want, got) {
				t.Errorf("hashJRing.GetTwo() got = %v, want %v", got, want)
			}
			if got1 != nil && reflect.DeepEqual(want, got1) {
				t.Errorf("hashJRing.GetTwo() got1 != %v, want %v", got1, want)
			}
			if got != nil && got1 != nil && got.GetAddr() == got1.GetAddr() {
				t.Errorf("hashJRing.GetTwo() got = %v, got1 %v", got, got1)
			}
		})
	}
}
