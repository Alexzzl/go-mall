package main

import (
	"OnlineShop/controllers"
	_ "OnlineShop/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
在创建数据库时同步执行这个函数
*/
func InitData() {
	orm.RunSyncdb("default", false, true)
    o := orm.NewOrm()

    //添加一个用户
	_, err := o.Raw("insert into user (id, nickname, password, email) values(?, ?, ?, ?)", 1, "test", "123456", "test@qq.com").Exec()
	if err != nil {
		fmt.Println("insert test user error", err)
	}
	//good表
	_, err_ := o.Raw("insert into good (id, name, description,  follows, type_code) value"+
		"(?, ?, ?, ?, ?)", 1, "Iphone 6s", "我喜欢这个牌子的服装", 999, 1).Exec()
	if err_ != nil {
		fmt.Println("insert good error", err_)
	}
	o.Raw("insert into good (id, name, description,  follows, type_code) value"+
		"(?, ?, ?, ?, ?)", 2, "Flatty Phone With Earphone", "good mobile", 888, 1).Exec()

	o.Raw("insert into good (id, name, description,  follows, type_code) value"+
		"(?, ?, ?, ?, ?)", 3, "恶魔的眼泪", "我喜欢这个牌子的服装", 777, 2).Exec()

	o.Raw("insert into good (id, name, description,  follows, type_code) value"+
		"(?, ?, ?, ?, ?)", 4, "希望之星", "我喜欢这个牌子的服装", 666, 2).Exec()

	o.Raw("insert into good (id, name, description,  follows, type_code) value"+
		"(?, ?, ?, ?, ?)", 5, "衣成天品", "我喜欢这个牌子的服装", 999, 5).Exec()
	o.Raw("insert into good (id, name, description,  follows, type_code) value"+
		"(?, ?, ?, ?, ?)", 6, "衣成天品", "我喜欢这个牌子的服装", 999, 6).Exec()
	o.Raw("insert into good (id, name, description,  follows, type_code) value"+
		"(?, ?, ?, ?, ?)", 7, "衣成天品", "我喜欢这个牌子的服装", 999, 7).Exec()

	//good_detail表
	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		1, 1, 1, "/static/img/mobile.jpg", 1234, 4099, "mobile").Exec()

	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		2, 2, 1, "/static/img/mobile2.jpg", 1234, 998, "mobile").Exec()

	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		3, 3, 1, "/static/img/zhubao1.jpg", 1234, 998, "diamonds").Exec()

	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		4, 4, 1, "/static/img/zhubao2.jpg", 1234, 998, "diamonds").Exec()

	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		5, 5, 1, "/static/img/sec1.jpg", 1234, 998, "sec").Exec()
	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		6, 6, 1, "/static/img/sec2.jpg", 1234, 998, "sec").Exec()
	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
        7, 7, 1, "/static/img/sec3.jpg", 1234, 998, "sec").Exec()
    o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
        8, 8, 1, "/static/img/clothes.jpg",1234,499, "clothes").Exec()

	//商品秒杀的数据
	o.Raw("insert into sec_kill_goods (good_id, created_time, end_time, discount_price, total_num) values(?, ?, ?, ?, ?)",
		5, "2019-10-30 09:23:23", "2019-11-30 22:23:23", 500, 100).Exec()
	o.Raw("insert into sec_kill_goods (good_id, created_time, end_time, discount_price, total_num) values(?, ?, ?, ?, ?)",
		6, "2019-10-30 09:23:23", "2019-11-30 22:23:23", 500, 100).Exec()
	o.Raw("insert into sec_kill_goods (good_id, created_time, end_time, discount_price, total_num) values(?, ?, ?, ?, ?)",
		7, "2019-10-30 09:23:23", "2019-11-30 22:23:23", 500, 100).Exec()

	o.Raw(" insert into shop_address (person_name, detail_address, phone, user_id) values(?, ?, ?, ?)",
		"天府吴彦祖", "四川省成都市高新区软件园", "13212345678", 1).Exec()
	//fmt.Println("denbug")
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/onlineshop?charset=utf8mb4&parseTime=True&loc=Local")
}

func main() {
	//StaticDir["/static"] = "static"
	// 第一次启动需要取消掉 InitData 函数的注释，因为第一次运行项目时需要执行 InitData 函数去数据库中创建相关的表，运行成功后就需要重新注释掉这个函数，否则会重复写入初始数据
	InitData()

	beego.SetStaticPath("/dist", "dist")
	// 自行扩展的功能，不能使用原生的beego
	beego.RegisterService(controllers.TradeSer)
	beego.Run()
}
