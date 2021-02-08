package main

import "fmt"

func main() {
	a := map[string]map[string]string{
		"hiero7": {
			"name":   "leo",
			"number": "F022",
		},
		"greatforti": {
			"name":   "leo",
			"number": "G008",
			"status": "leave",
		},
	}

	fmt.Println(a)

	x := make(map[string]int)

	fmt.Println(x)

	x["key"] = 10

	fmt.Println(x)

	y := make(map[int]string)

	y[1] = "leo"

	fmt.Println(y)

	z := make(map[int]int)
	z[0] = 123
	z[1] = 222
	fmt.Println("z", len(z), z)

	// 刪除 key
	delete(z, 0)
	fmt.Println("z", len(z), z)
	fmt.Println("z", len(z), z[4])

	// 判斷 map 裡的 key 是否存在
	name, status := z[4]
	fmt.Println(name, status)

	z[4] = 9999
	if name, status := z[4]; status {
		fmt.Println(name, status)
	}

	elements := map[string]map[string]string{
		"H": {
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

	if el, ok := elements["H"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	// fmt.Println(elements)

}
