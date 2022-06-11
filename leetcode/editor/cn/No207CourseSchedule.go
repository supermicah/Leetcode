package main

import "fmt"

//你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表
//示如果要学习课程 ai 则 必须 先学习课程 bi 。
//
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//
//
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：true
//解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//
// 示例 2：
//
//
//输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
//输出：false
//解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
//
//
//
// 提示：
//
//
// 1 <= numCourses <= 10⁵
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// prerequisites[i] 中的所有课程对 互不相同
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 1297 👎 0

func main() {
	//value := canFinish(2, [][]int{{1, 0}, {0, 1}})
	value := canFinish(2, [][]int{{1, 0}})
	no207Print("%+v", value)
}

func no207Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		// 有向边
		edges = make([][]int, numCourses)
		// 访问节点：0 未访问；1 访问中； 2 访问结束
		visited = make([]int, numCourses)
		// 是否有效
		valid = true
		// 递归函数
		canFinishDFS func(u int)
		// 课程顺序
		result []int
	)
	canFinishDFS = func(u int) {
		// 设置课程在进行中
		visited[u] = 1
		// 查询该元素的所有下游
		for _, v := range edges[u] {
			if visited[v] == 0 { // 如果节点未访问，则直接递归访问
				canFinishDFS(v)
				// 递归后，发现已经不符合，则退出
				if !valid {
					return
				}
			} else if visited[v] == 1 { //如果节点正在访问中，则直接失败
				valid = false
				return
			}
		}
		// 该节点访问结束
		visited[u] = 2
		// 返回结果收集，该结果倒叙，则是学习顺序
		result = append(result, u)
	}

	// 把课程依赖的所有节点都改为有向边，完成前面再完成后面
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	// 把每个课程进行访问
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			canFinishDFS(i)
		}

	}

	return valid
}

//leetcode submit region end(Prohibit modification and deletion)
