package main

import "fmt"

var students []*Students

type Students struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Grade int32 `json:"grade"`
}

func init() {
	fmt.Println("start students.go")
	students = append(students, &Students{
		Id:    1,
		Name:  "anggi",
		Grade: 4,
	})

	students = append(students, &Students{
		Id:    2,
		Name:  "irawan",
		Grade: 5,
	})

	students = append(students, &Students{
		Id:    3,
		Name:  "awan",
		Grade: 6,
	})
}

func GetStudents() []*Students  {
	return students
}

func SelectStudents(id int64)  *Students {
	for _, v := range students {
		if v.Id == id {
			return v
		}
	}

	return nil
}
