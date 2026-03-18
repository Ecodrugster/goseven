package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Course struct {
	Title         string
	MaxScore      int
	StudentsCount int
}

type Student struct {
	User   User
	Score  int
	Course Course
}

type Teacher struct {
	User    User
	Subject string
}

type Participant interface {
	GetInfo() string
	Act() string
}

func NewStudent(Name string, Age int, Score int, Course Course) Student {
	return Student{
		User:   NewUser(Name, Age),
		Score:  Score,
		Course: Course,
	}
}

func NewTeacher(Name string, Age int, Subject string) Teacher {
	return Teacher{
		User:    NewUser(Name, Age),
		Subject: Subject,
	}
}

func NewUser(Name string, Age int) User {
	return User{
		Name: Name,
		Age:  Age,
	}
}

func NewCourse(Title string, MaxScore int, StudentsCount int) Course {
	return Course{
		Title:         Title,
		MaxScore:      MaxScore,
		StudentsCount: StudentsCount,
	}
}

// Методы для структуры Student
func (s Student) GetInfo() string {
	return fmt.Sprintf("Student: %s, Score: %d, Course: %s", s.User.Name, s.Score, s.Course.Title)
}
func (s *Student) Study(maxScore int) string {
	s.Score += 10
	if s.Score >= maxScore {
		s.Score = maxScore
	}
	fmt.Println(s.User.Name, "Учится и получает балл", s.Score)
	return fmt.Sprintf("Балл: %d", s.Score)
}

func (s Student) Act() string {
	return s.Study(s.Course.MaxScore)
}

func AverageScore(students []*Student) float64 {
	if len(students) == 0 {
		return 0
	}
	total := 0
	for _, s := range students {
		total += s.Score
	}
	return float64(total) / float64(len(students))
}

// Методы для структуры Тичер
func (t Teacher) Teach() string {
	return fmt.Sprintf("%s преподает %s", t.User.Name, t.Subject)
}

func (t Teacher) GetInfo() string {
	return fmt.Sprintf("Teacher: %s, Subject: %s", t.User.Name, t.Subject)
}

func (t Teacher) GradeStudent(s *Student, points int, maxScore int) {
	s.Score += points
	if s.Score > maxScore {
		s.Score = maxScore
	}
	fmt.Println(t.User.Name, "оценивает", s.User.Name, "| Новый балл:", s.Score)
}

func (t Teacher) Act() string {
	return t.Teach()
}

// Методы для структуры Course
func (c Course) GetInfo() string {
	return fmt.Sprintf("Course: %s, MaxScore: %d, StudentsCount: %d", c.Title, c.MaxScore, c.StudentsCount)
}

func (c Course) AddStudent() string {
	return fmt.Sprintf("Студент добавлен в %s", c.Title)
}

// Методы для структуры User
func (u User) GetName() string {
	return u.Name
}

func (u User) GetAge() int {
	return u.Age
}

func (u User) Introduce() {
	fmt.Println("Меня зовут", u.Name, "и мне", u.Age, "лет")
}

func main() {
	course := NewCourse("Golang", 100, 30)
	student1 := NewStudent("Vova", 18, 5, course)
	student2 := NewStudent("Aurik", 18, 7, course)
	teacher := NewTeacher("Madiyar", 24, "Golang")


	participants := []Participant{
		&student1,
		&student2,
		&teacher,
	}

	fmt.Println("-----------")
	for _, p := range participants {
		fmt.Println(p.GetInfo())
		fmt.Println(p.Act())
		fmt.Println()
	}

	teacher.GradeStudent(&student1, 15, course.MaxScore)
	teacher.GradeStudent(&student2, 20, course.MaxScore)

	fmt.Println()
	students := []*Student{&student1, &student2}
	avg := AverageScore(students)

	fmt.Println("Средний балл:", avg)
}
