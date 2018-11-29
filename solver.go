package main

func filterPeople(set []description, filter func(description) bool) (result []description) {
	for _, v := range set {
		if filter(v) {
			result = append(result, v)
		}
	}
	return
}

func filterGroup(set [][]description, filter func([]description) bool) (result [][]description) {
	for _, v := range set {
		if filter(v) {
			result = append(result, v)
		}
	}
	return
}

func solve() []description {
	allPeople := allPossibilities()

	allowedPeople := filterPeople(allPeople, func(desc description) bool {
		for _, r := range rules {
			if !ruleDoesNotForbid(desc, r) {
				return false
			}
		}
		return true
	})

	allGroups := distinctGroups([]description{}, allowedPeople)

	allowedGroups := filterGroup(allGroups, func(group []description) bool {
		for _, r := range rules {
			if !ruleDoesNotForbid(group, r) {
				return false
			}
		}
		return true
	})

	return allowedGroups[0]
}
