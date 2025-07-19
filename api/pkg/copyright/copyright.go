package copyright

import (
	"gen_gin_tpl/pkg/variable"
	"time"
)

// _copyright 版权信息
type _copyright struct {
	Name      string `json:"name"`
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
}

func NewCopyright() _copyright {
	return _copyright{
		Name:      variable.WebTitle,
		StartTime: 2024,
		EndTime:   time.Now().Year(),
	}
}
