package verifycode

import (
	"fmt"
	"sync"

	"github.com/vinoMamba/gohub/pkg/helpers"
	"github.com/vinoMamba/gohub/pkg/redis"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

// 单例模式，生成唯一的VerifyCode实例
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   "verifycode:",
			},
		}
	})
	return internalVerifyCode
}
func (vc *VerifyCode) Set(email string) bool {
	code := vc.generateVerifyCode(email)
	//TODO: 发送邮件
	fmt.Println("发送邮件", email, code)
	return true
}

func (vc *VerifyCode) generateVerifyCode(key string) string {
	code := helpers.RandomNumber(6)
	vc.Store.Set(key, code)
	return code
}

func (vc *VerifyCode) CheckAnswer(key, answer string) bool {
	return vc.Store.Verify(key, answer)
}
