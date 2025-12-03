package main

import "fmt"

type Class struct {
	Level    string
	Teacher  *Teacher
	Students []Student
}

func NewClass(level string, students []Student, teacher *Teacher) Class {
	class := Class{
		Level:    level,
		Teacher:  teacher,
		Students: students,
	}
	return class
}

func (c *Class) GetStudents() []string {
	var allStudents []string
	for _, s := range c.Students {
		allStudents = append(allStudents, s.Name)
	}
	return allStudents
}

func (c *Class) GetClassData() {
	fmt.Println("----------Class Data----------")
	fmt.Printf("Level: %v\n", c.Level)
	fmt.Printf("Teacher: %v\n", c.Teacher.Name)
	fmt.Printf("Cert: %v\n", c.Teacher.GetCerts())
	fmt.Printf("Subjects: %v\n", c.Teacher.GetSubjects())
	fmt.Printf("Students: %v\n", c.GetStudents())
	fmt.Println("------------------------------")
}
