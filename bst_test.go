package bst

import (
	"reflect"
	"testing"
)

func Test_tree_DeepestNodes(test *testing.T) {
	testcases := []struct {
		name          string
		input         []int
		maxDepth      int
		deepestValues []int
	}{
		{
			name:          "single deepest node",
			input:         []int{12, 11, 90, 82, 7, 9},
			maxDepth:      3,
			deepestValues: []int{9},
		},
		{
			name:          "multiple deepest nodes",
			input:         []int{26, 82, 16, 92, 33},
			maxDepth:      2,
			deepestValues: []int{33, 92},
		},
		{
			name:          "no nodes",
			input:         []int{},
			maxDepth:      0,
			deepestValues: []int{},
		},
		{
			name:          "duplicate node values",
			input:         []int{1, 1, 1, 1, 1, 1, 2, 2, -1},
			maxDepth:      1,
			deepestValues: []int{-1, 2},
		},
	}
	for _, testcase := range testcases {
		test.Run(testcase.name, func(test *testing.T) {
			tree := NewIntBST(testcase.input)
			nodes, maxDepth := tree.DeepestNodes()
			nodeValues := IntegerNodeValues(nodes)
			if !reflect.DeepEqual(nodeValues, testcase.deepestValues) {
				test.Errorf("expected %v but got %v deepest Node values", testcase.deepestValues, nodeValues)
			}

			if maxDepth != testcase.maxDepth {
				test.Errorf("expected %d but got %d for max depth level", testcase.maxDepth, maxDepth)
			}
		})
	}
}

// running similar test cases to DeepestNodes, with outputs updated to require both Search and Delete
// functionality to be working as expected
func Test_tree_FindAndDelete(test *testing.T) {
	testcases := []struct {
		name          string
		input         []int
		deleteVal     int
		maxDepth      int
		deepestValues []int
	}{
		{
			name:          "single deepest node removal",
			input:         []int{12, 11, 90, 82, 7, 9},
			deleteVal:     9, // expecting that removing 9 aligns 7 and 82 at 2 levels deep
			maxDepth:      2,
			deepestValues: []int{7, 82},
		},
		{
			name:          "multiple deepest values with single value removal",
			input:         []int{26, 82, 16, 92, 33},
			deleteVal:     33, // expecting that removing 33 leaves 92 as the lone deepest Node now
			maxDepth:      2,
			deepestValues: []int{92},
		},
		{
			name:          "fork node removal",
			input:         []int{26, 82, 16, 92, 33},
			deleteVal:     82, // expecting that removing branch for 82 results in Node shift of its successor (92)
			maxDepth:      2,  // as replacement, leaving 33 as the deepest Node at the previous level
			deepestValues: []int{33},
		},
		{
			name:          "fork node removal with successor beyond parent level",
			input:         []int{26, 82, 16, 92, 33, 84},
			deleteVal:     82, // expecting that removing branch for 82 results in Node shift of its successor (84)
			maxDepth:      2,  // as replacement, leaving the second level with 33 and 92 in place
			deepestValues: []int{33, 92},
		},
	}
	for _, testcase := range testcases {
		test.Run(testcase.name, func(test *testing.T) {
			tree := NewIntBST(testcase.input)
			tree.Delete(tree.SearchInt(testcase.deleteVal))
			nodes, maxDepth := tree.DeepestNodes()
			nodeValues := IntegerNodeValues(nodes)
			if !reflect.DeepEqual(nodeValues, testcase.deepestValues) {
				test.Errorf("expected %v but got %v deepest Node values", testcase.deepestValues, nodeValues)
			}

			if maxDepth != testcase.maxDepth {
				test.Errorf("expected %d but got %d for max depth level", testcase.maxDepth, maxDepth)
			}
		})
	}
}

func Test_tree_MinMax(test *testing.T) {
	testcases := []struct {
		name        string
		input       []int
		expectedMin int
		expectedMax int
	}{
		{
			name:        "single deepest node",
			input:       []int{12, 11, 90, 82, 7, 9},
			expectedMax: 90,
			expectedMin: 7,
		},
		{
			name:        "multiple deepest nodes",
			input:       []int{26, 82, 16, 92, 33},
			expectedMax: 92,
			expectedMin: 16,
		},
	}
	for _, testcase := range testcases {
		test.Run(testcase.name, func(test *testing.T) {
			tree := NewIntBST(testcase.input)
			minValue := tree.Min().IntValue()
			maxValue := tree.Max().IntValue()

			if testcase.expectedMin != minValue {
				test.Errorf("expected %d but got %d for minimum IntValue", testcase.expectedMin, minValue)
			}
			if testcase.expectedMax != maxValue {
				test.Errorf("expected %d but got %d for maximum IntValue", testcase.expectedMax, maxValue)
			}
		})
	}
}

func Test_tree_Walk(test *testing.T) {
	testcases := []struct {
		name            string
		input           []int
		inOrderOutput   []int
		preOrderOutput  []int
		postOrderOutput []int
	}{
		{
			name:            "single deepest node",
			input:           []int{12, 11, 90, 82, 7, 9},
			inOrderOutput:   []int{11, 7, 9, 12, 90, 82},
			preOrderOutput:  []int{12, 11, 7, 9, 90, 82},
			postOrderOutput: []int{11, 7, 9, 90, 82, 12},
		},
		{
			name:            "multiple deepest nodes",
			input:           []int{26, 82, 16, 92, 33},
			inOrderOutput:   []int{16, 26, 82, 33, 92},
			preOrderOutput:  []int{26, 16, 82, 33, 92},
			postOrderOutput: []int{16, 82, 33, 92, 26},
		},
	}
	for _, testcase := range testcases {
		test.Run(testcase.name, func(test *testing.T) {
			tree := NewIntBST(testcase.input)

			inOrder := IntegerNodeValues(tree.WalkInOrder())

			if !reflect.DeepEqual(inOrder, testcase.inOrderOutput) {
				test.Errorf("expected %v but got %v for inOrder walk values", testcase.inOrderOutput, inOrder)
			}

			preOrder := IntegerNodeValues(tree.WalkPreOrder())

			if !reflect.DeepEqual(preOrder, testcase.preOrderOutput) {
				test.Errorf("expected %v but got %v for preOrder walk values", testcase.preOrderOutput, preOrder)
			}

			postOrder := IntegerNodeValues(tree.WalkPostOrder())

			if !reflect.DeepEqual(postOrder, testcase.postOrderOutput) {
				test.Errorf("expected %v but got %v for inOrder walk values", testcase.postOrderOutput, postOrder)
			}
		})
	}
}
