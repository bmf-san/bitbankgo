package types

// TypeCandle is a type for candle.
type TypeCandle string

const (
	// OneMin is a candle type for 1min.
	OneMin TypeCandle = "1min"
	// FiveMin is a candle type for 5min.
	FiveMin TypeCandle = "5min"
	// FifteenMin is a candle type for 15min.
	FifteenMin TypeCandle = "15min"
	// ThirtyMin is a candle type for 30min.
	ThirtyMin TypeCandle = "30min"
	// OneHour is a candle type for 1hour.
	OneHour TypeCandle = "1hour"
	// FourHour is a candle type for 4hour.
	FourHour TypeCandle = "4hour"
	// EightHour is a candle type for 8hour.
	EightHour TypeCandle = "8hour"
	// TweleveHour is a candle type for 12hour.
	TweleveHour TypeCandle = "12hour"
	// OneDay is a candle type for 1day.
	OneDay TypeCandle = "1day"
	// OneWeek is a candle type for 1week.
	OneWeek TypeCandle = "1week"
)
