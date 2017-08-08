

import React, {Component} from "react"
import {Radio, Table, Card, Button} from "antd";
const RadioGroup = Radio.Group
import "../../app/style/orderIndex.css"
import {StatusAction} from "./action"
import fetch from 'isomorphic-fetch'

function statusTrans(status) {
    let ret = null;
    if(status == 0) {
        ret = "立即支付"
    } else if (status == 1) {
        ret = "订单完成"
    } else if(status == 2) {
        ret = "订单结束"
    }
    return ret;
}

class OrderIndex2 extends Component {
    constructor() {
        super()

        this.filterHandler = this.filterHandler.bind(this)
        this.colums = [{
            title: '订单编号',
            dataIndex: 'orderIndex',
            key: 'orderIndex',
            render: text => <a href="#">{text}</a>,

        }, {
            title: '订单详情',
            dataIndex: 'orderDetail',
            key: 'orderDetail',
            render: shopDetail => {
                let content = shopDetail.map((item, index)=>{
                    return <p key={index}>{item}</p>
                })
                return <Card title="详情">
                    {content}
                </Card>
            }
        }, {
            title: '总价格',
            dataIndex: 'totalPrice',
            key: 'totalPrice',
        },
            {
                title: "状态",
                dataIndex: "status",
                key: "status",
                render: text => {

                    if (text.status < 2) {
                        return <Button onClick={() => {
                            StatusAction(text.OrderId, text.status)
                        }}> {statusTrans(text.status)}</Button>
                    } else {
                        return <Button disabled>{statusTrans(text.status)}</Button>
                    }

                }
            }
        ]
        this.state = {
            data: [],
            value: -1
        }
    }

    componentWillMount() {

        fetch("/api/user/orders", {
            credentials: 'same-origin'
        }).then(data=>data.json()).then(data=>{
            console.log("aaaaa", data.data)
            let result = data.data.map((item, index) =>{
                return {
                    key: index,
                    orderIndex: item.Id,
                    orderDetail: item.OrderDetail.map(it=>{
                        return "商品:" + it.GoodName + " 数量: " + it.GoodNum
                    }),
                    totalPrice: item.TotalPrice,
                    status: {
                        OrderId: item.Id,
                        status: item.Status
                    }
                }
            })
            console.log("order info: ", result)
            this.setState({
                data: result
            })})

    }
    filterHandler(e){
        let value = e.target.value

        let param = "?orderType=" + value
        console.log("\nparm: ", param)
        fetch("/api/user/orders" + param, {
            credentials: 'same-origin'
        }).then(data=>data.json()).then(data=>{
            console.log("aaaaa", data.data)
            let result = data.data.map((item, index) =>{
                return {
                    key: index,
                    orderIndex: item.Id,
                    orderDetail: item.OrderDetail.map(it=>{
                        return "商品:" + it.GoodName + " 数量: " + it.GoodNum
                    }),
                    totalPrice: item.TotalPrice,
                    status: {
                        OrderId: item.Id,
                        status: item.Status
                    }
                }
            })
            console.log("order info: ", result)
            this.setState({
                data: result,
                value: e.target.value
            })})

    }
    render(){
        return (
            <div className="main-block">
                <div className="title">
                    <span> 订单管理 </span>
                </div>

                <RadioGroup  value={this.state.value} size="large" onChange={this.filterHandler}>
                    <Radio  value={-1}>全部</Radio>
                    <Radio  value={0}>未付款</Radio>
                    <Radio  value={1}>已付款</Radio>
                    <Radio value={2}>已经完成</Radio>
                </RadioGroup>
                <div className="order-table">
                    <Table pagination={false} columns={this.colums} dataSource={this.state.data}></Table>
                </div>
            </div>
        );
    }
}
export default OrderIndex2;