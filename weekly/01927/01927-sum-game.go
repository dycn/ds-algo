// ============================================================================
// 题号: 1927
// 题目: Sum Game
// 难度: Medium
// 分类: Greedy, Math, String, Game Theory
// 链接: https://leetcode.cn/problems/sum-game/
// 每日一题日期: 2026-08-23
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// Alice 和 Bob 玩一个游戏，两人轮流行动，Alice 先手 。
//
// 给你一个 偶数长度 的字符串 num ，每一个字符为数字字符或者 '?' 。每一次操作中，如果 num 中至少有一个 '?' ，那么玩家可以执行以下操作：
//
// 选择一个下标 i 满足 num[i] == '?' 。
//
// 将 num[i] 用 '0' 到 '9' 之间的一个数字字符替代。
//
// 当 num 中没有 '?' 时，游戏结束。
//
// Bob 获胜的条件是 num 中前一半数字的和 等于 后一半数字的和。Alice 获胜的条件是前一半的和与后一半的和 不相等 。
//
// 比方说，游戏结束时 num = "243801" ，那么 Bob 获胜，因为 2+4+3 = 8+0+1 。如果游戏结束时 num = "243803" ，那么 Alice 获胜，因为 2+4+3 != 8+0+3 。
//
// 在 Alice 和 Bob 都采取 最优 策略的前提下，如果 Alice 获胜，请返回 true ，如果 Bob 获胜，请返回 false 。
//
// 示例 1：
//
// 输入：num = "5023"
// 输出：false
// 解释：num 中没有 '?' ，没法进行任何操作。
// 前一半的和等于后一半的和：5 + 0 = 2 + 3 。
//
// 示例 2：
//
// 输入：num = "25??"
// 输出：true
// 解释：Alice 可以将两个 '?' 中的一个替换为 '9' ，Bob 无论如何都无法使前一半的和等于后一半的和。
//
// 示例 3：
//
// 输入：num = "?3295???"
// 输出：false
// 解释：Bob 总是能赢。一种可能的结果是：
// - Alice 将第一个 '?' 用 '9' 替换。num = "93295???" 。
// - Bob 将后面一半中的一个 '?' 替换为 '9' 。num = "932959??" 。
// - Alice 将后面一半中的一个 '?' 替换为 '2' 。num = "9329592?" 。
// - Bob 将后面一半中最后一个 '?' 替换为 '7' 。num = "93295927" 。
// Bob 获胜，因为 9 + 3 + 2 + 9 = 5 + 9 + 2 + 7 。
//
// 提示：
//
// 2 <= num.length <= 105
//
// num.length 是 偶数 。
//
// num 只包含数字字符和 '?' 。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Alice and Bob take turns playing a game, with Alice starting first.
//
// You are given a string num of even length consisting of digits and '?' characters. On each turn, a player will do the following if there is still at least one '?' in num:
//
// Choose an index i where num[i] == '?'.
//
// Replace num[i] with any digit between '0' and '9'.
//
// The game ends when there are no more '?' characters in num.
//
// For Bob to win, the sum of the digits in the first half of num must be equal to the sum of the digits in the second half. For Alice to win, the sums must not be equal.
//
// For example, if the game ended with num = "243801", then Bob wins because 2+4+3 = 8+0+1. If the game ended with num = "243803", then Alice wins because 2+4+3 != 8+0+3.
//
// Assuming Alice and Bob play optimally, return true if Alice will win and false if Bob will win.
//
// Example 1:
//
// Input: num = "5023"
// Output: false
// Explanation: There are no moves to be made.
// The sum of the first half is equal to the sum of the second half: 5 + 0 = 2 + 3.
//
// Example 2:
//
// Input: num = "25??"
// Output: true
// Explanation: Alice can replace one of the '?'s with '9' and it will be impossible for Bob to make the sums equal.
//
// Example 3:
//
// Input: num = "?3295???"
// Output: false
// Explanation: It can be proven that Bob will always win. One possible outcome is:
// - Alice replaces the first '?' with '9'. num = "93295???".
// - Bob replaces one of the '?' in the right half with '9'. num = "932959??".
// - Alice replaces one of the '?' in the right half with '2'. num = "9329592?".
// - Bob replaces the last '?' in the right half with '7'. num = "93295927".
// Bob wins because 9 + 3 + 2 + 9 = 5 + 9 + 2 + 7.
//
// Constraints:
//
// 2 <= num.length <= 105
//
// num.length is even.
//
// num consists of only digits and '?'.
// ============================================================================

package leetcode

func sumGame(num string) bool {
    
}
