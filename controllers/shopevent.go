package controllers
import (
	"OnlineShop/models"
)
// 这里为了偷懒使用了指针类型，作为一个事件消息来说，
// 应该具备跨进程和跨机器的能力, 因为变量类型应该只能是基本类型
// 因为beego框架的限制，没有统一处理 controller的地方
// 也就是说在当前 `请求生命周期` 内，在任意地方应该可以获取当前`controller` 指针的能力，
// 如果具备了这样的能力，那么很容易在任意地方去 `make_response` (思想来源skynet框架中)。

type ShopEvent struct {
	//购物事件
	User *models.User
	GoodDetail *models.GoodDetail
	Num int
	// 当购物事件处理完成之后进行回调
	Controller *IndexController
}
func(shop_event *ShopEvent)LoadUser(){
	// 模型定义应该是UserId
	//通过LoadUser来实现user_id => UserInstance的转化
}