package gopgutil

import (
	"testing"
)

func init() {
	registerModels()
}

func TestBuildAliasFromRelationName(t *testing.T) {
	model := &testModel{}

	t.Run("relation not found", func(t *testing.T) {
		alias, err := BuildAliasFromRelationName(model, "TestUser")
		if err != ErrRelationNotFound {
			t.Errorf("Expected %v, got %v", ErrRelationNotFound, err)
		}
		if alias != "" {
			t.Errorf("Expected empty string, got %v", alias)
		}
	})

	t.Run("deep relation not found", func(t *testing.T) {
		alias, err := BuildAliasFromRelationName(model, "TestModel2.Story.Test2")
		if err != ErrRelationNotFound {
			t.Errorf("Expected %v, got %v", ErrRelationNotFound, err)
		}
		if alias != "" {
			t.Errorf("Expected empty string, got %v", alias)
		}
	})

	t.Run("success", func(t *testing.T) {
		tests := []struct {
			relationName   string
			expectedResult string
		}{
			{
				relationName:   "TestModel2",
				expectedResult: "test_model2",
			},
			{
				relationName:   "TestModel2.Story",
				expectedResult: "test_model2__story",
			},
			{
				relationName:   "TestModel2.Story.User",
				expectedResult: "test_model2__story__user",
			},
		}
		for _, test := range tests {
			alias, err := BuildAliasFromRelationName(model, test.relationName)
			if err != nil {
				t.Errorf("Expected nil, got %v", err)
			}
			if alias != test.expectedResult {
				t.Error("Incorrect alias")
			}
		}
	})
}
