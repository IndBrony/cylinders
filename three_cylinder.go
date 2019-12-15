package cylinders

import (
	"fmt"
	"strconv"
	"strings"
)

func LeftHandCountString(numsString string) (int, error) {
	stringSlice := strings.Split(numsString, " ")
	nums := make([]int, len(stringSlice))
	for index, char := range stringSlice {
		converted, err := strconv.Atoi(char)
		if err != nil {
			return -1, err
		}
		nums[index] = converted
	}

	return LeftHandCount(nums), nil
}

func LeftHandCount(nums []int) int {
	aCylinder := nums
	sCylinder := make([]int, len(nums))
	count := 0

	for i := 0; i < len(nums); i++ {
		leftHand := 0
		if i%2 == 0 {
			for _, numberInRightHand := range aCylinder {
				if numberInRightHand > leftHand {
					if leftHand != 0 {
						sCylinder = append([]int{leftHand}, sCylinder[:len(nums)]...)
						fmt.Println(numberInRightHand, leftHand)
						count++
					}
					leftHand = numberInRightHand
				} else {
					sCylinder = append([]int{numberInRightHand}, sCylinder[:len(nums)]...)
				}
			}
			fmt.Println(sCylinder)
			aCylinder = make([]int, len(nums))
		} else {
			for _, numberInRightHand := range sCylinder {
				if numberInRightHand > leftHand {
					if leftHand != 0 {
						aCylinder = append([]int{leftHand}, aCylinder[:len(nums)]...)
						fmt.Println(numberInRightHand, leftHand)
						count++
					}
					leftHand = numberInRightHand
				} else {
					aCylinder = append([]int{numberInRightHand}, aCylinder[:len(nums)]...)
				}
			}
			fmt.Println(aCylinder)
			sCylinder = make([]int, len(nums))
		}
	}
	return count
}
