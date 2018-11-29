package main

func allPossibilities() (results []description) {
	results = make([]description, 5*5*5*5*5*5)

	for p := 0; p < 5; p++ {
		for w := 0; w < 5; w++ {
			for c := 0; c < 5; c++ {
				for f := 0; f < 5; f++ {
					for d := 0; d < 5; d++ {
						for o := 0; o < 5; o++ {
							results[p*3125+w*625+c*125+f*25+d*5+o] = description{
								position: position(p),
								woman:    women(w),
								wearing:  colours(c),
								from:     hometown(f),
								drinking: drink(d),
								owns:     heirloom(o),
							}
						}
					}
				}
			}
		}
	}

	return
}

func noConflict(a, b description) bool {
	return a.position != b.position && a.woman != b.woman && a.wearing != b.wearing && a.from != b.from && a.drinking != b.drinking && a.owns != b.owns
}

func distinctGroups(group, allPeople []description) [][]description {
	if len(group) == 5 {
		return [][]description{group}
	}
	results := [][]description{}
	for _, p := range allPeople {
		valid := true
		for _, op := range group {
			if !noConflict(p, op) {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}

		// TODO exclude person from allPeople?
		for _, ng := range distinctGroups(append(group, p), allPeople) {
			results = append(results, ng)
		}
	}

	// TODO distinct?
	return results
}
