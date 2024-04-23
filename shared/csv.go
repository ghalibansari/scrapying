package shared

import (
	"encoding/csv"
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

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(header)
	if err != nil {
		return err
	}

	for _, v := range data {
		err = writer.Write(v)
		if err != nil {
			return err
		}
	}

	defer writer.Flush()

	return nil
}
