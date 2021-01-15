package Robber

import "fmt"

func Robber(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	h := make([]int, len(nums))
	h[0] = nums[0]
	max := h[0]
	for i := 1; i < len(nums); i++ {
		h[i] = nums[i]
		for j := 0; j <= i-2; j++ {
			if h[j]+nums[i] > h[i] {
				h[i] = h[j] + nums[i]
			}
		}
		if h[i] > max {
			max = h[i]
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Rob() {
	fmt.Println("Robber Assignment")
	num := []int{3, 3, 2}
	if len(num) == 0 {
		fmt.Println("0")
	}
	if len(num) == 1 {
		fmt.Println(num[0])
	}
	fmt.Println(Max(Robber(num[1:]), Robber(num[0:len(num)-1])))
}
