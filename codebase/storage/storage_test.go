package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func setUp() {
	defaultData := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	jsonData, _ := json.Marshal(defaultData)
	err := ioutil.WriteFile(JsonTestPath, jsonData, 0644)
	if err != nil {
		log.Panic(err)
	}
}

func tearDown() {
	testStorage, err := os.OpenFile(JsonTestPath, os.O_RDWR, 0644)
	defer testStorage.Close()
	if err != nil {
		log.Panic(err)
	}
	err = testStorage.Truncate(0)
	if err != nil {
		log.Panic(err)
	}
	_, err = testStorage.Seek(0, 0)
	if err != nil {
		log.Panic(err)
	}
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestSet(t *testing.T) {
	Set("key4", "value4", true)
	data := getStorage(true)
	if val, ok := data["key4"]; ok {
		if val != "value4" {
			t.Error("Value don't equals value4:", val)
		}
	} else {
		t.Error("Key doesn't exists:", "key4")
	}
}

func TestGet(t *testing.T) {
	value := Get("key1", true)
	if value != "value1" {
		t.Error("Value don't equals value1:", value)
	}
}

func TestDel(t *testing.T) {
	Del("key1", true)
	data := getStorage(true)
	if _, ok := data["key1"]; ok {
		t.Error("key1 exists in storage")
	}
}
