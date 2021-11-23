package ucb

import (
	"github.com/inspectorvitya/rotation_banners/internal/models"
	"math"
)

func SelectBanner(stats []models.Statistic) int {
	var totalShows int

	for _, s := range stats {
		if s.Shows == 0 {
			return s.BannerID
		}

		totalShows += s.Shows
	}

	var (
		maxWeight float64
		bannerID  int
	)

	for _, s := range stats {
		weight := float64(s.Clicks)/float64(s.Shows) + math.Sqrt(2*math.Log(float64(totalShows))/float64(s.Shows))
		if weight >= maxWeight {
			maxWeight = weight
			bannerID = s.BannerID
		}
	}

	return bannerID
}
