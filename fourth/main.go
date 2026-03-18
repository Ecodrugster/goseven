package main

import "fmt"
func main() {

	// =============Повторение=============
// nums := []int{1,2,3}
// nums = append(nums,4)
// fmt.Println(nums)
// 			0  1  2  3  4
// x := []int{10,20,30,40,50}
// a := x[1:4]   [start, end]
// fmt.Println(a)
		//    0 1 2 3 4
// x := []int{1,2,3,4,5}
// 			   0 1 2
// a := x[1:4] //[2 3 4]
// a[0] = 999
// fmt.Println(x)

// a = [1,2,3,4]
// b = copy(a)

// src := []int{1,2,3}
// dst := make([]int,len(src))
// copy(dst,src)

// dst[0] = 999
// fmt.Println(src)
// fmt.Println(dst)

// =============Практика=============
// 1) Срез и общая память
// nums := []int{5,10,15,20,25}
// part := nums[1:4]
// part[0] = 123
// fmt.Println(nums)
// fmt.Println(part)

// 2) Копирование
// nums := []int{5,10,15}
// copyNums := make([]int, len(nums))
// copy(copyNums, nums)

// copyNums[0] = 777
// fmt.Println(nums)
// fmt.Println(copyNums)

//============= RANGE =============

// nums := []int{5,10,15,20,25}
// Стандартная запись
// for i := 0; i < len(nums); i++{
// 	fmt.Println(nums[i])
// }
// Короткая запись
// for i, value := range nums{
// 	fmt.Println(i,value)
// }
// for i := 0; i<len(nums); i++{
// 	fmt.Println(i, nums[i])
// }
// for i := range nums{
// 	fmt.Println(i)
// }

// for _, value := range nums{
// 	fmt.Println(value)
// }

// for i, value := range nums{
// 	value *= 2
// 	fmt.Println(i, value)
// }
// s := "Qwerty"
// for i, ch := range s{
// 	fmt.Println(i, string(ch))
// }

// Найти суммы элементов массива
// Найти макс элемент внутри массива 
// Умножить все элементы массива на 2
// Фильтрация (вывести все элементы больше > 5)
// Инднекс первого числа кратного 3 %

// nums := [] int{1,5,10,12,14,25,35}
// sum := 0
// for _, value := range nums {
// 	sum += value
// }

// fmt.Println("Сумма элементов массива:", sum)

// max := nums[0]
// for _,value := range nums {
// 	if value > max {
// 		max = value
// 	}
// }
// fmt.Println("Макс элемент массива:", max)

// for i := range nums {
// 	value *= 2
// 	fmt.Println(i, value)
// }

// for _, value := range nums {
// 	if value > 5 {
// 		fmt.Println(value)
// 	}
// }

// for i, value := range nums {
// 	if value % 3 == 0 {
// 		fmt.Println("Индекс первого числа кратного 3:", i)
// 		break
// 	}
// }

// Задачи
// Задача 1 — name → age (5 людей) + вывести всех
// Задача 2 — найти возраст по имени (ok-check)
// Задача 3 — частота чисел
// Задача 4 — красиво вывести category -> []items


// people := map[string]int{
// 	"Aman": 25,
// 	"Adilet": 30,
// 	"Beka": 20,
// 	"Masha": 28,
// 	"Alina": 22,
// }
// // for name, age := range people {
// 	fmt.Println(name, "->", age)
// }

// age, ok := people["Aman"]
// if ok {
// 	fmt.Println("Возраст Амана:", age)
// }

// nums := []int{1,2,2,3,4,5,6,7,7}
// qwerty := make(map[int]int)
// for _, v := range nums {
// 	qwerty[v]++
// }
// fmt.Println("Частота чисел:", qwerty)

products := map[string][] string{
	"Fruits": {"Orange", "Apple", "Banana"},
	"Technology": {"Phone", "Laptop", "Tablet"},
}
for category, items := range products {
	fmt.Println (category, "->", items)
}

}