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
        protein INTEGER NOT NULL,
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
