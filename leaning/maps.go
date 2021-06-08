package main

import "fmt"

// Карта

// Карта (также известна как ассоциативный массив или словарь) — это неупорядоченная коллекция
// пар вида ключ-значение. Пример:

// var x map[string]int

// Карта представляется в связке с ключевым словом map,
// следующим за ним типом ключа в скобках и типом значения после скобок.
// Читается это следующим образом: «x — это карта string-ов для int-ов».

// Подобно массивам и срезам, к элементам карт можно обратиться с помощью скобок.

// Карта должна быть инициализирована перед тем, как будет использована. Надо написать так:

// x := make(map[string]int)
// x["key"] = 10
// fmt.Println(x["key"])

// Если выполнить эту программу, то вы должны увидеть 10.
// Выражение x["key"] = 10 похоже на те, что использовались при работе с массивами,
// но ключ тут не число, а строка (потому что в карте указан тип ключа string).
// Мы также можем создать карты с ключом типа int:

// x := make(map[int]int)
// x[1] = 10
// fmt.Println(x[1])

// Это выглядит очень похоже на массив, но существует несколько различий.
// Во-первых, длина карты (которую мы можем найти так: len(x)) может измениться,
// когда мы добавим в нее новый элемент. В самом начале при создании длина 0, после x[1] = 10
// она станет равна 1. Во-вторых, карта не является последовательностью.
// В нашем примере у нас есть элемент x[1], в случае массива должен быть и первый элемент x[0],
// но в картах это не так.

// Также мы можем удалить элементы из карты используя встроенную функцию delete: delete(x, 1)

func main() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
