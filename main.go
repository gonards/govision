// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"flag"
	"fmt"
)

func main() {
	pathFlag := flag.String("p", "./", "Image to process path")
	flag.Parse()
	fmt.Println("Image processing ...")
	results := GetLabels(*pathFlag)
	for name, result := range results {
		fmt.Println(name)
		fmt.Println("-------")
		for _, label := range result {
			fmt.Println(label)
		}
		fmt.Println("---------------------------------")
	}
}
