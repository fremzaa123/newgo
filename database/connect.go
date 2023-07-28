package connect

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "dbfin1.api6699.com"
	dbPort     = 5000
	dbUser     = "postgres89"
	dbPassword = "8!B8B8458dd544@$1356278639f7A0@9"
	dbName     = "lot368_db"
)

func ConnectDB() (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	// ทดสอบเชื่อมต่อฐานข้อมูล
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("เชื่อมต่อฐานข้อมูลสำเร็จ!")
	return db, nil
}

// // GetUsers คืนค่าข้อมูลผู้ใช้ทั้งหมด
// func GetUsers(db *sql.DB) ([]string, error) {
// 	rows, err := db.Query("SELECT * FROM bom_tb_agent")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []string
// 	for rows.Next() {
// 		var name string
// 		err := rows.Scan(&name)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, name)
// 	}

// 	return users, nil
// }
