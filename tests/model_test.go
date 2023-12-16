package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudent_MarshalJSON(t *testing.T) {
	student := &Student{
		ID:    1,
		Name:  "John Doe",
		Age:   20,
		Class: "Physics",
	}

	expectedJSON := {"id":1,"name":"John Doe","age":20,"class":"Physics"}

	actualJSON, err := json.Marshal(student)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(actualJSON))
}

func TestStudent_String(t *testing.T) {
	student := &Student{
		ID:    1,
		Name:  "John Doe",
		Age:   20,
		Class: "Physics",
	}

	expectedString := "\nId = 1\nName = John Doe\nAge = 20\nClass = Physics"

	actualString := student.String()
	assert.Equal(t, expectedString, actualString)
}