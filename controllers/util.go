package controllers

import (
	"OnlineShop/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
 获取商品对象
*/
func GetGoodById(good_id int) *models.Good {
	session := orm.NewOrm()
	good := models.Good{Id: good_id}
	err := session.Read(&good)
	if err == nil {
		_, err := session.LoadRelated(&good, "Comments", true)
		if err == nil {
			return &good
		}
	}
	return nil
}

/*
详细商品对象
*/
func GetGoodDetail(good_id int) *models.GoodDetail {
	session := orm.NewOrm()
	good := models.GoodDetail{Id: good_id}
	err := session.Read(&good)
	if err == nil {
		_, err := session.LoadRelated(&good, "Good", true)
		if err == nil {
			return &good
		}
		return &good
	}
	return nil
}

func GetUserByEmail(user_email string) *models.User {
	o := orm.NewOrm()
	user := models.User{Email: user_email}

	if err := o.Read(&user, "Email"); err == nil {
		return &user
	} else {
		return nil
	}
}

func GetUserById(user_id int) *models.User {
	/*
	   因为go指针的缘故，因此可以在这一步做优化, 将user缓存到内存中而不是每次从数据库中去查询,
	   可以优化不少数据库的查询, 因为是指针，也不会导致过期对象的存在
	*/
	o := orm.NewOrm()
	user := models.User{Id: user_id}
	if err := o.Read(&user); err == nil {
		o.LoadRelated(&user, "ShopAddress", true)
		o.LoadRelated(&user, "Orders", true)
		fmt.Println("888888:  ", user.Orders)
		for _, order := range user.Orders {
			o.LoadRelated(order, "OrderDetail", true)
		}
		return &user
	} else {
		return nil
	}
}

/*
获取用户的购物车
*/
func GetUserShopCart(user *models.User) (*models.ShopCart, error) {
	session := orm.NewOrm()
	cart := models.ShopCart{User: user}
	err := session.Read(&cart, "User")
	if err != nil {
		fmt.Println("没有购物车的")
		if _, err := session.Insert(&cart); err == nil {
			return &cart, nil
		} else {
			return nil, errors.New("创建购物车失败")
		}
	} else {
		fmt.Println("找到购物车")
		session.LoadRelated(&cart, "Goods", true)
		return &cart, nil
	}
}

/*
获取用户的收藏夹
*/
func GetUserCollection(user *models.User) (*models.History, error) {
	o := orm.NewOrm()
	history := models.History{User: user}
	err := o.Read(&history, "User")
	if err != nil {
		if _, err := o.Insert(&history); err == nil {
			return &history, nil
		} else {
			return nil, errors.New("创建收藏夹失败")
		}
	} else {
		return &history, nil
	}
}

func CheckoutPassword(user *models.User, OldPassword string) bool {
	if user.Password == OldPassword {
		return true
	}
	return false
}

func UserLogin(user_email string, password string) (bool, *models.User) {
	if user := GetUserByEmail(user_email); user == nil {
		return false, nil
	} else {
		return CheckoutPassword(user, password), user
	}
}

/*
 找到具体的商品信息
*/
func FindGoodDetail(good_id int, version int) *models.GoodDetail {
	o := orm.NewOrm()
	good := GetGoodById(good_id)
	if good == nil {
		return nil
	}
	good_detail := models.GoodDetail{Good: good, Version: version}
	err := o.Read(&good_detail, "Good", "Version")
	if err != nil {
		fmt.Println("没有找到具体的商品信息啊")
		return nil
	}
	return &good_detail
}

/*
将商品加入到购物车
*/
func UserBuyGood(user *models.User, good_id int, version int) (ret bool, err error) {
	o := orm.NewOrm()
	good := FindGoodDetail(good_id, version)
	if good == nil {
		return false, errors.New("商品不存在")
	}
	if shop_cart, err := GetUserShopCart(user); err == nil {
		m2m := o.QueryM2M(shop_cart, "Goods")
		if _, err := m2m.Add(good); err == nil {
			return true, nil
		} else {
			return false, err
		}
	}
	return false, err
}

/*
   加入收藏夹
*/
func GoodJoinCollection(user *models.User, good_id int) (bool, error) {
	o := orm.NewOrm()
	good := GetGoodById(good_id)
	if good == nil {
		return false, errors.New("商品不存在")
	}
	if history, err := GetUserCollection(user); err == nil {
		m2m := o.QueryM2M(history, "Goods")
		if m2m.Exist(good) { //已经收藏过一次那么就是取消收藏
			if _, err := m2m.Remove(good); err == nil {
				return true, nil
			}
		} else { //否者是加入收藏
			if _, err := m2m.Add(good); err == nil {
				return true, nil
			}
		}
	}
	return false, errors.New("收藏夹操作失败")
}

func GetUserShopCartNum(user_id int) (int, error) {
	o := orm.NewOrm()
	var cart models.ShopCart
	if err := o.QueryTable("shop_cart").Filter("User__Id", user_id).One(&cart); err == nil {
		o.LoadRelated(&cart, "Goods")
		return len(cart.Goods), err
	} else {
		return -1, errors.New("该用户不存在")
	}
}

// 喜欢 某个商品
func LikeGood(good_id int, like_status string) (bool, error) {
	o := orm.NewOrm()
	filter_ := orm.ColAdd
	if like_status == "unlike" {
		filter_ = orm.ColMinus
	}
	if _, err := o.QueryTable("good_detail").Filter("Id", good_id).Update(orm.Params{
		"follows": orm.ColValue(filter_, 1)}); err == nil {
		fmt.Println("关注成功")
		return true, nil
	} else {
		fmt.Println("\n\n关注失败: ", err.Error())
		return false, err
	}

}

// 根据商品类型
func FindShopByType(goodType string) []*models.GoodDetail {
	o := orm.NewOrm()
	var GoodDetails []*models.GoodDetail
	_, err := o.QueryTable("good_detail").Filter("TypeCode", goodType).RelatedSel().All(&GoodDetails)
	if err != nil {
		fmt.Println("\n\n 根据类型查询商品", err.Error())
	}
	return GoodDetails

}

//获取搜索的课程
func GetAllGoods(limit int) []*models.GoodDetail {
	o := orm.NewOrm()
	var GoodDetails []*models.GoodDetail
	_, err := o.QueryTable("good_detail").RelatedSel("Good").Limit(limit).All(&GoodDetails)
	if err != nil {
		return GoodDetails
	}
	return GoodDetails
}

//增加一条评论
func AddComment(good_id int, content string, user *models.User) (bool, error) {
	o := orm.NewOrm()
	good := GetGoodById(good_id)
	if good == nil {
		return false, errors.New("无效的good_id")
	}
	m2m := o.QueryM2M(good, "Comments")
	comment := models.Comment{
		User:        user,
		Content:     content,
		Type_:       1,
		CreatedTime: time.Now(),
	}
	err := o.Begin()
	if err != nil {
		return false, errors.New("事务开始失败")
	}
	if _, err := o.Insert(&comment); err != nil {
		return false, err
	}
	if _, err := m2m.Add(&comment); err == nil {
		o.Commit()
		return true, nil
	} else {
		o.Rollback()
		return false, err
	}
}

// 订单信息
func GetOrderById(orderID int) *models.Orders {
	o := orm.NewOrm()
	order := models.Orders{Id: orderID}
	if err := o.Read(&order); err == nil {
		//获取订单详情数据
		o.LoadRelated(&order, "OrderDetail", true)
		return &order
	} else {
		return nil
	}
}
