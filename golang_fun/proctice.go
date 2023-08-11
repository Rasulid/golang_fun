package main

import (
	"encoding/json"
	"fmt"
)

/*
Напишите общую функцию, которая берет срез сопоставимых объектов и сортирует их в любом направлении
(по возрастанию или по убыванию). использовать интерфейсы для выполнения общей функции функции.
вы должны использовать псевдоним типа, например, для сортировки целых чисел.
*/

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

//func main() {
//	var s IntSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
//	sort.Sort(s)
//	fmt.Println(s)
//	fmt.Println(s.Len())
//	s.Swap(5, 6)
//	fmt.Println(s)
//	fmt.Println(s.Less(1, 2))
//}

/*
	Представьте, что у вас есть структура Movie,
которая включает поле Runtime в int. Когда вы маршалируете объект Movie в json,
поле Runtime будет отображаться в целое число json. Например,
измените его на маршал на строку json «x mins». Вы должны изучить интерфейс json.Marshaller
*/

type Movie struct {
	Title   string
	Runtime int
}

func (m Movie) MarshalJSON() ([]byte, error) {
	custom := struct {
		Title   string
		Runtime string
	}{
		Title:   m.Title,
		Runtime: fmt.Sprintf("%d mins", m.Runtime),
	}
	return json.Marshal(custom)
}

func main() {
	movie := Movie{
		Title:   "Star Wars",
		Runtime: 120,
	}
	jsonData, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}
