#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
按题号顺序拉取 LeetCode 题目,生成到 leetcode/ 目录。

逻辑:
1. 扫描 leetcode/ 目录,找出已有最大题号 N(空目录则 N=0)
2. 从 N+1 开始,向后拉取 10 道免费题(跳过付费题与不存在的题号)
3. 文件路径、命名、内容格式与 weekly/ 完全一致

数据源:
- https://leetcode.cn/api/problems/all/ :题目索引(题号→titleSlug,付费标记)
- https://leetcode.cn/graphql question(titleSlug) :题目详情

输出:leetcode/{题号范围如1-100}/{题号}-{titleSlug}.go
"""
import os
import sys
import time
from datetime import date

import requests

# 复用 fetch_leetcode 模块的生成逻辑,保证文件格式与 weekly/ 一致
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from fetch_leetcode import (
    HEADERS,
    fetch_question_detail,
    extract_go_snippet,
    html_to_text,
    build_go_content,
)

REPO_ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
LEETCODE_DIR = os.path.join(REPO_ROOT, "leetcode")
COUNT = 10  # 每次拉取的题目数量
PROBLEMS_API = "https://leetcode.cn/api/problems/all/"


def group_dir(qid):
    """题号按 100 个一组,返回分组目录名,如 1-100、101-200。"""
    start = ((qid - 1) // 100) * 100 + 1
    end = start + 99
    return f"{start}-{end}"


def find_max_id(target_dir):
    """扫描目录下所有 .go 文件名(题号5位补零前缀),返回最大题号;目录不存在或为空返回 0。"""
    if not os.path.isdir(target_dir):
        return 0
    max_id = 0
    for root, _dirs, files in os.walk(target_dir):
        for fname in files:
            # 文件名格式:{题号5位补零}-{slug}.go
            prefix = fname.split("-", 1)[0]
            if prefix.isdigit():
                max_id = max(max_id, int(prefix))
    return max_id


def fetch_problems_index():
    """拉取全部题目索引,返回 {frontend_id(int): {"slug": str, "paid": bool}} 字典。"""
    resp = requests.get(PROBLEMS_API, headers=HEADERS, timeout=20)
    resp.raise_for_status()
    data = resp.json()
    index = {}
    for item in data.get("stat_status_pairs", []):
        stat = item.get("stat", {})
        fid = stat.get("frontend_question_id")
        slug = stat.get("question__title_slug")
        paid = item.get("paid_only", False)
        if fid is None or not slug:
            continue
        try:
            fid = int(fid)
        except (TypeError, ValueError):
            continue
        index[fid] = {"slug": slug, "paid": paid}
    return index


def pick_next_targets(index, start_id, count):
    """从 start_id 开始,挑出接下来 count 道免费题(题号, slug),跳过付费与缺失题号。"""
    targets = []
    max_known = max(index.keys()) if index else 0
    qid = start_id
    # 允许跳过一定数量的缺失题号,但设置上限避免无限循环
    while len(targets) < count and qid <= max_known + 200:
        info = index.get(qid)
        if info is None:
            qid += 1
            continue
        if info["paid"]:
            print(f"⏭  {qid} 付费题,跳过")
            qid += 1
            continue
        targets.append((qid, info["slug"]))
        qid += 1
    return targets


def main():
    max_id = find_max_id(LEETCODE_DIR)
    print(f"leetcode/ 目录当前最大题号: {max_id}")

    print("拉取题目索引...")
    try:
        index = fetch_problems_index()
    except Exception as e:
        print(f"❌ 拉取题目索引失败: {e}", file=sys.stderr)
        return 0
    print(f"索引共 {len(index)} 道题")

    targets = pick_next_targets(index, max_id + 1, COUNT)
    print(f"将拉取 {len(targets)} 道题: {[t[0] for t in targets]}")

    os.makedirs(LEETCODE_DIR, exist_ok=True)
    new_count = 0
    today_str = date.today().isoformat()
    for qid, slug in targets:
        group = group_dir(qid)
        target_dir = os.path.join(LEETCODE_DIR, group)
        filename = f"{qid}-{slug}.go"
        target_file = os.path.join(target_dir, filename)
        if os.path.exists(target_file):
            print(f"⏭  {group}/{filename} 已存在,跳过")
            continue

        detail = fetch_question_detail(slug)
        if not detail:
            print(f"⚠️  {qid} ({slug}) 详情获取失败,跳过", file=sys.stderr)
            continue

        go_code = extract_go_snippet(detail.get("codeSnippets"))
        title = detail.get("title") or slug
        difficulty = detail.get("difficulty") or "Unknown"
        tags = [
            t.get("name")
            for t in (detail.get("topicTags") or [])
            if t.get("name")
        ]
        desc_cn = html_to_text(detail.get("translatedContent"))
        desc_en = html_to_text(detail.get("content"))

        # 复用 weekly 的生成逻辑,日期字段改用"拉取日期"(非每日一题)
        content = build_go_content(
            qid, title, slug, difficulty, tags, today_str,
            go_code, desc_cn, desc_en, date_label="拉取日期",
        )
        os.makedirs(target_dir, exist_ok=True)
        with open(target_file, "w", encoding="utf-8") as f:
            f.write(content)
        print(f"✅ 新增 {group}/{filename}")
        new_count += 1
        time.sleep(0.5)  # 适当限速

    print(f"\n共新增 {new_count} 道题目")
    return new_count


if __name__ == "__main__":
    sys.exit(0 if main() is not None else 1)
