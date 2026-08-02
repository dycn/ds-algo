// ============================================================================
// 题号: 3014
// 题目: Minimum Number of Pushes to Type Word I
// 难度: Easy
// 分类: Greedy, Math, String
// 链接: https://leetcode.cn/problems/minimum-number-of-pushes-to-type-word-i/
// 每日一题日期: 2026-07-30
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个字符串 word，由 不同 小写英文字母组成。
//
// 电话键盘上的按键与 不同 小写英文字母集合相映射，可以通过按压按键来组成单词。例如，按键 2 对应 ["a","b","c"]，我们需要按一次键来输入 "a"，按两次键来输入 "b"，按三次键来输入 "c"。
//
// 现在允许你将编号为 2 到 9 的按键重新映射到 不同 字母集合。每个按键可以映射到 任意数量 的字母，但每个字母 必须 恰好 映射到 一个 按键上。你需要找到输入字符串 word 所需的 最少 按键次数。
//
// 返回重新映射按键后输入 word 所需的 最少 按键次数。
//
// 下面给出了一种电话键盘上字母到按键的映射作为示例。注意 1，*，# 和 0 不 对应任何字母。
//
// 示例 1：
//
// 输入：word = "abcde"
// 输出：5
// 解释：图片中给出的重新映射方案的输入成本最小。
// "a" -> 在按键 2 上按一次
// "b" -> 在按键 3 上按一次
// "c" -> 在按键 4 上按一次
// "d" -> 在按键 5 上按一次
// "e" -> 在按键 6 上按一次
// 总成本为 1 + 1 + 1 + 1 + 1 = 5 。
// 可以证明不存在其他成本更低的映射方案。
//
// 示例 2：
//
// 输入：word = "xycdefghij"
// 输出：12
// 解释：图片中给出的重新映射方案的输入成本最小。
// "x" -> 在按键 2 上按一次
// "y" -> 在按键 2 上按两次
// "c" -> 在按键 3 上按一次
// "d" -> 在按键 3 上按两次
// "e" -> 在按键 4 上按一次
// "f" -> 在按键 5 上按一次
// "g" -> 在按键 6 上按一次
// "h" -> 在按键 7 上按一次
// "i" -> 在按键 8 上按一次
// "j" -> 在按键 9 上按一次
// 总成本为 1 + 2 + 1 + 2 + 1 + 1 + 1 + 1 + 1 + 1 = 12 。
// 可以证明不存在其他成本更低的映射方案。
//
// 提示：
//
// 1 <= word.length <= 26
//
// word 仅由小写英文字母组成。
//
// word 中的所有字母互不相同。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a string word containing distinct lowercase English letters.
//
// Telephone keypads have keys mapped with distinct collections of lowercase English letters, which can be used to form words by pushing them. For example, the key 2 is mapped with ["a","b","c"], we need to push the key one time to type "a", two times to type "b", and three times to type "c" .
//
// It is allowed to remap the keys numbered 2 to 9 to distinct collections of letters. The keys can be remapped to any amount of letters, but each letter must be mapped to exactly one key. You need to find the minimum number of times the keys will be pushed to type the string word.
//
// Return the minimum number of pushes needed to type word after remapping the keys.
//
// An example mapping of letters to keys on a telephone keypad is given below. Note that 1, *, #, and 0 do not map to any letters.
//
// Example 1:
//
// Input: word = "abcde"
// Output: 5
// Explanation: The remapped keypad given in the image provides the minimum cost.
// "a" -> one push on key 2
// "b" -> one push on key 3
// "c" -> one push on key 4
// "d" -> one push on key 5
// "e" -> one push on key 6
// Total cost is 1 + 1 + 1 + 1 + 1 = 5.
// It can be shown that no other mapping can provide a lower cost.
//
// Example 2:
//
// Input: word = "xycdefghij"
// Output: 12
// Explanation: The remapped keypad given in the image provides the minimum cost.
// "x" -> one push on key 2
// "y" -> two pushes on key 2
// "c" -> one push on key 3
// "d" -> two pushes on key 3
// "e" -> one push on key 4
// "f" -> one push on key 5
// "g" -> one push on key 6
// "h" -> one push on key 7
// "i" -> one push on key 8
// "j" -> one push on key 9
// Total cost is 1 + 2 + 1 + 2 + 1 + 1 + 1 + 1 + 1 + 1 = 12.
// It can be shown that no other mapping can provide a lower cost.
//
// Constraints:
//
// 1 <= word.length <= 26
//
// word consists of lowercase English letters.
//
// All letters in word are distinct.
// ============================================================================

package leetcode

func minimumPushes(word string) int {
    
}
