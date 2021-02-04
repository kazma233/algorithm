package array

import (
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{4, 1, 2, 1, 2}}, want: 1},
		{name: "test2", args: args{nums: []int{2, 2, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumberByMap(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{4, 1, 2, 1, 2}}, want: 1},
		{name: "test2", args: args{nums: []int{2, 2, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberByMap(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberByMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
