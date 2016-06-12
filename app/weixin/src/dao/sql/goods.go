package sql

type Goods struct {
	Id           int64  `gorm:"primary_key"`
	Name         string `json:"-"`
	StoreId      int64
	StoreName    string
	Unit         string
	MinOrdernum  int64
	MaxOrdernum  int64
	Marketprice  int64
	Price        int64
	Costprice    int64
	Salenum      int64
	MonthSalenum int64
	WeekSalenum  int64
	State        int8
	Image        string
	Desc         string
	StcId        int64
	Sort         int64
	Barcode      string
	Storage      int32
	Addtime      int64
	Edittime     int64  `json:"-"`
	OpUser       string `json:"-"`
	Mtime        string `json:"-"`
}

func (this *Goods) TableName() string {
	return "goods"
}
