package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	classroom "github.com/italobbarros/go-testing-apis/pb/api"
	"google.golang.org/grpc"
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

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	server := grpc.NewServer()
	classroomServerVar := &classroomServer{
		students: make(map[string]*classroom.Student),
	}
	classroom.RegisterClassroomServiceServer(server, classroomServerVar)

	log.Printf("Server listening on port %s", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
