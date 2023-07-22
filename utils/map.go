package utils

import "fmt"

func MapExample() {
	map1 := map[string]string{
		"villain": "METRO",
		"hero":    "OCEAAN",
		"creep":   "REWW",
	}

	for k, v := range map1 {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}
	fmt.Printf("Super value: %#v", map1)
}
