package module

// exported struct type can be accessed from a outside package.
type Employee struct {
	Id   int    //exported field
	name string //unexported field
	City string
}
