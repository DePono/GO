package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Обьявляем структуру Record
type Record struct {
	Question string
	Answer   string
}

// Обьявляем необходимые переменные
var questionWithAnswer = "problems.csv"
var dataFromFile []Record
var answer string
var rightAnswer int
var wrongAnswer int
var usersFile string

func readCSVFile(filepath string) ([][]string, error) {
	//  проверяем, существует ли указанный путь к файлу и связан ли он с обычным файлом внутри функции.
	fileInfo, err := os.Stat(questionWithAnswer)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(questionWithAnswer, "error: not a regular file")
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	// Читаем файл csv весь сразу в тип данных lines - [][]string
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func shuffleData() {
	rand.NewSource(time.Now().Unix())
	rand.Shuffle(len(dataFromFile), func(i, j int) {
		dataFromFile[i], dataFromFile[j] = dataFromFile[j], dataFromFile[i]
	})
}

func main() {

	fmt.Println("Выберете путь к файлу, где будут лежать вопросы с ответами. Введите 0, если необходимо использовать файл по умолчанию")
	_, _ = fmt.Scan(&usersFile)

	if usersFile == "0" {
		// Считываем файл
		lines, _ := readCSVFile(questionWithAnswer)
		// Заполняем Record
		for _, line := range lines {
			temp := Record{
				Question: line[0],
				Answer:   line[1],
			}
			dataFromFile = append(dataFromFile, temp)
			shuffleData()
		}
	} else {
		// Считываем файл
		lines, _ := readCSVFile(usersFile)
		// Заполняем Record
		for _, line := range lines {
			temp := Record{
				Question: line[0],
				Answer:   line[1],
			}
			dataFromFile = append(dataFromFile, temp)
		}
	}
	fmt.Println("Читайте вопрос,отвечаете и жмете enter. Ответ состоит из одного слова")
	for _, record := range dataFromFile {
		fmt.Println(record.Question)
		_, err := fmt.Scan(&answer)
		if strings.ToLower(answer) == strings.ToLower(record.Answer) {
			rightAnswer++
		} else {
			wrongAnswer++
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("верных ответов %d\nневерных ответов %d\n", rightAnswer, wrongAnswer)
}
