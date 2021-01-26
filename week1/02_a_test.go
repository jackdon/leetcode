package week1

import "testing"

func NewNodeFromArr(arr []interface{}) *TreeNode {
	return nil
}

func TestGetTargetCopy(t *testing.T) {
	original := NewNodeFromArr(nil)
	getTargetCopy(original, nil, nil)
}
