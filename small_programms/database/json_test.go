package database

import (
	"testing"
	"fmt"
	"os"
)

type TestField2 struct {
	Name string `json:"name"`
}

var jsonDataBase DataBase

func TestMain(m *testing.M) {
	var err error
	jsonDataBase, err = NewJsonDataBase("./store/")
	if err != nil {
		panic(err)
	}
	retCode := m.Run()

	os.Exit(retCode)
}

func TestDataJsonBase_Create(t *testing.T) {
	var err error

	_, err = jsonDataBase.Create(TestField2{Name: "test_11"})
	if err != nil {
		t.Fatalf("error during created: %s", err)
	}

	_, err = jsonDataBase.Create(TestField2{Name: "test_22"})
	if err != nil {
		t.Fatalf("error during created: %s", err)
	}

	res, err := jsonDataBase.Get(TestField2{})
	if err != nil {
		t.Fatal(err)
	}

	entities := res.([]TestField2)
	if len(entities) < 1 {
		t.Fatalf("Get did not return entities")
	}

	for i, en := range entities {
		fmt.Println("en", en)
		en.Name = en.Name + "_updated"
		err = jsonDataBase.Update(i, en)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDataJsonBase_Update(t *testing.T) {
	jsonDataBase.Create(TestField2{
		Name: "test_1",
	})

	newName := "tttt"
	err := jsonDataBase.Update(0, TestField2{
		Name: newName,
	})
	if err != nil {
		t.Fatal(err)
	}

	tf, err := jsonDataBase.GetOne(0, TestField2{})
	if err != nil {
		t.Fatal(err)
	}
	if tf.(TestField2).Name != newName {
		t.Fatal("update: name not changed")
	}
}
