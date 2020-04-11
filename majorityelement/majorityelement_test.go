package majorityelement

import "testing"

func TestCompute(t *testing.T) {
	testData := []struct {
		arr         []int
		expectedVal int
	}{
		{[]int{1, 2, 3, 4, 5}, -1},
		{[]int{3, 3, 3, 1, 2}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{3, 1, 3, 3, 2}, 3},
	}

	for _, data := range testData {
		actual := Compute(data.arr)
		if actual == data.expectedVal {
			t.Log("SUCCESS")
		} else {
			t.Errorf("FAIL, expected %d, got %d", data.expectedVal, actual)
		}
	}
}
