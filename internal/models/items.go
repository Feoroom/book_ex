package models

import (
	"database/sql"
	"errors"
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

	stmt := `SELECT id, name, price FROM items
			WHERE id=$1`
	row := m.DB.QueryRow(stmt, id)

	item := &Item{}

	err := row.Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return item, nil
}

func (m *ItemModel) GetByName(name string) (*Item, error) {

	stmt := `SELECT id, name, price FROM items
			WHERE name=$1`

	row := m.DB.QueryRow(stmt, name)

	item := &Item{}
	err := row.Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return item, nil
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

	stmt := `SELECT id, name, price FROM items
			ORDER BY id`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []*Item{}

	for rows.Next() {

		item := &Item{}

		err = rows.Scan(&item.ID, &item.Name, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
