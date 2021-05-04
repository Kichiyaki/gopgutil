package gopgutil

import (
	"github.com/go-pg/pg/v10/orm"
	"strings"
	"testing"
)

func init() {
	registerModels()
}

func TestOrderAppender_Apply(t *testing.T) {
	t.Run("wrong relationName", func(t *testing.T) {
		_, err := (&OrderAppender{Orders: []string{"testModel.story.user.id"}}).Apply(orm.NewQuery(nil, &testModel{}))
		if err != ErrRelationNotFound {
			t.Errorf("expected \"%s\", got %v", ErrRelationNotFound, err)
		}
	})

	t.Run("max depth level", func(t *testing.T) {
		_, err := (&OrderAppender{Orders: []string{"testModel2.story.user.id"}, MaxDepth: 2}).Apply(orm.NewQuery(nil, &testModel{}))
		if err == nil || !strings.Contains(err.Error(), "depth") {
			t.Errorf("expected error about max depth level, got %v", err)
		}
	})

	t.Run("success", func(t *testing.T) {
		tests := []struct {
			shouldContain string
			orders        []string
		}{
			{
				shouldContain: `"test_model2"."id", "test_model"."id", "field1" DESC`,
				orders:        []string{"testModel2.id", "testModel.id", "field1 DESC"},
			},
			{
				shouldContain: `"test_model2__story"."id", "field1" DESC`,
				orders:        []string{"testModel2.story.id", "field1 DESC"},
			},
			{
				shouldContain: `"test_model2__story"."id"`,
				orders:        []string{"testModel2.story.user.id"},
			},
			{
				shouldContain: `"field1" DESC, "id" ASC`,
				orders:        []string{"field1 DESC", "id ASC"},
			},
		}
		for _, test := range tests {
			q, err := (&OrderAppender{Orders: test.orders}).Apply(orm.NewQuery(nil, &testModel{}))
			if err != nil {
				t.Errorf("expected nil, got %s", err)
			}
			query := orm.NewSelectQuery(q).String()
			if !strings.Contains(query, test.shouldContain) {
				t.Errorf(`query "%s" doesn't contain any fragment like "%s"`, query, test.shouldContain)
			}
		}
	})
}
