package Week_02

import "sort"

type pair struct {
	element int
	frequent int
}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	// 统计元素频率
	for _, n :=range nums {
		m[n]+=1
	}

	pairList := []pair{}

	// 按次数排序
	for ele, cnt := range m {
		pairList = append(pairList, pair{
			element:ele,
			frequent:cnt,
		})
	}

	// golang中排序的时间复杂度是什么？
	sort.Sort(pairs(pairList))

	// 时间复杂度怎么计算？
	// 如果用小顶堆，时间复杂度又该怎么计算？

	ans := []int{}

	// 取前k个元素
	for _, n := range pairList {
		if len(ans) >= k {
			break
		}
		ans = append(ans, n.element)
	}
	return ans
}

type pairs []pair
func (p pairs) Len() int {return len(p)}
func (p pairs) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func (p pairs) Less(i, j int) bool {return p[i].frequent > p[j].frequent } // 从大到小排序