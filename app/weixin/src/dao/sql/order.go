package sql

type Order struct {
	OrderId         int64 `gorm:"primary_key"`
	BuyerId         int64 `json:"-"`
	UserOrderTimes  int64 `json:"-"`
	AddressId       int64 `json:"-"`
	OrderTime       int64
	ExpectTime      int64
	ConfirmTime     int64
	FinishedTime    int64
	PaySn           int64 `json:"-"`
	PaymentTime     int64
	CostAmount      int
	Amount          int
	CostOrderAmount int
	OrderAmount     int
	RefundAmount    int64
	OrderState      int16
	CancelFlag      int8
	Ext             string `json:"-"`
	ExtInfo         string `gorm:"-"`
}
