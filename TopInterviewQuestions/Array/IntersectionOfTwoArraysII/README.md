#  Intersection of Two Arrays II

## Description
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

## Examples
Example 1
~~~
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
~~~

Example 2
~~~
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.
~~~

## Solution
The approach works as follows:    
A frequency map is built for nums1, where the key is the number and the value is its count.   
The second array nums2 is iterated through, and for each element, the map is checked for availability (count > 0).   
If present, the element is added to the result, and its count in the map is decremented.   
This ensures that each element appears in the result exactly as many times as it appears in both arrays.

Time Complexity: O(n + m)
Where n is the length of nums1 and m is the length of nums2.

Space Complexity: O(min(n, m))
The frequency map is built from the smaller array to optimize memory usage.
