package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gofr.dev/pkg/gofr"
	"go-rest-api/handler"
	"go-rest-api/model"
)

type MockStudentStore struct {
	mock.Mock
}

func (m *MockStudentStore) GetByID(ctx *gofr.Context, id string) (*model.Student, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Student), args.Error(1)
}

func (m *MockStudentStore) Create(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
	args := m.Called(ctx, student)
	return args.Get(0).(*model.Student), args.Error(1)
}

func (m *MockStudentStore) Update(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
	args := m.Called(ctx, student)
	return args.Get(0).(*model.Student), args.Error(1)
}

func (m *MockStudentStore) Delete(ctx *gofr.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockStudentStore) GetAll(ctx *gofr.Context) ([]model.Student, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Student), args.Error(1)
}

func TestGetAllHandler(t *testing.T) {
	mockStore := &MockStudentStore{}

	h := handler.New(mockStore)

	mockStore.On("GetAll", mock.Anything).Return([]model.Student{
		{ID: 1, Name: "John Doe", Age: 20, Class: "Math"},
		{ID: 2, Name: "Jane Doe", Age: 22, Class: "Science"},
	}, nil)

	mockContext := new(gofr.Context)

	result, err := h.GetAll(mockContext)

	expected := []model.Student{
		{ID: 1, Name: "John Doe", Age: 20, Class: "Math"},
		{ID: 2, Name: "Jane Doe", Age: 22, Class: "Science"},
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockStore.AssertExpectations(t)
}

func TestGetByIDHandler(t *testing.T) {
	mockStore := &MockStudentStore{}
	h := handler.New(mockStore)

	mockStore.On("GetByID", mock.Anything, "1").Return(&model.Student{ID: 1, Name: "John Doe", Age: 20, Class: "Math"}, nil)

	mockContext := new(gofr.Context)
s
	result, err := h.GetByID(mockContext)

	expected := &model.Student{ID: 1, Name: "John Doe", Age: 20, Class: "Math"}
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockStore.AssertExpectations(t)
}

func TestCreateHandler(t *testing.T) {
	mockStore := &MockStudentStore{}
	h := handler.New(mockStore)

	mockStore.On("Create", mock.Anything, mock.AnythingOfType("*model.Student")).Return(&model.Student{ID: 1, Name: "John Doe", Age: 20, Class: "Math"}, nil)

	mockContext := new(gofr.Context)

	result, err := h.Create(mockContext)

	expected := &model.Student{ID: 1, Name: "John Doe", Age: 20, Class: "Math"}
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockStore.AssertExpectations(t)
}

func TestUpdateHandler(t *testing.T) {
	mockStore := &MockStudentStore{}
	h := handler.New(mockStore)

	mockStore.On("Update", mock.Anything, mock.AnythingOfType("*model.Student")).Return(&model.Student{ID: 1, Name: "Updated John Doe", Age: 25, Class: "Updated Math"}, nil)

	mockContext := new(gofr.Context)

	result, err := h.Update(mockContext)

	expected := &model.Student{ID: 1, Name: "Updated John Doe", Age: 25, Class: "Updated Math"}
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockStore.AssertExpectations(t)
}

func TestDeleteHandler(t *testing.T) {

	mockStore := &MockStudentStore{}

	h := handler.New(mockStore)

	mockStore.On("Delete", mock.Anything, 1).Return(nil)

	mockContext := new(gofr.Context)

	result, err := h.Delete(mockContext)

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"message": "Deletion successful", "id": 1}, result)

	mockStore.AssertExpectations(t)
}