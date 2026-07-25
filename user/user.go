package user

type User struct {
	Name    string
	Surname string
	Email   string
	Age     int
}

func NewUser(name string, surname string, email string, age int) User {
	if name == "" {
		return User{}
	}
	if surname == "" {
		return User{}
	}
	if email == "" {
		return User{}
	}
	if age <= 0 || age >= 150 {
		return User{}
	}
	return User{
		Name:    name,
		Surname: surname,
		Email:   email,
		Age:     age,
	}
}

func (u *User) SetNewName(name string) {
	if name != "" {
		u.Name = name
	}
}

func (u *User) SetNewSurname(surname string) {
	if surname != "" {
		u.Surname = surname
	}
}

func (u *User) SetNewEmail(email string) {
	if email != "" {
		u.Email = email
	}
}

func (u *User) SetNewAge(age int) {
	if age > 0 && age < 150 {
		u.Age = age
	}
}

func (u *User) GetUser() (string, int) {
	return u.Name, u.Age
}
