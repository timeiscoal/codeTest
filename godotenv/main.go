package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Env 파일을 읽어서 나오는 내용들을 출력한다.
	mapData, err := godotenv.Read(".env")
	if err != nil {
		fmt.Println("Not found .env file or Can not Read file : ", err)
	}
	fmt.Println(mapData)

	fmt.Println(godotenv.Load(".env")) // nil
	errors := godotenv.Load(".env")
	if errors != nil {
		fmt.Println("Not found .env file : ", errors)
	}

	key := "CHECK"
	v := os.Getenv(key)
	if len(v) <= 0 {
		fmt.Printf("NOT FOUND env %s \n ", key)
	}
	fmt.Println(v) // YES
}
