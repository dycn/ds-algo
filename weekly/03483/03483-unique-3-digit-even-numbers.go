// ============================================================================
// 题号: 3483
// 题目: Unique 3-Digit Even Numbers
// 难度: Easy
// 分类: Recursion, Array, Hash Table, Enumeration
// 链接: https://leetcode.cn/problems/unique-3-digit-even-numbers/
// 每日一题日期: 2026-09-11
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个数字数组 digits，你需要从中选择三个数字组成一个三位偶数，你的任务是求出 不同 三位偶数的数量。
//
// 注意：每个数字在三位偶数中都只能使用 一次 ，并且 不能 有前导零。
//
// 示例 1：
//
// 输入： digits = [1,2,3,4]
//
// 输出： 12
//
// 解释： 可以形成的 12 个不同的三位偶数是 124，132，134，142，214，234，312，314，324，342，412 和 432。注意，不能形成 222，因为数字 2 只有一个。
//
// 示例 2：
//
// 输入： digits = [0,2,2]
//
// 输出： 2
//
// 解释： 可以形成的三位偶数是 202 和 220。注意，数字 2 可以使用两次，因为数组中有两个 2 。
//
// 示例 3：
//
// 输入： digits = [6,6,6]
//
// 输出： 1
//
// 解释： 只能形成 666。
//
// 示例 4：
//
// 输入： digits = [1,3,5]
//
// 输出： 0
//
// 解释： 无法形成三位偶数。
//
// 提示：
//
// 3 <= digits.length <= 10
//
// 0 <= digits[i] <= 9
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given an array of digits called digits. Your task is to determine the number of distinct three-digit even numbers that can be formed using these digits.
//
// Note: Each copy of a digit can only be used once per number, and there may not be leading zeros.
//
// Example 1:
//
// Input: digits = [1,2,3,4]
//
// Output: 12
//
// Explanation: The 12 distinct 3-digit even numbers that can be formed are 124, 132, 134, 142, 214, 234, 312, 314, 324, 342, 412, and 432. Note that 222 cannot be formed because there is only 1 copy of the digit 2.
//
// Example 2:
//
// Input: digits = [0,2,2]
//
// Output: 2
//
// Explanation: The only 3-digit even numbers that can be formed are 202 and 220. Note that the digit 2 can be used twice because it appears twice in the array.
//
// Example 3:
//
// Input: digits = [6,6,6]
//
// Output: 1
//
// Explanation: Only 666 can be formed.
//
// Example 4:
//
// Input: digits = [1,3,5]
//
// Output: 0
//
// Explanation: No even 3-digit numbers can be formed.
//
// Constraints:
//
// 3 <= digits.length <= 10
//
// 0 <= digits[i] <= 9
// ============================================================================

package leetcode

func totalNumbers(digits []int) int {
    
}
