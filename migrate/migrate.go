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
	// Steps looks at the currently active migration version. It will migrate up if n > 0, and down if n < 0.
	//m.Steps(2)

	// Down looks at the currently active migration version and will migrate all the way down (applying all down migrations).
	//m.Down()


	// Migrate all the way up ...
	if err := m.Up(); err != nil {
	    log.Fatal(err)
	}

	// Force 强制迁移到某个版本 并设置 dirty 为 false
	//m.Force(1512019560)

	// Migrate 根据当前版本，向前或向后迁移到某个指定的版本
	//m.Migrate(1512090940)

	// Drop 删除数据库中的所有东西
	//m.Drop()

	fmt.Println("迁移成功...")
}