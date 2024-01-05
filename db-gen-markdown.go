package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Table struct {
	TableSchema  string
	TableName    string
	TableComment string
}

type Column struct {
	ColumnName    string
	ColumnType    string
	IsNullable    string
	ColumnDefault string
	ColumnKey     string
	Extra         string
	ColumnComment string
}

const headerPattern = `
<details>
<summary>生成コード</summary>

~~~sh
go run main.go --uri="username:password@tcp(hostname:port)/database_name"
~~~

main.go
~~~go
%s
~~~

</details>
`

var header = fmt.Sprintf(strings.Replace(headerPattern, "~~~", "~~~~~~", -1), readSourceFile())

func main() {
	// Flags
	var (
		uri string
	)
	flag.StringVar(&uri, "uri", "", "parameter description") // "username:password@tcp(hostname:port)/database_name"
	flag.Parse()

	// DB
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	content := "\n# テーブル一覧\n" + header

	tables := tables(db)

	for _, t := range tables {
		fullTableName := t.TableSchema + "." + t.TableName

		tableHeader := `
# %s
%s

| name | type | null | key | default | extra | comment |
| --- | --- | --- | --- | --- | --- | --- |
`
		content += fmt.Sprintf(tableHeader, fullTableName, t.TableComment)

		cols := columns(db, t.TableSchema, t.TableName)
		for _, col := range cols {
			cols := []string{
				col.ColumnName,
				col.ColumnType,
				col.IsNullable,
				col.ColumnKey,
				col.ColumnDefault,
				col.Extra,
				col.ColumnComment,
			}
			content += "|" + strings.Join(cols, " | ") + "|\n"
		}
	}

	filename := fmt.Sprintf("tables-%s.md", time.Now().Format("2006-01-02"))
	err = os.WriteFile(filename, []byte(content), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func tables(db *gorm.DB) []Table {
	var tables []Table

	err := db.Raw(`select table_schema, table_name, table_comment from information_schema.tables WHERE table_schema <> 'information_schema';`).Scan(&tables).Error
	if err != nil {
		panic(err)
	}

	return tables
}

func columns(db *gorm.DB, schema, table string) []Column {
	var columns []Column

	err := db.Raw(`select column_name, column_type, is_nullable, column_default, column_key, extra, column_comment from information_schema.columns WHERE table_schema = ? and table_name = ?;`, schema, table).Scan(&columns).Error
	if err != nil {
		panic(err)
	}

	return columns
}

func readSourceFile() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("no source file")
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(b)
}
