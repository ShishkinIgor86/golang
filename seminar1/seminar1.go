package main

import (
	"fmt"
	"log"
	"os"
    "strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	var fileName, fileExt string
	// Напишите код, который выведет следующее
	pointIndex := strings.LastIndex(filePth,".")
	if pointIndex >= 0{
		fileName = filePth[strings.LastIndex(filePth,"/")+1: strings.LastIndex(filePth,".")]
		fileExt = filePth[pointIndex+1:]
		} else {
			fileName = filePth[strings.LastIndex(filePth,"/")+1:]
			fileExt = ""
		}
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения
    fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)

}
