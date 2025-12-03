package main

// Embeded structs, tags, anonymous structs, composition

type School struct {
	Classes []Class
}

func main() {
	teacher1 := NewTeacher("Sir Goodness", []string{"Maths", "Science", "Arts"}, []string{"Msc"})
	teacher2 := NewTeacher("Madam Mercy", []string{"English", "Reading"}, []string{"Bsc"})

	grade1 := NewClass("Grade One", []Student{}, teacher1)
	grade2 := NewClass("Grade Two", []Student{}, teacher2)

	s1 := NewStudent("Mike Goku", 10, &Class{Level: grade2.Level})
	s2 := NewStudent("Max Goku", 12, &Class{Level: grade2.Level})
	s3 := NewStudent("Austin Goku", 11, &Class{Level: grade1.Level})
	s4 := NewStudent("Tina Goku", 8, &Class{Level: grade1.Level})

	grade2.Students = append(grade2.Students, *s1)
	grade2.Students = append(grade2.Students, *s2)
	grade1.Students = append(grade1.Students, *s3)
	grade1.Students = append(grade1.Students, *s4)

	mySchool := &School{
		Classes: []Class{grade1, grade2},
	}

	mySchool.Classes[0].GetClassData()
	mySchool.Classes[1].GetClassData()
}
