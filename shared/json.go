package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJsonFile[T interface{}](filename string, v T) error {
	err := FileExtensionNotAllowed(filename)
	if err != nil {
		return err
	}

	jsonFile, err := os.Open(filename + ".json")
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

func WriteJsonFile[T interface{}](filename string, v T) error {
	err := FileExtensionNotAllowed(filename)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename+".json", data, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Data written to file: ", filename, "\n")

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
