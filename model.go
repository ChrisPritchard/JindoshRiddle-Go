package main

type position int

const (
	farLeft  position = 0
	left     position = 1
	centre   position = 2
	right    position = 3
	farRight position = 4
)

type women int

const (
	ladyWinslow    women = 0
	doctorMarcolla women = 1
	countessContee women = 2
	madamNatsiou   women = 3
	baronessFinch  women = 4
)

type colours int

const (
	purple colours = 0
	white  colours = 1
	red    colours = 2
	blue   colours = 3
	green  colours = 4
)

type heirloom int

const (
	ring        heirloom = 0
	birdPendant heirloom = 1
	diamond     heirloom = 2
	warMedal    heirloom = 3
	snuffTin    heirloom = 4
)

type drink int

const (
	beer     drink = 0
	whiskey  drink = 1
	rum      drink = 2
	absinthe drink = 3
	wine     drink = 4
)

type hometown int

const (
	dunwall  hometown = 0
	dabokva  hometown = 1
	baleton  hometown = 2
	fraeport hometown = 3
	karnaca  hometown = 4
)
