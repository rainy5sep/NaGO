package mysqltest

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlTest() {

	db, err := sql.Open("mysql", "root:rladndwo3@tcp(121.166.4.186:3152)/study?charset=utf8")

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success")
	defer db.Close() //main함수가 끝나면 db를 닫아라

	var name string
	var price int
	var flavor string
	var star int
	var seq int

	//rows, err := db.Query("INSERT INTO tacobell(name, price, flavor, star) VALUES ('타코 샐러드', 5900, 'Chicken', 4),('김치 치즈 후라이', 7100, 'kimchi', '4');")
	rows, err := db.Query("SELECT * FROM tacobell")
	//rows, err := db.Query("UPDATE tacobell SET price = 6200 WHERE seq = 1")
	//rows, err := db.Query("INSERT INTO tacobell(name, price, flavor, star) VALUES ('나쵸', 2700, 'Cheese', 4);")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&seq, &name, &price, &flavor, &star)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("rows", seq, name, price, flavor, star)
	}
}
