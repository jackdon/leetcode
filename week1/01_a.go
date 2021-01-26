package week1

func canFormArray(arr []int, pieces [][]int) bool {

	var intArrayEqual = func(a, b []int) bool {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	var flag1 = true
	for i := 0; i < len(pieces); i++ {
		piLen := len(pieces[i])
		var flag = false
		for j := 0; j < len(arr)-piLen+1; j++ {
			if intArrayEqual(arr[j:j+piLen], pieces[i]) {
				flag = true
				break
			}
		}
		if !flag {
			flag1 = false
			break
		}
	}
	return flag1
}
