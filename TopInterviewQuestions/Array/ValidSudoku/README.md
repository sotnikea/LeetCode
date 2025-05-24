# Valid Sudoku

## Description
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

- Each row must contain the digits 1-9 without repetition.
- Each column must contain the digits 1-9 without repetition.
- Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.   
Note:
- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

## Examples
Example 1
~~~
Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
~~~

Example 2
~~~
Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
~~~

## Solution
This solution validates a partially filled 9×9 Sudoku board by ensuring that:
- Each row contains the digits 1–9 without repetition.
- Each column contains the digits 1–9 without repetition.
- Each 3×3 sub-box contains the digits 1–9 without repetition.
The algorithm performs three separate passes:
- One for validating each row.
- One for validating each column.
- One for validating each 3×3 sub-grid.

A temporary map (map[byte]bool) is used in each check to track already seen digits. If a duplicate is detected during any check, the function returns false. Otherwise, it returns true after all validations succeed.

Implementation Notes
- 0x2e is the hexadecimal representation of the '.' character (ASCII code 46), used to skip empty cells.
- The 3×3 sub-boxes are traversed by iterating over their top-left corners using step size 3.

Time Complexity: O(1)   
The board size is fixed (9×9), so the total number of operations is constant.

Space Complexity: O(1)   
Although temporary maps are used, their size is bounded and constant due to the fixed board size.

