package main

import "fmt"

func main() {
	// weather := map[int]int{ /*map[Ключ]Значение*/
	// 	11: +3,
	// 	12: +6,
	// 	13: +9,
	// 	14: -4,
	// 	15: +1,
	// 	30: 0,
	// }
	// // weather2 := make(map[int]int, 10)

	// c, ok := weather[30]
	// fmt.Println(weather[11])
	// fmt.Println(c, ok)
	// if ok {
	// 	fmt.Println("Okay")
	// } else {
	// 	fmt.Println("Bad")
	// }

	// weather[20] = -10
	// fmt.Println(weather[20])

	// for k, v := range weather {
	// 	weather[k] += 1
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(weather)

	// нам нужно хранить данные о людях
	// нужно по имени человека сразу понимать был ли он судим
	criminal := map[string]bool{
		"Zhyldyz": true,
		"Kanagat": false,
		"Talshyn": false,
		"Katya":   true,
	}
	c, ok := criminal["Talshyn"]
	if !ok {
		fmt.Println("Человек нет в базе!")
		return
	}
	fmt.Println("Человек найден в базе!")
	if c {
		fmt.Println("Человек судим!")
	} else {
		fmt.Println("Человек не судим!")
	}

	fmt.Println(criminal)
}
