package courseregistry

import "testing"

func TestRegistryBasics(t *testing.T) {
    r := NewRegistry()
    // Add 5 students
    if err := r.AddStudent(Student{ID: 1, Name: "Alice", Courses: []string{"Go", "Databases"}}); err != nil {
        t.Fatalf("AddStudent failed: %v", err)
    }
    if err := r.AddStudent(Student{ID: 2, Name: "Bob", Courses: []string{"Go"}}); err != nil {
        t.Fatalf("AddStudent failed: %v", err)
    }
    if err := r.AddStudent(Student{ID: 3, Name: "Charlie", Courses: []string{}}); err != nil {
        t.Fatalf("AddStudent failed: %v", err)
    }
    if err := r.AddStudent(Student{ID: 4, Name: "Diana", Courses: []string{"Python", "Go"}}); err != nil {
        t.Fatalf("AddStudent failed: %v", err)
    }
    if err := r.AddStudent(Student{ID: 5, Name: "Eve", Courses: []string{"Databases", "Security"}}); err != nil {
        t.Fatalf("AddStudent failed: %v", err)
    }

    students := r.ListStudents()
    if len(students) != 5 {
        t.Fatalf("expected 5 students, got %d", len(students))
    }

    stats := r.CoursesCount()
    if stats["Go"] != 3 {
        t.Fatalf("expected Go -> 3, got %d", stats["Go"])
    }
    if stats["Databases"] != 2 {
        t.Fatalf("expected Databases -> 2, got %d", stats["Databases"])
    }
    if stats["Python"] != 1 {
        t.Fatalf("expected Python -> 1, got %d", stats["Python"])
    }
    if stats["Security"] != 1 {
        t.Fatalf("expected Security -> 1, got %d", stats["Security"])
    }
}
