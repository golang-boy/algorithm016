package Week_02

import "testing"

func TestGroupAnagrams(t *testing.T) {

	list := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	expected := [][]string{
		[]string{"ate","eat","tea"},
		[]string{"nat","tan"},
		[]string{"bat"},
	}

	got := groupAnagrams(list)
	if len(got) != len(expected) {
		t.Errorf("groupAnagrams beyond the expected error: %+v", got)
	}


}