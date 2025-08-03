package captcha

import (
	"fmt"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils/array"
	"gen_gin_tpl/pkg/variable"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"strings"
)

var captchaList []*base64Captcha.Captcha

// 生成字符串验证码
//
// 参数：
//   - width 验证码宽度
//   - height 验证码高度
//   - noiseCount 干扰线条数量
//   - length 验证码长度
//   - bgColor 背景颜色
//   - fonts 字体列表
//   - lineOpt 线条选项
//
// 返回值：
//   - *base64Captcha.Captcha 验证码对象
func newString(height, width, noiseCount, length int, bgColor *color.RGBA, fonts []string, lineOpt int) *base64Captcha.Captcha {
	source := fmt.Sprintf("%s%s", base64Captcha.TxtNumbers, base64Captcha.TxtAlphabet)
	driver := base64Captcha.NewDriverString(height, width, noiseCount, lineOpt, length, source, bgColor, nil, fonts)
	return base64Captcha.NewCaptcha(driver, variable.CaptchaStore)
}

// 生成中文验证码
//
// 参数：
//   - width 验证码宽度
//   - height 验证码高度
//   - noiseCount 干扰线条数量
//   - bgColor 背景颜色
//   - lineOpt 线条选项
//
// 返回值：
//   - *base64Captcha.Captcha 验证码对象
func newChinese(width, height, noiseCount, length int, bgColor *color.RGBA, lineOpt int) *base64Captcha.Captcha {
	fonts := []string{
		"wqy-microhei.ttc",
	}
	driver := base64Captcha.NewDriverChinese(height, width, noiseCount, lineOpt, length, base64Captcha.TxtChineseCharaters, bgColor, nil, fonts)
	return base64Captcha.NewCaptcha(driver, variable.CaptchaStore)
}

// 生成数学验证码
//
// 参数：
//   - width 验证码宽度
//   - height 验证码高度
//   - noiseCount 干扰线条数量
//   - bgColor 背景颜色
//   - fonts 字体列表
//   - lineOpt 线条选项
//
// 返回值：
//   - *base64Captcha.Captcha 验证码对象
func newMath(height, noiseCount int, bgColor *color.RGBA, fonts []string, lineOpt int) *base64Captcha.Captcha {
	driver := base64Captcha.NewDriverMath(height, height*4, noiseCount, lineOpt, bgColor, nil, fonts)
	return base64Captcha.NewCaptcha(driver, variable.CaptchaStore)
}

// 生成数字验证码
//
// 参数：
//   - width 验证码宽度
//   - height 验证码高度
//   - length 验证码长度
//
// 返回值：
//   - *base64Captcha.Captcha 验证码对象
func newDigit(width, height, length int) *base64Captcha.Captcha {
	driver := base64Captcha.NewDriverDigit(height, width, length, 0.7, 80)
	return base64Captcha.NewCaptcha(driver, variable.CaptchaStore)
}

// 处理字体
func captchaFonts(captchaFonts []string) []string {
	var fonts []string
	for _, font := range captchaFonts {
		if strings.TrimSpace(font) == "wqy-microhei" {
			fonts = append(fonts, "wqy-microhei.ttc")
			continue
		}
		fonts = append(fonts, fmt.Sprintf("%s.ttf", font))
	}
	return fonts
}

// 计算验证码宽度和高度
//
// 参数：
//   - length 验证码长度
//
// 返回值：
//   - width 验证码宽度
//   - height 验证码高度
func computationWidth(length int) (width, height int) {
	height = 40
	// 每个字符宽度 = 高度 * 0.6
	charWidth := int(float64(height) * 0.6)
	space := 5    // 字符间隔
	padding := 10 // 左右 padding

	// 动态计算验证码宽度
	width = padding*2 + (charWidth+space)*length - space
	return width, height
}

// InitImageCaptcha 根据提供的参数初始化图片验证码。
// 该函数会生成多种类型的验证码，包括字母数字、中文、数学和数字验证码。
// 如果提供了验证码类型列表，则选择对应的验证码；否则，程序会触发 panic。
//
// 参数：
//   - length     验证码的长度。
//   - noiseCount 验证码图片中的干扰线条数量。
//   - showLine   包含线条选项的整数切片。
//   - fontsList  包含字体名称的字符串切片。
//   - typeList   包含要生成的验证码类型的字符串切片。
func InitImageCaptcha(length, noiseCount int, showLine []int, fontsList, typeList []string) {
	fonts := captchaFonts(fontsList)
	lineOpt := array.RandomArray[int](showLine)
	width, height := computationWidth(length)
	bgColor := &color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(0)}

	tempCaptchaMap := map[string]*base64Captcha.Captcha{
		"alphanum": newString(height, width, noiseCount, length, bgColor, fonts, lineOpt),
		"chinese":  newChinese(width, height, noiseCount, length, bgColor, lineOpt),
		"math":     newMath(height, noiseCount, bgColor, fonts, lineOpt),
		"digit":    newDigit(width, height, length),
	}

	if len(typeList) > 0 {
		for _, key := range typeList {
			captchaList = append(captchaList, tempCaptchaMap[strings.ToLower(key)])
		}
	} else {
		log.Panic().Msg("请至少选择一个验证码类型")
	}
}
