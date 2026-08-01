package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Введите команду: добавить лук
// вы хотите добавить лук
// Введите команду: удалить морковь
// вы кажется хотите удалить морковь

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите команду: ")

		// scanner.Scan()
		// if !ok {
		// 	fmt.Println("Ошибка ввода!")
		// 	return
		// }

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода!")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text)
		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели!")
			return
		}

		fmt.Println("text:", text)

		fmt.Println("Слова:", fields)
		fmt.Println("Команда:", fields[0])

		cmd := fields[0]

		if cmd == "выйти" {
			fmt.Println("До скорого!")
			return
		}
		if cmd == "добавить" {
			fmt.Println("Вы хотите добавить:", fields[1:])
		} else if cmd == "удалить" {
			fmt.Println("Вы кажется хотите удалить:", fields[1:])
		} else if cmd == "help" {
			fmt.Println("Команда: добавить {что нужно добавить}")
			fmt.Println("-- эта команда позволяет добавлять что-то")
			fmt.Println("")
			fmt.Println("Команда: добавудалитьить {что нужно удалить}")
			fmt.Println("-- эта команда позволяет удалять что-то")
			fmt.Println("")
		} else {
			fmt.Println("Вы ввели неизвестную команду")
		}
	}
}
