# Plus One

## Description
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.

## Examples
Example 1
~~~
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
~~~

Example 2
~~~
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
~~~

Example 3
~~~
Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].
~~~

## Solution
The algorithm works as follows:   
1. It creates a copy of the original digits slice to avoid modifying the input.
2. Starting from the last digit, it checks if the digit is 9. If so, it sets it to 0 and continues left.
3. If a digit is less than 9, it is incremented and the loop breaks.
4. If all digits were 9, a 1 is prepended to the slice (e.g., [9, 9] becomes [1, 0, 0]).

The use of append ensures correct memory handling when prepending the leading 1.

Time Complexity: O(n)   
Worst-case scenario is when all digits are 9, requiring a full pass.

Space Complexity: O(n)   
A new slice is allocated as a copy of the original, and potentially extended by one element.