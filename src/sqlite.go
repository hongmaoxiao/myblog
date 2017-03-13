package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func openDB(dsn string) (db *sql.DB, err error) {
	if !strings.Contains(dsn, "?") {
		dsn = dsn + "?"
	}
	if !strings.Contains(dsn, "parseTime=") {
		dsn = dsn + "&parseTime=true"
	}
	if !strings.Contains(dsn, "loc=") {
		dsn = dsn + "&loc=auto"
	}

	return sql.Open("sqlite3", dsn)
}

type columndef struct {
	Name, Type string
	NotNull    bool
	Dflt_value *string
	PK         bool
}

var (
	id_def         = columndef{Name: "id", Type: "INTEGER", PK: true}
	updated_at_def = columndef{Name: "updated_at", Type: "TIMESTAMP"}
	created_at_def = columndef{Name: "created_at", Type: "TIMESTAMP"}
)

func (p *columndef) sameas(t *columndef) bool {
	if p.Dflt_value != nil && t.Dflt_value != nil {
		if *p.Dflt_value != *t.Dflt_value {
			return false
		}
	} else if !(p.Dflt_value == nil && t.Dflt_value == nil) {
		return false
	}

	return p.Type == t.Type && p.NotNull == t.NotNull && p.PK == t.PK
}

func (p *columndef) to_add_column_sql(table string) string {
	s := `ALTER TABLE ` + table + ` ADD COLUMN `
	s += "`" + p.Name + "` " + p.Type
	if p.NotNull {
		s += " NOT NULL"
	}
	if p.Dflt_value != nil {
		s += " DEFAULT " + *p.Dflt_value
	}
	return s
}

func to_create_table_sql(table string, columns []columndef) string {
	var pks []string
	for _, cdef := range columns {
		if cdef.PK {
			pks = append(pks, cdef.Name)
		}
	}

	s := `CREATE TABLE IF NOT EXISTS ` + table + `(`
	for _, cdef := range columns {
		s += "`" + cdef.Name + "` " + cdef.Type
		if cdef.NotNull {
			s += " NOT NULL"
		}
		if cdef.Dflt_value != nil {
			s += " DEFAULT " + *cdef.Dflt_value
		}
		if cdef.PK && len(pks) == 1 {
			s += " PRIMARY KEY"
		}
		s += ","
	}

	if len(pks) > 1 {
		s += "PRIMARY KEY(" + strings.Join(pks, ", ") + "),"
	}
	s = strings.TrimSuffix(s, ",")

	s += `);`
	return s
}

func check_schema(db *sql.DB, table string, columns []columndef) {
	s := to_create_table_sql(table, columns)
	_, err := db.Exec(s)
	if err != nil {
		log.Printf("%q: %s\n", err, s)
		return
	}

	rows, err := db.Query("PRAGMA table_info(" + table + ")")
	if err != nil {
		log.Println(err)
		return
	}

	var onbeach []columndef
	for rows.Next() {
		var t columndef
		i := 0
		if err := rows.Scan(&i, &t.Name, &t.Type, &t.NotNull, &t.Dflt_value, &t.PK); err != nil {
			log.Println(err)
			continue
		}
		onbeach = append(onbeach, t)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	rows.Close()

	var change []columndef
	for _, t := range columns {
		ismiss := true
		ischange := false
		for _, cdef := range onbeach {
			if cdef.Name == t.Name {
				ismiss = false
				ischange = !cdef.sameas(&t)
				break
			}
		}
		if ismiss {
			change = append(change, t)
			log.Println("miss", t.Name, t.Type, t.NotNull, t.Dflt_value, t.PK)
		} else if ischange {
			change = append(change, t)
			log.Println("change", t.Name, t.Type, t.NotNull, t.Dflt_value, t.PK)
		}
	}

	for _, cdef := range change {
		s := cdef.to_add_column_sql(table)
		_, err := db.Exec(s)
		log.Printf("%q: %s\n", err, s)
	}
}
