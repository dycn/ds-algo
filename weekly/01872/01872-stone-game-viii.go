// ============================================================================
// 题号: 1872
// 题目: Stone Game VIII
// 难度: Hard
// 分类: Minimax, Array, Math, Dynamic Programming, Game Theory, Prefix Sum, 零和博弈
// 链接: https://leetcode.cn/problems/stone-game-viii/
// 每日一题日期: 2026-08-24
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// Alice 和 Bob 玩一个游戏，两人轮流操作， Alice 先手 。
//
// 总共有 n 个石子排成一行。轮到某个玩家的回合时，如果石子的数目 大于 1 ，他将执行以下操作：
//
// 选择一个整数 x > 1 ，并且 移除 最左边的 x 个石子。
//
// 将 移除 的石子价值之 和 累加到该玩家的分数中。
//
// 将一个 新的石子 放在最左边，且新石子的值为被移除石子值之和。
//
// 当只剩下 一个 石子时，游戏结束。
//
// Alice 和 Bob 的 分数之差 为 (Alice 的分数 - Bob 的分数) 。 Alice 的目标是 最大化 分数差，Bob 的目标是 最小化 分数差。
//
// 给你一个长度为 n 的整数数组 stones ，其中 stones[i] 是 从左边起 第 i 个石子的价值。请你返回在双方都采用 最优 策略的情况下，Alice 和 Bob 的 分数之差 。
//
// 示例 1：
//
// 输入：stones = [-1,2,-3,4,-5]
// 输出：5
// 解释：
// - Alice 移除最左边的 4 个石子，得分增加 (-1) + 2 + (-3) + 4 = 2 ，并且将一个价值为 2 的石子放在最左边。stones = [2,-5] 。
// - Bob 移除最左边的 2 个石子，得分增加 2 + (-5) = -3 ，并且将一个价值为 -3 的石子放在最左边。stones = [-3] 。
// 两者分数之差为 2 - (-3) = 5 。
//
// 示例 2：
//
// 输入：stones = [7,-6,5,10,5,-2,-6]
// 输出：13
// 解释：
// - Alice 移除所有石子，得分增加 7 + (-6) + 5 + 10 + 5 + (-2) + (-6) = 13 ，并且将一个价值为 13 的石子放在最左边。stones = [13] 。
// 两者分数之差为 13 - 0 = 13 。
//
// 示例 3：
//
// 输入：stones = [-10,-12]
// 输出：-22
// 解释：
// - Alice 只有一种操作，就是移除所有石子。得分增加 (-10) + (-12) = -22 ，并且将一个价值为 -22 的石子放在最左边。stones = [-22] 。
// 两者分数之差为 (-22) - 0 = -22 。
//
// 提示：
//
// n == stones.length
//
// 2 <= n <= 105
//
// -104 <= stones[i] <= 104
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Alice and Bob take turns playing a game, with Alice starting first.
//
// There are n stones arranged in a row. On each player's turn, while the number of stones is more than one, they will do the following:
//
// Choose an integer x > 1, and remove the leftmost x stones from the row.
//
// Add the sum of the removed stones' values to the player's score.
//
// Place a new stone, whose value is equal to that sum, on the left side of the row.
//
// The game stops when only one stone is left in the row.
//
// The score difference between Alice and Bob is (Alice's score - Bob's score). Alice's goal is to maximize the score difference, and Bob's goal is the minimize the score difference.
//
// Given an integer array stones of length n where stones[i] represents the value of the ith stone from the left, return the score difference between Alice and Bob if they both play optimally.
//
// Example 1:
//
// Input: stones = [-1,2,-3,4,-5]
// Output: 5
// Explanation:
// - Alice removes the first 4 stones, adds (-1) + 2 + (-3) + 4 = 2 to her score, and places a stone of
//   value 2 on the left. stones = [2,-5].
// - Bob removes the first 2 stones, adds 2 + (-5) = -3 to his score, and places a stone of value -3 on
//   the left. stones = [-3].
// The difference between their scores is 2 - (-3) = 5.
//
// Example 2:
//
// Input: stones = [7,-6,5,10,5,-2,-6]
// Output: 13
// Explanation:
// - Alice removes all stones, adds 7 + (-6) + 5 + 10 + 5 + (-2) + (-6) = 13 to her score, and places a
//   stone of value 13 on the left. stones = [13].
// The difference between their scores is 13 - 0 = 13.
//
// Example 3:
//
// Input: stones = [-10,-12]
// Output: -22
// Explanation:
// - Alice can only make one move, which is to remove both stones. She adds (-10) + (-12) = -22 to her
//   score and places a stone of value -22 on the left. stones = [-22].
// The difference between their scores is (-22) - 0 = -22.
//
// Constraints:
//
// n == stones.length
//
// 2 <= n <= 105
//
// -104 <= stones[i] <= 104
// ============================================================================

package leetcode

func stoneGameVIII(stones []int) int {
    
}
