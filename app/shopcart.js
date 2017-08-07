

import React, {Component} from "react"
import {Table} from "antd"
import ReactDOM from 'react-dom';
import "../app/style/shopcart_main.css"
import {Button} from "antd"
import fetch from 'isomorphic-fetch'

class ShopCart extends  Component {
    constructor(){
        super()
        this.colums = [{
            title: 'Name',
            dataIndex: 'name',
            key: 'name',
            render: text => <a href="#">{text}</a>,
            }, {
            title: 'Price',
            dataIndex: 'price',
            key: 'price',
            }, {
            title: 'Num',
            dataIndex: 'num',
            key: 'num',
            }
        ]
        this.total_price = 0

        this.state = {
            data: [
            ],
            address: {

            },
            total_price : 0,
        }
    }

    componentWillMount(){
        let total_price = 0
         fetch("/api/user/good_cart").then(res => res.json()).then(data => {
               function getdata(data_) {
                console.log("购物车里面的东西在哪里咧: ", data)
                let res =  data_ .map((item, index)=>{
                    total_price += item.Price * item.Num
                    console.log(total_price, item.Price, item.Num)
                    return  {
                        key: index,
                        name: item.Name,
                        price: item.Price,
                        num: item.Num
                    }
                })

                return res
               }

                this.setState({
                     data:    getdata(data.data),
                     address: {
                        name: data.address.name,
                        phone: data.address.phone,
                        address: data.address.address
                     } ,
                     total_price: total_price
                  })
                })
                
    }

    render (){
        let total_price = this.state.total_price
        return (
            <div className="box">
                <div className="title">
                    <span className="span">购物车结算</span>
                </div>
                <div className="main">


                    <div className="address-div">
                        <span className="address-name">{ this.state.address.name }</span>
                        <span className="address-phone">{ this.state.address.phone }</span>
                        <span className="address-detail">{ this.state.address.address } </span>
                    </div>
                    <Table pagination={false} columns={this.colums} dataSource={this.state.data}></Table>

                </div>
                <div className="button-parent">
                        <span className="total-price">
                            总价格: ￥{total_price}
                        </span>
                    <div className="button">
                        <Button>购买</Button>
                    </div>
                </div>
            </div>
        );
    }
}


let rootElement = document.getElementById('root')
ReactDOM.render(<ShopCart/>, rootElement)