# Valid Palindrome

## Description
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

## Examples
Example 1:
~~~
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
~~~

Example 2:
~~~
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.    
~~~

Example 3:
~~~
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
~~~

## Constraints:

1 <= s.length <= 2 * 105    
s consists only of printable ASCII characters.

## Solution
This solution processes the string in two phases:

1. Filtering and Normalization:
- Converts all uppercase letters to lowercase.
- Skips all non-alphanumeric characters.
- Stores the normalized characters in the forward slice.
2. Palindrome Check:
- Builds a backward slice by reversing the forward one.
- Compares both slices using slices.Equal() to check if they are identical.

Time Complexity: O(n)    
Space Complexity: O(n) due to use of two slices.

