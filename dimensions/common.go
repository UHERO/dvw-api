package dimensions

type Dimension interface{
	Handle() string
	NameW() string
	NameT() string
	Parent() Dimension
}
