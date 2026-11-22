// ============================================================================
// 题号: 2914
// 题目: Minimum Number of Changes to Make Binary String Beautiful
// 难度: Medium
// 分类: String
// 链接: https://leetcode.cn/problems/minimum-number-of-changes-to-make-binary-string-beautiful/
// 每日一题日期: 2026-11-18
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个长度为偶数下标从 0 开始的二进制字符串 s 。
//
// 如果可以将一个字符串分割成一个或者更多满足以下条件的子字符串，那么我们称这个字符串是 美丽的 ：
//
// 每个子字符串的长度都是 偶数 。
//
// 每个子字符串都 只 包含 1 或 只 包含 0 。
//
// 你可以将 s 中任一字符改成 0 或者 1 。
//
// 请你返回让字符串 s 美丽的 最少 字符修改次数。
//
// 示例 1：
//
// 输入：s = "1001"
// 输出：2
// 解释：我们将 s[1] 改为 1 ，且将 s[3] 改为 0 ，得到字符串 "1100" 。
// 字符串 "1100" 是美丽的，因为我们可以将它分割成 "11|00" 。
// 将字符串变美丽最少需要 2 次修改。
//
// 示例 2：
//
// 输入：s = "10"
// 输出：1
// 解释：我们将 s[1] 改为 1 ，得到字符串 "11" 。
// 字符串 "11" 是美丽的，因为它已经是美丽的。
// 将字符串变美丽最少需要 1 次修改。
//
// 示例 3：
//
// 输入：s = "0000"
// 输出：0
// 解释：不需要进行任何修改，字符串 "0000" 已经是美丽字符串。
//
// 提示：
//
// 2 <= s.length <= 105
//
// s 的长度为偶数。
//
// s[i] 要么是 '0' ，要么是 '1' 。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a 0-indexed binary string s having an even length.
//
// A string is beautiful if it's possible to partition it into one or more substrings such that:
//
// Each substring has an even length.
//
// Each substring contains only 1's or only 0's.
//
// You can change any character in s to 0 or 1.
//
// Return the minimum number of changes required to make the string s beautiful.
//
// Example 1:
//
// Input: s = "1001"
// Output: 2
// Explanation: We change s[1] to 1 and s[3] to 0 to get string "1100".
// It can be seen that the string "1100" is beautiful because we can partition it into "11|00".
// It can be proven that 2 is the minimum number of changes needed to make the string beautiful.
//
// Example 2:
//
// Input: s = "10"
// Output: 1
// Explanation: We change s[1] to 1 to get string "11".
// It can be seen that the string "11" is beautiful because we can partition it into "11".
// It can be proven that 1 is the minimum number of changes needed to make the string beautiful.
//
// Example 3:
//
// Input: s = "0000"
// Output: 0
// Explanation: We don't need to make any changes as the string "0000" is beautiful already.
//
// Constraints:
//
// 2 <= s.length <= 105
//
// s has an even length.
//
// s[i] is either '0' or '1'.
// ============================================================================

package leetcode

func minChanges(s string) int {
    
}
