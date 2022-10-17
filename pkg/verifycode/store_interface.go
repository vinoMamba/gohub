package verifycode

type Store interface {
	//保存验证码
	Set(id string, verifycode string) bool
	//获取验证码
	Get(id string) string
	//检查验证码
	Verify(id string, answer string) bool
}
