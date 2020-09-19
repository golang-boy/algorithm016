package Week_02

import "sort"

func groupAnagrams(strs []string) [][]string {

	listMap := make(map[string][]string)

	ans := [][]string{}

	for _, str := range strs {
		// 先排序
		r := runes(str)
		sort.Sort(r)
        newStr := string(r)
		// 放入map中进行比较
		if v, ok := listMap[newStr]; !ok {
			listMap[newStr] = []string{str}
		} else {
			listMap[newStr] = append(v, str)
		}
	}

	for _, v := range listMap {
		ans = append(ans, v)
	}

	return ans
}

type runes []rune

func (s runes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s runes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s runes) Len() int{
	return len(s)
}
