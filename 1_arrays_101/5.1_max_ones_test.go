package arrays101

import "testing"

func Test_findMaxConsecutiveOnes2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]int{1, 0, 1, 1, 0, 1}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes2(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes2() = %v, want %v", got, tt.want)
			}
		})
	}
}
