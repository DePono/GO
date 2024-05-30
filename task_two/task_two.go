package main

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"os"
	"sort"
)

type Candidate struct {
	Name  string
	Votes int
}

type Students struct {
	Students []Student `json:"students"`
}

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}
type Objects struct {
	Objects []Object
}

type Object struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Results struct {
	Results []Result `json:"results"`
}
type Result struct {
	ObjectId  int `json:"object_id"`
	StudentId int `json:"student_id"`
	Result    int `json:"result"`
}

// Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int. Каждый элемент в пересечении должен быть уникальным.
// Слайс-результат должен быть отсортирован в восходящем порядке.Примеры:
// 1) Если на вход подается только 1 слайс [1, 2, 3, 2], результатом должен быть слайс [1, 2, 3].
// 2) Вход: 2 слайса [1, 2, 3, 2] и [3, 2], результат - [2, 3].
// 3) Вход: 3 слайса [1, 2, 3, 2], [3, 2] и [], результат - [].

func crossUnique(slices ...[]int) []int {
	numberSlice := len(slices)
	resultSlice := make([]int, 0)
	resultSliceTwo := make([]int, 0)
	result := make(map[int]int, numberSlice)
	//Прверяем на длину поступивших массивов
	for _, slice := range slices {
		if len(slice) != 0 {
			for _, element := range slice {
				resultSlice = append(resultSlice, element)
			}
		} else {
			// Если хоть одна длина массива равно 0, то возвращаем пустой массив
			return nil
		}
	}
	if numberSlice == 0 {
		return nil
	} else if numberSlice == 1 {
		for i := 0; i < len(resultSlice); i++ {
			result[resultSlice[i]] += 1
		}
		for i, v := range result {
			if v > 0 {
				resultSliceTwo = append(resultSliceTwo, i)
			}
		}
		sort.Ints(resultSliceTwo)
		return resultSliceTwo
	} else if numberSlice > 1 {
		for i := 0; i < len(resultSlice); i++ {
			result[resultSlice[i]] += 1
		}
		for i, v := range result {
			if v > 1 {
				resultSliceTwo = append(resultSliceTwo, i)
			}
		}
		sort.Ints(resultSliceTwo)
		return resultSliceTwo
	}
	return resultSliceTwo
}

func countVoices(arrayNames []string) []Candidate {
	arrayCandidates := make([]Candidate, 0)
	candidates := make(map[string]int)
	for i := 0; i < len(arrayNames); i++ {
		candidates[arrayNames[i]] += 1
	}
	for k, v := range candidates {
		temp := Candidate{
			Name:  k,
			Votes: v,
		}
		arrayCandidates = append(arrayCandidates, temp)
	}
	return arrayCandidates
}

func readJSON(filename string) (students Students, objects Objects, results Results) {
	//Открываем файл по пути
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening file:", err)
	}
	// Анмаршеллим данные

	err = json.Unmarshal(content, &students)
	err = json.Unmarshal(content, &objects)
	err = json.Unmarshal(content, &results)
	if err != nil {
		log.Fatal("Error when unmarshalling:", err)
	}
	return students, objects, results
}

func PrintTable(students Students, objects Objects, results Results) {
	tab := table.NewWriter()
	tab.AppendHeader(table.Row{"Student name", "Grade", "Object", "Result"})
	for _, object := range objects.Objects {
		for _, result := range results.Results {
			for _, student := range students.Students {

				tab.AppendRows([]table.Row{{student.Name, student.Grade, object.Name, result.Result}})
			}
		}
	}
	fmt.Println(tab.Render())
}

func SummaryTable(students Students, objects Objects, results Results) {

}

