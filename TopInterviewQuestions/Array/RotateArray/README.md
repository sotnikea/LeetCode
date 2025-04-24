# Remove Duplicates from Sorted Array

## Description
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with O(1) extra space?

## Examples
Example 1
~~~
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
~~~

Example 2
~~~
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation: 
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
~~~

## Solution

Solution 1
```go
func rotate1(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	tmpArr := make([]int, k)

	// Save k elements
	for i := 0; i < k; i++ {
		tmpArr[i] = nums[len(nums)-1-i]
	}

	// Change element position
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}

	// Set first k element from tmp array
	for i := 0; i < k; i++ {
		nums[i] = tmpArr[k-i-1]
	}
}
```
This solution rotates an integer slice nums to the right by k steps. If k is greater than the length of the array, it is reduced modulo the array length to avoid unnecessary full rotations.

The approach consists of three main steps:
- Backup the last k elements into a temporary slice (tmpArr), storing them in reverse order.
- Shift the remaining elements in-place k positions to the right, starting from the end of the array.
- Restore the saved elements from tmpArr to the beginning of the array, reversing them back to maintain the original order.

This method ensures that the rotation is performed in-place except for the temporary buffer of size k.

Time Complexity: O(n)
All elements are visited a constant number of times.

Space Complexity: O(k)
A temporary slice of size k is used to store the elements that would be overwritten during the shift.

Solution 2
```go
func rotate2(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	for i := 0; i < k; i++ {

		j := len(nums) - 1
		tmp := nums[j]

		for ; j > 0; j-- {
			nums[j] = nums[j-1]
		}

		nums[0] = tmp
	}
}
```
This solution rotates an integer slice nums to the right by k steps using a naive iterative approach. If k is greater than the length of the array, it is reduced modulo the array length to avoid redundant full rotations.

The algorithm performs the rotation by repeating a one-step right shift k times. For each step:
- Save the last element of the array (tmp).
- Shift all elements one position to the right.
- Insert tmp at the first position.

While simple and easy to understand, this approach can be inefficient for large values of k.

Time Complexity: O(n × k)
Each of the k steps involves shifting n elements.

Space Complexity: O(1)
Only one temporary variable is used; the operation is performed entirely in-place.

Solution 3
```go
func rotate3(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	tmpArr := make([]int, len(nums))

	tmpPos := 0
	for i := len(nums) - k; i < len(nums); i++ {
		tmpArr[tmpPos] = nums[i]
		tmpPos++
	}

	for i := 0; i < len(nums)-k; i++ {
		tmpArr[tmpPos] = nums[i]
		tmpPos++
	}

	copy(nums, tmpArr)
}
```
This solution rotates an integer slice nums to the right by k steps using an auxiliary array of the same size. If k is greater than the array length, it is reduced modulo len(nums) to eliminate unnecessary full rotations.

The logic consists of:
- Copying the last k elements (which will become the new prefix of the rotated array) into a temporary buffer.
- Copying the remaining elements (from the beginning of the array up to the rotation point) into the buffer.
- Copying the reordered buffer back into the original slice using copy.

This method ensures the correct order after rotation while keeping the implementation simple and safe.

Time Complexity: O(n)
Every element is copied exactly twice.

Space Complexity: O(n)
An additional slice of the same length is used to hold the rotated result.