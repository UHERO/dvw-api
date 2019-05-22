package dimensions

type Dimension interface{
	Handle() string
	NameP() string
	NameW() string
	NameT() string
	Parent() Dimension
}

type ModuleDimensions struct {
	Dimensions	[]string  `json:"dimensions"`
}

