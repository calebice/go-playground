package queue

import (
	"reflect"
	"testing"
)

func TestQueue_IsEmpty(t *testing.T) {
	type fields struct {
		items []Item
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Empty list",
			fields: fields{items: []Item{}},
			want:   true,
		},
		{
			name:   "None empty queue",
			fields: fields{items: []Item{1}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Queue{
				items: tt.fields.items,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("Queue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Front(t *testing.T) {
	type fields struct {
		items []Item
	}
	tests := []struct {
		name   string
		fields fields
		want   Item
	}{
		{
			name:   "Return from empty list",
			fields: fields{items: []Item{}},
			want:   nil,
		},
		{
			name:   "Return from single list",
			fields: fields{items: []Item{1}},
			want:   1,
		},
		{
			name:   "Return multi list",
			fields: fields{items: []Item{1, 2}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Queue{
				items: tt.fields.items,
			}
			if got := s.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Size(t *testing.T) {
	type fields struct {
		items []Item
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Single entry",
			fields: fields{items: []Item{1}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Queue{
				items: tt.fields.items,
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("Queue.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	type fields struct {
		items []Item
	}
	tests := []struct {
		name      string
		fields    fields
		want      Item
		wantQueue *Queue
	}{
		{
			name:      "Remove single item",
			fields:    fields{items: []Item{1}},
			want:      1,
			wantQueue: &Queue{items: []Item{}},
		},
		{
			name:      "Remove from multi list",
			fields:    fields{items: []Item{1, 2}},
			want:      1,
			wantQueue: &Queue{items: []Item{2}},
		},
		{
			name:      "Remove from empty list",
			fields:    fields{items: []Item{}},
			want:      nil,
			wantQueue: &Queue{items: []Item{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Queue{
				items: tt.fields.items,
			}
			if got := s.Dequeue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(s, tt.wantQueue) {
				t.Errorf("Queue.Dequeue() = %v, want %v", s, tt.wantQueue)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type fields struct {
		items []Item
	}
	type args struct {
		item Item
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantQueue *Queue
	}{
		{
			name:      "Add to empty list",
			fields:    fields{items: []Item{}},
			args:      args{1},
			wantQueue: &Queue{items: []Item{1}},
		},
		{
			name:      "Add to existing list",
			fields:    fields{items: []Item{1}},
			args:      args{2},
			wantQueue: &Queue{items: []Item{1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Queue{
				items: tt.fields.items,
			}
			s.Enqueue(tt.args.item)
			if !reflect.DeepEqual(s, tt.wantQueue) {
				t.Errorf("Queue.Enqueue %v, want %v", s, tt.wantQueue)
			}
		})
	}
}
