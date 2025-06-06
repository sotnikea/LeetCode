# Rotate Image

## Description
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

## Examples
Example 1
~~~
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
~~~

Example 2
~~~
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
~~~

## Solution
This solution performs an in-place 90-degree clockwise rotation of an n x n matrix, representing an image. It rotates the matrix layer by layer, starting from the outermost layer and moving inward.

For each layer:
- Elements from the top, right, bottom, and left sides are referenced using pointers
- A 4-way cyclic swap is performed: top ← left, right ← top, bottom ← right, left ← bottom.
- This allows the rotation to be done without allocating extra space and without relying on matrix transposition or row reversal.

Implementation Notes   
- The loop runs over each layer (or "ring") of the matrix.
- For each element within a layer, its four corresponding sides are identified and rotated in place using pointer assignments.
- The core rotation relies on computing symmetrical positions relative to the current layer and the matrix size.

Time Complexity: O(n^2/4)
Space Complexity: O(1)
