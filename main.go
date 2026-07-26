package main

import (
	"fmt"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	user1 := User{
		Name:    "Almaz",
		Rating:  5.5,
		Premium: true,
	}
	user2 := User{
		Name:    "Shyngys",
		Rating:  5.5,
		Premium: false,
	}
	user3 := User{
		Name:    "Muhamedzhan",
		Rating:  5.5,
		Premium: true,
	}
	user4 := User{
		Name:    "Dias",
		Rating:  5.5,
		Premium: false,
	}

	users := [4]User{user1, user2, user3, user4}

	for index, user := range users {
		fmt.Println(index, user)

		if user.Premium {
			users[index].Rating += 1.0
		}

		fmt.Println(users)
		fmt.Println("")
	}

	for _, user := range users {
		fmt.Println(user)
	}
	// for i := 0; i < len(users); i++ {

	// 	fmt.Println(users[i])

	// 	if users[i].Premium {
	// 		users[i].Rating += 1.0
	// 	}

	// 	fmt.Println(users[i])
	// 	fmt.Println("")
	// }
}
