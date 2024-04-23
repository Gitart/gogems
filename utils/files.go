package utils

import (
	"fmt"
	"os"
	"time"
)

// Создание файла и
// запись в него построчно

func WriteFile(filename, text string) {

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	tm := time.Now().Format("02.01.2006 15:04:05 ")
	if _, err = f.WriteString(tm + text + "\n"); err != nil {
		fmt.Println(err)
	}
}
