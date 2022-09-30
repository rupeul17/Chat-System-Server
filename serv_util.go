package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func input_number() int {
	rd := bufio.NewReader(os.Stdin)

	TmpInt, err := rd.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	TmpInt = strings.TrimSpace(TmpInt)
	input_num, err := strconv.Atoi(TmpInt)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return input_num
}
