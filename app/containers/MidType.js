
import React, {Component, PropTypes} from 'react';

import {connect} from "react-redux"
import Time from '../../app/containers/time'
import '../../app/style/show_content.css'
import {Button} from "antd";
import {shop_filter, sec_shop} from '../../app/action'


class ShopShow extends Component {
    constructor(props) {
        super(props)
    }

    shop_filter(type_code) {
        // 商品筛选
        let {dispatch} = this.props
        dispatch(shop_filter(type_code))
    }
    componentWillMount(){
        let {dispatch} = this.props
        dispatch(sec_shop())
    }

    sec_goods_render(goods){
        console.log("\n\n\nthis--------")
        console.log(goods[0])
        let can_sectime_buy = true;
        if (goods.length > 0) {
            let now = Date.parse(new Date())
            let end = Date.parse(new Date(goods[0].EndTime))

            if (now > end)
                can_sectime_buy = false
        }
        return goods.map((item, index)=> {

                console.log("\n\n\n1224535464g5u5hjg4g4g34g4g43fg4")
                return(
                <li key={index} className="sectime-item-li">
                    <img className="sectime-item-li-img" src={item.Good.Image}></img>
                    <span className="sectime-item-li-span">${item.Good.Price}</span>
                    <span className="sectime-item-li-span-origin"> ${item.DiscountPrice}</span>
                    {
                        can_sectime_buy ?
                            <Button>立即抢购</Button>
                            :
                            <Button disabled>立即抢购</Button>
                    }
                </li>)
            }
        )
    }
    render() {
        const {dispatch, shop_type, sec_shop} = this.props

        let sec_goods_com = this.sec_goods_render(sec_shop.sec_goods)
        console.log("\n\n\nwhatis this")
        console.log(sec_goods_com)
        const pic_url = 'https://m.360buyimg.com/mobilecms/s80x80_jfs/t5773/256/159664156/17475/742fec7e/591d9475Na810c2eb.png'
        let end_time = Date.parse(new Date());
        if (sec_shop.sec_goods.length > 0 ){
            end_time = Date.parse(new Date(sec_shop.sec_goods[0].EndTime))
        }
        const can_sectime_buy = false;
        let buy_classname = null;
        if(can_sectime_buy){
            buy_classname = 'can_buy-btn'
        }
        else{
            buy_classname = 'can-not-buy-btn'
        }

        return (
            <div className="border">
                <div className="main-block">
                    <nav className="shop-type">
                        {
                            Array.from(shop_type.keys()).map((ele) => (
                                <a href="javascript:void(0)" onClick={()=> this.shop_filter(shop_type[ele].type_code)} key={ele} className="shop-type-detail">
                                    <div className="quick-box">
                                        <img src={pic_url} style={{width: 58, height: 58}}></img>
                                        <span style={{color: '#000000'}}>{shop_type[ele].name}</span>
                                    </div>
                                </a>)
                            )
                        }
                    </nav>
                </div>
                <div>
                    <div className="hot-goods">
                        <div style={{ display: 'inline-flex'}}>
                            <img className="sectime-img" src="https://m.360buyimg.com/mobilecms/jfs/t3451/307/73678054/7807/dd9134d/57fdff2eNb7cd186f.png"></img>
                        </div>
                        <div style={{display: 'inline-flex'}}>
                            <Time end_time={end_time} />
                        </div>

                    </div>
                    <div className="sectime-item">
                        <ul className="sectime-item-ul">
                            {sec_goods_com}
                        </ul>

                    </div>
                </div>
            </div>
        );
    }
}

function select(state) {
    return {
        sec_shop: state.sec_shop_redux
    }
}
// 包装 component ，注入 dispatch 和 state 到其默认的 connect(select)(App) 中；
export default connect(select)(ShopShow)