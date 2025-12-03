package main

type Student struct {
	Name  string
	Level *Class
	Age   int
}

func NewStudent(name string, age int, level *Class) *Student {
	student := &Student{
		Name:  name,
		Age:   age,
		Level: level,
	}
	return student
}

