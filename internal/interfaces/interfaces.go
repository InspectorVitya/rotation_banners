package interfaces

import (
	"context"

	"github.com/inspectorvitya/rotation_banners/internal/models"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	Fatal(msg string)
	Warn(msg string)
}

type Storage interface {
	AddBanner(ctx context.Context, bannerID, slotID int) error
	DeleteBannerRotation(ctx context.Context, bannerID, slotID int) error
	IncrementViews(ctx context.Context, slotID, bannerID, userGroupID int) error
	IncrementClick(ctx context.Context, slotID, bannerID, userGroupID int) error
	SelectBanners(ctx context.Context, slotID, userGroupID int) ([]models.Statistic, error)
}
