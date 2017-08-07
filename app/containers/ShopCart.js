

import React, {Component, PropTypes} from 'react';
import {connect} from 'react-redux'
import '../../app/style/shop_cart.css'
import {Icon} from 'antd'

class ShopCart extends Component {
    constructor(props){
        super(props)
    }

    render(){
        const icon_style = {
            fontSize: "29px",
            color: "#ffffff",
            position: "absolute",
            zIndex:50,
            left: "5.5px",
            top: "5px"
        }
        const {GoodNum} = this.props
        return (
             <div>
                 <a href="/user/shop_cart">
                    <p className="shop-icon">
                        <Icon style={icon_style} type="shopping-cart"></Icon>
                        {
                            GoodNum.GoodNum != -1?
                            <span className="shop-num">{GoodNum.GoodNum}</span>
                                :null
                        }
                    </p>
                 </a>
             </div>
        );
    }
}

function select(state) {
    return {
        GoodNum: state.shop_cart_manager
    }
}
// 包装 component ，注入 dispatch 和 state 到其默认的 connect(select)(App) 中；
export default connect(select)(ShopCart)