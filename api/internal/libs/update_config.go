package libs

import (
	"encoding/json"
	"fmt"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/enums/boolean"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/variable"
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	ContextIsEncrypted *models.Config `json:"IS_ENCRYPTED_RESPONSE"`
	WebTitle           *models.Config `json:"WEB_TITLE"`
	PublicPEM          *models.Config `json:"PUBLIC_PEM"`
	PrivatePEM         *models.Config `json:"PRIVATE_PEM"`
	Countdown          *models.Config `json:"COUNTDOWN"`
}

var (
	WebConfig *Config
)

func init() {
	WebConfig = &Config{}
}

func toListMap(v []*models.Config) ([]map[string]any, error) {
	var result []map[string]any
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateWebConfig 初始化Web配置
// 参数：
//   - configList: 配置列表
//
// 返回值：
//   - 无
//
// 说明：
//   - 初始化Web配置，将配置列表转换为JSON格式，并设置到变量中
func UpdateWebConfig(configList []*models.Config) {
	toJSON, err := toListMap(configList)
	if err != nil {
		log.Error().Err(err).Msg("Web 配置初始化失败")
		os.Exit(-1)
	}
	WebConfig.update(toJSON)
}

func (r *Config) update(configList []map[string]any) {
	v := reflect.ValueOf(r).Elem() // Config struct
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		for _, cfg := range configList {
			if name, ok := cfg["Name"].(string); ok && name == jsonTag {
				// 创建 ModelConfig 指针
				model := &models.Config{}
				// 遍历 map 设置 ModelConfig 字段
				modelV := reflect.ValueOf(model).Elem()
				for k, val := range cfg {
					f := modelV.FieldByName(k)
					if f.IsValid() && f.CanSet() {
						fv := reflect.ValueOf(val)
						if !fv.IsValid() {
							continue // 值是 nil，跳过
						}

						// 如果目标字段是指针且 fv 类型可赋值给 Elem
						if f.Kind() == reflect.Ptr {
							if fv.Type().AssignableTo(f.Type().Elem()) {
								ptr := reflect.New(f.Type().Elem())
								ptr.Elem().Set(fv)
								f.Set(ptr)
							}
						} else if fv.Type().AssignableTo(f.Type()) {
							f.Set(fv)
						}
					}
				}
				v.Field(i).Set(reflect.ValueOf(model))
			}
		}
	}

	variable.WebTitle = r.GetWebTitle()
	variable.PublicPEM = r.GetPublicPEM()
	variable.PrivatePEM = r.GetPrivatePEM()
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

func (r *Config) GetConfigList() []models.Config {
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
