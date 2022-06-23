package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。
//
// 请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。
//
// 如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。
//
// 军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。
//
//
//
// 示例 1：
//
//
//输入：mat =
//[[1,1,0,0,0],
// [1,1,1,1,0],
// [1,0,0,0,0],
// [1,1,0,0,0],
// [1,1,1,1,1]],
//k = 3
//输出：[2,0,3]
//解释：
//每行中的军人数目：
//行 0 -> 2
//行 1 -> 4
//行 2 -> 1
//行 3 -> 2
//行 4 -> 5
//从最弱到最强对这些行排序后得到 [2,0,3,1,4]
//
//
// 示例 2：
//
//
//输入：mat =
//[[1,0,0,0],
// [1,1,1,1],
// [1,0,0,0],
// [1,0,0,0]],
//k = 2
//输出：[0,2]
//解释：
//每行中的军人数目：
//行 0 -> 1
//行 1 -> 4
//行 2 -> 1
//行 3 -> 1
//从最弱到最强对这些行排序后得到 [0,2,3,1]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 2 <= n, m <= 100
// 1 <= k <= m
// matrix[i][j] 不是 0 就是 1
//
// Related Topics 数组 二分查找 矩阵 排序 堆（优先队列） 👍 170 👎 0

func main() {
	value := 1
	no1337Print("%+v", value)
}

func no1337Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
type pair struct{ pow, idx int }

func kWeakestRows(mat [][]int, k int) []int {
	m := len(mat)
	pairs := make([]pair, m)
	for i, row := range mat {
		pow := sort.Search(len(row), func(j int) bool { return row[j] == 0 })
		pairs[i] = pair{pow, i}
	}
	rand.Seed(time.Now().UnixNano())
	randomizedSelected(pairs, 0, m-1, k)
	pairs = pairs[:k]
	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i], pairs[j]
		return a.pow < b.pow || a.pow == b.pow && a.idx < b.idx
	})
	ans := make([]int, k)
	for i, p := range pairs {
		ans[i] = p.idx
	}
	return ans
}

func randomizedSelected(a []pair, l, r, k int) {
	if l >= r {
		return
	}
	pos := randomPartition(a, l, r)
	num := pos - l + 1
	if k == num {
		return
	}
	if k < num {
		randomizedSelected(a, l, pos-1, k)
	} else {
		randomizedSelected(a, pos+1, r, k-num)
	}
}

func randomPartition(a []pair, l, r int) int {
	i := rand.Intn(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return kWeakestRowsPartition(a, l, r)
}

func kWeakestRowsPartition(a []pair, l, r int) int {
	pivot := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j].pow < pivot.pow || a[j].pow == pivot.pow && a[j].idx <= pivot.idx {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

//leetcode submit region end(Prohibit modification and deletion)
