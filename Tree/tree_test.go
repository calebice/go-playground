package tree

import (
	"reflect"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	type fields struct {
		root *Node
		size int
	}
	type args struct {
		data Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BST
	}{
		{
			name:   "Insert one item into empty list",
			fields: fields{root: nil},
			args:   args{data: 1},
			want:   &BST{root: &Node{data: 1}, size: 1},
		},
		{
			name:   "Insert item into none-empty list (greater than)",
			fields: fields{root: &Node{data: 1}, size: 1},
			args:   args{data: 2},
			want:   &BST{root: &Node{data: 1, right: &Node{data: 2}}, size: 2},
		},
		{
			name:   "Insert item into none-empty list (less than)",
			fields: fields{root: &Node{data: 2}, size: 1},
			args:   args{data: 1},
			want:   &BST{root: &Node{data: 2, left: &Node{data: 1}}, size: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			bst.Insert(tt.args.data)
			if !reflect.DeepEqual(tt.want, bst) {
				t.Errorf("BST.Insert() = %v, want %v", bst, tt.want)
			}

		})
	}
}

func TestBST_Search(t *testing.T) {
	type fields struct {
		root *Node
		size int
	}
	type args struct {
		data Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "search field from tree (3 item less than)",
			fields: fields{root: &Node{data: 4, left: &Node{data: 1, right: &Node{data: 3}}}, size: 3},
			args:   args{3},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := bst.Search(tt.args.data); got != tt.want {
				t.Errorf("BST.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Remove(t *testing.T) {
	type fields struct {
		root *Node
		size int
	}
	type args struct {
		data Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBST *BST
	}{
		{
			name:    "delete root from tree (1 item)",
			fields:  fields{root: &Node{data: 1}, size: 1},
			args:    args{data: 1},
			wantBST: &BST{root: nil},
		},
		{
			name:    "delete root from tree (2 item greater than)",
			fields:  fields{root: &Node{data: 1, right: &Node{data: 2}}, size: 2},
			args:    args{data: 1},
			wantBST: &BST{root: &Node{data: 2}, size: 1},
		},
		{
			name:    "delete root from tree (2 item less than)",
			fields:  fields{root: &Node{data: 2, left: &Node{data: 1}}, size: 2},
			args:    args{data: 2},
			wantBST: &BST{root: &Node{data: 1}, size: 1},
		},
		{
			name:    "delete root from tree (3 item less than)",
			fields:  fields{root: &Node{data: 4, left: &Node{data: 1, right: &Node{data: 3}}}, size: 3},
			args:    args{data: 4},
			wantBST: &BST{root: &Node{data: 1, right: &Node{data: 3}}, size: 2},
		},
		{
			name:    "delete root from tree (3 item greater than)",
			fields:  fields{root: &Node{data: 1, right: &Node{data: 4, left: &Node{data: 3}}}, size: 3},
			args:    args{data: 1},
			wantBST: &BST{root: &Node{data: 4, left: &Node{data: 3}}, size: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			bst.Remove(tt.args.data)
			if !reflect.DeepEqual(bst, tt.wantBST) {
				t.Errorf("BST.Remove() BST = %v, want %v", bst, tt.wantBST)
			}
		})
	}
}
