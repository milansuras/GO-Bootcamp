package main

import "fmt"

type City_List struct {
	city    string
	state   string
	country string
}

func main() {
	var c1 City_List = City_List{"Vattiyoorkavu", "kerala", "India"}
	fmt.Println(c1.state)
	var c2 City_List = City_List{city: "Mumbai", state: "Maharashtra", country: "India"}
	fmt.Println(c2)
	fmt.Println(c2.country)
	var c3 City_List
	c4 := City_List{"Sydney", "New South Wales", "Australia"}
	fmt.Println(c4)

	fmt.Println(c3.city)
}
