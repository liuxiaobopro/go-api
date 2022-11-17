package model

// Demo demo表
type Demo struct {
	Id            int    `xorm:"pk autoincr not null int(11) 'id'" json:"id"`
	Title         string `xorm:"not null comment('名称') varchar(200) 'title'" json:"title"`
	Logo          string `xorm:"default '' comment('logo') varchar(300) 'logo'" json:"logo" binding:"required"`
	Pic           string `xorm:"default '' comment('店内照片') varchar(900) 'pic'" json:"pic"`
	ProvCode      int    `xorm:"not null comment('省份id') int(11) 'prov_code'" json:"provCode"`
	CityCode      int    `xorm:"not null comment('城市id') int(11) 'city_code'" json:"cityCode"`
	Address       string `xorm:"not null comment('详细地址') varchar(200) 'address'" json:"address"`
	Notice        string `xorm:"comment('公告') text 'notice'" json:"notice"`
	HeadName      string `xorm:"not null comment('负责人名称') varchar(10) 'head_name'" json:"headName"`
	HeadPhone     string `xorm:"not null comment('负责人手机号') char(11) 'head_phone'" json:"headPhone"`
	StaffName     string `xorm:"default '' comment('员工名称') varchar(10) 'staff_name'" json:"staffName"`
	StaffPhone    string `xorm:"default '' comment('员工手机号') char(11) 'staff_phone'" json:"staffPhone"`
	IsStart       int    `xorm:"default 1 comment('是否营业(正常1  关闭2)') tinyint(1) 'is_start'" json:"isStart"`
	WorkTime      int    `xorm:"default 1 comment('营业时间(早1  中2  晚3)') tinyint(1) 'work_time'" json:"workTime"`
	OrderPhone    string `xorm:"default '' comment('接单电话') char(11) 'order_phone'" json:"orderPhone"`
	DisType       int    `xorm:"default 1 comment('配送费类型(免费1  份2  单3)') tinyint(1) 'dis_type'" json:"disType"`
	DisPrice      int    `xorm:"default 0 comment('配送费金额') int(11) 'dis_price'" json:"disPrice"`
	MealboxPrice  int    `xorm:"default 0 comment('餐盒费') int(11) 'mealbox_price'" json:"mealboxPrice"`
	ReserveStatus int    `xorm:"default 1 comment('预定功能(开启1  关闭2)') tinyint(1) 'reserve_status'" json:"reserveStatus"`
	IsFrozen      int    `xorm:"default 1 comment('是否冻结账户(开启1  冻结2)') tinyint(1) 'is_frozen'" json:"isFrozen"`
	Commission    string `xorm:"comment('抽佣规则') text 'commission'" json:"commission"`
	CreateTime    int    `xorm:"default NULL int(10) 'create_time' created" json:"createTime"`
	UpdateTime    int    `xorm:"default NULL int(10) 'update_time' updated" json:"updateTime"`
	DeleteTime    int    `xorm:"default NULL int(10) 'delete_time' deleted" json:"deleteTime"`
	UpdateUid     int    `xorm:"default 0 int(11) 'update_uid'" json:"updateUid"`
}

type DemoModel struct {
	Name string `json:"name"`
}

type DemoSearchModel struct {
	Name string `json:"name"`
}

type DemoListModel struct {
	Name string `json:"name"`
}
