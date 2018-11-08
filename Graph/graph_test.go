package graph

import (
	"reflect"
	"testing"
)

func TestGraph_AddEdge(t *testing.T) {
	testNodes := []*Node{&Node{1}, &Node{2}}
	testMap := map[Node][]*Node{Node{1}: []*Node{&Node{2}}, Node{2}: []*Node{&Node{1}}}
	type fields struct {
		nodes []*Node
		edges map[Node][]*Node
	}
	type args struct {
		n1 *Node
		n2 *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Graph
	}{
		{
			name:   "Add edge",
			fields: fields{nodes: testNodes},
			args:   args{n1: &Node{1}, n2: &Node{2}},
			want:   &Graph{nodes: testNodes, edges: testMap},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				nodes: tt.fields.nodes,
				edges: tt.fields.edges,
			}
			g.AddEdge(tt.args.n1, tt.args.n2)
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("g.AddEdge() wanted: %v, got %v", tt.want, g)
			}
		})
	}
}

func TestGraph_AddNode(t *testing.T) {
	type fields struct {
		nodes []*Node
		edges map[Node][]*Node
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Graph
	}{
		{
			name:   "Add node to empty list",
			fields: fields{},
			args:   args{&Node{1}},
			want:   &Graph{nodes: []*Node{&Node{1}}},
		},
		{
			name:   "Add node to list with items",
			fields: fields{nodes: []*Node{&Node{1}}},
			args:   args{&Node{2}},
			want:   &Graph{nodes: []*Node{&Node{1}, &Node{2}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				nodes: tt.fields.nodes,
				edges: tt.fields.edges,
			}
			g.AddNode(tt.args.node)
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("g.AddNode() wanted: %v, got %v", tt.want, g)
			}
		})
	}
}

func TestMakeNode(t *testing.T) {
	type args struct {
		item Item
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Create a node",
			args: args{item: 1},
			want: &Node{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeNode(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
