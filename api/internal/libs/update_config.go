package libs

import (
	"fmt"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/enums/boolean"
	"gen_gin_tpl/pkg/enums/dtype"
	"gen_gin_tpl/pkg/variable"
	"reflect"
	"strconv"
	"sync"
)

type Config struct {
	ContextIsEncrypted *models.Config `json:"IS_ENCRYPTED_RESPONSE"`
	WebTitle           *models.Config `json:"WEB_TITLE"`
	PublicPEM          *models.Config `json:"PUBLIC_PEM"`
	PrivatePEM         *models.Config `json:"PRIVATE_PEM"`
	SessionKey         *models.Config `json:"SESSION_KEY"`
	Countdown          *models.Config `json:"COUNTDOWN"`
}

var (
	ones      sync.Once
	WebConfig *Config
)

func InitializeWebConfig() {
	ones.Do(func() {
		InitializeUpdateWebConfig()
	})
}

func InitializeUpdateWebConfig() {
	WebConfig = &Config{
		ContextIsEncrypted: &models.Config{Name: constants.ContextIsEncrypted, Value: boolean.False.Key(), Default: boolean.False.Key(), ValueT: dtype.Bool, Show: boolean.True, Description: "是否开启加密模式"},
		WebTitle:           &models.Config{Name: constants.WebTitle, Value: variable.WebTitle, Default: variable.WebTitle, ValueT: dtype.String, Show: boolean.True, Description: "站点标题"},
		PublicPEM:          &models.Config{Name: constants.PublicPEM, Value: string(variable.PublicPEM), Default: string(variable.PublicPEM), ValueT: dtype.String, Show: boolean.True, Description: "加密公钥"},
		PrivatePEM:         &models.Config{Name: constants.PrivatePEM, Value: string(variable.PrivatePEM), Default: string(variable.PrivatePEM), ValueT: dtype.String, Show: boolean.True, Description: "加密私钥"},
		SessionKey:         &models.Config{Name: constants.SessionKey, Value: string(variable.SessionKey), Default: string(variable.SessionKey), ValueT: dtype.String, Show: boolean.True, Description: "会话密钥"},
		Countdown:          &models.Config{Name: constants.Countdown, Value: "60", Default: "60", ValueT: dtype.Int, Show: boolean.True, Description: "统一倒计时时间，单位秒"},
	}
}

func (r *Config) UpdateWebConfig() {
	configs, err := query.Config.Find()
	if err != nil {
		return
	}
	v := reflect.ValueOf(r).Elem() // Config struct
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		for _, model := range configs {
			if model.Name == jsonTag {
				v.Field(i).Set(reflect.ValueOf(model))
			}
		}
	}

	variable.WebTitle = r.GetWebTitle()
	variable.PublicPEM = r.GetPublicPEM()
	variable.PrivatePEM = r.GetPrivatePEM()
	variable.SessionKey = r.GetSessionKey()
}

func (r *Config) IsContextIsEncrypted() bool {
	if r.ContextIsEncrypted.Value == "" {
		return r.ContextIsEncrypted.Default == boolean.True.Key()
	}
	return r.ContextIsEncrypted.Value == boolean.True.Key()
}

func (r *Config) GetWebTitle() string {
	if r.WebTitle.Value == "" {
		return r.WebTitle.Default
	}
	return r.WebTitle.Value
}

func (r *Config) GetPublicPEM() []byte {
	if r.PublicPEM.Value == "" {
		return []byte(r.PublicPEM.Default)
	}
	return []byte(r.PublicPEM.Value)
}

func (r *Config) GetPrivatePEM() []byte {
	if r.PrivatePEM.Value == "" {
		return []byte(r.PrivatePEM.Default)
	}
	return []byte(r.PrivatePEM.Value)
}

func (r *Config) GetSessionKey() []byte {
	if r.SessionKey.Value == "" {
		return []byte(r.SessionKey.Default)
	}
	return []byte(r.SessionKey.Value)
}

func (r *Config) GetCountdown() int {
	if r.Countdown.Value == "" {
		n, err := strconv.Atoi(r.Countdown.Default)
		if err != nil {
			fmt.Println("GetCountdown.Countdown.Default 转换失败:", err)
			return 60
		}
		return n
	}
	n, err := strconv.Atoi(r.Countdown.Value)
	if err != nil {
		fmt.Println("GetCountdown.Countdown.Value 转换失败:", err)
		return 60
	}
	return n
}

func (r *Config) GetModelList() []models.Config {
	var mc []models.Config
	v := reflect.ValueOf(r).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			// 字段本身就是 ModelConfig 值类型
			mc = append(mc, field.Interface().(models.Config))
		case reflect.Ptr:
			// 字段是 *ModelConfig 指针类型
			if !field.IsNil() {
				mc = append(mc, field.Elem().Interface().(models.Config))
			}
		default:
			panic("unhandled default case")
		}
	}
	return mc
}
