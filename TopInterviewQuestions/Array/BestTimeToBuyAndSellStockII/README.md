# Best Time to Buy and Sell Stock II

## Description
You are given an integer array prices where prices[i] is the price of a given stock on the i-th day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.

## Examples
Example 1:
~~~
Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.
~~~

Example 2:
~~~
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.
~~~

Example 3:
~~~
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
~~~

## Solution
The approach scans the prices array to find every local minimum followed by a local maximum:

It first advances minIndex to the next valley (where price is lower than following prices).

Then it advances maxIndex to the next peak (where price is higher or equal to previous ones).   
The profit for each such valley-to-peak pair is added to the total.   
The search continues from the next index after the last sell day.   
The loop breaks when no more valid buy-sell pairs can be found.

Time Complexity
Time Complexity: O(n)
Each index is visited at most twice.
