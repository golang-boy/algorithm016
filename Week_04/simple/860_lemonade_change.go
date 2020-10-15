package simple

// https://leetcode-cn.com/problems/lemonade-change/

var fiveDollar int = 5
var tenDollar int = 10
var twentyDollar int = 20

func lemonadeChange(bills []int) bool {

	bankNotes := map[int]int{
		fiveDollar:   0,
		tenDollar:    0,
		twentyDollar: 0,
	}
	var IsChanged = false
	for _, bill := range bills {
		if fiveDollar == bill {
			bankNotes[fiveDollar]++
			IsChanged = true
			continue
		}
		if tenDollar == bill {
			if bankNotes[fiveDollar] != 0 {
				IsChanged = true
				bankNotes[fiveDollar]--
				bankNotes[tenDollar]++
			} else {
				IsChanged = false
				goto OUT
			}
		}

		if twentyDollar == bill {
			if bankNotes[tenDollar] != 0 && bankNotes[fiveDollar] != 0 {
				bankNotes[tenDollar]--
				bankNotes[fiveDollar]--
				IsChanged = true
			} else if bankNotes[tenDollar] == 0 && bankNotes[fiveDollar] >= 3 {
				bankNotes[fiveDollar] -= 3
				IsChanged = true
			} else {
				IsChanged = false
				goto OUT
			}
		}
	}

OUT:
	return IsChanged
}
