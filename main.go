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
	_, err := o.Raw("insert into user (id, nickname, password, email) values(?, ?, ?, ?)", 1, "zxl", "0000", "zz@qq.com").Exec()
	if err != nil {
		fmt.Println("11111", err)
	}
	//good表
	_, err_ := o.Raw("insert into good (id, name, description,  follows, type_code) value"+
		"(?, ?, ?, ?, ?)", 1, "Iphone 6s", "我喜欢这个牌子的服装", 999, 1).Exec()
	if err_ != nil {
		fmt.Println("333", err_)
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
		1, 1, 1, "http://127.0.0.1:8080/static/img/mobile.jpg", 1234, 4099, "mobile").Exec()

	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		2, 2, 1, "http://127.0.0.1:8080/static/img/mobile2.jpg", 1234, 998, "mobile").Exec()

	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		3, 3, 1, "http://127.0.0.1:8080/static/img/zhubao1.jpg", 1234, 998, "diamonds").Exec()

	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		4, 4, 1, "http://127.0.0.1:8080/static/img/zhubao2.jpg", 1234, 998, "diamonds").Exec()

	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		5, 5, 1, "http://127.0.0.1:8080/static/img/sec1.jpg", 1234, 998, "sec").Exec()
	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		6, 6, 1, "http://127.0.0.1:8080/static/img/sec2.jpg", 1234, 998, "sec").Exec()
	o.Raw("insert into good_detail (id, good_id, version, image, stock, price, type_code) values (?, ?, ?, ?, ?, ?, ?)",
		7, 7, 1, "http://127.0.0.1:8080/static/img/sec3.jpg", 1234, 998, "sec").Exec()

	//商品秒杀的数据
	o.Raw("insert into sec_kill_goods (good_id, created_time, end_time, discount_price, total_num) values(?, ?, ?, ?, ?)",
		5, "2017-07-28 14:23:23", "2017-07-25 22:23:23", 500, 100).Exec()
	o.Raw("insert into sec_kill_goods (good_id, created_time, end_time, discount_price, total_num) values(?, ?, ?, ?, ?)",
		6, "2017-07-28 14:23:23", "2017-07-25 22:23:23", 500, 100).Exec()
	o.Raw("insert into sec_kill_goods (good_id, created_time, end_time, discount_price, total_num) values(?, ?, ?, ?, ?)",
		7, "2017-07-28 14:23:23", "2017-07-25 22:23:23", 500, 100).Exec()

	o.Raw(" insert into shop_address (person_name, detail_address, phone, user_id) values(?, ?, ?, ?)",
		"天府吴彦祖", "四川省成都市高新区软件园", "13212345678", 1).Exec()
	fmt.Println("denbug")
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:zxl@tcp(127.0.0.1:3306)/onlineShop")

}

func main() {
	//StaticDir["/static"] = "static"
	//InitData()

	beego.SetStaticPath("/dist", "dist")
	// 自行扩展的功能，不能使用原生的beego
	beego.RegisterService(controllers.TradeSer)
	beego.Run()
}
