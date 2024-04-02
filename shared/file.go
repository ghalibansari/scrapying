package shared

import (
	"fmt"
	"os"
)

func FileExtensionNotAllowed(filename string) error {
	if filename[len(filename)-5:] == ".json" {
		return fmt.Errorf("file extension not allowed")
	}
	return nil
}

// TODO
func ConvertJsonToCsv(v map[string]interface{}) ([]string, [][]string, error) {
	if len(v) == 0 {
		return nil, nil, fmt.Errorf("json data is empty")
	}

	for _, value := range v {
		if _, ok := value.(map[string]interface{}); ok {
			break
		}
		return nil, nil, fmt.Errorf("json data is not nested")
	}

	//get the header.
	var header []string
	var data [][]string

	for key, value := range v {
		if _, ok := value.(map[string]interface{}); ok {
			header = append(header, key)
			break
		}
	}

	//get the data
	for _, value := range v {
		var row []string
		for _, val := range value.(map[string]interface{}) {
			row = append(row, val.(string))
		}
		data = append(data, row)
	}

	return header, data, nil

}

func CreateFileIfNotExist(fileName string) (*os.File, error) {
	if !FileExists(fileName) {
		file, err := os.Create(fileName)
		if err != nil {
			return nil, fmt.Errorf("could not create file: %w", err)
		}
		return file, nil
	}

	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	return file, nil
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}
