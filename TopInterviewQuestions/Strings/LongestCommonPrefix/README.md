# Longest Common Prefix

## Description
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

## Examples
Example 1:
~~~
Input: strs = ["flower","flow","flight"]
Output: "fl"
~~~

Example 2:
~~~
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
~~~

# Constraints:

- 1 <= strs.length <= 200
- 0 <= strs[i].length <= 200
- strs[i] consists of only lowercase English letters if it is non-empty.

## Solution
This implementation handles edge cases such as empty arrays or single-element arrays.    
It iteratively compares each character of the first string with the corresponding characters of the remaining strings.    
As soon as a mismatch is found or the shortest string ends, the loop stops and returns the prefix accumulated so far.   
Time complexity: O(n * m) — where n is the number of strings and m is the length of the shortest string.    
Space complexity: O(1) — no extra space is used beyond variables.