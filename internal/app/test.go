package app

import (
	"fmt"
	"log"
	"os"
)

func File() {
	files, err := os.ReadDir("/etc/letsencrypt/live")
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i].Name())
	}
}
