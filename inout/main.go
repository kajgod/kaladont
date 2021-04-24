package inout

import (
	"bufio"
	"log"
	"os"
)

// Read imports word data from a file called rijeci.txt
func Read() []string {
	r := []string{}
	f, err := os.Open("./rijeci.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		r = append(r, s.Text())
	}
	return r
}
