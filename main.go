package main

import (
	"fmt"
	CourseRegistry "example.com/CheckingTask/courseregistry"
)

func main() {
	registry := CourseRegistry.NewRegistry()

	// Initial test data
	registry.AddStudent(CourseRegistry.Student{ID: 1, Name: "Alice", Courses: []string{"Go", "Databases"}})
	registry.AddStudent(CourseRegistry.Student{ID: 2, Name: "Bob", Courses: []string{"Go"}})
	registry.AddStudent(CourseRegistry.Student{ID: 3, Name: "Charlie", Courses: []string{}})

	for {
		fmt.Println("\nSTUDENT COURSE REGISTRY")
		fmt.Println("1. Add Student")
		fmt.Println("2. Enroll Course")
		fmt.Println("3. Remove Course")
		fmt.Println("4. List Students")
		fmt.Println("5. Course Statistics")
		fmt.Println("6. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id uint64
			var name string
			fmt.Print("ID: ")
			fmt.Scan(&id)
			fmt.Print("Name: ")
			fmt.Scan(&name)

			err := registry.AddStudent(CourseRegistry.Student{ID: id, Name: name})
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 2:
			var id uint64
			var course string
			fmt.Print("Student ID: ")
			fmt.Scan(&id)
			fmt.Print("Course: ")
			fmt.Scan(&course)

			err := registry.EnrollCourse(id, course)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 3:
			var id uint64
			var course string
			fmt.Print("Student ID: ")
			fmt.Scan(&id)
			fmt.Print("Course: ")
			fmt.Scan(&course)

			err := registry.RemoveCourse(id, course)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 4:
			students := registry.ListStudents()
			for _, s := range students {
				fmt.Printf("ID: %d | Name: %s | Courses: %v\n", s.ID, s.Name, s.Courses)
			}

		case 5:
			stats := registry.CoursesCount()
			for course, count := range stats {
				fmt.Printf("%s → %d\n", course, count)
			}

		case 6:
			return
		}
	}
}
