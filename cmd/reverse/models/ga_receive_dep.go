package models

type GaReceiveDep struct {
	Id    string `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	Title string `xorm:"not null default '' comment('名称') unique VARCHAR(50)"`
}

func (m *GaReceiveDep) TableName() string {
	return "ga_receive_dep"
}

func (m *GaReceiveDep) TableComment() string {
	return "ga_receive_dep"
}
