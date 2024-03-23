package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJsonFile[T interface{}](filename string, v T) error {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	json.Unmarshal(byteValue, v)

	return nil
}

func PrintJson[T interface{}](v T) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling json data:", err)
		return
	}
	fmt.Println(string(data))
}
