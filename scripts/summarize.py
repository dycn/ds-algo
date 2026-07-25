#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
整理 leetcode/ 与 weekly/ 下的题解情况,汇总后追加到 README.md 末尾。

归纳内容:
- 总览(总题数 / 已解 / 未解)
- 按来源统计(weekly / leetcode)
- 按难度统计(Easy / Medium / Hard)
- 明细表格(题号、题目、难度、分类、来源、状态、链接)

更新方式:
- 以 "## 题解统计(自动生成)" 标题为锚点实现幂等:首次运行追加到 README 末尾,
  后续运行删除该标题到文件末尾的旧内容后重新追加,不会累加。
- 不使用 HTML 注释标记(部分预览器会把标记间内容当注释隐藏,导致预览看不到)。

"已解"判定:去掉注释、空行、package/func 签名、单独的 `}` 后,若仍有代码行则视为已解。
"""
import os
import re
import sys

REPO_ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
README = os.path.join(REPO_ROOT, "README.md")
DIRS = ["weekly", "leetcode"]
SUMMARY_HEADER = "## 题解统计(自动生成)"
# 旧版用 HTML 注释标记,部分预览器会隐藏标记间内容,此处仅用于清理迁移
LEGACY_MARKS = ["<!-- LEETCODE-SUMMARY-START -->", "<!-- LEETCODE-SUMMARY-END -->"]


def is_solved(content):
    """判断题解是否已实现:函数体内是否有实质代码。"""
    # 去掉块注释 /* ... */(LeetCode 常把 ListNode 等类型定义放在 /** ... */ 注释里)
    content = re.sub(r'/\*.*?\*/', '', content, flags=re.DOTALL)
    meaningful = []
    for line in content.split("\n"):
        s = line.strip()
        if not s or s.startswith("//"):
            continue
        meaningful.append(s)
    # 去掉 package 行、func 签名行、单独的 `}` 行,剩下即为函数体实现代码
    code_lines = [
        l for l in meaningful
        if not l.startswith("package")
        and not l.startswith("func")
        and l != "}"
    ]
    return len(code_lines) > 0


def parse_go_file(filepath):
    """解析 .go 文件头注释,返回题目录字典。"""
    with open(filepath, encoding="utf-8") as f:
        content = f.read()
    meta = {
        "qid": "", "title": "", "difficulty": "",
        "tags": "", "link": "", "date": "",
    }
    for line in content.split("\n"):
        s = line.strip()
        if not s.startswith("//"):
            break
        # 去掉 "//" 前缀
        body = s[2:].strip()
        # 各字段格式 "键: 值",链接含冒号,用 split(":", 1) 只切第一个
        if body.startswith("题号:"):
            meta["qid"] = body.split(":", 1)[1].strip()
        elif body.startswith("题目:"):
            meta["title"] = body.split(":", 1)[1].strip()
        elif body.startswith("难度:"):
            meta["difficulty"] = body.split(":", 1)[1].strip()
        elif body.startswith("分类:"):
            meta["tags"] = body.split(":", 1)[1].strip()
        elif body.startswith("链接:"):
            meta["link"] = body.split(":", 1)[1].strip()
        elif body.startswith("每日一题日期:") or body.startswith("拉取日期:"):
            meta["date"] = body.split(":", 1)[1].strip()
    meta["solved"] = is_solved(content)
    return meta


def scan_dir(name):
    """扫描指定目录下所有 .go 文件,返回题目录列表。"""
    base = os.path.join(REPO_ROOT, name)
    if not os.path.isdir(base):
        return []
    results = []
    for root, _dirs, files in os.walk(base):
        for f in sorted(files):
            if f.endswith(".go"):
                full = os.path.join(root, f)
                meta = parse_go_file(full)
                meta["source"] = name
                # 相对仓库根的路径,用于在 README 中生成可点击的题解链接
                meta["code_path"] = os.path.relpath(full, REPO_ROOT)
                results.append(meta)
    return results


def sort_key(e):
    try:
        return (0, int(e["qid"]))
    except (TypeError, ValueError):
        return (1, 0)


def build_summary(entries):
    """构建汇总 markdown 文本:leetcode 与 weekly 各一个表格,格式参考练习进度表。"""
    entries.sort(key=sort_key)

    # 难度英文转中文(与练习进度表一致)
    diff_map = {"Easy": "简单", "Medium": "中等", "Hard": "困难"}

    # 按来源分组
    by_source = {"leetcode": [], "weekly": []}
    for e in entries:
        by_source.setdefault(e["source"], []).append(e)

    lines = []
    lines.append("## 题解统计(自动生成)")
    lines.append("")
    lines.append(f"> 由 `scripts/summarize.py` 生成,最后更新于 {date_today()}。请勿手动编辑此区块。")
    lines.append("")

    for source in ["leetcode", "weekly"]:
        items = by_source.get(source, [])
        if not items:
            continue
        total = len(items)
        solved = sum(1 for e in items if e["solved"])
        lines.append(f'<h1 style="text-align:center">{source}({solved}/{total})</h1>')
        lines.append("")
        lines.append("| 题号 | 题目 | 类别 | 难度 | 状态 |")
        lines.append("|:---|:---|:---:|:---:|:---:|")
        for e in items:
            status = "✅" if e["solved"] else "⬜"
            title = e["title"] or "-"
            title_linked = f"[{title}]({e['link']})" if e["link"] else title
            # 类别用 <br/> 分隔,与练习进度表一致
            tags = e["tags"].replace(", ", "<br/>") if e["tags"] else ""
            diff = diff_map.get(e["difficulty"], e["difficulty"] or "")
            lines.append(
                f"| {e['qid']} | {title_linked} | {tags} | {diff} | {status} |"
            )
        lines.append("")

    return "\n".join(lines)


def date_today():
    from datetime import date
    return date.today().isoformat()


def update_readme(summary_text):
    """幂等更新 README:以汇总标题为锚点,有则替换到末尾,无则追加到末尾。"""
    with open(README, encoding="utf-8") as f:
        content = f.read()
    # 清理可能残留的旧版 HTML 注释标记(会致部分预览器隐藏内容)
    for mark in LEGACY_MARKS:
        content = content.replace(mark + "\n", "")
        content = content.replace(mark, "")
    if SUMMARY_HEADER in content:
        # 删除旧汇总块(从标题到文件末尾),保留其前的内容
        idx = content.index(SUMMARY_HEADER)
        content = content[:idx].rstrip() + "\n\n"
        action = "替换"
    else:
        if not content.endswith("\n"):
            content += "\n"
        content += "\n"
        action = "追加"
    content += summary_text + "\n"
    with open(README, "w", encoding="utf-8") as f:
        f.write(content)
    return action


def main():
    entries = []
    for d in DIRS:
        items = scan_dir(d)
        print(f"扫描 {d}/: {len(items)} 道题")
        entries.extend(items)

    print(f"\n共 {len(entries)} 道题,生成汇总...")
    summary = build_summary(entries)
    action = update_readme(summary)
    print(f"✓ 已{action}到 {os.path.relpath(README, REPO_ROOT)}")


if __name__ == "__main__":
    main()
