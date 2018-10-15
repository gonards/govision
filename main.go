// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"fmt"
)

func main() {

	results := getLabels("./test/tete.JPG")
	for name, result := range results {
		fmt.Println(name)
		fmt.Println("-------")
		for _, label := range result {
			fmt.Println(label)
		}
		fmt.Println("---------------------------------")
	}
}
