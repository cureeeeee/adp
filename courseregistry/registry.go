package courseregistry

import (
	"errors"
	"sort"
)

type Student struct {
	ID      uint64
	Name    string
	Courses []string
}

type Registry struct {
	Students map[uint64]Student
}

func NewRegistry() *Registry {
	return &Registry{Students: make(map[uint64]Student)}
}

// AddStudent adds a new student with validation
func (r *Registry) AddStudent(student Student) error {
	if student.Name == "" {
		return errors.New("student name cannot be empty")
	}
	if _, exists := r.Students[student.ID]; exists {
		return errors.New("student ID already exists")
	}
	r.Students[student.ID] = student
	return nil
}

// EnrollCourse enrolls a student in a course
func (r *Registry) EnrollCourse(studentID uint64, course string) error {
	if course == "" {
		return errors.New("course name cannot be empty")
	}

	student, exists := r.Students[studentID]
	if !exists {
		return errors.New("student not found")
	}

	for _, c := range student.Courses {
		if c == course {
			return errors.New("course already enrolled")
		}
	}

	student.Courses = append(student.Courses, course)
	r.Students[studentID] = student
	return nil
}

// RemoveCourse removes a course from a student
func (r *Registry) RemoveCourse(studentID uint64, course string) error {
	student, exists := r.Students[studentID]
	if !exists {
		return errors.New("student not found")
	}

	index := -1
	for i, c := range student.Courses {
		if c == course {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("course not found for student")
	}

	student.Courses = append(student.Courses[:index], student.Courses[index+1:]...)
	r.Students[studentID] = student
	return nil
}

// ListStudents returns all students as a slice
func (r *Registry) ListStudents() []Student {
	result := make([]Student, 0, len(r.Students))
	for _, student := range r.Students {
		result = append(result, student)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result
}

// CoursesCount returns course statistics
func (r *Registry) CoursesCount() map[string]int {
	stats := make(map[string]int)
	for _, student := range r.Students {
		for _, course := range student.Courses {
			stats[course]++
		}
	}
	return stats
}
