// ============================================================================
// 题号: 3622
// 题目: Check Divisibility by Digit Sum and Product
// 难度: Easy
// 分类: Math
// 链接: https://leetcode.cn/problems/check-divisibility-by-digit-sum-and-product/
// 每日一题日期: 2026-08-22
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个正整数 n。请判断 n 是否可以被以下两值之和 整除：
//
// n 的 数字和（即其各个位数之和）。
//
// n 的 数字积（即其各个位数之积）。
//
// 如果 n 能被该和整除，返回 true；否则，返回 false。
//
// 示例 1：
//
// 输入： n = 99
//
// 输出： true
//
// 解释：
//
// 因为 99 可以被其数字和 (9 + 9 = 18) 与数字积 (9 * 9 = 81) 之和 (18 + 81 = 99) 整除，因此输出为 true。
//
// 示例 2：
//
// 输入： n = 23
//
// 输出： false
//
// 解释：
//
// 因为 23 无法被其数字和 (2 + 3 = 5) 与数字积 (2 * 3 = 6) 之和 (5 + 6 = 11) 整除，因此输出为 false。
//
// 提示：
//
// 1 <= n <= 106
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a positive integer n. Determine whether n is divisible by the sum of the following two values:
//
// The digit sum of n (the sum of its digits).
//
// The digit product of n (the product of its digits).
//
// Return true if n is divisible by this sum; otherwise, return false.
//
// Example 1:
//
// Input: n = 99
//
// Output: true
//
// Explanation:
//
// Since 99 is divisible by the sum (9 + 9 = 18) plus product (9 * 9 = 81) of its digits (total 99), the output is true.
//
// Example 2:
//
// Input: n = 23
//
// Output: false
//
// Explanation:
//
// Since 23 is not divisible by the sum (2 + 3 = 5) plus product (2 * 3 = 6) of its digits (total 11), the output is false.
//
// Constraints:
//
// 1 <= n <= 106
// ============================================================================

package leetcode

func checkDivisibility(n int) bool {
    
}
