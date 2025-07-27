package main

import "fmt"

func userInput() int {
	userInp := 0
	fmt.Println(`- 1. Посмотреть закладки
- 2. Добавить закладку
- 3. Удалить закладку
- 4. Выход`)
	fmt.Println("-----------------")
	fmt.Print("Ввод:")
	fmt.Scan(&userInp)
	fmt.Println("-----------------")
	return userInp
}

func logica(myMap map[string]string) {
	flag := 0
	for flag != 4 {
		userInp := userInput()
		switch {
		case userInp == 1:
			if len(myMap) == 0 {
				fmt.Println("Пока нет закладок")
			}
			flag = 1
			for key, value := range myMap {
				fmt.Println(key, ":", value)
			}
		case userInp == 2:
			key := ""
			value := ""
			flag = 2
			fmt.Print("Введите ключ:")
			fmt.Scan(&key)
			fmt.Print("Введите значение:")
			fmt.Scan(&value)
			myMap[key] = value
			fmt.Printf("Значение %s добавленно!\n", value)
		case userInp == 3:
			key := ""
			flag = 3
			for key, value := range myMap {
				fmt.Println(key, ":", value)
			}
			fmt.Print("Введите ключ для удаления:")
			fmt.Scan(&key)
			delete(myMap, key)
			fmt.Printf("Значение %s удаленно!\n", key)
		case userInp == 4:
			flag = 4
		}
	}
}

func main() {
	myMap := map[string]string{}
	logica(myMap)
}
