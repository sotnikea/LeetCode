# Reverse Integer

## Description
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

## Examples
Example 1:
~~~
Input: x = 123
Output: 321
~~~

Example 2:
~~~
Input: x = -123
Output: -321
~~~

Example 3:
~~~
Input: x = 120
Output: 21
~~~

## Solution
This solution implements integer reversal using arithmetic operations and ensures the result fits within the 32-bit signed integer range.

Key Steps:
- Define 32-bit bounds using math.Pow:    
maxInt32 = 2³¹ - 1    
minInt32 = -2³¹

- Extract digits from x using modulus and division.
- Build the reversed integer by shifting previous result and appending the current digit.
- Check overflow after the full reversal using the defined bounds.
