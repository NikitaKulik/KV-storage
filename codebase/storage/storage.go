/**

KV storage into JSON format

*/

package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const JsonPath = "./storage/storage.json"
const JsonTestPath = "./storage/storage_test.json"

// Function for convert data from JSON to map
func getStorage(isTestMode bool) map[string]string {
	storagePath := JsonPath
	if isTestMode {
		storagePath = JsonTestPath
	}

	jsonFile, err := os.Open(storagePath)
	if err != nil {
		log.Panic(err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panic(err)
	}

	data := make(map[string]string)
	if len(bytes) > 0 {
		err = json.Unmarshal([]byte(bytes), &data)
		if err != nil {
			log.Panic(err)
		}
	}

	return data
}

// Function for "SET" handling
func Set(key string, value string, isTestMode bool) {
	storagePath := JsonPath
	if isTestMode {
		storagePath = JsonTestPath
	}

	storage := getStorage(isTestMode)
	storage[key] = value
	jsonData, _ := json.Marshal(storage)
	err := ioutil.WriteFile(storagePath, jsonData, 0644)
	if err != nil {
		log.Panic(err)
	}
}

// Function for "GET" handling
func Get(key string, isTestMode bool) string {
	storage := getStorage(isTestMode)
	value := storage[key]
	log.Printf(value)
	return value
}

// Function for "DEL" handling
func Del(key string, isTestMode bool) {
	storagePath := JsonPath
	if isTestMode {
		storagePath = JsonTestPath
	}

	storage := getStorage(isTestMode)
	delete(storage, key)
	jsonData, _ := json.Marshal(storage)
	err := ioutil.WriteFile(storagePath, jsonData, 0644)
	if err != nil {
		log.Panic(err)
	}
}
