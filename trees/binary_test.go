package trees

import (
	"testing"
)

var root *Node
var n1 = &Node{Val: 1}
var n2 = &Node{Val: 2}
var n3 = &Node{Val: 3}
var n4 = &Node{Val: 9}
var n5 = &Node{Val: 5}
var n6 = &Node{Val: 6}
var n7 = &Node{Val: 7}
var n8 = &Node{Val: 8}
var n9 = &Node{Val: 4}
var n10 = &Node{Val: 10}

func TestMain(m *testing.M) {
	/**
		  	       0
		        1      2
		    3      9
		  5   6  7   8
			4
		 10
	**/
	root = NewNode(0, n1, n2)
	n1.Left = n3
	n1.Right = n4
	n3.Left = n5
	n3.Right = n6
	n4.Left = n7
	n4.Right = n8
	n6.Left = n9
	n9.Left = n10

	m.Run()
}

func Test_findDepLength(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "寻找最大深度", args: args{root}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDepLength(tt.args.node); got != tt.want {
				t.Errorf("findDepLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDepVal(t *testing.T) {
	maxDepVal := Node{}
	depQueue := []*Node{n4}
	findDepVal(depQueue, &maxDepVal)
	t.Logf("%v 寻找最左边,最深的值为: %v\n", depQueue[0].Val, maxDepVal.Val)
}

func Test_findMaxVal(t *testing.T) {
	type args struct {
		queue  []*Node
		oldVal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "寻找最大值", args: args{[]*Node{root}, 0}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxVal(tt.args.queue, tt.args.oldVal); got != tt.want {
				t.Errorf("findMaxVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
