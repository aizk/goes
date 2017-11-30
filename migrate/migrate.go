package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate/database/mysql"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/source/file"
	"log"
	"fmt"
)
// 数据库可能会冲突，修改 github.com/go-sql-driver/mysql <init> 函数中的数据连接名
func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goes?multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		// file:///absolute/path
		// file://relative/path
		"file://./",
		// 数据库名
		"mysql",
		driver,
	)
	defer m.Close()
	if err != nil {
		log.Fatal(err)
	}

	version, dirty, err := m.Version()

	fmt.Println("Now Version:", version, dirty)
	//m.Steps(2)

	//m.Down()
	// Migrate all the way up ...
	if err := m.Up(); err != nil {
	    log.Fatal(err)
	}

	fmt.Println("迁移成功...")
}