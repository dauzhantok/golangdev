package postgres
import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gomidka/internal/models"
	"gomidka/internal/store"
)

func (db DB) Breads() store.BreadsRepository {
	if db.breads == nil {
		db.breads = NewBreadsRepository(db.conn)
	}

	return db.breads
}

type BreadsRepository struct {
	conn *sqlx.DB
}

func NewBreadsRepository(conn *sqlx.DB) store.BreadsRepository {
	return &BreadsRepository{conn: conn}
}

func (c BreadsRepository) Create(ctx context.Context, bread *models.Bread) error {
	_, err := c.conn.Exec("INSERT INTO bread(name) VALUES ($1)", bread.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c BreadsRepository) All(ctx context.Context, filter *models.BreadsFilter) ([]*models.Bread, error) {
	breads := make([]*models.Bread, 0)
	basicQuery := "SELECT * FROM bread"

	if filter.Query != nil {
		basicQuery = fmt.Sprintf("%s WHERE name ILIKE $1", basicQuery)

		if err := c.conn.Select(&breads, basicQuery, "%"+*filter.Query+"%"); err != nil {
			return nil, err
		}

		return breads, nil
	}

	if err := c.conn.Select(&breads, basicQuery); err != nil {
		return nil, err
	}

	return breads, nil
}

func (c BreadsRepository) ByID(ctx context.Context, id int) (*models.Bread, error) {
	bread := new(models.Bread)
	if err := c.conn.Get(bread, "SELECT id, name FROM bread WHERE id=$1", id); err != nil {
		return nil, err
	}

	return bread, nil
}

func (c BreadsRepository) Update(ctx context.Context, bread *models.Bread) error {
	_, err := c.conn.Exec("UPDATE bread SET name = $1 WHERE id = $2", bread.Name, bread.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c BreadsRepository) Delete(ctx context.Context, id int) error {
	_, err := c.conn.Exec("DELETE FROM bread WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
