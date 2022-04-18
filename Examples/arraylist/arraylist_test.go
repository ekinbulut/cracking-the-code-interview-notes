package arraylist

import (
	"reflect"
	"testing"
)

func TestArrayList_Remove(t *testing.T) {
	type fields struct {
		data []interface{}
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{

		{
			name: "remove first item",
			fields: fields{
				data: []interface{}{1, 2, 3, 4, 5},
			},
			args: args{
				index: 0,
			},
		},
		{
			name: "remove last item",
			fields: fields{
				data: []interface{}{1, 2, 3, 4, 5},
			},
			args: args{
				index: 4,
			},
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList{
				data: tt.fields.data,
			}
			a.Remove(tt.args.index)
		})
	}
}

func TestNewArrayList(t *testing.T) {
	tests := []struct {
		name string
		want *ArrayList
	}{
		// TODO: Add test cases.
		// fill in the test cases
		{
			name: "create new arraylist",
			want: &ArrayList{}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayList() = %v, want %v", got, tt.want)
			}
		})
	}
}
