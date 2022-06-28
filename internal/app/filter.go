package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Filter(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(str, "server_name") && strings.Contains(str, ".") {
			strSlice := strings.Fields(str)
			str = strings.Trim(strSlice[1], ";")
			fmt.Println(str)
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("scan error: ", err)
		return
	}
}
