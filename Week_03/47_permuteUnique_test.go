package Week_03

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"abc",
			args{[]int{1, 1, 2}},
			[][]int{
				[]int{1, 1, 2},
				[]int{1, 2, 1},
				[]int{2, 1, 1},
			},
		},
		{
			"abcd",
			args{[]int{3, 3, 0, 3}},
			[][]int{
				[]int{0, 3, 3, 3},
				[]int{3, 0, 3, 3},
				[]int{3, 3, 0, 3},
				[]int{3, 3, 3, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
