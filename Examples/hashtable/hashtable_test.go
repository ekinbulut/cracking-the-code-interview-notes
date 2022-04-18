package hashtable

import (
	"reflect"
	"testing"
)

func TestNewHashTable(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *HashTable
	}{
		// TODO: Add test cases.
		// fill in the test cases
		{
			name: "New Hash Table",
			args: args{
				size: 10,
			},
			want: &HashTable{
				size: 10,
				data: make([]*LinkedList, 10),
			},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashTable(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Add(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		h    *HashTable
		args args
	}{
		// TODO: Add test cases.
		// fill in 
		{
			name: "Add to new hashtable",
			h: NewHashTable(10),
			args: {
				"key",
				"value",
			},
		}
	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Add(tt.args.key, tt.args.value)
		})
	}
}

func TestHashTable_hash(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		h    *HashTable
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.hash(tt.args.key); got != tt.want {
				t.Errorf("HashTable.hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		h    *HashTable
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "",
			h: NewHashTable(10),
			args: {
				"key"
			},
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashTable.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
