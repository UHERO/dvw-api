package data

var ModDimList = map[string][]string{
	"ade": {"markets", "destinations", "indicators"},
	"airseat": {"markets", "destinations", "indicators"},
	"hotel": {"categories", "indicators"},
	"char": {"groups", "indicators"},
	"exp": {"groups", "categories", "indicators"},
}

type Dimension struct {
	string
}
