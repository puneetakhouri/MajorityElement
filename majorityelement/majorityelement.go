package majorityelement

import "fmt"

//Start does the main calculations
func Start() {
	//arr := []int{3, 3, 1, 2}
	arr := []int{3, 1, 3, 3, 2}
	//fmt.Printf("%d", Compute(arr))
	num := Compute(arr)

	fmt.Printf("%d", num)
}

//Compute test
func Compute(arr []int) int {
	num, count := ComputeBoyerMooreAlgorithm(arr)
	if count > len(arr)/2 {
		//fmt.Printf("%d", num)
		return num
	} else {
		count := 0
		for i := 0; i < len(arr); i++ {
			if arr[i] == num {
				count++
				if count > len(arr)/2 {
					//fmt.Printf("%d", num)
					return num
				}
			}
		}
	}
	return -1
}

//ComputeBoyerMooreAlgorithm test
func ComputeBoyerMooreAlgorithm(arr []int) (int, int) {
	i, m := 0, 0
	for j := 0; j < len(arr); j++ {
		if i == 0 {
			m = arr[j]
			i = 1
		} else if m == arr[j] {
			i = i + 1
		} else {
			i = i - 1
		}
	}
	return m, i
}
