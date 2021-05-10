package model

import (
	"database/sql/driver"

	"gopkg.in/guregu/null.v4"
)

type Int64 struct {
	null.Int
}

type String struct {
	null.String
}

func (n Int64) Value() (driver.Value, error) {
	if !n.Valid {
		return 0, nil
	}
	return n.Int64, nil
}

func (n String) Value() (driver.Value, error) {
	if !n.Valid {
		return ``, nil
	}
	return n.String, nil
}
