package Week_02

import "testing"

func TestIsAnagram(t *testing.T) {
	list1 := map[string]string {
		"anagram": "nagaram",
		"中国":   "国中",
		"中国人": "人中国",
		"china": "inach",
	}

	for s, m := range list1 {
	    got := isAnagram(s, m)
	    if !got {
	    	t.Errorf("isAnagram(%s, %s) = %t", s, m, got)
	    }
	} 

	list2 := map[string]string {
		"anagramn": "nagaramm",
		"中国":   "国国",
		"中国人": "人人中国",
		"chiana": "inachn",
		"aacc": "cccc",
	}

	for s, m := range list2 {
	    got := isAnagram(s, m)
	    if got {
	    	t.Errorf("isNotAnagram(%s, %s) = %t", s, m, got)
	    }
	} 
}