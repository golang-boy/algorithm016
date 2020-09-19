package Week_02

func isAnagram(s, t string) bool {

	if len([]rune(s)) != len([]rune(t)) {
		return false
	}

	o := map[rune]int64{}
	for _, ss := range []rune(s) {
		o[ss] += 1
	}

	for _, ss := range []rune(t) {
		if _, ok := o[ss]; !ok {
			return false
		} else {
			o[ss] -= 1
			if o[ss] < 0 {
				return false
			}
		}
	}
	return true
}