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

func ruleDoesNotForbid(testable, rule interface{}) bool {
	switch testable.(type) {
	case description:
		{
			desc := testable.(description)
			switch rule.(type) {
			case isTrue:
				{
					if applies(rule.(isTrue).subject, desc) {
						return applies(rule.(isTrue).targetOrFact, desc)
					} else {
						return !applies(rule.(isTrue).targetOrFact, desc)
					}
				}
			case notTrue:
				{
					return !applies(rule.(notTrue).subject, desc) || !applies(rule.(notTrue).targetOrFact, desc)
				}
			default:
				return true
			}
		}
	case []description:
		return true
	default:
		return true
	}
}
