package Week_03

import "strings"

var butten = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result := []string{}
	for _, d := range digits {
		if len(result) == 0 {
			result = butten[string(d)]
			continue
		}
		result = combinee(result, butten[string(d)])
	}
	return result

}

func combinee(a, b []string) (c []string) {
	for _, aa := range a {
		for _, bb := range b {
			c = append(c, strings.Join([]string{aa, bb}, ""))
		}
	}
	return
}
