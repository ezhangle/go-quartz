package csm

type CommonNode struct {
	value  int
	min    int
	max    int
	values []int
}

var _ csmNode = (*CommonNode)(nil)

func NewCommonNode(value, lowerBound, upperBound int, values []int) *CommonNode {
	return &CommonNode{value, lowerBound, upperBound, values}
}

func (n *CommonNode) Value() int {
	return n.value
}

func (n *CommonNode) Reset() {
	n.value = n.max
	n.Next()
}

func (n *CommonNode) Next() (overflowed bool) {
	if n.hasRange() {
		return n.nextInRange()
	}

	return n.next()
}

func (n *CommonNode) findForward() result {
	if !n.isValid() {
		if n.Next() {
			return overflowed
		}
		return advanced
	}
	return unchanged
}

func (n *CommonNode) hasRange() bool {
	return len(n.values) != 0
}

func (n *CommonNode) next() bool {
	n.value++
	if n.value > n.max {
		n.value = n.min
		return true
	}
	return false
}

func (n *CommonNode) nextInRange() bool {
	// find the next value in the range (assuming n.values is sorted)
	for _, value := range n.values {
		if value > n.value {
			n.value = value
			return false
		}
	}

	// the end of the values array is reached; set to the first valid value
	n.value = n.values[0]
	return true
}

func (n *CommonNode) isValid() bool {
	withinLimits := n.value >= n.min && n.value <= n.max
	if n.hasRange() {
		withinLimits = withinLimits && contains(n.values, n.value)
	}
	return withinLimits
}
