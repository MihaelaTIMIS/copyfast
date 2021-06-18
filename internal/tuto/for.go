package tuto

import "fmt"

func createFor() {

	// var i = 0
	// while
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// for (i = 0)
	// for
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// do while
	for {
		fmt.Println("ok")
		if true {
			break
		}
	}

	names := []string{"Florent", "Toto", "Tata"}
	// foreach
	for index, name := range names {
		fmt.Println(index, name) // Florent, Toto, Tata
	}

	for _, name := range names {
		fmt.Println(name) // Florent, Toto, Tata
	}

	for index := range names {
		fmt.Println(index) // 0,1,2
	}

}
