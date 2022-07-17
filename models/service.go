package models

import (
	"time"
)

type Service struct {
	Id   int    `json:"id" xorm:"not null pk index INT"`
	Name string `json:"name" xorm:"not null pk comment('1. 水蜡洗车  
2. 精品打蜡
3. 内饰深度清洗 （包括：仪表盘清洗打蜡、座椅精洗上光、车内顶蓬精洗、门板精洗养护、地板精洗、后备箱精洗、消毒）
4. 轮毂翻新
5. 发动机机舱清洗
6. 摩托车清洗') VARCHAR(255)"`
	Description string    `json:"description" xorm:"VARCHAR(4096)"`
	Photo       []byte    `json:"photo" xorm:"LONGBLOB"`
	Price       float32   `json:"price" xorm:"not null default 0 FLOAT"`
	CreateAt    time.Time `json:"create_at" xorm:"TIMESTAMP"`
}
