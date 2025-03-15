package consumer

import "gomall/app/email/biz/consumer/email"

// 用于初始化所有消费者
func Init() {
	email.ConsumerInit()
}
