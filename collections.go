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
