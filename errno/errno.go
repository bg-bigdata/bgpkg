package errno

import "fmt"

type Errno struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Errno) Error() string {
	return fmt.Sprintf("code:%d Message:%s", e.Code, e.Message)
}

var (
	OK                     = Errno{Code: 200, Message: "ok"}
	InternalServerError    = Errno{Code: 100001, Message: "系统异常"}
	DataFailError          = Errno{Code: 100002, Message: "数据获取失败"}
	AccountNotForeendError = Errno{Code: 100003, Message: "该账户非前台账号不能登录"}
	AccountNotEnableError  = Errno{Code: 100004, Message: "该账户未启用,不能登录"}
	RequireLoginError      = Errno{Code: 100005, Message: "登录失效"}
	LogoutError            = Errno{Code: 100006, Message: "退出失败"}
	DeleteDataFailError    = Errno{Code: 100007, Message: "数据删除失败"}
	// =
	ParamsParserError       = Errno{Code: 200001, Message: "参数解析失败"}
	ParamsEmptyError        = Errno{Code: 200002, Message: "参数不能为空"}
	MobileEmptyError        = Errno{Code: 200003, Message: "手机号不能为空"}
	VerifyCodeEmptyError    = Errno{Code: 200004, Message: "验证码不能为空"}
	VerifyCodeValidError    = Errno{Code: 200005, Message: "验证码错误或已失效"}
	LoginFailError          = Errno{Code: 200006, Message: "登录失败，请稍后重试"}
	SendSmsFailError        = Errno{Code: 200007, Message: "验证码发送失败，请稍后重试"}
	VerifyCodeFormatError   = Errno{Code: 200008, Message: "验证码格式错误"}
	MobileFormatError       = Errno{Code: 200009, Message: "手机号格式错误"}
	RoleSelectedError       = Errno{Code: 200010, Message: "已选择角色,不能再次选择"}
	RoleSelectFailError     = Errno{Code: 200011, Message: "角色选择失败，请重试"}
	UnknowRoleErr           = Errno{Code: 200012, Message: "未知角色"}
	ParamsValidError        = Errno{Code: 200013, Message: "无效参数"}
	RegisterEndError        = Errno{Code: 200014, Message: "您已经登记过基础信息"}
	RegisterFailError       = Errno{Code: 200015, Message: "基础信息登记失败,请稍后重试"}
	UpdateFailError         = Errno{Code: 200016, Message: "更新失败,请稍后重试"}
	CompanyNotExistsError   = Errno{Code: 200017, Message: "所选公司不存在"}
	OftenMeshNotExistsError = Errno{Code: 200018, Message: "所选常用农机不存在"}
	UserNamePwdError        = Errno{Code: 200101, Message: "用户名或密码错误"}
	UserNameError           = Errno{Code: 200102, Message: "请输入正确的用户名"}
	UserPwdError            = Errno{Code: 200103, Message: "请输入密码"}
)

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case Errno:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	}
	fmt.Printf("响应错误信息被拦截: %s\n", err.Error())
	return InternalServerError.Code, InternalServerError.Message
}

func NewErrNo(code int, message string) Errno {
	return Errno{
		Code:    code,
		Message: message,
	}
}
