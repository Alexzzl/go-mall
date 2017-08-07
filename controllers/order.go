package controllers

import (
	"OnlineShop/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func ProduceOrderFact(event ShopEvent) (bool, error) {
	o := orm.NewOrm()
	//user := utils.GetUserById(event.UserId)
	//if user == nil {
	//	return false, errors.New("用户不存在")
	//}
	order_detail := models.OrderDetail{
		GoodName:      event.GoodDetail.Good.Name,
		GoodNum:       1,
		GoodPrice:     event.GoodDetail.Price,
		AddressName:   event.User.ShopAddress[0].PersonName,
		AddressPhone:  event.User.ShopAddress[0].Phone,
		AddressDetail: event.User.ShopAddress[0].DetailAddress,
	}
	orders := []*models.OrderDetail{&order_detail}
	fmt.Println(orders)
	new_order := models.Orders{User: event.User, Status: 0, OrderDetail: orders}

	if _, err := o.Insert(&new_order); err == nil {
		return true, nil
	} else {
		fmt.Println("order insert error", err.Error())
		return false, errors.New("插入数据库失败")
	}

}

type OrderInterface interface {
	//生成一个订单
	ProduceOrder(ShopEvent) (bool, error)
}

type OrderService struct {
}

func (b *OrderService) ProduceOrder(event ShopEvent) (bool, error) {
	return ProduceOrderFact(event)
}

//
type OrderStatusFunc func(OrderID int)

//第一个参数时订单编号， 第二个是订单参数
type Payment interface {
	OrderPay(int, OrderStatusFunc) (bool, error)
}

type WeChartPayment struct {
	//具备支付需要的一些配置
}

func (c WeChartPayment) PaymentAPi() (bool, error) {
	// 具体的支付实现，比如支付宝或者微信的接口调用在这里实现
	return true, nil
}

func OrderFunc(orderID int) {
	order := GetOrderById(orderID)
	action := models.OrderAction{ActionType: 1, OrderId: orderID}
	order.Dispatch(action)
}

func (c WeChartPayment) OrderPay(orderId int, orderFunc OrderStatusFunc) (bool, error) {
	order := GetOrderById(orderId)
	if order == nil {
		return false, errors.New("无效的orderId")
	}
	if _, err := c.PaymentAPi(); err != nil {
		return false, err
	}
	orderFunc(orderId)
	return true, nil
}
