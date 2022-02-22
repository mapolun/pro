package fn

import (
	"pro/config"
	"regexp"
	"strings"
)

//验证图片合法性
func ValidateImage(filePath string) bool {
	var host string
	host = config.Get.Upload.Host
	return strings.Contains(filePath, host)
}

//验证手机号
func ValidateMobile(mobile string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	regs := regexp.MustCompile(reg)
	return regs.MatchString(mobile)
}
