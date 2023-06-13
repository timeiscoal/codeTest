package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	a1 := now.AddDate(0, 0, 0)
	a2 := now.AddDate(0, 0, 1)
	a3 := now.AddDate(0, 1, 0)
	a4 := now.AddDate(1, 0, 0)

	fmt.Printf("기준 : %v\n 하루 증가 : %v\n 한달 증가 : %v\n 일년 증가 %v\n", a1, a2, a3, a4)
	// 년월일, 시간 분 초,나노세컨드,지역
	fmt.Println(time.Date(a1.Year(), a1.Month(), a1.Day(), 0, 0, 0, 0, time.Local))

	
}
