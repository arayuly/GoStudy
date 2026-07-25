package main

import (
	"fmt"
	"study/greeting"
)

type User struct {
	Name    string
	Age     int
	Phone   string
	IsClose bool
	Ratind  float64
}

func NewUser(name string, age int, phone string, isClose bool, rating float64) User {
	if name == "" {
		return User{}
	}
	if age <= 0 || age >= 150 {
		return User{}
	}
	if phone == "" {
		return User{}
	}
	if rating < 0.0 || rating > 10 {
		return User{}
	}
	return User{
		Name:    name,
		Age:     age,
		Phone:   phone,
		IsClose: isClose,
		Ratind:  rating,
	}
}

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) ChangePhone(newPhone string) {
	if newPhone != "" {
		u.Phone = newPhone
	}
}

func (u *User) CloseAccount() {
	u.IsClose = true
}

func (u *User) OpenAccount() {
	u.IsClose = false
}

func Greeting(u *User) {
	fmt.Println("Всем привет!")
	fmt.Println("Меня зовут,", u.Name)
	fmt.Println("Мой рейтинг:", u.Ratind)
	fmt.Println("")
	u.Name = "Saske"
}

func (u *User) RatingUp(rating float64) {
	if u.Ratind+rating <= 10.0 {
		u.Ratind += rating
		fmt.Println("Вы добавили рейтинг к пользователю:", u.Name)
	} else {
		fmt.Println("Вы не прошли валидацию")
	}
}

func (u *User) DownRating(rating float64) {
	if u.Ratind-rating >= 0 {
		u.Ratind -= rating
	}
}

func main() {
	user := NewUser(
		"Itachi",
		1300,
		"45565465575756",
		true,
		9.7,
	)
	fmt.Println(user)
	fmt.Println("")

	Greeting(&user)
	fmt.Println(user.Name)

	user.RatingUp(0.1)
	fmt.Println(user.Ratind)

	greeting.SayHello()
}
