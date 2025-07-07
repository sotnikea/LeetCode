# Remove Nth Node From End of List

## Description
Given the head of a linked list, remove the nth node from the end of the list and return its head.

## Examples
Example 1:
~~~
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
~~~

Example 2:
~~~
Input: head = [1], n = 1
Output: []
~~~

Example 3:
~~~
Input: head = [1,2], n = 1
Output: [1]
~~~

## Constraints:
- The number of nodes in the list is sz.
- 1 <= sz <= 30
- 0 <= Node.val <= 100
- 1 <= n <= sz

# Follow up:   
Could you do this in one pass?

## Solution   
This implementation removes the nth node from the end of a singly-linked list in two passes within a single loop.   
It traverses the list while tracking the current position and dynamically determines when to start following the node that precedes the target node to be removed.   
If the node to delete is the head of the list, the algorithm correctly returns the new head.   
Edge cases such as single-element lists and removing the last node are explicitly handled.   
Time complexity: O(sz) — one full traversal of the list.   
Space complexity: O(1) — constant extra space.   
