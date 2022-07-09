package bst

import (
	"reflect"
	"testing"
)

func Test_tree_DeepestNodes(test *testing.T) {
	tests := []struct {
		name          string
		input         []int
		maxDepth      int
		deepestValues []int
	}{
		{
			name:          "single deepest value",
			input:         []int{12, 11, 90, 82, 7, 9},
			maxDepth:      3,
			deepestValues: []int{9},
		},
		{
			name:          "multiple deepest values",
			input:         []int{26, 82, 16, 92, 33},
			maxDepth:      2,
			deepestValues: []int{33, 92},
		},
		{
			name:          "no values",
			input:         []int{},
			maxDepth:      0,
			deepestValues: []int{},
		},
		{
			name:          "duplicate values",
			input:         []int{1, 1, 1, 1, 1, 1, 2, 2, -1},
			maxDepth:      1,
			deepestValues: []int{-1, 2},
		},
	}
	for _, testcase := range tests {
		test.Run(testcase.name, func(test *testing.T) {
			tree := NewFromIntSlice(testcase.input)
			nodes, maxDepth := tree.DeepestNodes()
			nodeValues := NodesToIntSlice(nodes)
			if !reflect.DeepEqual(nodeValues, testcase.deepestValues) {
				test.Errorf("expected %v but got %v deepest Node values", testcase.deepestValues, nodeValues)
			}

			if maxDepth != testcase.maxDepth {
				test.Errorf("expected %d but got %d for max depth level", testcase.maxDepth, maxDepth)
			}
		})
	}
}

// running similar test cases to DeepestNodes, with outputs updated to require both Find and Delete
// functionality to be working as expected
func Test_tree_FindAndDelete(test *testing.T) {
	tests := []struct {
		name          string
		input         []int
		deleteVal     int
		maxDepth      int
		deepestValues []int
	}{
		{
			name:          "single deepest value",
			input:         []int{12, 11, 90, 82, 7, 9},
			deleteVal:     9,
			maxDepth:      2,
			deepestValues: []int{7, 82},
		},
		{
			name:          "multiple deepest values",
			input:         []int{26, 82, 16, 92, 33},
			deleteVal:     33,
			maxDepth:      2,
			deepestValues: []int{92},
		},
	}
	for _, testcase := range tests {
		test.Run(testcase.name, func(test *testing.T) {
			tree := NewFromIntSlice(testcase.input)
			tree.Delete(tree.Find(testcase.deleteVal))
			nodes, maxDepth := tree.DeepestNodes()
			nodeValues := NodesToIntSlice(nodes)
			if !reflect.DeepEqual(nodeValues, testcase.deepestValues) {
				test.Errorf("expected %v but got %v deepest Node values", testcase.deepestValues, nodeValues)
			}

			if maxDepth != testcase.maxDepth {
				test.Errorf("expected %d but got %d for max depth level", testcase.maxDepth, maxDepth)
			}
		})
	}
}
