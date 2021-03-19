/**
 * @Author: gaoshiyao
 * @Description: bgres
 * @File:  bgres
 * @Date: 2021/03/19 18:04
 */

package bgres

type BgRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	OK                      = BgRes{Code: 200, Message: "ok"}
	InternalServerError     = BgRes{Code: 100001, Message: "系统异常"}
	DataFailError           = BgRes{Code: 100002, Message: "数据获取失败"}
	AccountNotForeendError  = BgRes{Code: 100003, Message: "该账户非前台账号不能登录"}
	AccountNotEnableError   = BgRes{Code: 100004, Message: "该账户未启用,不能登录"}
	RequireLoginError       = BgRes{Code: 100005, Message: "登录失效"}
	LogoutError             = BgRes{Code: 100006, Message: "退出失败"}
	DeleteDataFailError     = BgRes{Code: 100007, Message: "数据删除失败"}
	ParamsParserError       = BgRes{Code: 200001, Message: "参数解析失败"}
	ParamsEmptyError        = BgRes{Code: 200002, Message: "参数不能为空"}
	MobileEmptyError        = BgRes{Code: 200003, Message: "手机号不能为空"}
	VerifyCodeEmptyError    = BgRes{Code: 200004, Message: "验证码不能为空"}
	VerifyCodeValidError    = BgRes{Code: 200005, Message: "验证码错误或已失效"}
	LoginFailError          = BgRes{Code: 200006, Message: "登录失败，请稍后重试"}
	SendSmsFailError        = BgRes{Code: 200007, Message: "验证码发送失败，请稍后重试"}
	VerifyCodeFormatError   = BgRes{Code: 200008, Message: "验证码格式错误"}
	MobileFormatError       = BgRes{Code: 200009, Message: "手机号格式错误"}
	RoleSelectedError       = BgRes{Code: 200010, Message: "已选择角色,不能再次选择"}
	RoleSelectFailError     = BgRes{Code: 200011, Message: "角色选择失败，请重试"}
	UnknowRoleErr           = BgRes{Code: 200012, Message: "未知角色"}
	ParamsValidError        = BgRes{Code: 200013, Message: "无效参数"}
	RegisterEndError        = BgRes{Code: 200014, Message: "您已经登记过基础信息"}
	RegisterFailError       = BgRes{Code: 200015, Message: "基础信息登记失败,请稍后重试"}
	UpdateFailError         = BgRes{Code: 200016, Message: "更新失败,请稍后重试"}
	CompanyNotExistsError   = BgRes{Code: 200017, Message: "所选公司不存在"}
	OftenMeshNotExistsError = BgRes{Code: 200018, Message: "所选常用农机不存在"}
	UserNamePwdError        = BgRes{Code: 200101, Message: "用户名或密码错误"}
	UserNameError           = BgRes{Code: 200102, Message: "请输入正确的用户名"}
	UserPwdError            = BgRes{Code: 200103, Message: "请输入密码"}
)
