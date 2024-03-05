package main

import "fmt"

type User struct {
	id          int
	name        string
	age         int
	visitDays   int
	total_likes int
	followers   int
}

var users []User = []User{

	User{101, "A", 18, 20, 109, 515},
	User{102, "B", 25, 12, 567, 777},
	User{103, "C", 30, 12, 891, 1111},
}

type By func(user User) bool

func (by By) Filter(s1 []User) []User {
	filtered := make([]User, 0)
	for _, ele := range s1 {
		if by(ele) {
			filtered = append(filtered, ele)
		}
	}
	return filtered
}

func main() {
	//frequent user

	frequent := func(user User) bool {
		return user.visitDays > 100
	}

	//liked user
	liked := func(user User) bool {
		return user.total_likes > 500
	}

	// large followers
	high := func(user User) bool {
		return user.followers > 1000
	}

	//age

	matured := func(user User) bool {
		return user.age >= 30
	}

	frequentUsers := By(frequent).Filter(users)
	fmt.Println(frequentUsers)

	liked_users := By(liked).Filter(users)
	fmt.Println(liked_users)

	highend_users := By(high).Filter(users)
	fmt.Println(highend_users)

	maturedusers := By(matured).Filter(users)
	fmt.Println(maturedusers)

}
