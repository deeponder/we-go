/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Here are few examples.
[1,3,5,6], 5 → 2
[1,3,5,6], 2 → 1
[1,3,5,6], 7 → 4
[1,3,5,6], 0 → 0
 */
package array

func searchInsertPosition(nums []int, val int) int {
	i := 0
	for i < len(nums) && nums[i] <= val {
		if nums[i] == val {
			return i
		}

		i++
	}
	return i
}

func main() {
	println(searchInsertPosition([]int{1,3,5,6}, 5))
}
