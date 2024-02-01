package tree

import (
	"testing"
)

func Test_tree_insertFix(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		n *node
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
			tr := &tree{
				root: tt.fields.root,
			}
			tr.insertFix(tt.args.n)
		})
	}
}

func Test_tree_Search(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		key keyType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   valType
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &tree{
				root: tt.fields.root,
			}
			got, got1 := tr.Search(tt.args.key)
			if got != tt.want {
				t.Errorf("tree.Search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("tree.Search() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
