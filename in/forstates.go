package main

import "fmt"

func main() {
	elefantes := 11

	for number := 1; number <= elefantes; number++ {
		complementElefant := ""
		complementBalance := ""
		if number > 1 {
			complementElefant = "s"
			complementBalance = "n"
		}
		fmt.Println(
			number,
			"elefante"+complementElefant+
				" se balancea"+complementBalance+
				" sobre la tela de una araña")
	}
}
