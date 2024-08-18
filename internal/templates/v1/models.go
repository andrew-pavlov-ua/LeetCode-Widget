package v1

const (
	EasyMaxValue   int64 = 820
	MediumMaxValue int64 = 1710
	HardMaxValue   int64 = 732
	BarWidthValue  int64 = 240
)

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
