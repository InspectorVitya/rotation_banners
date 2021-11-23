package ucb

import (
	"testing"

	"github.com/inspectorvitya/rotation_banners/internal/models"
)

func TestSelectBanner(t *testing.T) {
	t.Run("select banner with 0 shows", func(t *testing.T) {
		var expectedBannerID = 2
		stats := []models.Statistic{
			{
				BannerID: 1,
				SlotID:   1,
				GroupID:  1,
				Clicks:   10,
				Shows:    10,
			},
			{
				BannerID: expectedBannerID,
				SlotID:   1,
				GroupID:  0,
				Clicks:   0,
				Shows:    0,
			},
		}

		actualBannerID := SelectBanner(stats)
		require.Equal(t, expectedBannerID, actualBannerID)
	})

	t.Run("select banner with max weight", func(t *testing.T) {
		var expectedBannerID = 2
		stats := []models.Statistic{
			{
				BannerID: 1,
				SlotID:   1,
				GroupID:  1,
				Clicks:   19,
				Shows:    45,
			},
			{
				BannerID: expectedBannerID,
				SlotID:   1,
				GroupID:  1,
				Clicks:   30,
				Shows:    40,
			},
			{
				BannerID: 3,
				SlotID:   1,
				GroupID:  1,
				Clicks:   7,
				Shows:    50,
			},
		}

		actualBannerID := SelectBanner(stats)
		require.Equal(t, expectedBannerID, actualBannerID)
	})
}
