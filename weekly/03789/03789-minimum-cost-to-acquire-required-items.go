// ============================================================================
// 题号: 3789
// 题目: Minimum Cost to Acquire Required Items
// 难度: Medium
// 分类: Greedy, Math
// 链接: https://leetcode.cn/problems/minimum-cost-to-acquire-required-items/
// 每日一题日期: 2026-11-11
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你五个整数 cost1, cost2, costBoth, need1 和 need2。
//
// 有三种类型的物品可以购买：
//
// 类型 1 的物品花费 cost1，并仅满足类型 1 的需求 1 个单位。
//
// 类型 2 的物品花费 cost2，并仅满足类型 2 的需求 1 个单位。
//
// 类型 3 的物品花费 costBoth，同时满足类型 1 和类型 2 的需求各 1 个单位。
//
// 你需要购买足够的物品，使得：
//
// 满足类型 1 的总需求 至少 为 need1。
//
// 满足类型 2 的总需求 至少 为 need2。
//
// 返回满足这些需求的 最小 可能总花费。
//
// 示例 1：
//
// 输入： cost1 = 3, cost2 = 2, costBoth = 1, need1 = 3, need2 = 2
//
// 输出： 3
//
// 解释：
//
// 购买 3 个类型 3 的物品，总花费为 3 * 1 = 3，此时类型 1 的总需求满足 3（>= need1 = 3），类型 2 的总需求满足 3（>= need2 = 2）。
//
// 任何其他有效的购买方案都会花费更多，因此最小总花费为 3。
//
// 示例 2：
//
// 输入： cost1 = 5, cost2 = 4, costBoth = 15, need1 = 2, need2 = 3
//
// 输出： 22
//
// 解释：
//
// 购买 need1 = 2 个类型 1 的物品和 need2 = 3 个类型 2 的物品，总花费为 2 * 5 + 3 * 4 = 10 + 12 = 22。
//
// 任何其他有效的购买方案都会花费更多，因此最小总花费为 22。
//
// 示例 3：
//
// 输入： cost1 = 5, cost2 = 4, costBoth = 15, need1 = 0, need2 = 0
//
// 输出： 0
//
// 解释：
//
// 由于不需要任何物品（need1 = need2 = 0），因此无需购买，总花费为 0。
//
// 提示：
//
// 1 <= cost1, cost2, costBoth <= 106
//
// 0 <= need1, need2 <= 109
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given five integers cost1, cost2, costBoth, need1, and need2.
//
// There are three types of items available:
//
// An item of type 1 costs cost1 and contributes 1 unit to the type 1 requirement only.
//
// An item of type 2 costs cost2 and contributes 1 unit to the type 2 requirement only.
//
// An item of type 3 costs costBoth and contributes 1 unit to both type 1 and type 2 requirements.
//
// You must collect enough items so that the total contribution toward type 1 is at least need1 and the total contribution toward type 2 is at least need2.
//
// Return an integer representing the minimum possible total cost to achieve these requirements.
//
// Example 1:
//
// Input: cost1 = 3, cost2 = 2, costBoth = 1, need1 = 3, need2 = 2
//
// Output: 3
//
// Explanation:
//
// After buying three type 3 items, which cost 3 * 1 = 3, the total contribution to type 1 is 3 (>= need1 = 3) and to type 2 is 3 (>= need2 = 2).
//
// Any other valid combination would cost more, so the minimum total cost is 3.
//
// Example 2:
//
// Input: cost1 = 5, cost2 = 4, costBoth = 15, need1 = 2, need2 = 3
//
// Output: 22
//
// Explanation:
//
// We buy need1 = 2 items of type 1 and need2 = 3 items of type 2: 2 * 5 + 3 * 4 = 10 + 12 = 22.
//
// Any other valid combination would cost more, so the minimum total cost is 22.
//
// Example 3:
//
// Input: cost1 = 5, cost2 = 4, costBoth = 15, need1 = 0, need2 = 0
//
// Output: 0
//
// Explanation:
//
// Since no items are required (need1 = need2 = 0), we buy nothing and pay 0.
//
// Constraints:
//
// 1 <= cost1, cost2, costBoth <= 106
//
// 0 <= need1, need2 <= 109
// ============================================================================

package leetcode

func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
    
}
