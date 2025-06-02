# Reverse String

## Description
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

## Examples
Example 1:
~~~
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
~~~

Example 2:
~~~
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
~~~

## Solution
This implementation uses the classic two-pointer technique to reverse the array in-place.
It defines two pointers — left starting at the beginning and right at the end of the array.
On each iteration, the characters at these positions are swapped, and the pointers move toward each other.

This approach ensures:

In-place reversal without allocating any extra memory

O(n) time complexity, where n is the length of the array

O(1) space complexity

