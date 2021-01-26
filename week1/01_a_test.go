package week1

import "testing"

func TestCanFormArray(t *testing.T) {
	// Example1
	{
		var (
			arr    = []int{85}
			pieces = [][]int{{85}}
		)
		t.Log(canFormArray(arr, pieces))
	}
	// Example2
	{
		var (
			arr    = []int{15, 88}
			pieces = [][]int{{88}, {15}}
		)
		t.Log(canFormArray(arr, pieces))
	}
	// Example3
	{
		var (
			arr    = []int{49, 18, 16}
			pieces = [][]int{{16, 18, 49}}
		)
		t.Log(canFormArray(arr, pieces))
	}
	// Example4
	{
		var (
			arr    = []int{91, 4, 64, 78}
			pieces = [][]int{{78}, {4, 64}, {91}}
		)
		t.Log(canFormArray(arr, pieces))
	}
	// Example5
	{
		var (
			arr    = []int{1, 3, 5, 7}
			pieces = [][]int{{2, 4, 6, 8}}
		)
		t.Log(canFormArray(arr, pieces))
	}
}
