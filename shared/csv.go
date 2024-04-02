package shared

import (
	"fmt"
	"os"
)

func WriteCsvFile(filename string, header []string, data [][]string) error {
	err := FileExtensionNotAllowed(filename)
	if err != nil {
		return err
	}

	if len(header) != len(data[0]) {
		return fmt.Errorf("header and data length mismatch")
	}

	file, err := os.Create(filename + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(ToCsv(header))
	if err != nil {
		return err
	}

	for _, v := range data {
		_, err = file.WriteString(ToCsv(v))
		if err != nil {
			return err
		}
	}

	return nil
}

func ToCsv(data []string) string {
	var csv string
	for i, v := range data {
		csv += v
		if i != len(data)-1 {
			csv += ","
		}
	}
	csv += "\n"
	return csv
}
