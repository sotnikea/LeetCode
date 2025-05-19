# Two Sum

## Description
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## Examples
Example 1
~~~
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
~~~

Example 2
~~~
Input: nums = [3,2,4], target = 6
Output: [1,2]
~~~

Example 3
~~~
Input: nums = [3,3], target = 6
Output: [0,1]
~~~

Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

## Solution
This solution solves the Two Sum problem by using a map to track all indices of each number in the input slice nums. The algorithm works as follows:

Preprocessing step:
Build a map allNums where each key is a number from the array, and each value is a slice of all indices where that number occurs.

Main loop:
For each element val at index i, compute its complement target - val and check if it exists in the map.

If it does, iterate through all stored indices of the complement.

If a matching index is found and it's not equal to i, store the result.

The loop exits early once a valid pair of indices is found.

This method ensures that even if numbers repeat, multiple indices are considered correctly.

Time Complexity: O(n²) in the worst case (due to nested loops when handling repeated elements)

Space Complexity: O(n)