package sqlstorage

import (
	"context"
	"fmt"

	"github.com/inspectorvitya/rotation_banners/internal/interfaces"
	"github.com/inspectorvitya/rotation_banners/internal/models"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	isExistErrCode = 23505
)

//var ErrBannerinSlotIsNotExist = errors.New("in slot there is no banners")

type StorageInDB struct {
	db     *sqlx.DB
	logger interfaces.Logger
}

func New(ctx context.Context, dbURL string, logger interfaces.Logger) (*StorageInDB, error) {
	conn, err := sqlx.ConnectContext(ctx, "pgx", dbURL)

	if err != nil {
		return nil, err
	}
	s := &StorageInDB{
		db:     conn,
		logger: logger,
	}

	return s, nil
}

func (s *StorageInDB) Close(ctx context.Context) {
	s.logger.Info("connection to DB is closing gracefully...")
	if err := s.db.Close(); err != nil {
		s.logger.Error("failed to stop DB: " + err.Error())
	} else {
		s.logger.Info("connection to DB is closed gracefully...")
	}
}

func (s *StorageInDB) AddBanner(ctx context.Context, bannerID, slotID int) error {
	//добавить транзакцию

	bannerExist, err := s.isExistInBannerInSlotsCheck(slotID, bannerID)
	if err != nil {
		return fmt.Errorf("add banner storage: %w", err)
	}
	if !bannerExist {
		_, err := s.db.ExecContext(ctx, "INSERT INTO rotation (banner_id, slot_id) VALUES ($1, $2)",
			bannerID,
			slotID,
		)
		if err != nil {
			return fmt.Errorf("add banner storage: %w", err)
		}
	}

	return nil
}

func (s *StorageInDB) DeleteBannerRotation(ctx context.Context, bannerID, slotID int) error {
	_, err := s.db.ExecContext(ctx, "DELETE from rotation where banner_id = $1 and slot_id = $2",
		bannerID,
		slotID,
	)
	if err != nil {
		return fmt.Errorf("delete banner storage: %w", err)
	}
	return nil
}

func (s *StorageInDB) SelectBanners(ctx context.Context, slotID, groupID int) ([]models.Statistic, error) {
	query := `select r.banner_id, r.slot_id, s.count_show, s.count_click
	from rotation r
	left join "statistics" s on r.banner_id = s.banner_id and r.slot_id = s.slot_id and r.slot_id = $1
	where  s.group_id = $2 and r.status = false`

	var stat []models.Statistic
	err := s.db.Select(&stat, query, slotID, groupID)
	if err != nil {
		return nil, fmt.Errorf("delete banner storage: %w", err)
	}

	return stat, nil
}

func (s *StorageInDB) IncrementViews(ctx context.Context, slotID, bannerID, userGroupID int) error {
	query := `INSERT INTO "statistics" (slot_id, banner_id, group_id, count_show)
	VALUES ($1, $2, $3)
	ON CONFLICT (banner_id, slot_id, group_id) DO UPDATE
	SET count_show = "statistics".count_show + 1;`

	_, err := s.db.ExecContext(ctx, query, slotID, bannerID, userGroupID)
	if err != nil {
		return fmt.Errorf("delete banner storage: %w", err)
	}
	return nil
}

func (s *StorageInDB) IncrementClick(ctx context.Context, slotID, bannerID, userGroupID int) error {
	query := `update "statistics" set count_click = count_click +1 where banner_id = $1 and slot_id = $2 and group_id = $3;`
	_, err := s.db.ExecContext(ctx, query, slotID, bannerID, userGroupID)
	if err != nil {
		return fmt.Errorf("delete banner storage: %w", err)
	}
	return nil
}

func (s *StorageInDB) isExistInBannerInSlotsCheck(slotID, bannerID int) (bool, error) {
	query := `select exists(select FROM rotation WHERE slot_id = $1 and banner_id = $2)`
	var countRow bool
	err := s.db.QueryRowx(query, slotID, bannerID).Scan(&countRow)

	if err != nil {
		return false, fmt.Errorf("delete banner storage: %w", err)
	}

	return countRow, nil
}
