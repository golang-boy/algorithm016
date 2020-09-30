package Week_03

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"abc",
			args{n: 4, k: 2},
			[][]int{
				[]int{1, 2},
				[]int{1, 3},
				[]int{1, 4},
				[]int{2, 3},
				[]int{2, 4},
				[]int{3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
