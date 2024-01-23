package main

import (
	"context"
	"fmt"
	"sync"

	classroom "github.com/italobbarros/go-testing-apis/api"
)

type classroomServer struct {
	students map[string]*classroom.Student
	mu       sync.Mutex
	classroom.UnimplementedClassroomServiceServer
}

func (s *classroomServer) CreateStudent(ctx context.Context, req *classroom.CreateStudentRequest) (*classroom.CreateStudentResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := fmt.Sprintf("%d", len(s.students)+1)
	student := &classroom.Student{
		Id:     id,
		Name:   req.Name,
		Age:    req.Age,
		Gender: req.Gender,
	}

	s.students[id] = student

	return &classroom.CreateStudentResponse{Student: student}, nil
}

func (s *classroomServer) UpdateStudent(ctx context.Context, req *classroom.UpdateStudentRequest) (*classroom.UpdateStudentResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	student, exists := s.students[req.Id]
	if !exists {
		return nil, fmt.Errorf("student with ID %s not found", req.Id)
	}

	student.Name = req.Name
	student.Age = req.Age
	student.Gender = req.Gender

	return &classroom.UpdateStudentResponse{Student: student}, nil
}

func (s *classroomServer) DeleteStudent(ctx context.Context, req *classroom.DeleteStudentRequest) (*classroom.DeleteStudentResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.students[req.Id]
	if !exists {
		return nil, fmt.Errorf("student with ID %s not found", req.Id)
	}

	delete(s.students, req.Id)

	return &classroom.DeleteStudentResponse{Id: req.Id}, nil
}
