package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//create a file and check for err
	f, err := os.Create("/tmp/testlog.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	//SetOutput takes a writer type and sets the output destination
	//since type file implements the writer interface, we can pass a *os.File to it (i.e. the `f` created above)
	log.SetOutput(f)

	f2, err := os.Open("nofile.txt")
	if err != nil {
		//if you used fmt.Println here, nothing would get written the log
		log.Println("Error happened", err)
	}
	defer f2.Close()

	//just a reminder to check the directory since there's no error in the console output
	fmt.Println("Check the /tmp/testlog.txt file in the directory")

}
