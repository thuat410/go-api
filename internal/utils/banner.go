package utils

import (
	"fmt"
	"log"
	"runtime"
)

func PrintSuccessBanner(port string) {
	cyan := "\033[36m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	fmt.Println(cyan + `
   ______   ____       ___     ____     ____
  / ____/  / __ \     /   |   / __ \   /  _/
 / / __   / / / /    / /| |  / /_/ /   / /
/ /_/ /  / /_/ /    / ___ | / ____/  _/ /
\____/   \____/    /_/  |_|/_/      /___/
                                             ` + reset)

	fmt.Println(green + "==================================================================" + reset)
	log.Printf("%s🚀 API Gateway Server stated SUCCESSFULLY!%s\n", green, reset)
	log.Printf("💻 Enviroment : %s\n", runtime.GOOS+" ("+runtime.GOARCH+")")
	log.Printf("🐹 Go version : %s\n", runtime.Version())
	log.Printf("📡 Base url   : %shttp://localhost:%s%s\n", yellow, port, reset)
	fmt.Println(green + "==================================================================" + reset)
}
