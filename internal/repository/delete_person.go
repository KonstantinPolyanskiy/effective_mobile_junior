package repository

import (
	"context"
	"go.uber.org/zap"
)

func (r Repository) SoftDeletePerson(id int) (bool, error) {
	var isDeleted bool

	deletePersonQuery := `
	UPDATE Person SET is_deleted = true 
	WHERE person_id=$1
	RETURNING is_deleted
`

	r.log.Debug("delete query",
		zap.String("query", deletePersonQuery),
	)

	err := r.db.QueryRow(context.Background(), deletePersonQuery, id).Scan(&isDeleted)
	if err != nil {
		r.log.Debug("error soft delete person",
			zap.Int("id person", id),
			zap.String("error", err.Error()),
		)

		return false, err
	}

	return isDeleted, nil
}
