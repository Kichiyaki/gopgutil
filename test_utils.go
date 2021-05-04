package gopgutil

import "github.com/go-pg/pg/v10/orm"

type user struct {
	ID      int
	Name    string
	Emails  []string
	Stories []*story `pg:"rel:has-many"`
}

type story struct {
	ID     int
	Title  string
	UserID int
	User   *user `pg:"rel:has-one"`
}

type testModel2 struct {
	ID      int
	StoryID int
	Story   *story `pg:"rel:has-one"`
}

type testModel struct {
	ID           int
	Field1       string
	TestModel2ID int
	TestModel2   *testModel2 `pg:"rel:has-one"`
}

func registerModels() {
	models := []interface{}{
		&testModel{},
		&testModel2{},
		&story{},
		&user{},
	}
	for _, model := range models {
		orm.RegisterTable(model)
	}
}
