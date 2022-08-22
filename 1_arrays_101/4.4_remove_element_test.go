package arrays101

import "testing"

func Test_removeElement2(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement2(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement2() = %v, want %v", got, tt.want)
			}
		})
	}
}
