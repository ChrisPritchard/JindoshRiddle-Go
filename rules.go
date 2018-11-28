package main

var rules = []interface{}{
	isTrue{madamNatsiou, purple},
	isTrue{countessContee, farLeft},
	isTrue{left, white},
	leftOf{red, blue},
	isTrue{red, beer},
	isTrue{dunwall, green},
	nextTo{birdPendant, dunwall},

	isTrue{doctorMarcolla, ring},
	notTrue{doctorMarcolla, dabokva},
	isTrue{dabokva, warMedal},
	nextTo{baleton, snuffTin},
	nextTo{baleton, whiskey},
	isTrue{ladyWinslow, rum},
	isTrue{fraeport, absinthe},
	isTrue{centre, wine},
	isTrue{baronessFinch, karnaca},
}
