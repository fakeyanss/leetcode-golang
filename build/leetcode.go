package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// 题目信息结构体（新增 IsLCP/LCPNum 字段）
type LeetCodeProblem struct {
	Num        int    // 普通题序号
	LCPNum     int    // LCP题序号
	IsLCP      bool   // 是否是LCP题
	Title      string // 题目名称
	URL        string // 题目链接
	FilePath   string // 代码文件相对路径
	Difficulty string // 难度（简单/中等/困难）
	Solution   string // 解法描述
}

const (
	// 配置项
	codeDir          = "leetcode"  // 刷题代码目录
	readmePath       = "README.md" // README路径
	tableStartMarker = "<!-- LEETCODE_TABLE_START -->"
	tableEndMarker   = "<!-- LEETCODE_TABLE_END -->"
	repoRoot         = "." // 仓库根目录
)

// 遍历目录解析所有题目文件
func parseAllProblems() ([]LeetCodeProblem, error) {
	var normalProblems []LeetCodeProblem // 普通LC题
	var lcpProblems []LeetCodeProblem    // LCP题

	// 遍历指定目录下的所有.go文件
	err := filepath.Walk(codeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 跳过目录、测试文件
		if info.IsDir() || strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}
		// 只处理.go文件
		if filepath.Ext(info.Name()) != ".go" {
			return nil
		}

		// 解析单个文件
		problem, err := parseProblemFile(path)
		if err != nil {
			fmt.Printf("解析文件失败 %s: %v\n", path, err)
			return nil // 跳过解析失败的文件
		}

		// 区分普通题和LCP题
		if problem.IsLCP {
			lcpProblems = append(lcpProblems, problem)
		} else if problem.Num > 0 {
			normalProblems = append(normalProblems, problem)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("遍历目录失败: %w", err)
	}

	// 普通题按数字升序排序
	sort.Slice(normalProblems, func(i, j int) bool {
		return normalProblems[i].Num < normalProblems[j].Num
	})

	// LCP题按LCPNum升序排序
	sort.Slice(lcpProblems, func(i, j int) bool {
		return lcpProblems[i].LCPNum < lcpProblems[j].LCPNum
	})

	// 合并：普通题在前，LCP题在后
	allProblems := append(normalProblems, lcpProblems...)
	return allProblems, nil
}

// 解析单个题目文件（新增LCP解析逻辑）
func parseProblemFile(filePath string) (LeetCodeProblem, error) {
	var problem LeetCodeProblem
	fileName := filepath.Base(filePath)

	// 处理代码文件路径
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		problem.FilePath = filePath // 兜底使用原始路径
	} else {
		rootAbs, _ := filepath.Abs(repoRoot)
		relPath, _ := filepath.Rel(rootAbs, absPath)
		problem.FilePath = relPath // 存储相对于仓库根目录的路径
	}

	// ========== 1. 优先解析LCP题 ==========
	lcpNumRegex := regexp.MustCompile(`LCP\s+(\d+)`)
	lcpNumMatches := lcpNumRegex.FindStringSubmatch(fileName)
	if lcpNumMatches != nil && len(lcpNumMatches) >= 2 {
		problem.IsLCP = true
		_, _ = fmt.Sscanf(lcpNumMatches[1], "%d", &problem.LCPNum)
		// 从文件夹名二次验证（如lcLCP 20）
		if problem.LCPNum == 0 {
			dirName := filepath.Base(filepath.Dir(filePath))
			dirLcpMatches := lcpNumRegex.FindStringSubmatch(dirName)
			if dirLcpMatches != nil && len(dirLcpMatches) >= 2 {
				_, _ = fmt.Sscanf(dirLcpMatches[1], "%d", &problem.LCPNum)
			}
		}
		if problem.LCPNum == 0 {
			return problem, fmt.Errorf("未提取到LCP题目序号")
		}
	} else {
		// ========== 2. 解析普通LC题 ==========
		numRegex := regexp.MustCompile(`lc(\d+)|(\d+)\.`)
		numMatches := numRegex.FindStringSubmatch(fileName)
		if numMatches != nil {
			var numStr string
			if numMatches[1] != "" {
				numStr = numMatches[1]
			} else if numMatches[2] != "" {
				numStr = numMatches[2]
			}
			_, _ = fmt.Sscanf(numStr, "%d", &problem.Num)
		}
		if problem.Num == 0 {
			return problem, fmt.Errorf("未提取到题目序号")
		}
	}

	// ========== 3. 读取文件内容解析注释 ==========
	file, err := os.Open(filePath)
	if err != nil {
		return problem, fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return problem, fmt.Errorf("读取文件内容失败: %w", err)
	}
	fileContent := content.String()

	// ========== 4. 提取题目标题 ==========
	// 优先匹配 // [151] 反转字符串中的单词 或 // [LCP 20] 快速公交 格式
	titleRegex1 := regexp.MustCompile(`\[([^\]]+)\]\s*([^\n]+)`)
	titleMatches1 := titleRegex1.FindStringSubmatch(fileContent)
	if titleMatches1 != nil && len(titleMatches1) >= 3 {
		problem.Title = strings.TrimSpace(titleMatches1[2])
	} else {
		// 备用：从文件名提取
		if problem.IsLCP {
			// LCP文件名：LCP 20.快速公交.go → 快速公交
			titleRegex2 := regexp.MustCompile(`LCP\s+\d+\.([^.]+)\.go`)
			titleMatches2 := titleRegex2.FindStringSubmatch(fileName)
			if titleMatches2 != nil && len(titleMatches2) >= 2 {
				problem.Title = strings.TrimSpace(titleMatches2[1])
			} else {
				problem.Title = "未知标题"
			}
		} else {
			// 普通文件名：151.反转字符串中的单词.go → 反转字符串中的单词
			titleRegex2 := regexp.MustCompile(`\d+\.([^.]+)\.go`)
			titleMatches2 := titleRegex2.FindStringSubmatch(fileName)
			if titleMatches2 != nil && len(titleMatches2) >= 2 {
				problem.Title = strings.TrimSpace(titleMatches2[1])
			} else {
				problem.Title = "未知标题"
			}
		}
	}

	// ========== 5. 提取题目难度 ==========
	difficultyMap := map[string]string{
		"Easy":   "Easy",
		"Medium": "Medium",
		"Hard":   "Hard",
		"简单":     "Easy",
		"中等":     "Medium",
		"困难":     "Hard",
	}
	difficultyRegex := regexp.MustCompile(`(Easy|Medium|Hard|简单|中等|困难)`)
	difficultyMatches := difficultyRegex.FindStringSubmatch(fileContent)
	if difficultyMatches != nil {
		problem.Difficulty = difficultyMap[difficultyMatches[1]]
	} else {
		problem.Difficulty = "未知"
	}

	// ========== 6. 提取题目URL ==========
	urlRegex := regexp.MustCompile(`https://leetcode\.cn/problems/[^(\n)]+`)
	urlMatches := urlRegex.FindStringSubmatch(fileContent)
	if urlMatches != nil {
		problem.URL = strings.TrimSpace(urlMatches[0])
	} else {
		problem.URL = ""
	}

	// ========== 7. 提取解法描述 ==========
	solutionRegex1 := regexp.MustCompile(`//\s*(解法|思路)[:：]\s*([^\n]+)`)
	solutionMatches1 := solutionRegex1.FindStringSubmatch(fileContent)
	if solutionMatches1 != nil && len(solutionMatches1) >= 3 {
		problem.Solution = strings.TrimSpace(solutionMatches1[2])
	} else {
		algorithmRegex := regexp.MustCompile(`//\s*(栈|递归|动态规划|贪心|二叉堆|前缀和|字典树|LRU|DP|BFS|DFS)[^(\n)]+`)
		algorithmMatches := algorithmRegex.FindStringSubmatch(fileContent)
		if algorithmMatches != nil {
			problem.Solution = algorithmMatches[1]
		} else {
			problem.Solution = ""
		}
	}

	return problem, nil
}

// 生成Markdown表格（适配LCP序号展示）
func generateMarkdownTable(problems []LeetCodeProblem) string {
	var table strings.Builder
	// 表格头部
	table.WriteString("| 序号 | 题目 | 文件 | 难度 | 解法 |\n")
	table.WriteString("|---|---|---|---|---|\n")

	// 表格行
	for _, p := range problems {
		// 处理序号展示：普通题显示数字，LCP题显示 LCP {数字}
		var numCell string
		if p.IsLCP {
			numCell = fmt.Sprintf("LCP %d", p.LCPNum)
		} else {
			numCell = fmt.Sprintf("%d", p.Num)
		}

		// 处理题目链接
		titleCell := fmt.Sprintf("[%s](%s)", p.Title, p.URL)
		// 处理文件链接
		fileCell := fmt.Sprintf("[Link](%s)", p.FilePath)

		// 拼接行
		table.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
			numCell, titleCell, fileCell, p.Difficulty, p.Solution))
	}

	return table.String()
}

// 更新README中的表格
func updateReadme(tableContent string) error {
	// 读取README内容
	readmeContent, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("读取README失败: %w", err)
	}

	// 替换标记间的内容
	contentStr := string(readmeContent)
	startIdx := strings.Index(contentStr, tableStartMarker)
	endIdx := strings.Index(contentStr, tableEndMarker)
	if startIdx == -1 || endIdx == -1 || startIdx > endIdx {
		return fmt.Errorf("未找到表格标记（%s / %s）", tableStartMarker, tableEndMarker)
	}

	// 拼接新内容
	newContent := fmt.Sprintf(
		"%s\n%s\n%s",
		contentStr[:startIdx+len(tableStartMarker)],
		tableContent,
		contentStr[endIdx:],
	)

	// 写入README
	err = os.WriteFile(readmePath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("写入README失败: %w", err)
	}

	return nil
}

func main() {
	// 1. 解析所有题目
	problems, err := parseAllProblems()
	if err != nil {
		fmt.Printf("解析题目失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("成功解析 %d 道题目（普通题+LCP题）\n", len(problems))

	// 2. 生成表格
	table := generateMarkdownTable(problems)

	// 3. 更新README
	err = updateReadme(table)
	if err != nil {
		fmt.Printf("更新README失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ README表格更新成功！")
}
