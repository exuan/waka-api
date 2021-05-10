package model

import (
	"database/sql"

	"gorm.io/datatypes"
)

type (
	Heartbeat struct {
		Id            int64          `gorm:"AUTOINCREMENT;primaryKey;not null;"`
		UserId        int64          `gorm:"type:int not null default 0;"`
		Time          float64        `gorm:"type:float not null default 0.0;"`
		Entity        string         `gorm:"type:varchar(500) not null default '';"`
		Type          string         `gorm:"type:varchar(64) not null default '';"`
		Category      String         `gorm:"type:varchar(64) not null default '';"`
		IsWrite       sql.NullBool   `gorm:"default:true"`
		Project       string         `gorm:"type:varchar(128) not null default '';"`
		Branch        String         `gorm:"type:varchar(128) not null default '';"`
		Language      string         `gorm:"type:varchar(128) not null default '';"`
		Dependencies  datatypes.JSON `gorm:"type:json not null;"`
		Lines         Int64          `gorm:"type:int not null default 0;"`
		Lineno        Int64          `gorm:"type:int not null default 0;"`
		Cursorpos     Int64          `gorm:"type:int not null default 0;"`
		Editor        string         `gorm:"type:varchar(32)not null default '';"`
		EditorVersion string         `gorm:"type:varchar(32) not null default '';"`
		Machine       string         `gorm:"type:varchar(128) not null default '';"`
		Platform      string         `gorm:"type:varchar(32) not null default '';"`
		Kernel        string         `gorm:"type:varchar(128) not null default '';"`
		UserAgent     string         `gorm:"type:varchar(300) not null default '';"`
		Ip            string         `gorm:"type:varchar(128) not null default '';"`
		Timezone      string         `gorm:"type:varchar(128) not null default '';"`
		UpdateTime    int64          `gorm:"type:int not null default 0;autoUpdateTime"`
		CreateTime    int64          `gorm:"type:int not null default 0;autoCreateTime"`
	}
)
