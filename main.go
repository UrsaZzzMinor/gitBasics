package main

import "fmt"

func find(arr []int, k int) []int {
	m := make(map[int]int)
	for i, a := range arr {
		if j, ok := m[k-a]; ok {
			return []int{i, j}
		}
		m[a] = i
	}

	return nil
}

func removeDuplicates(input []string) []string {
	output := make([]string, 0)
	inputSet := make(map[string]struct{}, len(input))
	for _, v := range input {
		if _, ok := inputSet[v]; !ok {
			output = append(output, v)
		}
		inputSet[v] = struct{}{}
	}
	return output
}

func main() {

	input := []string{
		"dog",
		"cat",
		"dog",
		"parrot",
		"cat",
		"bird",
	}

	fmt.Println(input)

	fmt.Println(removeDuplicates(input))

	/*arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 5
	fmt.Println(find(arr, k))*/
	/*sentence := "Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου"

	frequency := make(map[rune]int)
	for _, v := range sentence {
		frequency[v]++
	}

	for k, v := range frequency {
		fmt.Printf("Sign %c was met %d times\n", k, v)
	}

	//	another task

	pricelist := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200,
		"сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	total := 0
	fmt.Printf("List of yummy things:")

	for k, v := range pricelist {
		if v > 500 {
			fmt.Println(k)
		}
	}

	for _, v := range order {
		total += pricelist[v]
	}
	fmt.Printf("The order value is %d", total)*/
}
