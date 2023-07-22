package loop

import "fmt"

func FourTypeOfFor() {
	// Condition for
	// init; condition; post-process
	for i := 0; i < 100; i++ {
		fmt.Printf("%v\n", i)
	}

	arr := []string{"you", "are", "the", "go", "dev"}
	for index, value := range arr {
		fmt.Printf("At %v: %v\n", index, value)
	}

	besties := map[string]string{
		"dart": "my cutie, my sweetness and my eternal love",
		"go":   "awesome guy with some what real useful",
		"js":   "I hate you",
	}

	for key, value := range besties {
		fmt.Printf("%+v: %+v\n", key, value)
	}

	//// for-ever
	// for true {}
	//// equalvalent with for true
	// for {}
}
