package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	// 1
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите текст: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Вы ввели:", text)

	file, err := os.OpenFile("diary.txt", os.O_RDWR|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		panic(err)
	}
	fmt.Println("Текст успешно записан в файл diary.txt")

	// 2
	data, err := os.ReadFile("diary.txt")
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Содержимое файла diary.txt:", string(data))
}