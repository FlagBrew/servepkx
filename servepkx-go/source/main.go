// PKSM uploader by Allen Lydiard
// PKSM is developed by Bernardo Giordano
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/schollz/progressbar"
)

func main() {
	// check to see if we have any arguments first
	if len(os.Args) > 1 {
		fmt.Println("Looking for 3DS IP, please wait!")
		ip := getIP()
		file, err := os.Stat(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		if file.IsDir() {
			pokemon, err := ioutil.ReadDir(os.Args[1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Uploading " + strconv.Itoa(len(pokemon)) + " Pokemon please wait!")
			bar := progressbar.New(len(pokemon))
			for _, pkmn := range pokemon {
				if strings.HasSuffix(pkmn.Name(), "pk6") || strings.HasSuffix(pkmn.Name(), "pk7") ||
					strings.HasSuffix(pkmn.Name(), "wc7full") || strings.HasSuffix(pkmn.Name(), "wc6full") || strings.HasSuffix(pkmn.Name(), "wc6") || strings.HasSuffix(pkmn.Name(), "wc7") {
					readpkmn(pkmn, fmt.Sprintf("%s/%s", os.Args[1], pkmn.Name()), ip)
					bar.Add(1)
				}
			}
			bar.Finish()
			fmt.Println("")
		} else {
			if strings.HasSuffix(file.Name(), "pk6") || strings.HasSuffix(file.Name(), "pk7") ||
				strings.HasSuffix(file.Name(), "wc7full") || strings.HasSuffix(file.Name(), "wc6full") || strings.HasSuffix(file.Name(), "wc6") || strings.HasSuffix(file.Name(), "wc7") {
				fmt.Println("Uploading " + file.Name() + " please wait!")
				readpkmn(file, os.Args[1], ip)
			} else {
				fmt.Println("This is not a valid pokemon file!")
				os.Exit(1)
			}
		}
		fmt.Println("Upload Complete!")
		fmt.Println("Have a nice day!")
		os.Exit(0)
	}
	fmt.Println("You must choose a pokemon file to upload!")
}
