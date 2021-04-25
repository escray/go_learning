// Week 02 Homework
// 在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
// 是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 我觉的可以考虑 Wrap 这个 error，代码如下：

package homework_week02

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type Person struct {
	ID int
	Name string
	Age int
}

func (p *Person) Execute(db *sql.DB, query string, args ...interface{})	error {
		_, err := db.Exec(query, args...)
		if err != nil {

			return errors.Wrap(err, fmt.Sprintf("exec: '%s' %v failed", query, args))
		}
	return nil
}

func (p *Person) Query(db *sql.DB, query string, args ...interface{}) ([]Person, error) {
	persons := make([]Person, 0, 10)
	rows, err := db.Query(query, args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return persons, nil
		}
		return nil, errors.Wrap(err, fmt.Sprintf("exec: '%s' %+v failed", query, args))
	}

	defer rows.Close()

	for rows.Next() {
		var p Person
		err := rows.Scan(&p.ID, &p.Name, &p.Age)
		if err != nil {
			return nil, errors.Wrap(err, "rows.Scan: ")
		}
		persons = append(persons, p)
	}
	return persons, nil
}
