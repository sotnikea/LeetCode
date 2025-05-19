# Move Zeroes

## Description
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

## Examples
Example 1
~~~
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
~~~

Example 2
~~~
Input: nums = [0]
Output: [0]
~~~

## Solution
The approach scans the array from left to right. For each zero encountered, it searches for the next non-zero element and swaps them. If no more non-zero elements remain, it breaks early.

Key steps:
1. Iterate over the array up to the second-to-last element.
2. When a zero is found, search forward for the next non-zero element.
3. If found, swap the zero with that non-zero element.
4. If no swap occurs (remaining elements are all zero), terminate early.

Time Complexity: O(n²) in the worst case   
For every zero, it may need to scan the rest of the array to find a non-zero.

Space Complexity: O(1)   
All operations are performed in-place using only a constant number of variables.