package gopgutil

import (
	"strings"
	"testing"
)

func init() {
	registerModels()
}

func TestBuildAliasFromRelationName(t *testing.T) {
	t.Run("relation not found", func(t *testing.T) {
		alias, err := BuildAliasFromRelationName(&testModel{}, "TestUser")
		if err != ErrRelationNotFound {
			t.Errorf("Expected %v, got %v", ErrRelationNotFound, err)
		}
		if alias != "" {
			t.Errorf("Expected empty string, got %v", alias)
		}
	})

	t.Run("deep relation not found", func(t *testing.T) {
		alias, err := BuildAliasFromRelationName(&testModel{}, "TestModel2.Story.Test2")
		if err != ErrRelationNotFound {
			t.Errorf("Expected %v, got %v", ErrRelationNotFound, err)
		}
		if alias != "" {
			t.Errorf("Expected empty string, got %v", alias)
		}
	})

	t.Run("should return an error when model = nil", func(t *testing.T) {
		alias, err := BuildAliasFromRelationName(nil, "TestUser")
		if err == nil || !strings.Contains(err.Error(), "Invalid model") {
			t.Errorf("Expected error about invalid model, got %v", err)
		}
		if alias != "" {
			t.Errorf("Expected empty string, got %v", alias)
		}
	})

	t.Run("success", func(t *testing.T) {
		tests := []struct {
			relationName   string
			expectedResult string
			model          interface{}
		}{
			{
				relationName:   "TestModel2",
				expectedResult: "test_model2",
				model:          &testModel{},
			},
			{
				relationName:   "TestModel2.Story",
				expectedResult: "test_model2__story",
				model:          &testModel{},
			},
			{
				relationName:   "TestModel2.Story.User",
				expectedResult: "test_model2__story__user",
				model:          &testModel{},
			},
			{
				relationName:   "User",
				expectedResult: "user",
				model:          &story{},
			},
		}
		for _, test := range tests {
			alias, err := BuildAliasFromRelationName(test.model, test.relationName)
			if err != nil {
				t.Errorf("Expected nil, got %v", err)
			}
			if alias != test.expectedResult {
				t.Error("Incorrect alias")
			}
		}
	})
}
