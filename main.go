package main

import (
//	"bufio"
	"fmt"
	"log"
//	"io"
	"io/ioutil"
//	"os"
	"flag"
	"regexp"
)

func readArgs() (action int, file_id string, mail bool ) {
	learning := flag.Bool("l", false, "set learning mode")
	rfc2822 := flag.Bool("r", false, "set rfc2822 mode - that is: strip headers")
	print := flag.Bool("p", false, "set printing mode")

	flag.Parse()
	
	if (*print) {
		action = 3
	} else if (*learning) {
		action = 2
	} else {
		action = 1
	}
	log.Printf("[FILTER] mode %d", action)

	mail = *rfc2822;	
	log.Printf("[FILTER] mail %s", mail)

	file_id = flag.Arg(0)
	if (file_id == "") {
		file_id = "/tmp/test"
	}
	log.Printf("[FILTER] reading file_id %s", file_id)
	return;
}

func print(file_id string, mail bool) {
	dat, err  := ioutil.ReadFile(file_id)
	if (err != nil) {
		log.Printf("[FILTER] can't open file %s", file_id)
	} else {
		if (mail) {
			re := regexp.MustCompile("(\n\n|\n\r\n\r)")
			index := re.FindIndex(dat);
			if (index != nil) {
				dat = dat[index[1]:];
			}
		}
		fmt.Print(string(dat))
	}
}

func filter(file_id string, mail bool) {

}

func main() {
	action, file_id, mail := readArgs()

	switch action {
		case 1:
			filter(file_id, mail);
		case 2:
			log.Printf("[FILTER] learning not implemented yet")
		case 3: 
			print(file_id, mail);
	}
}
