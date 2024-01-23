package main

import (
	"context"
	"fmt"
	"log"

	classroom "github.com/italobbarros/go-testing-apis/api"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := classroom.NewClassroomServiceClient(conn)

	// Test CreateStudent
	createResp, err := client.CreateStudent(context.Background(), &classroom.CreateStudentRequest{
		Name:   "John Doe",
		Age:    25,
		Gender: "Male",
	})
	if err != nil {
		log.Fatalf("Error creating student: %v", err)
	}
	fmt.Printf("Created Student: %+v\n", createResp.Student)

	// Test UpdateStudent
	updateResp, err := client.UpdateStudent(context.Background(), &classroom.UpdateStudentRequest{
		Id:     createResp.Student.Id,
		Name:   "Updated Name",
		Age:    30,
		Gender: "Female",
	})
	if err != nil {
		log.Fatalf("Error updating student: %v", err)
	}
	fmt.Printf("Updated Student: %+v\n", updateResp.Student)

	// Test DeleteStudent
	deleteResp, err := client.DeleteStudent(context.Background(), &classroom.DeleteStudentRequest{
		Id: createResp.Student.Id,
	})
	if err != nil {
		log.Fatalf("Error deleting student: %v", err)
	}
	fmt.Printf("Deleted Student with ID: %s\n", deleteResp.Id)
}
