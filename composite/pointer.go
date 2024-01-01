package composite

func Addition(left, right int) int {
	// var leftPointer *int = &left
	leftPointer := &left
	// var rightPointer *int = &right
	rightPointer := &right
	sum := *leftPointer + *rightPointer
	return sum
}
