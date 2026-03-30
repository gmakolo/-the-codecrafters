
// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: Makolo Eleojo Gift
// Squad: Alpha Squad

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var smallWords = map[string]bool{
	"a": true, "an": true, "the": true, "and": true, "but": true, "or": true,
	"for": true, "nor": true, "on": true, "at": true, "to": true, "by": true,
	"in": true, "of": true, "up": true, "as": true, "is": true, "it": true,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" { continue }
		parts := strings.Fields(input)
		cmd := strings.ToLower(parts[0])
		text := strings.Join(parts[1:], " ")
		if cmd == "exit" { fmt.Println("Goodbye."); break }
		if text == "" { fmt.Printf("✗ No text provided. Usage: %s <text>\n", cmd); continue }
		switch cmd {
		case "upper": fmt.Println(strings.ToUpper(text))
		case "lower": fmt.Println(strings.ToLower(text))
		case "cap": fmt.Println(capWords(text))
		case "title": fmt.Println(titleCase(text))
		case "snake": fmt.Println(snakeCase(text))
		case "reverse": fmt.Println(reverseWords(text))
		default:
			fmt.Printf("✗ Unknown command: \"%s\"\nValid: upper, lower, cap, title, snake, reverse, exit\n", cmd)
		}
	}
}

func capWords(s string) string {
	ws := strings.Fields(strings.ToLower(s))
	for i,w := range ws { if w!=""{r:=[]rune(w); r[0]=unicode.ToUpper(r[0]); ws[i]=string(r)} }
	return strings.Join(ws," ")
}

func titleCase(s string) string {
	ws := strings.Fields(strings.ToLower(s))
	for i,w := range ws { if i==0||!smallWords[w]{r:=[]rune(w); r[0]=unicode.ToUpper(r[0]); ws[i]=string(r)} }
	return strings.Join(ws," ")
}

func snakeCase(s string) string {
	var b strings.Builder
	s = strings.ToLower(s)
	for _,r:=range s {
		if unicode.IsLetter(r)||unicode.IsDigit(r) { b.WriteRune(r) }
		if unicode.IsSpace(r)||r=='-' { b.WriteRune('_') }
	}
	return strings.ReplaceAll(b.String(),"__","_")
}

func reverseWords(s string) string {
	ws := strings.Fields(s)
	for i,w:=range ws{ r:=[]rune(w); for l,h:=0,len(r)-1;l<h;l,h=l+1,h-1{r[l],r[h]=r[h],r[l]} ws[i]=string(r) }
	return strings.Join(ws," ")
}