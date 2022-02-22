package request

import validate2 "github.com/Pis0sion/components/validate"

type WechatComplainRequest struct {
	Mobile   string   `form:"mobile" json:"mobile" binding:"required,min=11,max=11" reject:"请输入手机号"`
	Reason   string   `form:"reason" json:"reason" binding:"required"`
	ClientIp string   `form:"clientIp" json:"client_ip"`
	Content  string   `form:"content" json:"content" binding:"required"`
	Image    []string `form:"image" json:"image"`
}

func (v *WechatComplainRequest) validate() error {
	return validate2.NewValidator(v).Validate()
}
