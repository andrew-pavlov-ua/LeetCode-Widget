package v1

type LcStats struct {
	Username    string
	Rank        int64
	Lvl         int64
	Experience  int64
	EasyCount   int64
	MediumCount int64
	HardCount   int64
	TotalCount  int64
}

type BarsWidth struct {
	EasyWidth   float64
	MediumWidth float64
	HardWidth   float64
}

type maxStats struct {
	EasyMax   int64
	MediumMax int64
	HardMax   int64
	BarWidth  int64
}

func NewMaxStats() maxStats {
	return maxStats{EasyMax: 820, MediumMax: 1710, HardMax: 732, BarWidth: 240}
}
