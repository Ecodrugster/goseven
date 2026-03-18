package main

import "fmt"


type Engine struct {
	Power int
	Type string
}

func (e Engine) Start() {
	fmt.Println("Двигатель Запущен:", e.Type, e.Power, "л.с")
}

type Car struct{
	Brand string
	Model string
	Engine Engine
}

// type Teacher struct {
// 	Name    string
// 	Age     int
// 	Subject string
// }

// type Car struct {
// 	Brand string
// 	Year  int
// }

// func (c Car) Drive() {
// 	fmt.Println("Brand:", c.Brand)
// 	fmt.Println("Year:", c.Year)
// }

// func (t Teacher) ShowInfo() {
// 	fmt.Println("Name:", t.Name)
// 	fmt.Println("Age:", t.Age)
// 	fmt.Println("Subject:", t.Subject)
// }

// type Printer interface {
// 	Print()
// }

// func (t Teacher) Print() {
// 	fmt.Println("Teacher:", t.Name, "| Subject:", t.Subject)
// }

// type Student struct{
// 	Name string
// 	Age int
// 	Group string
// 	GPA float64
// }

// type BankAccount struct{
// 	owner string
// 	balance float64
// }

// type Speaker interface{
// 	Speak()
// }
// type Dog struct{
// 	Name string
// }
// type Cat struct{
// 	Name string
// }

// func (c Cat) Speak() {
// 	fmt.Println(c.Name, "Мяукает")
// }

// func (d Dog) Speak() {
// 	fmt.Println(d.Name, "Гавкает")
// }

// func MakeSpeak(s Speaker) {
// 	s.Speak()
// }

// func NewBankAccount(owner string, balance float64) BankAccount {
// 	return BankAccount{
// 		owner: owner,
// 		balance: balance,
// 	}
// }

// func (b BankAccount) GetBalance() float64 {
// 	return b.balance
// }

// func (b *BankAccount) Deposit(amount float64) {
// 	if amount > 0 {
// 		b.balance += amount
// 	}
// }
// func (s Student) Info() {
// 	fmt.Println("Name:", s.Name)
// 	fmt.Println("Age:", s.Age)
// 	fmt.Println("Group:", s.Group)
// 	fmt.Println("GPA:", s.GPA)
// }

func main() {
	// t1 := Teacher{
	// 	Name:    "Alex",
	// 	Age:     25,
	// 	Subject: "Math",
	// }
	// t1.ShowInfo()
	// fmt.Println()

	// var p Printer = t1
	// p.Print()

	// c1 := Car{
	// 	Brand: "Honda",
	// 	Year:  2005,
	// }
	// c1.Drive()

	// s1 := Student{
	// 	Name: "Amina",
	// 	Age: 20,
	// 	Group: "CS101",
	// 	GPA: 3.8,
	// }
	// s1.Info()

	// acc := NewBankAccount("Alex", 10000)
	// acc.Deposit(5000)

	// fmt.Println("Balance:", acc.GetBalance())

	// d := Dog{Name: "Rex"}
	// c := Cat{Name: "Marusya"}

	// MakeSpeak(d)
	// MakeSpeak(c)
	

	car := Car{
		Brand: "Toyota",
		Model: " Camry",
		Engine: Engine{
			Power: 181,
			Type: "Бензин",
		},
	}
	fmt.Println("Марка: ", car.Brand)
	fmt.Println("Модель: ", car.Model)

	car.Engine.Start()
}
