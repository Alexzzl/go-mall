package routers

import (
	"OnlineShop/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 首页数据
	// beego.RESTRouter("/", &controllers.IndexController{}, "GET:Index")
	// 购物接口
	//beego.RESTRouter("/buy", &controllers.IndexController{}, "POST:Buy")
	// api 接口全部以 api 作为链接前缀
	ns := beego.NewNamespace("/api",
		beego.NSRouter("/", &controllers.IndexController{}, "GET:GoodJoinCart"),
		beego.NSRouter("/good/joincart", &controllers.IndexController{}, "POST:GoodJoinCart"),
		beego.NSRouter("/user/login", &controllers.IndexController{}, "POST:UserLogin"),
		beego.NSRouter("/good/collection", &controllers.IndexController{}, "POST:JoinCollection"),
		beego.NSRouter("/good/seckill", &controllers.IndexController{}, "POST:SeckillBuy"),
		beego.NSRouter("/user_info", &controllers.IndexController{}, "GET:UserInfo"),
		beego.NSRouter("/good/:good_id([0-9]+)/like", &controllers.IndexController{}, "PUT:GoodLike"),
		beego.NSRouter("/good/good_type", &controllers.IndexController{}, "GET:GoodsFilter"),
		beego.NSRouter("/good/:good_id([0-9]+)/comments", &controllers.IndexController{}, "GET:GetComments;POST:AddComment"),
		beego.NSRouter("/good/:good_id([0-9]+)/info", &controllers.IndexController{}, "GET:GoodDetail"),
		beego.NSRouter("/good/seckill", &controllers.IndexController{}, "GET:SecKillGoods"),
		beego.NSRouter("/user/good_cart", &controllers.IndexController{}, "GET:CartSummary"),
		beego.NSRouter("/user/orders", &controllers.IndexController{}, "GET:ShowOrders"),
		beego.NSRouter("/user/cart", &controllers.IndexController{}, "GET:CartProduceOrder"),
	)
	beego.AddNamespace(ns)
	// 页面路由
	beego.Router("/", &controllers.IndexController{}, "GET:Index")
	//访问商品详情页面
	beego.Router("/good/:good_id([0-9]+)", &controllers.IndexController{}, "GET:GoodDetailIndex")
	// 访问用户购物车页面
	beego.Router("/user/shop_cart", &controllers.IndexController{}, "GET:ShopCartDetail")
	// 访问订单页面
	beego.Router("/order", &controllers.IndexController{}, "GET:OrderIndex")
}
