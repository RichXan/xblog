import os
import re
from pathlib import Path
from datetime import datetime

def get_leetcode_files(directory):
    """获取所有的 leetcode 题解文件"""
    leetcode_dir = Path(directory)
    return sorted([f for f in leetcode_dir.glob("*.md") if f.name != "README.md"])

def parse_problem_info(file_path):
    """解析题目文件的信息"""
    filename = file_path.name
    # 匹配文件名中的题号和标题
    match = re.match(r"(\d{4})\s+(.+)\.md", filename)
    if not match:
        return None
    
    number, title = match.groups()
    
    # 读取文件内容获取更多信息
    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()
        # 获取文件的修改时间
        mod_time = os.path.getmtime(file_path)
        mod_date = datetime.fromtimestamp(mod_time).strftime('%Y-%m-%d')
        
        # 判断题目类型
        if "数组" in content or "Array" in content:
            category = "数组"
        elif "链表" in content or "Linked List" in content:
            category = "链表"
        elif "队列" in content or "Queue" in content:
            category = "队列"
        elif "栈" in content or "Stack" in content:
            category = "栈"
        elif "字符串" in content or "String" in content:
            category = "字符串"
        elif "二叉树" in content or "Binary Tree" in content:
            category = "二叉树"
        elif "动态规划" in content or "Dynamic Programming" in content:
            category = "动态规划"
        elif "回溯" in content or "Backtracking" in content:
            category = "回溯"
        elif "贪心" in content or "Greedy" in content:
            category = "贪心"
        elif "排序" in content or "Sorting" in content:
            category = "排序"
        elif "堆" in content or "Heap" in content:
            category = "堆"
        elif "图" in content or "Graph" in content:
            category = "图"
        elif "动态规划" in content or "Dynamic Programming" in content:
            category = "动态规划"
        else:
            category = "其他"
            
    return {
        "number": number,
        "title": title,
        "category": category,
        "path": f"/PersonalGrowth/CS/Algorithm/leetcode/{filename}",
        "date": mod_date
    }

def update_daily_checkin(problems):
    """生成每日打卡内容"""
    # 按日期分组
    problems_by_date = {}
    for prob in problems:
        date = prob["date"]
        if date not in problems_by_date:
            problems_by_date[date] = []
        problems_by_date[date].append(prob)
    
    # 生成打卡表格内容
    checkin_rows = []
    checkin_rows.append("## 每日打卡")
    checkin_rows.append("> 记录每天的刷题数量和重点题目，保持刷题习惯\n")
    checkin_rows.append("| 日期 | 题目数 | 重点题目 | 备注 |")
    checkin_rows.append("|------|--------|----------|------|")
    
    # 按日期倒序排列
    for date in sorted(problems_by_date.keys(), reverse=True):
        daily_problems = problems_by_date[date]
        problem_count = len(daily_problems)
        problem_numbers = [f"#{p['number']}" for p in daily_problems]
        category = daily_problems[0]["category"]  # 使用第一题的分类作为备注
        
        row = f"| {date} | {problem_count} | {', '.join(problem_numbers)} | {category} |"
        checkin_rows.append(row)
    
    return "\n".join(checkin_rows)

def update_readme(problems, readme_path):
    """更新 README.md 文件"""
    # 按类别组织题目
    problems_by_category = {}
    for prob in problems:
        category = prob["category"]
        if category not in problems_by_category:
            problems_by_category[category] = []
        problems_by_category[category].append(prob)
    
    # 读取原文件内容
    with open(readme_path, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # 定位到已完成题目部分
    start_marker = "## 已完成题目"
    end_marker = "## 每日打卡"
    
    pre_content = content.split(start_marker)[0]
    remaining_content = content.split(end_marker)[1].split("\n", 1)[1]
    
    # 生成新的已完成题目列表
    new_completed = [f"{start_marker}\n"]
    for category, probs in problems_by_category.items():
        new_completed.append(f"### {category}")
        for prob in sorted(probs, key=lambda x: x["number"]):
            new_completed.append(f"- [x] [{prob['number']} {prob['title']}]({prob['path']})")
        new_completed.append("")
    
    # 生成每日打卡内容
    daily_checkin = update_daily_checkin(problems)
    
    # 组合新内容
    new_content = (
        pre_content + 
        "\n".join(new_completed) +
        "\n" + daily_checkin +
        "\n" + remaining_content
    )
    
    # 写入文件
    with open(readme_path, 'w', encoding='utf-8') as f:
        f.write(new_content)

def main():
    base_dir = "PersonalGrowth/CS/Algorithm/leetcode"
    readme_path = os.path.join(base_dir, "README.md")
    
    # 获取所有题解文件
    leetcode_files = get_leetcode_files(base_dir)
    
    # 解析每个文件的信息
    problems = []
    for file_path in leetcode_files:
        info = parse_problem_info(file_path)
        if info:
            problems.append(info)
    
    # 更新 README.md
    update_readme(problems, readme_path)
    print(f"Updated {len(problems)} problems in README.md")

if __name__ == "__main__":
    main() 