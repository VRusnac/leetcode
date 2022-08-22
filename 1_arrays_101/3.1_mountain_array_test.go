package arrays101

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{[]int{0, 3, 2, 1}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
