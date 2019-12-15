package cylinders

import (
	"strconv"
	"strings"
)

//LeftHandCountString String wrapper for LeftHandCount function
//input should be separated by a space. ex: "1 2 3"
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

//LeftHandCount to count how many times the ball (nums) in lefthand get replaced
func LeftHandCount(nums []int) int {
	//variable declaration
	aCylinder := nums
	sCylinder := make([]int, len(nums))
	count := 0

	//for each num in nums do a loop
	for i := 0; i < len(nums); i++ {
		//lefthand represent the highest of the known value
		//because the constaint of this problem is 1-1000 the 0 is its min value
		leftHand := 0

		//change the origin cylinder for a count
		//first iteration always starts from aCylinder
		if i%2 == 0 {
			//foreach number in aCylinder, put the current number in the right hand
			for _, numberInRightHand := range aCylinder {
				//hold the highest number in the left hand
				//and prepend the other the number to the sCylinder
				//count everytime number on the left hand got replaced
				if numberInRightHand > leftHand {
					if leftHand != 0 {
						sCylinder = append([]int{leftHand}, sCylinder[:len(nums)]...)
						count++
					}
					leftHand = numberInRightHand
				} else {
					sCylinder = append([]int{numberInRightHand}, sCylinder[:len(nums)]...)
				}
			}
			//emptied the used cylinder
			aCylinder = make([]int, len(nums))
		} else {
			//same proccess as above but flip the cylinder
			for _, numberInRightHand := range sCylinder {
				if numberInRightHand > leftHand {
					if leftHand != 0 {
						aCylinder = append([]int{leftHand}, aCylinder[:len(nums)]...)
						count++
					}
					leftHand = numberInRightHand
				} else {
					aCylinder = append([]int{numberInRightHand}, aCylinder[:len(nums)]...)
				}
			}
			sCylinder = make([]int, len(nums))
		}
	}
	return count
}
