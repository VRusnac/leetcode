package arrays101

import "testing"

func TestFindNumbers(t *testing.T) {
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
			args: args{[]int{555, 901, 482, 1771}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.nums); got != tt.want {
				t.Errorf("FindNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
