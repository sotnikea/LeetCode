# Valid Anagram

## Description
Given two strings s and t, return true if t is an anagram of s, and false otherwise.   
Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

## Examples
Example 1:
~~~
Input: s = "anagram", t = "nagaram"

Output: true
~~~

Example 2:
~~~
Input: s = "rat", t = "car"

Output: false
~~~

## Constraints    
1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.

## Solution
This solution uses a hash map (map[rune]int) to count character frequencies in the first string and then adjusts the counts using the second string.

Key Steps:
- Count the frequency of each character in string s using a map.
- Iterate through string t, decrementing the corresponding counts in the map.
- If a character in t does not exist in the map or its count becomes zero too early, return false.
- Once a character’s count reaches zero, remove it from the map.
- If the map is empty after processing both strings, return true.

Supports full Unicode input thanks to use of rune.   
Time complexity: O(n), where n is the length of the strings.    
Space complexity: O(k) where k is the number of unique characters in s.