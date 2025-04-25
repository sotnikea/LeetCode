# Contains Duplicate

## Description
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

## Examples
Example 1
~~~
Input: nums = [1,2,3,1]
Output: true
Explanation:
The element 1 occurs at the indices 0 and 3.
~~~

Example 2
~~~
Input: nums = [1,2,3,4]
Output: false
Explanation:
All elements are distinct.
~~~

Example 3
~~~
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
~~~

## Solution

This solution determines whether an integer slice nums contains any duplicate elements. It uses a hash map (map[int]int) to count the occurrences of each number.    
The algorithm proceeds in two phases:    
- Count each element: Iterate over the array, incrementing the count for each number in the map.
- Check for duplicates: Iterate over the map values and return true if any value is greater than 1, indicating a duplicate.
- If no duplicates are found, the function returns false.
    
Time Complexity: O(n)
One pass to build the map and one pass to check for duplicates.

Space Complexity: O(n)
Up to n keys could be stored in the map, where n is the number of elements in the array.

