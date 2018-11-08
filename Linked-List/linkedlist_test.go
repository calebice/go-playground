package linkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	type fields struct {
		head *Node
		size int
	}
	type args struct {
		data Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LinkedList
	}{
		{
			name:   "Starting a fresh list",
			fields: fields{head: nil},
			args:   args{data: 1},
			want:   &LinkedList{head: &Node{data: 1, next: nil}, size: 1},
		},
		{
			name:   "Adding to existing list",
			fields: fields{head: &Node{data: 1}, size: 1},
			args:   args{data: 2},
			want:   &LinkedList{head: &Node{data: 1, next: &Node{data: 2}}, size: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			ll.Append(tt.args.data)
			if !reflect.DeepEqual(ll, tt.want) {
				t.Errorf("Append() expected: %v, got: %v", tt.want, ll)
			}
		})
	}
}

func TestLinkedList_Insert(t *testing.T) {

	middleLink := &Node{data: 0, next: &Node{data: 2}}
	middleList := &Node{data: 0, next: &Node{data: 1, next: &Node{data: 2}}}
	endItemLink := &Node{data: 0, next: &Node{data: 2}}

	endList := &Node{data: 0, next: &Node{data: 2, next: &Node{data: 3}}}

	type fields struct {
		head *Node
		size int
	}
	type args struct {
		i    int
		data Item
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     Item
		wantErr  error
		wantList *LinkedList
	}{
		{
			name:     "Insert at beginning",
			fields:   fields{head: nil, size: 0},
			args:     args{i: 0, data: 1},
			want:     1,
			wantErr:  nil,
			wantList: &LinkedList{head: &Node{data: 1}, size: 1},
		},
		{
			name:     "Insert in middle",
			fields:   fields{head: middleLink, size: 2},
			args:     args{i: 1, data: 1},
			want:     1,
			wantErr:  nil,
			wantList: &LinkedList{head: middleList, size: 3},
		},
		{
			name:     "Insert at end",
			fields:   fields{head: endItemLink, size: 2},
			args:     args{i: 2, data: 3},
			want:     3,
			wantErr:  nil,
			wantList: &LinkedList{head: endList, size: 3},
		},
		{
			name:     "Failed index",
			fields:   fields{head: endItemLink, size: 3},
			args:     args{i: 10, data: 5},
			want:     nil,
			wantErr:  fmt.Errorf("Invalid index provided: 10"),
			wantList: &LinkedList{head: endItemLink, size: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			got, err := ll.Insert(tt.args.i, tt.args.data)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("LinkedList.Insert() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.Insert() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(ll, tt.wantList) {
				t.Errorf("LinkedList.Insert() = %v, wantList %v", ll, tt.wantList)
			}
		})
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	endRemoveNode := &Node{data: 0, next: &Node{data: 2}}
	endRemoveList := &Node{data: 0}
	startRemoveNode := &Node{data: 0, next: &Node{data: 2}}
	startRemoveList := &Node{data: 2}

	type fields struct {
		head *Node
		size int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     Item
		wantErr  bool
		wantList *LinkedList
	}{
		{
			name:     "Remove from end of list",
			fields:   fields{head: endRemoveNode, size: 2},
			args:     args{1},
			want:     2,
			wantErr:  false,
			wantList: &LinkedList{head: endRemoveList, size: 1},
		},
		{
			name:     "Remove start of list",
			fields:   fields{head: startRemoveNode, size: 2},
			args:     args{0},
			want:     0,
			wantErr:  false,
			wantList: &LinkedList{head: startRemoveList, size: 1},
		},
		{
			name:     "Invalid remove",
			fields:   fields{head: startRemoveNode, size: 1},
			args:     args{10},
			want:     nil,
			wantErr:  true,
			wantList: &LinkedList{head: startRemoveList, size: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			got, err := ll.RemoveAt(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.RemoveAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.RemoveAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	testNode := &Node{data: 0, next: &Node{data: 2}}

	type fields struct {
		head *Node
		size int
	}
	type args struct {
		data Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "Contains value",
			fields: fields{head: testNode, size: 2},
			args:   args{data: 2},
			want:   1,
		},
		{
			name:   "Not containing value",
			fields: fields{head: testNode, size: 2},
			args:   args{data: 10},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ll := &LinkedList{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if got := ll.IndexOf(tt.args.data); got != tt.want {
				t.Errorf("LinkedList.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Size(t *testing.T) {
	testNode := &Node{data: 0, next: &Node{data: 2}}

	type fields struct {
		head *Node
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "return size",
			fields: fields{head: testNode, size: 2},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := LinkedList{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if got := ll.Size(); got != tt.want {
				t.Errorf("LinkedList.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
