# Single Number

## Description
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.    
You must implement a solution with a linear runtime complexity and use only constant extra space.

## Examples
Example 1
~~~
Input: nums = [2,2,1]
Output: 1
~~~

Example 2
~~~
Input: nums = [4,1,2,1,2]
Output: 4
~~~

Example 3
~~~
Input: nums = [1]
Output: 1
~~~

## Solution
This solution finds the unique number in an integer slice where every element appears exactly twice, except for one element that appears only once.

The algorithm uses the bitwise XOR (^) operation and relies on its key properties:

a ^ a = 0 — any number XORed with itself cancels out   
a ^ 0 = a — XOR with zero leaves the number unchanged.

XOR is associative and commutative, so the order of operations doesn't matter.

By XORing all elements of the array, all duplicate values cancel each other out, and only the unique number remains

Time Complexity: O(n)
One pass through the array    
Space Complexity: O(1)
Only one integer variable is used for accumulation.