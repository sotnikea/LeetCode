# First Unique Character in a String

## Description
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

## Examples
Example 1:
~~~
Input: s = "leetcode"

Output: 0

Explanation:

The character 'l' at index 0 is the first character that does not occur at any other index.
~~~

Example 2:
~~~
Input: s = "loveleetcode"

Output: 2
~~~

Example 3:
~~~
Input: s = "aabb"

Output: -1
~~~

## Constraints:

1 <= s.length <= 105     
s consists of only lowercase English letters.

## Solution
This solution uses a nested loop approach with a map to track already-checked characters, minimizing redundant comparisons.

Key Steps:
Create a map[byte]bool to track characters that have already been identified as duplicates.

Iterate over the string:
- If the current character was already checked, skip it.
- Otherwise, scan the rest of the string to see if it repeats.
- If it does, mark it and any duplicate as checked.
- If it doesn’t repeat, return its index.

If no unique character is found, return -1.

Time complexity is O(n²) in the worst case due to nested loops.    
Memory usage is O(1), since the alphabet size is limited (lowercase English letters only).
