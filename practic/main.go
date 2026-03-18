package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Group string
	Grade float64
}

type Teacher struct {
	Name       string
	Age        int
	Subject    string
	Experience int
}

type Speaker interface {
	Speak() string
}

func PrintSpeech(s Speaker) {
	fmt.Println(s.Speak())
}

func (s Student) Info() {
	fmt.Println("Student: ", s.Name)
	fmt.Println("Age: ", s.Age)
	fmt.Println("Group: ", s.Group)
	fmt.Println("Grade: ", s.Grade)
}

func (s Student) Study(hours int) {
	fmt.Println(s.Name, " studies ", hours, "hours")
}

func (s *Student) UpdateGrade(newGrade float64) {
	s.Grade = newGrade
}

func (s Student) Speak() string {
	return "Hello, I am a student" + s.Name
}

func (t Teacher) Info() {
	fmt.Println("Teacher: ", t.Name)
	fmt.Println("Age: ", t.Age)
	fmt.Println("Subject: ", t.Subject)
	fmt.Println("Experience: ", t.Experience)
}

func (t Teacher) Teach(topic string) {
	fmt.Println(t.Name, "is teaching ", topic)
}

func (t *Teacher) UpdateExperience(years int) {
	t.Experience += years
}

func (t Teacher) Speak() string {
	return "Hello, I am a teacher" + t.Name
}

func main() {
	s1 := Student{
		Name:  "Aman",
		Age:   21,
		Group: "A-123",
		Grade: 3.6,
	}

	s2 := Student{
		Name:  "Alibek",
		Age:   20,
		Group: "A-122",
		Grade: 4.0,
	}

	fmt.Println(s1)
	fmt.Println(s2)

	teacher := Teacher{
		Name:       "Alina",
		Age:        32,
		Subject:    "Math",
		Experience: 10,
	}

	fmt.Println(teacher)

	fmt.Println("---Methods---")

	s1.Info()
	s2.Info()

	teacher.Info()

	s1.Study(2)

	s1.UpdateGrade(4.5)
	fmt.Println("Update grade: ", s1.Grade)

	teacher.Info()
	teacher.Teach("Interfaces")

	teacher.UpdateExperience(2)
	fmt.Println("New Experience:", teacher.Experience)

	fmt.Println("\n--- Interface ---")

	PrintSpeech(s1)
	PrintSpeech(teacher)
}
