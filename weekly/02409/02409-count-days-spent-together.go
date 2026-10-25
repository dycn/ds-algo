// ============================================================================
// 题号: 2409
// 题目: Count Days Spent Together
// 难度: Easy
// 分类: Math, String
// 链接: https://leetcode.cn/problems/count-days-spent-together/
// 每日一题日期: 2026-10-24
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// Alice 和 Bob 计划分别去罗马开会。
//
// 给你四个字符串 arriveAlice ，leaveAlice ，arriveBob 和 leaveBob 。Alice 会在日期 arriveAlice 到 leaveAlice 之间在城市里（日期为闭区间），而 Bob 在日期 arriveBob 到 leaveBob 之间在城市里（日期为闭区间）。每个字符串都包含 5 个字符，格式为 "MM-DD" ，对应着一个日期的月和日。
//
// 请你返回 Alice和 Bob 同时在罗马的天数。
//
// 你可以假设所有日期都在 同一个 自然年，而且 不是 闰年。每个月份的天数分别为：[31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31] 。
//
// 示例 1：
//
// 输入：arriveAlice = "08-15", leaveAlice = "08-18", arriveBob = "08-16", leaveBob = "08-19"
// 输出：3
// 解释：Alice 从 8 月 15 号到 8 月 18 号在罗马。Bob 从 8 月 16 号到 8 月 19 号在罗马，他们同时在罗马的日期为 8 月 16、17 和 18 号。所以答案为 3 。
//
// 示例 2：
//
// 输入：arriveAlice = "10-01", leaveAlice = "10-31", arriveBob = "11-01", leaveBob = "12-31"
// 输出：0
// 解释：Alice 和 Bob 没有同时在罗马的日子，所以我们返回 0 。
//
// 提示：
//
// 所有日期的格式均为 "MM-DD" 。
//
// Alice 和 Bob 的到达日期都 早于或等于 他们的离开日期。
//
// 题目测试用例所给出的日期均为 非闰年 的有效日期。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Alice and Bob are traveling to Rome for separate business meetings.
//
// You are given 4 strings arriveAlice, leaveAlice, arriveBob, and leaveBob. Alice will be in the city from the dates arriveAlice to leaveAlice (inclusive), while Bob will be in the city from the dates arriveBob to leaveBob (inclusive). Each will be a 5-character string in the format "MM-DD", corresponding to the month and day of the date.
//
// Return the total number of days that Alice and Bob are in Rome together.
//
// You can assume that all dates occur in the same calendar year, which is not a leap year. Note that the number of days per month can be represented as: [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31].
//
// Example 1:
//
// Input: arriveAlice = "08-15", leaveAlice = "08-18", arriveBob = "08-16", leaveBob = "08-19"
// Output: 3
// Explanation: Alice will be in Rome from August 15 to August 18. Bob will be in Rome from August 16 to August 19. They are both in Rome together on August 16th, 17th, and 18th, so the answer is 3.
//
// Example 2:
//
// Input: arriveAlice = "10-01", leaveAlice = "10-31", arriveBob = "11-01", leaveBob = "12-31"
// Output: 0
// Explanation: There is no day when Alice and Bob are in Rome together, so we return 0.
//
// Constraints:
//
// All dates are provided in the format "MM-DD".
//
// Alice and Bob's arrival dates are earlier than or equal to their leaving dates.
//
// The given dates are valid dates of a non-leap year.
// ============================================================================

package leetcode

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
    
}
