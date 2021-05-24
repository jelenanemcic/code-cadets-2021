package sqlite

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

// BetRepository provides methods that operate on cals_bets SQLite database.
type BetRepository struct {
	dbExecutor DatabaseExecutor
	betMapper  BetMapper
}

// NewBetRepository creates and returns a new BetRepository.
func NewBetRepository(dbExecutor DatabaseExecutor, betMapper BetMapper) *BetRepository {
	return &BetRepository{
		dbExecutor: dbExecutor,
		betMapper:  betMapper,
	}
}

// InsertBetCalculated inserts the provided calculated bet into the database. An error is returned if the operation
// has failed.
func (r *BetRepository) InsertBetCalculated(ctx context.Context, bet domainmodels.BetCalculated) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryInsertBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "calc_bets repository failed to insert a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetRepository) queryInsertBet(ctx context.Context, bet storagemodels.BetCalculated) error {
	// If payout is 0, do not insert it.
	if bet.Payment == 0 {
		insertBetSQL := "INSERT INTO bets(id, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?)"
		statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
		if err != nil {
			return err
		}

		_, err = statement.ExecContext(ctx, bet.Id, bet.SelectionId, bet.SelectionCoefficient, bet.Payment)
		return err
	}

	insertBetSQL := "INSERT INTO bets(id, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?)"
	statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.Id, bet.SelectionId, bet.SelectionCoefficient, bet.Payment)
	return err
}

// UpdateBetCalculated updates the provided calculated bet in the database. An error is returned if the operation
// has failed.
func (r *BetRepository) UpdateBetCalculated(ctx context.Context, bet domainmodels.BetCalculated) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryUpdateBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "calc_bets repository failed to update a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetRepository) queryUpdateBet(ctx context.Context, bet storagemodels.BetCalculated) error {
	updateBetSQL := "UPDATE bets SET selection_id=?, selection_coefficient=?, payment=? WHERE id=?"

	statement, err := r.dbExecutor.PrepareContext(ctx, updateBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.Payment, bet.Id)
	return err
}

// GetBetCalculatedByID fetches a calculated bet from the database and returns it. The second returned value indicates
// whether the bet exists in DB. If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBetCalculatedByID(ctx context.Context, id string) (domainmodels.BetCalculated, bool, error) {
	storageBet, err := r.queryGetBetByID(ctx, id)
	if err == sql.ErrNoRows {
		return domainmodels.BetCalculated{}, false, nil
	}
	if err != nil {
		return domainmodels.BetCalculated{}, false, errors.Wrap(err, "calc_bets repository failed to get a bet with id "+id)
	}

	domainBet := r.betMapper.MapStorageBetToDomainBet(storageBet)
	return domainBet, true, nil
}

func (r *BetRepository) queryGetBetByID(ctx context.Context, id string) (storagemodels.BetCalculated, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.BetCalculated{}, err
	}
	defer row.Close()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	row.Next()

	var selectionId string
	var selectionCoefficient int
	var paymentSql sql.NullInt64

	err = row.Scan(&id, &selectionId, &selectionCoefficient, &paymentSql)
	if err != nil {
		return storagemodels.BetCalculated{}, err
	}

	var payment int
	if paymentSql.Valid {
		payment = int(paymentSql.Int64)
	}

	return storagemodels.BetCalculated{
		Id:                   id,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
	}, nil
}

// GetBetCalculatedBySelectionID fetches all calculated bets from the database that have provided selectionId
// and returns them. The second returned value indicates whether the bet exists in DB.
// If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBetCalculatedBySelectionID(ctx context.Context, selectionId string) ([]domainmodels.BetCalculated, bool, error) {
	storageBets, err := r.queryGetBetBySelectionID(ctx, selectionId)
	if err == sql.ErrNoRows {
		return []domainmodels.BetCalculated{}, false, nil
	}
	if err != nil {
		return []domainmodels.BetCalculated{}, false, errors.Wrap(err, "calc_bets repository failed to get a bet with selection id "+selectionId)
	}

	var domainBets []domainmodels.BetCalculated
	for _, storageBet := range storageBets {
		domainBets = append(domainBets, r.betMapper.MapStorageBetToDomainBet(storageBet))
	}

	return domainBets, true, nil
}

func (r *BetRepository) queryGetBetBySelectionID(ctx context.Context, selectionId string) ([]storagemodels.BetCalculated, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE selection_id='"+selectionId+"';")
	if err != nil {
		return []storagemodels.BetCalculated{}, err
	}
	defer row.Close()

	var betsCalculated []storagemodels.BetCalculated

	// A loop over all returned rows.
	for row.Next() {

		var id string
		var selectionCoefficient int
		var paymentSql sql.NullInt64

		err = row.Scan(&id, &selectionId, &selectionCoefficient, &paymentSql)
		if err != nil {
			return []storagemodels.BetCalculated{}, err
		}

		var payment int
		if paymentSql.Valid {
			payment = int(paymentSql.Int64)
		}

		betsCalculated = append(betsCalculated, storagemodels.BetCalculated{
			Id:                   id,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
		})
	}

	return betsCalculated, nil
}
