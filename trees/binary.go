package trees

// 二叉树

// Node 树节点
type Node struct {
	Val   int   // 节点的值
	Left  *Node // 左树
	Right *Node // 右树
}

// NewNode 新建一个节点
func NewNode(val int, left, right *Node) *Node {
	return &Node{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func findDepLength(node *Node) int {
	if node == nil {
		return 0
	}

	left := findDepLength(node.Left)
	right := findDepLength(node.Right)

	if left >= right {
		return left + 1
	}
	return right + 1
}

func findDepVal(queue []*Node, maxLeft *Node) {
	if len(queue) <= 0 {
		return
	}

	val := queue[0]
	queue = queue[1:]
	if val.Left != nil {
		queue = append(queue, val.Left)
		*maxLeft = *val.Left
	}
	if val.Right != nil {
		queue = append(queue, val.Right)
	}

	findDepVal(queue, maxLeft)
}

func findMaxVal(queue []*Node, oldVal int) int {
	if len(queue) <= 0 {
		return oldVal
	}

	val := queue[0]
	queue = queue[1:]
	if val.Left != nil {
		queue = append(queue, val.Left)
	}
	if val.Right != nil {
		queue = append(queue, val.Right)
	}

	if oldVal < val.Val {
		oldVal = val.Val
	}
	return findMaxVal(queue, oldVal)
}
