package input

import (
    "fmt"
    "os"
)

func GetInput(filename string) string {
  data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	} else {
    return string(data)
  }
}
