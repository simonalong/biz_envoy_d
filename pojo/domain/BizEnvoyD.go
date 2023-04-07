package domain

type BizEnvoyD struct {
	Id    int64
	Times string
}

func (BizEnvoyD) TableName() string {
	return "biz-d-service"
}
