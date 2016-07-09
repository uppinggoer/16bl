package sql

type OrderGoods struct {
	Id               int64 `gorm:"primary_key" json:"-"`
	OrderId          int64 `json:"-"`
	BuyerId          int64 `json:"-"`
	GoodsId          int64
	GoodsName        string
	GoodsNorms       string
	GoodsImage       string
	GoodsUnit        string
	GoodsNum         int64
	GoodsPrice       int
	GoodsMarketprice int
	GoodsCostprice   int `json:"-"`
}
