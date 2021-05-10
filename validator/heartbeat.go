package validator

import (
	"github.com/exuan/waka-api/model"
	"gorm.io/datatypes"
)

type (
	Heartbeat struct {
		Time          float64 `json:"time" validate:"omitempty"`
		UserId        int64
		Entity        string         `json:"entity" validate:"omitempty"`
		Type          string         `json:"type" validate:"omitempty"`
		Category      string         `json:"category" validate:"omitempty"`
		IsWrite       bool           `json:"is_write" validate:"omitempty"`
		Project       string         `json:"project" validate:"omitempty"`
		Branch        string         `json:"branch" validate:"omitempty"`
		Language      string         `json:"language" validate:"omitempty"`
		Dependencies  datatypes.JSON `gorm:"type:json;default:[];"`
		Lines         model.Int64    `json:"lines" validate:"omitempty" gorm:"default:0"`
		Lineno        model.Int64    `json:"lineno" validate:"omitempty"`
		Cursorpos     model.Int64    `json:"cursorpos" validate:"omitempty"`
		Editor        string
		EditorVersion string
		Machine       string
		Platform      string
		Kernel        string
		Timezone      string
		Ip            string
		UserAgent     string `json:"user_agent" validate:"omitempty"`
		UpdateTime    int64  `gorm:"type:int not null default 0;autoUpdateTime"`
		CreateTime    int64  `gorm:"type:int not null default 0;autoCreateTime"`
	}
)
