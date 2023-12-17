package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gofr.dev/pkg/gofr"
	"go-rest-api/datastore"
	"go-rest-api/model"
	"testing"
)

func TestStore_Create(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO student").WithArgs(sqlmock.AnyArg(), "John Doe", 12, "Maths").WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT id,name,age,class FROM student").WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "class"}).AddRow(1, "John Doe", 12, "SST"))

	ctx := &gofr.Context{}
	s := datastore.New()

	testStudent := model.Student{
		Name:  "John Doe",
		Age:   13,
		Class: "Science",
	}

	createdStudent, err := s.Create(ctx, &testStudent)
	if err != nil {
		t.Errorf("Error creating student: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

	if createdStudent.ID == 0 {
		t.Errorf("Expected a valid ID for the created student, got 0")
	}

	if createdStudent.Name != "John Doe" {
		t.Errorf("Expected CustomerName to be 'John Doe', got '%s'", createdStudent.Name)
	}
}

func TestStore_Update(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()


	mock.ExpectExec("UPDATE students").WithArgs(1, "John Doe", 13, "Maths").WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := &gofr.Context{}
	s := datastore.New()

	
	testStudent := model.Student{
		Name:  "John Doe",
		Age:   13,
		Class: "Science",
	}

	updatedStudent, err := s.Update(ctx, &testStudent)
	if err != nil {
		t.Errorf("Error updating student: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

	if updatedStudent.ID != testStudent.ID {
		t.Errorf("Expected updated student ID to be %d, got %d", testStudent.ID, updatedStudent.ID)
	}
	
}
func TestStore_Delete(t *testing.T) {
	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM students").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := &gofr.Context{}
	s := datastore.New()
        
	teststudentID := 1

	err = s.Delete(ctx, teststudentID)
	if err != nil {
		t.Errorf("Error deleting student: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
