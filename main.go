package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nkanaev/numb/pkg/runtime"
)

func main() {
	log.SetFlags(0)

	flag.Parse()

	expr := flag.Arg(0)
	if expr == "" {
		log.Fatal("need 1 arg")
	}

	rt := runtime.NewRuntime()
	rt.LoadBuiltins()

	const stdinMarker = "{}"
	if strings.Contains(expr, stdinMarker) {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			line := sc.Text()
			expr := strings.ReplaceAll(expr, stdinMarker, line)

			r, err := rt.Eval(expr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(r)
		}
		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}
		return
	}

	r, err := rt.Eval(expr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
