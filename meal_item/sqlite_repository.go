package mealitem

import (
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
)

type Repository interface {
	Migrate() error
	Create(mealitem MealItem) (*MealItem, error)
	All() ([]MealItem, error)
	GetByName(name string) (*MealItem, error)
	Update(id int64, updated MealItem) (*MealItem, error)
	Delete(id int64) error
}

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS mealitems(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        calories INTEGER NOT NULL,
        carbs INTEGER NOT NULL,
        fat INTEGER NOT NULL,
        protein INTEGER NOT NULL
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(mealitem MealItem) (*MealItem, error) {
	res, err := r.db.Exec("INSERT INTO mealitems(name, calories, carbs, fat, protein) values(?,?,?,?,?)",
		mealitem.Name, mealitem.Calories, mealitem.Carbs, mealitem.Fat, mealitem.Protein)

	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	mealitem.ID = id

	return &mealitem, nil
}

func (r *SQLiteRepository) All() ([]MealItem, error) {
	rows, err := r.db.Query("SELECT * FROM mealitems")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []MealItem
	for rows.Next() {
		var item MealItem
		if err := rows.Scan(&item.ID, &item.Name, &item.Calories, &item.Carbs,
			&item.Fat, &item.Protein); err != nil {
			return nil, err
		}
		all = append(all, item)
	}

	return all, nil
}

func (r *SQLiteRepository) GetByName(name string) (*MealItem, error) {
	row := r.db.QueryRow("SELECT * FROM mealitems WHERE name = ?", name)

	var item MealItem
	if err := row.Scan(&item.ID, &item.Name, &item.Calories, &item.Carbs,
		&item.Fat, &item.Protein); err != nil {
		return nil, ErrNotExists
	}
	return &item, nil
}

func (r *SQLiteRepository) Update(id int64, updated MealItem) (*MealItem, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}

	res, err := r.db.Exec(`UPDATE mealitems SET name = ?, calories = ?, carbs =
    ?, fat = ?, protein = ? WHERE id = ?`, updated.Name, updated.Calories,
		updated.Carbs, updated.Fat, updated.Protein, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *SQLiteRepository) Delete(id int64) error {
	res, err := r.db.Exec(`DELETE FROM mealitems WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return nil
}
