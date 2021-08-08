package sample

import (
	"database/sql"
	"fmt"
	"log"
)

func Sample() {

	db, err := sql.Open("mysql", "")

	if err != nil {
		panic(err)
	}

	defer db.Close() //defer은 이 함수가 모두 끝나고 나서 실행됨, 중간에 적더라도 무조건 함수 끝나구 실행, 까먹을까봐 중간에 적는거임

	var name string

	rows, err := db.Query("SELECT domain FROM setting_domain LIMIT 15")
	if err != nil {
		log.Fatal(err) //log타입 panic타입 조금 다름
	}

	defer rows.Close() //DB랑 rows는 꼭 닫기 꼭꼭꼭

	for rows.Next() {
		err := rows.Scan(&name) // 반환되는 값이 다 에러.. 다 에러..래..
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("rows", name)
	}

}
