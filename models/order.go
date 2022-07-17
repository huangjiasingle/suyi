package models

import (
	"time"
)

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
	CreateAt         time.Time `json:"create_at" xorm:"TIMESTAMP"`
	Status           string    `json:"status" xorm:"comment('1. 待接单 2. 已接单 3. 已完成') VARCHAR(255)"`
	UpdateAt         time.Time `json:"update_at" xorm:"TIMESTAMP"`
	Cost             float32   `json:"cost" xorm:"FLOAT(10,2)"`
}
