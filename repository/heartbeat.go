package repository

import (
	"github.com/exuan/waka-api/validator"
)

func (r *repository) BatchHeartbeat(vhs *[]*validator.Heartbeat) (int64, error) {
	tx := r.db.Create(vhs)
	return tx.RowsAffected, tx.Error
}
