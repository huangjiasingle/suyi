package models

import "time"

// Order 订单信息表
type Order struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT"`
	UserName         string    `json:"user_name" xorm:"VARCHAR(255)"`
	UserPhone        string    `json:"user_phone" xorm:"not null VARCHAR(255)"`
	CarType          string    `json:"car_type" xorm:"not null VARCHAR(255)"`
	CarNum           string    `json:"car_num" xorm:"not null VARCHAR(255)"`
	CarLocation      string    `json:"car_location" xorm:"VARCHAR(4096)"`
	CarLongitude     string    `json:"car_longitude" xorm:"VARCHAR(255)"`
	CarLatitude      string    `json:"car_latitude" xorm:"VARCHAR(4096)"`
	ServiceType      int       `json:"service_type" xorm:"not null index INT"`
	ExternalServices int       `json:"external_services" xorm:"index INT"`
	CreateAt         time.Time `json:"create_at" xorm:"created TIMESTAMP"`
	Status           string    `json:"status" xorm:"comment('1. 待接单 2. 已接单 3. 已完成 4. 已取消') VARCHAR(255)"`
	UpdateAt         time.Time `json:"update_at" xorm:"TIMESTAMP"`
	Cost             float32   `json:"cost" xorm:"FLOAT(10,2)"`
}

/*
Service 服务项目表
1. 水蜡洗车
2. 精品打蜡
3. 内饰深度清洗 （包括：仪表盘清洗打蜡、座椅精洗上光、车内顶蓬精洗、门板精洗养护、地板精洗、后备箱精洗、消毒）
4. 轮毂翻新
5. 发动机机舱清洗
6. 摩托车清洗
*/
type Service struct {
	Id          int       `json:"id" xorm:"not null pk index INT"`
	Name        string    `json:"name" xorm:"not null pk VARCHAR(255)"`
	Description string    `json:"description" xorm:"VARCHAR(4096)"`
	Photo       []byte    `json:"photo" xorm:"LONGBLOB"`
	Price       float32   `json:"price" xorm:"not null default 0 FLOAT"`
	CreateAt    time.Time `json:"create_at" xorm:"TIMESTAMP"`
}

type Pagination struct {
	PageSise int         `json:"pageSize"`
	PageNum  int         `json:"pageNum"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}
