package models

/*
用户model 定义
*/

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"sync"
	"time"
)

const (
	//商品分类的常量
	MenWear       = 1 // 男装
	WenTrousers   = 2 // 男裤
	WomanWear     = 3 // 女装
	WomanTrousers = 4 // 女裤
	Shoes         = 5 // 鞋子
)

type User struct {
	Id int `orm:"pk;auto"`
	//昵称
	Nickname string `orm:unique`
	//密码不直接暴露出去
	Password string
	//用户头像
	Image string

	Email string `orm:unique`
	//购物车
	ShopCart *ShopCart `orm:"null;reverse(one)"`
	//收藏的商品
	History  *History   `orm:"null;reverse(one)"`
	Comments []*Comment `orm:"reverse(many)"`
	//orders
	Orders      []*Orders      `orm:"reverse(many)"`
	ShopAddress []*ShopAddress `orm:"reverse(many)"`
}

//收货地址
type ShopAddress struct {
	Id            int `orm:"pk;auto"`
	PersonName    string
	DetailAddress string
	Phone         string //收货人的电话号码
	User          *User  `orm:"rel(fk)"`
}

//购物车
type ShopCart struct {
	Id    int `orm:"pk;auto"`
	Name  int
	Goods []*GoodDetail `orm:"null;rel(m2m)"`
	User  *User         `orm:"null;rel(one);on_delete(cascade)"`
}

//收藏历史
type History struct {
	Id    int     `orm:"pk;auto"`
	Goods []*Good `orm:"rel(m2m)"`
	User  *User   `orm:"rel(one);on_delete(cascade)"`
}

//商品
type Good struct {
	Id          int `orm:"pk;auto"`
	Name        string
	Description string
	// 这个价格是不准确的， 具体应该以GoodDetail里的价格为准
	Price float32
	//关注数量
	Follows int64 `orm:default(0)`
	//分类
	TypeCode int32
	//库存剩余
	Stock int64
	//评论
	Comments []*Comment `orm:"rel(m2m)"`
	//型号s,m,
	version int
}

type GoodDetail struct {
	Id   int   `orm:"pk;auto"`
	Good *Good `orm:"rel(one)"`

	// 颜色分类
	ColorType string `orm:"null"`
	// 尺码大小
	Size string `orm:"null"`
	//扩展字段
	Other string `orm:"null"`
	// 同一个商品下不同颜色和尺码属于不同的version
	Version int
	Price   float32
	//图片
	Image string
	// 库存
	Stock int64
	//订单
	//Orders []*Orders `orm:"reverse(many)"`
	// 关注数量
	Follows int64 `orm:"default(0)"`
	// 分类
	TypeCode string
}

type Comment struct {
	Id      int   `orm:"pk;auto"`
	User    *User `orm:"rel(fk)"`
	Content string
	//是直接评论（1）还是回复(2)
	Type_ int `orm:"default(1)"`
	//创建时间
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	//上一条评论
	Comment *Comment `orm:"null;rel(one)"`
}

//秒杀活动
type SecKillGoods struct {
	Id   int         `orm:"pk;auto"`
	Good *GoodDetail `orm:"rel(one)"`
	// 秒杀活动开始的时间
	CreatedTime time.Time `orm:"type(datetime)"`
	// 截止时间
	EndTime time.Time `orm:"type(datetime)"`
	// 优惠之后的价格
	DiscountPrice float32
	// 总的抢购数量
	TotalNum int64
}

/*
订单
支付功能 省略， 提供接口，可以自行扩展实现
订单取消功能 省略，可自行扩展实现
*/

type Orders struct {
	Id   int   `orm:"pk;auto"`
	User *User `orm:"rel(fk)"`
	//Good *GoodDetail `orm:"rel(fk)"`
	// 总价格
	TotalPrice float32
	//具体详情
	OrderDetail []*OrderDetail `orm:"reverse(many)"`
	// 0刚下单但是未支付， 1表示已支付， 2订单完成， 其他更多状态自行扩展
	Status      int       `orm:"default(0)"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`

	reduceFuncPoint reduceFunc
}

//订单详情
type OrderDetail struct {
	Id int `orm:"pk;auto"`
	//商品名字快照
	GoodName string
	//商品数量
	GoodNum int
	//价格
	GoodPrice float32
	// 收货人名字
	AddressName string
	//收货人电话
	AddressPhone string
	// 收货人详细地址
	AddressDetail string
	// 关联的订单
	Order *Orders `orm:"rel(fk)"`
}

type reduceFunc func(int, OrderAction) (int, error)

// 省略redux中的createstore函数， 直接将reduce函数进行绑定
func (order *Orders) registerReduce(reduceFunc reduceFunc) {
	order.reduceFuncPoint = reduceFunc
}

func (order *Orders) GetState() int {
	return order.Status
}
func (order *Orders) Dispatch(action OrderAction) (bool, error) {
	if order.reduceFuncPoint == nil {
		order.registerReduce(reduceStatus) //注册reduceState函数
	}
	//订单状态的修改需要加锁
	lock := new(sync.Mutex)
	lock.Lock()
	defer lock.Unlock()

	newestStatus, _ := order.reduceFuncPoint(order.GetState(), action)
	order.Status = newestStatus
	o := orm.NewOrm()
	if _, err := o.Update(order, "Status"); err == nil {
		return true, nil
	} else {
		fmt.Println("更新状态失败: ", err.Error())
		return false, err
	}
}

// 订单状态发生变化的action
type OrderAction struct {
	// 1表示支付action， 2表示订单完成的action
	ActionType int
	OrderId    int
}

/*
 reduce函数来同一处理订单的状态管理，使状态的变化变得清晰
*/
func reduceStatus(previewStatus int, action OrderAction) (int, error) {
	//当前订单的状态 和 接收action 是有一定的匹配关系的

	switch {
	case previewStatus == 0 && action.ActionType == 1:
		return 1, nil
	case previewStatus == 1 && action.ActionType == 2:
		return 2, nil
	default:
		fmt.Println("订单状态不能发生变化")
		return previewStatus, nil
	}
}

func init() {
	//注册model
	orm.RegisterModel(new(User), new(ShopCart), new(History), new(Good), new(Comment), new(GoodDetail), new(Orders),
		new(SecKillGoods), new(ShopAddress), new(OrderDetail))
}
