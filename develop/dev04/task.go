package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю. 


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.


Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого, 
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру. 
В результате каждое слово должно встречаться только один раз.
*/

func main() {
	dictionary := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	result := anagrams(&dictionary)

	for key, value := range result {
		if len(value) > 1 {
			fmt.Printf("Собственное множество для анаграмм %s: %v\n", key, value)
		}
	}
}

func anagrams(words *[]string) map[string][]string {
	plentyAnagram := make(map[string][]string)

	for _, word := range *words {
		wordLower := sortString(strings.ToLower(word))

		if set, found := plentyAnagram[wordLower]; found {
			plentyAnagram[wordLower] = append(set, word)

			continue
		}

			plentyAnagram[wordLower] = []string{word}
	}

	for key, value := range plentyAnagram {
		if len(value) <= 1 {
			delete(plentyAnagram, key)

			continue
		}

			sort.Strings(plentyAnagram[key])
	}

	return plentyAnagram
}

func sortString(s string) string {
	sortRunes := []rune(s)
	sort.Slice(sortRunes, func(i, j int) bool {
		return sortRunes[i] < sortRunes[j]
	})

	return string(sortRunes)
}