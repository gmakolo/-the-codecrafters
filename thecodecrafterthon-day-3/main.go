package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter a comand. \n 1. upper [string] \n 2. lower [string] \n 3. cap [string] \n 4. title [string] \n 5. snake [string] \n 6. reverse [string] \n 7. exit")
		sc.Scan()
		t := strings.TrimSpace(sc.Text())
		if t == "" {
			continue
		}

		p := strings.SplitN(t, " ", 2)
		cmd := strings.ToLower(p[0])
		if cmd == "exit" {
			return
		}
		if len(p) < 2 {
			fmt.Println("No text")
			continue
		}

		s := p[1]

		switch cmd {
		case "upper":
			fmt.Println(strings.ToUpper(s))
		case "lower":
			fmt.Println(strings.ToLower(s))
		case "cap":
			fmt.Println(cap(s))
		case "title":
			fmt.Println(title(s))
		case "snake":
			fmt.Println(snake(s))
		case "reverse":
			fmt.Println(rev(s))
		default:
			fmt.Println("Unknown")
		}
	}
}

func cap(s string) string {
	w := strings.Fields(strings.ToLower(s))
	for i := range w {
		w[i] = strings.ToUpper(w[i][:1]) + w[i][1:]
	}
	return strings.Join(w, " ")
}

func title(s string) string {
	w := strings.Fields(strings.ToLower(s))
	for i := range w {
		if i != 0 && (w[i] == "a" || w[i] == "an" || w[i] == "the" || w[i] == "of" || w[i] == "in" || w[i] == "to") {
			continue
		}
		w[i] = strings.ToUpper(w[i][:1]) + w[i][1:]
	}
	return strings.Join(w, " ")
}

func snake(s string) string {
	var r []rune
	for _, c := range strings.ToLower(s) {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			r = append(r, c)
		}
		if unicode.IsSpace(c) {
			r = append(r, '_')
		}
	}
	return string(r)
}

func rev(s string) string {
	w := strings.Fields(s)
	for i := range w {
		r := []rune(w[i])
		for l, h := 0, len(r)-1; l < h; l, h = l+1, h-1 {
			r[l], r[h] = r[h], r[l]
		}
		w[i] = string(r)
	}
	return strings.Join(w, " ")
}
