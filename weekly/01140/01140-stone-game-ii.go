// ============================================================================
// 题号: 1140
// 题目: Stone Game II
// 难度: Medium
// 分类: Minimax, Array, Math, Dynamic Programming, Game Theory, Prefix Sum, 零和博弈
// 链接: https://leetcode.cn/problems/stone-game-ii/
// 每日一题日期: 2026-08-09
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// Alice 和 Bob 继续他们的石子游戏。许多堆石子 排成一行，每堆都有正整数颗石子 piles[i]。游戏以谁手中的石子最多来决出胜负。
//
// Alice 和 Bob 轮流进行，Alice 先开始。最初，M = 1。
//
// 在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。
//
// 游戏一直持续到所有石子都被拿走。
//
// 假设 Alice 和 Bob 都发挥出最佳水平，返回 Alice 可以得到的最大数量的石头。
//
// 示例 1：
//
// 输入：piles = [2,7,9,4,4]
// 输出：10
// 解释：如果一开始 Alice 取了一堆，Bob 取了两堆，然后 Alice 再取两堆。Alice 可以得到 2 + 4 + 4 = 10 堆。
// 如果 Alice 一开始拿走了两堆，那么 Bob 可以拿走剩下的三堆。在这种情况下，Alice 得到 2 + 7 = 9 堆。返回 10，因为它更大。
//
// 示例 2:
//
// 输入：piles = [1,2,3,4,5,100]
// 输出：104
//
// 提示：
//
// 1 <= piles.length <= 100
//
// 1 <= piles[i] <= 104
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Alice and Bob continue their games with piles of stones. There are a number of piles arranged in a row, and each pile has a positive integer number of stones piles[i]. The objective of the game is to end with the most stones.
//
// Alice and Bob take turns, with Alice starting first.
//
// On each player's turn, that player can take all the stones in the first X remaining piles, where 1 <= X <= 2M. Then, we set M = max(M, X). Initially, M = 1.
//
// The game continues until all the stones have been taken.
//
// Assuming Alice and Bob play optimally, return the maximum number of stones Alice can get.
//
// Example 1:
//
// Input: piles = [2,7,9,4,4]
//
// Output: 10
//
// Explanation:
//
// If Alice takes one pile at the beginning, Bob takes two piles, then Alice takes 2 piles again. Alice can get 2 + 4 + 4 = 10 stones in total.
//
// If Alice takes two piles at the beginning, then Bob can take all three piles left. In this case, Alice get 2 + 7 = 9 stones in total.
//
// So we return 10 since it's larger.
//
// Example 2:
//
// Input: piles = [1,2,3,4,5,100]
//
// Output: 104
//
// Constraints:
//
// 1 <= piles.length <= 100
//
// 1 <= piles[i] <= 104
// ============================================================================

package leetcode

func stoneGameII(piles []int) int {
    
}
