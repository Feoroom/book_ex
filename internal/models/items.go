package models

import (
	"database/sql"
)

type Item struct {
	ID    int
	Name  string
	Price int
}

type ItemModel struct {
	DB *sql.DB
}

func (m *ItemModel) Get(id int) (*Item, error) {
	return nil, nil
}

func (m *ItemModel) Insert(name string, price int) (int, error) {
	stmt := `insert into items (name, price)
			values ($1, $2) RETURNING id`
	id := 0
	err := m.DB.QueryRow(stmt, name, price).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ItemModel) GetAll() ([]*Item, error) {
	return nil, nil
}
