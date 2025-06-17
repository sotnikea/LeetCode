# Implement strStr()

## Description
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

## Examples
Example 1:
~~~
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
~~~

Example 2:
~~~
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
~~~

# Constraints:

1 <= haystack.length, needle.length <= 104    
haystack and needle consist of only lowercase English characters.


## Solution
- Iterate through haystack to find a character that matches the first character of needle.
- When a match is found, try to match the entire substring.
- If mismatch occurs mid-way, reset and continue from the next character after the previous start.
- Return the index if full match succeeds, otherwise -1.

Time Complexity:   
Worst-case: O(n × m), where n is the length of haystack and m is the length of needle.
