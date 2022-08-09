package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/nkanaev/numb/pkg/runtime"
	"golang.org/x/term"
)

var prompt = "> "
var prefix = "  "

type poserror interface {
	error
	Pos() (int, int)
}

func repl(rt *runtime.Runtime) {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), state)

	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	terminal := term.NewTerminal(screen, prompt)
	terminal.Write([]byte("enter `q` to quit\n"))

	for {
		line, err := terminal.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			os.Stderr.WriteString(line + "\n")
			os.Exit(1)
		}
		if line == "q" {
			break
		}
		out, err := rt.Eval(line)
		if err != nil {
			switch err.(type) {
			case poserror:
				err := err.(poserror)
				start, end := err.Pos()
				out = strings.Repeat(" ", start)
				out += strings.Repeat("^", end-start+1) + "\n"
				out += prefix + err.Error()
			default:
				out = err.Error()
			}
		}
		if out != "" {
			terminal.Write([]byte(prefix + out + "\n"))
		}
	}
}

func read(rt *runtime.Runtime, r io.Reader) {
	var maxqwidth, maxawidth int
	qlines := make([]string, 0)
	alines := make([]string, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		qline := runtime.Clean(scanner.Text())
		qlines = append(qlines, qline)
		if len(qline) > maxqwidth {
			maxqwidth = len(qline)
		}

		aline, err := rt.Eval(qline)
		if err != nil {
			alines = append(alines, "! "+err.Error())
			continue
		}
		alines = append(alines, aline)

		awidth := len(aline)
		spaceidx := strings.Index(aline, " ")
		if spaceidx != -1 {
			awidth = spaceidx
		}

		if awidth > maxawidth {
			maxawidth = awidth
		}
	}

	for i := 0; i < len(qlines); i++ {
		q, a := qlines[i], alines[i]

		var apad int
		if len(a) > 0 && a[0] != '!' {
			apad = maxawidth - len(a)
			spaceidx := strings.Index(a, " ")
			if spaceidx != -1 {
				apad = maxawidth - spaceidx
			}
		}
		if len(a) > 0 {
			fmt.Printf("%s%s    | %s%s\n",
				q, strings.Repeat(" ", maxqwidth-len(q)),
				strings.Repeat(" ", apad), a,
			)
		} else {
			fmt.Println(q)
		}
	}
}

func main() {
	rt := runtime.NewRuntime()
	rt.LoadBuiltins()

	flag.Parse()

	rcfile := path.Join(os.Getenv("HOME"), ".numbrc")
	if info, err := os.Stat(rcfile); err == nil && info.Mode().IsRegular() {
		rt.LoadFile(rcfile)
	}

	if flag.NArg() == 1 {
		line, err := rt.Eval(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(line)
		return
	}

	if term.IsTerminal(0) {
		repl(rt)
	} else {
		read(rt, os.Stdin)
	}
}
