package main

import (
	"sort"
)

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
			person := testable.(description)
			switch rule.(type) {
			case isTrue:
				{
					r := rule.(isTrue)
					if applies(r.subject, person) {
						return applies(r.target, person)
					} else {
						return !applies(r.target, person)
					}
				}
			case notTrue:
				{
					r := rule.(notTrue)
					return !applies(r.subject, person) || !applies(r.target, person)
				}
			case leftOf:
				{
					r := rule.(leftOf)
					return ruleDoesNotForbid(testable, notTrue{r.subject, r.target})
				}
			case rightOf:
				{
					r := rule.(rightOf)
					return ruleDoesNotForbid(testable, notTrue{r.subject, r.target})
				}
			case nextTo:
				{
					r := rule.(nextTo)
					return ruleDoesNotForbid(testable, notTrue{r.subject, r.target})
				}
			default:
				return true
			}
		}
	case []description:
		{
			switch rule.(type) {
			case leftOf:
				{
					group := testable.([]description)
					sort.Slice(group, func(a, b int) bool {
						return group[a].position < group[b].position
					})

					r := rule.(leftOf)
					for i := 0; i < len(group)-1; i++ {
						if applies(r.subject, group[i]) && applies(r.target, group[i+1]) {
							return true
						}
					}
					return false
				}
			case rightOf:
				{
					r := rule.(rightOf)
					return ruleDoesNotForbid(testable, leftOf{r.target, r.subject})
				}
			case nextTo:
				{
					r := rule.(nextTo)
					return ruleDoesNotForbid(testable, leftOf{r.subject, r.target}) || ruleDoesNotForbid(testable, leftOf{r.target, r.subject})
				}
			default:
				return true
			}
		}
	default:
		return true
	}
}
