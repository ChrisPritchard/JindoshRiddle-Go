package main

func applies(subject interface{}, description description) bool {
	switch subject.(type) {
	case women:
		return description.woman == subject
	case position:
		return description.position == subject
	case heirloom:
		return description.owns == subject
	case drink:
		return description.drinking == subject
	case colours:
		return description.wearing == subject
	case hometown:
		return description.from == subject
	default:
		return false
	}
}
