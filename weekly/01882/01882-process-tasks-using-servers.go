// ============================================================================
// 题号: 1882
// 题目: Process Tasks Using Servers
// 难度: Medium
// 分类: Array, Heap (Priority Queue)
// 链接: https://leetcode.cn/problems/process-tasks-using-servers/
// 每日一题日期: 2026-11-07
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你两个 下标从 0 开始 的整数数组 servers 和 tasks ，长度分别为 n​​​​​​ 和 m​​​​​​ 。servers[i] 是第 i​​​​​​​​​​ 台服务器的 权重 ，而 tasks[j] 是处理第 j​​​​​​ 项任务 所需要的时间（单位：秒）。
//
// 你正在运行一个仿真系统，在处理完所有任务后，该系统将会关闭。每台服务器只能同时处理一项任务。第 0 项任务在第 0 秒可以开始处理，相应地，第 j 项任务在第 j 秒可以开始处理。处理第 j 项任务时，你需要为它分配一台 权重最小 的空闲服务器。如果存在多台相同权重的空闲服务器，请选择 下标最小 的服务器。如果一台空闲服务器在第 t 秒分配到第 j 项任务，那么在 t + tasks[j] 时它将恢复空闲状态。
//
// 如果没有空闲服务器，则必须等待，直到出现一台空闲服务器，并 尽可能早 地处理剩余任务。 如果有多项任务等待分配，则按照 下标递增 的顺序完成分配。
//
// 如果同一时刻存在多台空闲服务器，可以同时将多项任务分别分配给它们。
//
// 构建长度为 m 的答案数组 ans ，其中 ans[j] 是第 j 项任务分配的服务器的下标。
//
// 返回答案数组 ans​​​​ 。
//
// 示例 1：
//
// 输入：servers = [3,3,2], tasks = [1,2,3,2,1,2]
// 输出：[2,2,0,2,1,2]
// 解释：事件按时间顺序如下：
// - 0 秒时，第 0 项任务加入到任务队列，使用第 2 台服务器处理到 1 秒。
// - 1 秒时，第 2 台服务器空闲，第 1 项任务加入到任务队列，使用第 2 台服务器处理到 3 秒。
// - 2 秒时，第 2 项任务加入到任务队列，使用第 0 台服务器处理到 5 秒。
// - 3 秒时，第 2 台服务器空闲，第 3 项任务加入到任务队列，使用第 2 台服务器处理到 5 秒。
// - 4 秒时，第 4 项任务加入到任务队列，使用第 1 台服务器处理到 5 秒。
// - 5 秒时，所有服务器都空闲，第 5 项任务加入到任务队列，使用第 2 台服务器处理到 7 秒。
//
// 示例 2：
//
// 输入：servers = [5,1,4,3,2], tasks = [2,1,2,4,5,2,1]
// 输出：[1,4,1,4,1,3,2]
// 解释：事件按时间顺序如下：
// - 0 秒时，第 0 项任务加入到任务队列，使用第 1 台服务器处理到 2 秒。
// - 1 秒时，第 1 项任务加入到任务队列，使用第 4 台服务器处理到 2 秒。
// - 2 秒时，第 1 台和第 4 台服务器空闲，第 2 项任务加入到任务队列，使用第 1 台服务器处理到 4 秒。
// - 3 秒时，第 3 项任务加入到任务队列，使用第 4 台服务器处理到 7 秒。
// - 4 秒时，第 1 台服务器空闲，第 4 项任务加入到任务队列，使用第 1 台服务器处理到 9 秒。
// - 5 秒时，第 5 项任务加入到任务队列，使用第 3 台服务器处理到 7 秒。
// - 6 秒时，第 6 项任务加入到任务队列，使用第 2 台服务器处理到 7 秒。
//
// 提示：
//
// servers.length == n
//
// tasks.length == m
//
// 1 <= n, m <= 2 * 105
//
// 1 <= servers[i], tasks[j] <= 2 * 105
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given two 0-indexed integer arrays servers and tasks of lengths n​​​​​​ and m​​​​​​ respectively. servers[i] is the weight of the i​​​​​​th​​​​ server, and tasks[j] is the time needed to process the j​​​​​​th​​​​ task in seconds.
//
// Tasks are assigned to the servers using a task queue. Initially, all servers are free, and the queue is empty.
//
// At second j, the jth task is inserted into the queue (starting with the 0th task being inserted at second 0). As long as there are free servers and the queue is not empty, the task in the front of the queue will be assigned to a free server with the smallest weight, and in case of a tie, it is assigned to a free server with the smallest index.
//
// If there are no free servers and the queue is not empty, we wait until a server becomes free and immediately assign the next task. If multiple servers become free at the same time, then multiple tasks from the queue will be assigned in order of insertion following the weight and index priorities above.
//
// A server that is assigned task j at second t will be free again at second t + tasks[j].
//
// Build an array ans​​​​ of length m, where ans[j] is the index of the server the j​​​​​​th task will be assigned to.
//
// Return the array ans​​​​.
//
// Example 1:
//
// Input: servers = [3,3,2], tasks = [1,2,3,2,1,2]
// Output: [2,2,0,2,1,2]
// Explanation: Events in chronological order go as follows:
// - At second 0, task 0 is added and processed using server 2 until second 1.
// - At second 1, server 2 becomes free. Task 1 is added and processed using server 2 until second 3.
// - At second 2, task 2 is added and processed using server 0 until second 5.
// - At second 3, server 2 becomes free. Task 3 is added and processed using server 2 until second 5.
// - At second 4, task 4 is added and processed using server 1 until second 5.
// - At second 5, all servers become free. Task 5 is added and processed using server 2 until second 7.
//
// Example 2:
//
// Input: servers = [5,1,4,3,2], tasks = [2,1,2,4,5,2,1]
// Output: [1,4,1,4,1,3,2]
// Explanation: Events in chronological order go as follows:
// - At second 0, task 0 is added and processed using server 1 until second 2.
// - At second 1, task 1 is added and processed using server 4 until second 2.
// - At second 2, servers 1 and 4 become free. Task 2 is added and processed using server 1 until second 4.
// - At second 3, task 3 is added and processed using server 4 until second 7.
// - At second 4, server 1 becomes free. Task 4 is added and processed using server 1 until second 9.
// - At second 5, task 5 is added and processed using server 3 until second 7.
// - At second 6, task 6 is added and processed using server 2 until second 7.
//
// Constraints:
//
// servers.length == n
//
// tasks.length == m
//
// 1 <= n, m <= 2 * 105
//
// 1 <= servers[i], tasks[j] <= 2 * 105
// ============================================================================

package leetcode

func assignTasks(servers []int, tasks []int) []int {
    
}
