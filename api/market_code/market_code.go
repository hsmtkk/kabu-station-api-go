package market_code

type MarketCode string

const (
	Tokyo    MarketCode = "1"  //東証
	WholeDay            = "2"  // 日通し
	Day                 = "23" // 日中
	Night               = "24" // 夜間
)
