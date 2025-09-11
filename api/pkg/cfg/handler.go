package cfg

import (
	"fmt"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/pkg/utils/network"
	"gen_gin_tpl/pkg/variable"
	"github.com/mojocn/base64Captcha"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

var (
	C *Config
)

func init() {
	C = &Config{
		SQLite: SQLite{
			DbName: fmt.Sprintf("%s.db", constants.ProjectName),
		},
		MySQL: MySQL{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "",
			DbName:   constants.ProjectName,
		},
		Redis: Redis{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			DB:       0,
		},
		Email: Email{
			SMTPHost: "smtp.163.com",
			SMTPPort: 587,
		},
		Web: Web{
			IP:   network.GetLanIP(),
			Port: 9527,
			Mode: "release",
		},
		Logger: Logger{
			Level:      zerolog.InfoLevel.String(),
			MaxSize:    50,
			LogPath:    "logs",
			FormatJSON: false,
			Compress:   true,
			MaxAge:     7,
			MaxBackups: 5,
		},
		Captcha: Captcha{
			Length:     5,
			NoiseCount: 30,
			Fonts:      []string{"3Dumb", "actionj", "ApothecaryFont", "chromohv", "Comismsh", "DeborahFancyDress", "DENNEthree-dee", "Flim-Flam", "RitaSmith", "wqy-microhei"},
			Type:       []string{"alphaNum", "chinese", "math", "digit"},
			ShowLine: []int{
				base64Captcha.OptionShowHollowLine, // 空心线（中空的曲线，像钢笔划出来的线条）
				base64Captcha.OptionShowSlimeLine,  // 粘稠线（细长黏糊糊的线，比较流动性）
				base64Captcha.OptionShowSineLine,   // 正弦波线（类似波浪线，波形起伏）

				// 空心线 + 粘稠线，两种线叠加，增强干扰
				base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
				//空心线 + 正弦波线，中空曲线和波浪线叠加
				base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSineLine,
				// 粘稠线 + 正弦波线，黏糊线+波浪线，视觉干扰力较强
				base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowSineLine,

				// 空心线 + 粘稠线 + 正弦波线，全体上线，满满的视觉干扰
				base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowSineLine,
			},
		},
		Other: Other{
			IsEmail:      true,
			DbType:       constants.SQLite,
			IsInitialize: false,
			CacheType:    constants.Memory,
			DataPath:     "data",
		},
	}

	// ini 覆盖
	path, err := file.GetFileAbsPath(".", fmt.Sprintf("%s.ini", constants.ProjectName))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	} else {
		variable.ConfigPath = path
	}

	if _file, err := ini.Load(variable.ConfigPath); err == nil {
		_ = _file.MapTo(C)
	} else {
		if _file, err := ini.Load("test/gen_gin_tpl.ini"); err == nil {
			_ = _file.MapTo(C)
		}
	}

	if C.Logger.Level == "no" {
		C.Logger.Level = ""
	}
	C.Logger.Level = strings.ToLower(C.Logger.Level)
	C.Other.DbType = constants.GetDbName(C.Other.DbType)
	C.Other.CacheType = constants.GetCacheName(C.Other.CacheType)
}

// SaveToIni 保存配置到 ini 文件
func SaveToIni() {
	_file := ini.Empty()
	err := _file.ReflectFrom(C)
	if err != nil {
		log.Fatal().Err(err).Msg("配置保存，序列化成ini失败")
	}

	for _, name := range []string{constants.SQLite, constants.MySQL} {
		if strings.EqualFold(C.Other.DbType, name) {
			continue
		}
		_file.DeleteSection(name)
	}

	if strings.EqualFold(C.Other.CacheType, constants.Memory) {
		for _, name := range []string{constants.Redis} {
			_file.DeleteSection(name)
		}
	}
	if !C.Other.IsEmail {
		_file.DeleteSection("Email")
	}

	if _file.SaveTo(variable.ConfigPath) != nil {
		log.Fatal().Err(err).Msg("配置文件保存失败")
	}
}
