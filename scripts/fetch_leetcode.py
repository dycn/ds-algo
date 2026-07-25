#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
从 LeetCode 中国站拉取最近 7 天的每日一题,生成 Go 题目模板文件。

数据源:https://leetcode.cn/graphql
- dailyQuestionRecords(year, month):获取某月每日一题列表
- question(titleSlug):获取题目详情(含 Go 代码模板、题目描述)

输出:weekly/{题号5位补0}/{题号5位补0}-{titleSlug}.go
"""
import os
import sys
import time
import html as html_module
from datetime import date, timedelta
from html.parser import HTMLParser

import requests

GRAPHQL_URL = "https://leetcode.cn/graphql"
HEADERS = {
    "Content-Type": "application/json",
    "User-Agent": (
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 "
        "(KHTML, like Gecko) Chrome/120.0 Safari/537.36"
    ),
    "Origin": "https://leetcode.cn",
    "Referer": "https://leetcode.cn/",
}

REPO_ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
SOLUTIONS_DIR = os.path.join(REPO_ROOT, "weekly")
DAYS_BACK = 7

# 按月查询每日一题记录
DAILY_QUERY = """
query dailyQuestionRecords($year: Int!, $month: Int!) {
  dailyQuestionRecords(year: $year, month: $month) {
    date
    question {
      questionFrontendId
      title
      titleSlug
      difficulty
      topicTags { name slug }
    }
  }
}
"""

# 查询单题详情(含各语言代码模板、中英文题目描述)
QUESTION_QUERY = """
query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionFrontendId
    title
    titleSlug
    difficulty
    topicTags { name slug }
    codeSnippets { lang langSlug code }
    content
    translatedContent
  }
}
"""


class HTMLToText(HTMLParser):
    """将 LeetCode 题目描述 HTML 转成可读纯文本,保留段落与代码块换行。"""

    BLOCK_TAGS = {
        "p", "div", "br", "li", "tr", "h1", "h2", "h3", "h4", "h5", "h6",
        "blockquote", "ul", "ol", "table",
    }

    def __init__(self):
        super().__init__()
        self.parts = []
        self.in_pre = 0

    def handle_starttag(self, tag, attrs):
        if tag == "pre":
            self.in_pre += 1
            self.parts.append("\n")
        elif tag in self.BLOCK_TAGS:
            self.parts.append("\n")

    def handle_endtag(self, tag):
        if tag == "pre" and self.in_pre > 0:
            self.in_pre -= 1
            self.parts.append("\n")
        elif tag in self.BLOCK_TAGS:
            self.parts.append("\n")

    def handle_data(self, data):
        self.parts.append(data)

    def get_text(self):
        text = html_module.unescape("".join(self.parts))
        lines = text.split("\n")
        result = []
        blank = False
        for line in lines:
            stripped = line.rstrip()
            if not stripped:
                if not blank:
                    result.append("")
                    blank = True
            else:
                result.append(stripped)
                blank = False
        return "\n".join(result).strip()


def html_to_text(html_str):
    """HTML 转纯文本,空输入返回空串。"""
    if not html_str:
        return ""
    parser = HTMLToText()
    parser.feed(html_str)
    return parser.get_text()


def comment_block(text, header_line):
    """把多行文本包装成 // 注释块,带分隔标题行。"""
    lines = text.split("\n") if text else ["(无)"]
    out = ["// ============================================================================",
           f"// {header_line}",
           "// ============================================================================"]
    for line in lines:
        out.append(f"// {line}" if line else "//")
    out.append("// ============================================================================")
    return "\n".join(out)


def gql(query, variables):
    """发送 GraphQL 请求,返回 data 字段。"""
    resp = requests.post(
        GRAPHQL_URL,
        json={"query": query, "variables": variables},
        headers=HEADERS,
        timeout=15,
    )
    resp.raise_for_status()
    payload = resp.json()
    if "errors" in payload:
        raise RuntimeError(payload["errors"])
    return payload.get("data") or {}


def fetch_daily_for_month(year, month):
    """获取某月每日一题记录列表,失败时返回空列表。"""
    try:
        data = gql(DAILY_QUERY, {"year": year, "month": month})
        return data.get("dailyQuestionRecords") or []
    except Exception as e:
        print(f"⚠️  查询 {year}-{month:02d} 每日一题失败: {e}", file=sys.stderr)
        return []


def fetch_question_detail(title_slug):
    """获取单题详情(含 Go 代码模板与题目描述),失败时返回 None。"""
    try:
        data = gql(QUESTION_QUERY, {"titleSlug": title_slug})
        return data.get("question")
    except Exception as e:
        print(f"⚠️  获取题目 {title_slug} 详情失败: {e}", file=sys.stderr)
        return None


def extract_go_snippet(code_snippets):
    """从 codeSnippets 中提取 Go 代码模板。"""
    if not code_snippets:
        return ""
    for s in code_snippets:
        if s.get("lang") == "Go" or s.get("langSlug") == "golang":
            return s.get("code", "") or ""
    return ""


def pad_qid(qid):
    """题号补零至 5 位,保证目录排序顺序从 1 开始递增。"""
    return str(qid).zfill(5)


def normalize_date(raw):
    """将日期字符串规范化为 YYYY-MM-DD(补零),便于排序。"""
    if not raw:
        return ""
    parts = raw.strip().split("-")
    if len(parts) == 3:
        return f"{parts[0]}-{int(parts[1]):02d}-{int(parts[2]):02d}"
    return raw.strip()


def build_go_content(qid, title, slug, difficulty, tags, date_str,
                     go_code, desc_cn, desc_en, date_label="每日一题日期"):
    """构造 .go 文件内容:头注释 + 题目描述(中英) + Go 代码模板。"""
    tags_str = ", ".join(tags) if tags else "Uncategorized"
    header = (
        "// ============================================================================\n"
        f"// 题号: {qid}\n"
        f"// 题目: {title}\n"
        f"// 难度: {difficulty}\n"
        f"// 分类: {tags_str}\n"
        f"// 链接: https://leetcode.cn/problems/{slug}/\n"
        f"// {date_label}: {date_str}\n"
        "// ============================================================================\n\n"
    )
    desc_section = ""
    if desc_cn:
        desc_section += comment_block(desc_cn, "题目描述(中文)") + "\n\n"
    if desc_en:
        desc_section += comment_block(desc_en, "题目描述(英文)") + "\n\n"

    # LeetCode 返回的 Go 模板通常不含 package 声明,统一补上
    if "package " in go_code:
        body = go_code.rstrip() + "\n"
    elif go_code.strip():
        body = "package leetcode\n\n" + go_code.rstrip() + "\n"
    else:
        body = "package leetcode\n\n// TODO: 在此实现题目要求的方法\n"
    return header + desc_section + body


def main():
    today = date.today()
    # 查询当月与上月,覆盖跨月场景
    last_month_day = today.replace(day=1) - timedelta(days=1)
    months_to_query = {
        (today.year, today.month),
        (last_month_day.year, last_month_day.month),
    }

    # 收集所有每日一题记录
    records = []
    for y, m in months_to_query:
        records.extend(fetch_daily_for_month(y, m))

    # 规范化日期并按日期降序排序,取最近 DAYS_BACK 条
    for r in records:
        r["date"] = normalize_date(r.get("date", ""))
    records = [r for r in records if r.get("question")]
    records.sort(key=lambda r: r.get("date", ""), reverse=True)
    recent = records[:DAYS_BACK]

    print(f"最近 {DAYS_BACK} 天找到 {len(recent)} 道每日一题")

    os.makedirs(SOLUTIONS_DIR, exist_ok=True)
    new_count = 0
    for r in recent:
        q = r.get("question") or {}
        qid = q.get("questionFrontendId")
        slug = q.get("titleSlug")
        if not qid or not slug:
            continue

        qid_padded = pad_qid(qid)
        target_dir = os.path.join(SOLUTIONS_DIR, qid_padded)
        if os.path.exists(target_dir):
            print(f"⏭  {qid_padded}/{slug} 已存在,跳过")
            continue

        # 拉取详情以获取 Go 代码模板与题目描述(列表接口可能不含这些字段)
        detail = fetch_question_detail(slug) or q
        go_code = extract_go_snippet(detail.get("codeSnippets"))
        title = detail.get("title") or q.get("title") or slug
        difficulty = detail.get("difficulty") or q.get("difficulty") or "Unknown"
        tags = [
            t.get("name")
            for t in (detail.get("topicTags") or q.get("topicTags") or [])
            if t.get("name")
        ]
        date_str = r.get("date", "unknown")
        # 中英文题目描述(HTML 转纯文本)
        desc_cn = html_to_text(detail.get("translatedContent"))
        desc_en = html_to_text(detail.get("content"))

        content = build_go_content(
            qid, title, slug, difficulty, tags, date_str,
            go_code, desc_cn, desc_en,
        )
        os.makedirs(target_dir, exist_ok=True)
        filename = f"{qid_padded}-{slug}.go"
        with open(os.path.join(target_dir, filename), "w", encoding="utf-8") as f:
            f.write(content)
        print(f"✅ 新增 {qid_padded}/{filename}")
        new_count += 1
        time.sleep(0.5)  # 适当限速,降低被风控概率

    print(f"\n共新增 {new_count} 道题目")
    return new_count


if __name__ == "__main__":
    sys.exit(0 if main() is not None else 1)
