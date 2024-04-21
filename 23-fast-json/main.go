package main

import (
	"fmt"
	"github.com/valyala/fastjson"
	"slices"
)

func main() {
	var parser fastjson.Parser
	jsonData := `{"nome": "Breno", "idade": 23, "bool": true, "arr":[12,11,10]}`

	parseValue, err := parser.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Nome: %s\n", parseValue.GetStringBytes("nome"))
	fmt.Printf("Idade: %d\n", parseValue.GetInt("idade"))
	fmt.Printf("Bool: %v\n", parseValue.GetBool("bool"))

	array := parseValue.GetArray("arr")
	for index, value := range array {
		fmt.Printf("Index: %d - Value: %s\n", index, value)
	}

	fmt.Printf("\n")
	userJsonData := `{"user": {"name": "Breno", "age": 31, "email": "breno@email.com", "role": "admin"}}`
	parseValue, err = parser.Parse(userJsonData)
	if err != nil {
		panic(err)
	}

	user := parseValue.GetObject("user")
	fmt.Printf("Name: %s\n", user.Get("name"))
	fmt.Printf("Email: %s\n", user.Get("email"))
	fmt.Printf("Role: %s\n", user.Get("role"))
	fmt.Printf("Age: %s\n", user.Get("age"))

	fmt.Printf("JSON: %s\n", user.String())

	arrayString := [][]string{
		{"a", "b", "c"},
		{"d", "e", "a"},
		{"g", "o", "i"},
	}

	fmt.Println(CheckVowelMatrix(arrayString))
}

func CheckVowelMatrix(matrix [][]string) string {
	for i := 0; i < (len(matrix) - 1); i++ {
		line := matrix[i]
		for j := 0; j < (len(line) - 1); j++ {
			isFirstVowel := isVowel(line[j])

			if isFirstVowel {
				isSecondLetterAVowel := isVowel(matrix[i][j+1])
				if !isSecondLetterAVowel {
					continue
				}

				isThirdLetterAVowel := isVowel(matrix[i+1][j])
				if !isThirdLetterAVowel {
					continue
				}

				if isVowel(matrix[i+1][j+1]) {
					return fmt.Sprintf("%d-%d", i, j)
				}
			}
		}
	}
	return "not found"
}

func isVowel(letter string) bool {
	vowels := []string{
		"a", "e", "i", "o", "u",
	}

	return slices.Contains(vowels, letter)
}
