package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	userText, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	userText = strings.ToLower(userText)
	words := strings.Fields(userText)
	byte_str := []byte(userText)

	for _, x := range words {
		if strings.HasPrefix(x, "http://") {
			n := strings.Index(string(byte_str), "http://") + utf8.RuneCountInString("http://") -1
			m := strings.Index(string(byte_str), "http://") + utf8.RuneCountInString(x) - 1
			for i := n; i <= m; i++ {
				byte_str[i] = 42
			}
		}
	}
	fmt.Println(string(byte_str))
}

// Here's my spammy page: hTTp://hehefouls.netHAHAHA see you hTTp://hehefouls.netHAHAHA hTTp://hehefouls.netHAHAHA
