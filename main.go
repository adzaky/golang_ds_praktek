package main

import "fmt"

func doLoop(any interface{}) {
	switch v := any.(type) {
	case []string:
		for key, value := range v {
			fmt.Println("Ini key : ", key+1)
			fmt.Println("Ini value : ", value)
		}
		fmt.Println()
	case map[string]string:
		for key, value := range v {
			fmt.Println("Ini key : ", key)
			fmt.Println("Ini value : ", value)
		}
		fmt.Println()
	case [4]string:
		for key, value := range v {
			fmt.Println("Ini key : ", key+1)
			fmt.Println("Ini value : ", value)
		}
		fmt.Println()
	}
}

func main() {
	names := [4]string{
		"Muhammad", "Adlim", "Dzaky", "Pratama",
	}

	doLoop(names)

	var grade []int
	grade = append(grade, 100)
	grade = append(grade, 90)
	grade = append(grade, 80)

	doLoop(grade)

	mk := make (map[string]string)
	mk["Kelas A"] = "Matematika"
	mk["Kelas B"] = "IPA"
	mk["Kelas C"] = "Sejarah"

	doLoop(mk)
}