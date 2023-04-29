package mysql

import (
	"context"
	"database/sql"
	"project/src/gen"
	"project/src/pkg/utils"
	"project/src/services/financial/configs"
)

type Repository struct {
	Db *sql.DB
}

func New() (*Repository, error) {
	cfg, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.Conn)
	if err != nil {
		return nil, err
	}
	return &Repository{Db: db}, nil
}

func (r *Repository) InsertInvoice(ctx context.Context, params *gen.AddInvoiceRequest) (*gen.BasicResponse, error) {
	tx, err := r.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	newUUID, err := utils.GenerateNewUUid()
	if err != nil {
		return nil, err
	}
	rq, err := tx.ExecContext(
		ctx,
		`insert into invoices (uuid, value, client_id, due_date) values (?,?,?,?)`,
		newUUID,
		params.Value,
		params.ClientId,
		params.DueDate)

	if err != nil {
		return nil, err
	}
	lastInsertedId, err := rq.LastInsertId()
	if err != nil {
		return nil, err
	}
	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &gen.BasicResponse{
		ReferenceId: int32(lastInsertedId),
		Success:     true,
		Message:     "",
	}, nil
}
