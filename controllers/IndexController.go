package controllers

import (
	"OnlineShop/requests"
	"OnlineShop/response"
	"encoding/json"
	// "fmt"
	"OnlineShop/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"runtime"
	"strconv"
	"time"
)

type IndexController struct {
	beego.Controller
}

// 用户登录
// api /api/user/login
func (c *IndexController) UserLogin() {

	defer func() {
		c.ServeJSON()
	}()
	var body requests.UserLogin
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.Data["json"] = response.FailRes("数据格式不正确")
	} else {
		username, password := body.UserName, body.Password
		username = "zz@qq.com"
		password = "0000"
		if code, user := UserLogin(username, password); code == true {
			c.SetSession("session", user.Id)
			shop_cart_nums, _ := GetUserShopCartNum(user.Id)
			c.Data["json"] = map[string]interface{}{"ok": true, "data": map[string]interface{}{
				"UserName":    user.Nickname,
				"image":       user.Image,
				"UserId":      user.Id,
				"ShopCartNum": shop_cart_nums,
			}}
		} else {
			c.Data["json"] = response.FailRes("用户名或者密码错误")
		}
	}
}

// 获取用户信息，如果失败，那么说明未登录
func (c *IndexController) UserInfo() {
	defer c.ServeJSON()

	user_id := c.GetSession("session")
	if user_id == nil {
		c.Data["json"] = map[string]interface{}{"ok": false, "message": "no auth"}
	} else {
		if user := GetUserById(user_id.(int)); user != nil {
			shop_cart_nums, _ := GetUserShopCartNum(user.Id)
			user_info := map[string]interface{}{
				"UserName":    user.Nickname,
				"image":       user.Image,
				"UserId":      user.Id,
				"ShopCartNum": shop_cart_nums,
			}
			c.Data["json"] = map[string]interface{}{"ok": true, "data": user_info}

		} else {
			c.Data["json"] = map[string]interface{}{"ok": false, "message": "server error"}
		}
	}
}

// 商城首页
func (c *IndexController) Index() {
	// 设置cookie来防止
	c.TplName = "index.tpl"
}

// 商品详情页
func (c *IndexController) GoodDetailIndex() {
	c.TplName = "good_detail.tpl"
}

//购物车结算页面
func (c *IndexController) ShopCartDetail() {
	c.TplName = "good_cart.html"
}

func (c *IndexController) OrderIndex() {
	c.TplName = "orderIndex.html"
}

/*
加入购物车操作
*/
type JoinCart struct {
	GoodId int
}

func (c *IndexController) GoodJoinCart() {

	defer func() {
		c.ServeJSON()
	}()
	v := c.GetSession("session")
	var version int = 1
	if v == nil {
		//c.Redirect("/user/login", 302)
		//c.SetSession("session_user_id", 1)
		c.Data["json"] = map[string]interface{}{"ok": false, "message": "no auth"}
	} else {
		var body JoinCart
		var good_id int = -1
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err == nil {
			if v := GetGoodDetail(body.GoodId); v != nil {
				good_id = body.GoodId
			}
		}
		if good_id == -1 {
			c.Data["json"] = map[string]interface{}{"ok": false, "message": "无效的商品Id"}
			return
		}
		if user := GetUserById(v.(int)); user != nil {
			if _, err := UserBuyGood(user, good_id, version); err == nil {
				c.Data["json"] = map[string]interface{}{"ok": true, "message": "success"}
			} else {
				c.Data["json"] = map[string]interface{}{"ok": false, "message": err.Error()}
			}
		} else {
			c.Data["json"] = map[string]interface{}{"ok": false, "message": "非法用户"}
		}
	}
}

//商品是否喜欢
func (c *IndexController) GoodLike() {
	defer func() {
		fmt.Println(c.Data["json"])
		c.ServeJSON()
	}()

	good_id := c.Ctx.Input.Param(":good_id")
	like_status := c.GetString("status")
	num, err := strconv.Atoi(good_id)
	if err != nil || (like_status != "like" && like_status != "unlike") {
		c.Data["json"] = map[string]interface{}{"ok": false, "message": "无效的good_id"}
	} else {
		LikeGood(num, like_status)
		c.Data["json"] = map[string]interface{}{"ok": true}
	}
}

// 商品类型筛选
func (c *IndexController) GoodsFilter() {
	defer func() {
		c.ServeJSON()
	}()

	goodType := c.GetString("good_type")
	var goods []*models.GoodDetail
	if goodType != "all" {
		goods = FindShopByType(goodType)
	} else {
		goods = GetAllGoods(10)
	}
	c.Data["json"] = map[string]interface{}{"ok": true, "data": goods}
}

/*
加入到收藏夹
*/
func (c *IndexController) JoinCollection() {
	defer func() {
		c.ServeJSON()
	}()
	//user_id := c.GetSession("session_user_id").(int)
	c.Data["json"] = response.SuccRes()
}

//秒杀抢购的请求
func (c *IndexController) SeckillBuy() {
	defer func() {
		// 初始状态
		st := time.Now().Unix()
		for {
			end := time.Now().Unix()
			if _, ok := c.GetData("json"); ok {
				c.ServeJSON()
				break
			} else if end-st >= 2 { //大于2秒直接返回失败
				c.Data["json"] = response.FailRes("超时，购买失败")
				c.ServeJSON()
				break
			} else {
				//让出cpu
				runtime.Gosched()
			}
		}
	}()
	/*
	   var v requests.BuyRequest
	   json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	*/
	user := GetUserById(1)
	good_detail := GetGoodDetail(1)
	if good_detail == nil || user == nil {
		fmt.Println("\n数据库查找失败")
	} else {
		event := ShopEvent{User: user, GoodDetail: good_detail, Num: 1, Controller: c}
		if _, err := TradeSer.AddSecEvent(&event); err == nil {
			fmt.Println("加入消息队列成功")
		}
	}

}

// 获取商品的评论
type coments_response_struct struct {
	Comment     string
	Image       string
	Nickname    string
	CreatedTime time.Time
}

func comments_response(goods *models.Good) []coments_response_struct {
	result := make([]coments_response_struct, len(goods.Comments))
	for index, item := range goods.Comments {
		result[index] = coments_response_struct{
			Comment:     item.Content,
			Image:       item.User.Image,
			Nickname:    item.User.Nickname,
			CreatedTime: item.CreatedTime}
	}
	return result
}

func (c *IndexController) GetComments() {
	defer c.ServeJSON()
	//注意这个是Good.Id不是GoodDetail.Id
	good_id := c.Ctx.Input.Param(":good_id")
	good_id_num, err := strconv.Atoi(good_id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"ok": false, "message": "无效的good_id"}
		return
	}
	good_detail := GetGoodDetail(good_id_num)
	good := GetGoodById(good_detail.Good.Id)
	data := comments_response(good)
	c.Data["json"] = map[string]interface{}{"ok": true, "data": data}
}

// 增加一次评论
type add_comment_struct struct {
	Comment string
}

func (c *IndexController) AddComment() {
	defer c.ServeJSON()

	v := c.GetSession("session")
	if v == nil {
		c.Data["json"] = map[string]interface{}{"ok": false, "message": "no auth"}
		return
	}
	user := GetUserById(v.(int))
	good_id := c.Ctx.Input.Param(":good_id")
	good_id_num, _ := strconv.Atoi(good_id)
	var body add_comment_struct
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err == nil {
		if _, err := AddComment(good_id_num, body.Comment, user); err == nil {
			c.Data["json"] = map[string]interface{}{"ok": true, "data": map[string]interface{}{"Username": "zxl", "Image": "zxxx"}}
			return
		} else {
			fmt.Println("add comment eror: ", err.Error())
		}
	}
	c.Data["json"] = map[string]interface{}{"ok": false}
}

//商品详情
func (c *IndexController) GoodDetail() {
	defer c.ServeJSON()
	id := c.Ctx.Input.Param(":good_id")
	GoodDetailIdNum, _ := strconv.Atoi(id)
	good_detail := GetGoodDetail(GoodDetailIdNum)
	c.Data["json"] = map[string]interface{}{"ok": true, "data": good_detail}
}

// 秒杀商品数据
func (c *IndexController) SecKillGoods() {
	defer c.ServeJSON()

	o := orm.NewOrm()
	var goods []models.SecKillGoods
	if _, err := o.QueryTable("sec_kill_goods").RelatedSel().All(&goods); err == nil {
		c.Data["json"] = map[string]interface{}{"ok": true, "data": goods}
	} else {
		c.Data["json"] = map[string]interface{}{"ok": true, "data": goods}
	}
}

//购物车商品展示
type goodcart_goods struct {
	Name  string
	Num   int
	Price float32
}

func (c *IndexController) CartSummary() {
	defer c.ServeJSON()

	v := c.GetSession("session")
	if v == nil {
		c.Data["json"] = map[string]interface{}{"ok": false, "message": "no auth"}
		return
	}

	user := GetUserById(v.(int))
	//user := GetUserById(1)
	cart, _ := GetUserShopCart(user)
	// 根据商品id进行聚合， 返回name, num, price
	if len(cart.Goods) > 0 {
		nums := 0
		var GroupBy = make(map[int]*goodcart_goods)
		//根据商品id聚合
		for _, item := range cart.Goods {
			if _, ok := GroupBy[item.Id]; ok {
				GroupBy[item.Id].Num = GroupBy[item.Id].Num + 1
			} else {
				nums += 1
				GroupBy[item.Id] = &goodcart_goods{Name: item.Good.Name, Num: 1, Price: item.Price}
			}
		}
		var result = make([]goodcart_goods, nums)
		nums = 0
		for _, item := range GroupBy {
			result[nums] = goodcart_goods{Name: item.Name, Num: item.Num, Price: item.Price}
			nums += 1
		}
		var result_json = make(map[string]interface{})
		result_json["ok"] = true
		result_json["data"] = result
		if len(user.ShopAddress) > 0 {
			result_json["address"] = map[string]interface{}{
				"name":    user.ShopAddress[0].PersonName,
				"phone":   user.ShopAddress[0].Phone,
				"address": user.ShopAddress[0].DetailAddress,
			}
		}
		c.Data["json"] = result_json
	} else {
		c.Data["json"] = map[string]interface{}{"ok": true, "data": [...]int{}}
	}

}

//获取所有的订单信息
func (c *IndexController) ShowOrders() {
	defer c.ServeJSON()

	userIDStr := c.GetSession("session")
	if userIDStr == nil {
		c.Data["json"] = map[string]interface{}{
			"ok":      false,
			"message": " no auth",
		}
		return
	}
	userID, _ := userIDStr.(int)
	user := GetUserById(userID)

	orderType, err := c.GetInt("orderType")
	if err != nil || orderType == -1 {
		c.Data["json"] = map[string]interface{}{
			"ok":   true,
			"data": user.Orders,
		}
	} else {
		var result = make([]*models.Orders, len(user.Orders))
		index := 0
		for _, item := range user.Orders {
			if item.Status == orderType {
				result[index] = item
				index += 1
			}
		}
		c.Data["json"] = map[string]interface{}{
			"ok":   true,
			"data": result[:index],
		}
	}
}

// 购物车生成订单
func (c *IndexController) CartProduceOrder() {
	defer c.ServeJSON()

	Uid := c.GetSession("session")
	if Uid == nil {
		c.Data["json"] = map[string]interface{}{
			"ok":      false,
			"message": "no auth",
		}
		return
	}
	user := GetUserById(Uid.(int))
	cart, _ := GetUserShopCart(user)
	nums := 0
	var GroupBy = make(map[int]*goodcart_goods)
	//根据商品id聚合
	for _, item := range cart.Goods {
		if _, ok := GroupBy[item.Id]; ok {
			GroupBy[item.Id].Num = GroupBy[item.Id].Num + 1
		} else {
			nums += 1
			GroupBy[item.Id] = &goodcart_goods{Name: item.Good.Name, Num: 1, Price: item.Price}
		}
	}
	var result = make([]goodcart_goods, nums)
	nums = 0
	totalprice := float32(0)
	for _, item := range GroupBy {
		result[nums] = goodcart_goods{Name: item.Name, Num: item.Num, Price: item.Price}
		totalprice += item.Price * (float32)(item.Num)
		nums += 1
	}
	var result_json = make(map[string]interface{})
	result_json["ok"] = true
	result_json["data"] = result
	if len(user.ShopAddress) > 0 {
		result_json["address"] = map[string]string{
			"name":    user.ShopAddress[0].PersonName,
			"phone":   user.ShopAddress[0].Phone,
			"address": user.ShopAddress[0].DetailAddress,
		}
	}
	// 用户下单操作
	var goodDetails = make([]models.OrderDetail, nums)

	ormSession := orm.NewOrm()
	ormSession.Begin()
	userOrder := models.Orders{User: user, Status: 0, TotalPrice: totalprice}

	for index, item := range result {
		goodDetails[index] = models.OrderDetail{
			GoodName: item.Name, GoodNum: item.Num, GoodPrice: item.Price,
			AddressName:   user.ShopAddress[0].PersonName,
			AddressPhone:  user.ShopAddress[0].Phone,
			AddressDetail: user.ShopAddress[0].DetailAddress,
			Order:         &userOrder,
		}
	}

	if _, err := ormSession.Insert(&userOrder); err != nil {
		fmt.Println("orders error", err.Error())
		c.Data["json"] = map[string]interface{}{
			"ok":      false,
			"message": "执行失败",
		}
		return
	}

	if _, err := ormSession.InsertMulti(nums, goodDetails); err != nil {
		ormSession.Rollback()
		c.Data["json"] = map[string]interface{}{
			"ok":      false,
			"message": "执行失败2",
		}
		fmt.Println("order detail error")
		return
	}
	ormSession.Commit()
	c.Data["json"] = map[string]interface{}{
		"ok": true,
	}
}

func (c *IndexController) OrderFilter() {
	defer c.ServeJSON()

	Uid := c.GetSession("session")
	if Uid == nil {
		c.Data["json"] = map[string]interface{}{
			"ok":      false,
			"message": "no auth",
		}
		return
	}
	user := GetUserById(Uid.(int))
	orderType, _ := c.GetInt("orderType")
	var result = make([]*models.Orders, len(user.Orders))
	index := 0
	for _, item := range user.Orders {
		if item.Status == orderType {
			result[index] = item
			index += 1
		}
	}
	c.Data["json"] = map[string]interface{}{
		"ok":   true,
		"data": result[:index],
	}
}

//订单状态更新
type OrderUpdate struct {
	OrderID int
	Status  string
}

func (c *IndexController) OrderStatusUpdate() {
	defer c.ServeJSON()

	var body OrderUpdate
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err == nil {
		if v := GetOrderById(body.OrderID); v != nil {
			var action models.OrderAction

			if body.Status == "orderPay" {
				action.ActionType = 1
				action.OrderId = body.OrderID
			} else if body.Status == "orderComplete" {
				action.ActionType = 2
				action.OrderId = body.OrderID
			}
			if ret, _ := v.Dispatch(action); ret == true {
				c.Data["json"] = map[string]interface{}{
					"ok": true}
				return
			}
		}
	}
	c.Data["json"] = map[string]interface{}{"ok": false}
}
