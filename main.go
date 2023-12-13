package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		log.Fatalln("Please provide a file.")
	}

	f, err := ioutil.ReadFile(args[1])

	if err != nil {
		log.Fatalln("Error: " + err.Error())
	}

	code := string(f)

	base := "local cells = {0}"

	cur := 1
	var set = make(map[int]bool)
	set[1] = true

	instructions := strings.Split(code, "")

	for _, instruction := range instructions {
		switch instruction {
		case "+":
			base += fmt.Sprintf("cells[%v] = cells[%v] + 1;\n", cur, cur)
		case "-":
			base += fmt.Sprintf("cells[%v] = cells[%v] - 1;\n", cur, cur)
		case ">":
			cur++
			if !set[cur] {
				base += fmt.Sprintf("cells[%v] = 0;\n", cur)
			}
		case "<":
			cur--
			if cur < 1 {
				log.Fatalln("ERROR: Grid escape detected at instruction " + instruction)
			}
		case "[":
			base += fmt.Sprintf("while cells[%v] ~= 0 do \n", cur)
		case "]":
			base += "end;\n"
		case ".":
			base += fmt.Sprintf("print(string.char(cells[%v]));\n", cur)
		case ",":
			base += fmt.Sprintf("cells[%v] = io.read();\nif #cells[%v] > 1 then\n error('Expected string of length 1')\nend;\ncells[%v] = string.byte(cells[%v]);\n", cur, cur, cur, cur)
		}
	}
	ioutil.WriteFile("out.lua", []byte(base), 0644)

	fmt.Println("Outputed to out.lua")
}
