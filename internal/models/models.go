package models

type (
	Statistic struct {
		BannerID int `db:"banner_id"`
		SlotID   int `db:"slot_id"`
		Clicks   int `db:"count_click"`
		Shows    int `db:"count_show"`
		GroupID  int
	}
)
