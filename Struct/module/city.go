package module

// unexported struct accessible inside package module only
type city struct {
	city_name string //unexported field accessible inside package module
	state     string
}
