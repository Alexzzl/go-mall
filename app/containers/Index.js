import React, { Component, PropTypes } from 'react'
import { connect } from 'react-redux'
import {Icon, Affix, Button, Input} from 'antd'
import classNames from 'classnames'

import ShopShow from '../../app/containers/MidType'
import UserInfoComponent from '../../app/containers/UserInfo'
import ShopListDetail from '../../app/containers/ShopDetail'
import  ShopCart from  '../../app/containers/ShopCart'

import "../../app/style/index.css"

import {searchAction_, AsyncGetUserInfo } from '../../app/action'

const Search = Input.Search;

class App extends  Component{
    constructor(props){
        super(props);
        this.on_fcourse= this.on_fcourse.bind(this)
    }

    on_fcourse (display_mode) {
        let {dispatch} = this.props
        //console.log("current state: " + display_mode, dispatch)
        dispatch(searchAction_(display_mode))
    }


    render(){
        const { dispatch, goods, shopcart_count, current_page, display_mode, modal_status, user_info} = this.props
        const shop_type = [
            {
                "name": "手机",
                type_code: "mobile"

            },
            {
                "name": "钻石",
                type_code: "diamonds"
            },
            {
                "name": "服装",
                type_code: "clothes"
            },
            {
                "name": "水果",
                "link": 'fegeg'
            },
            {
                "name": "家居",
                "link": 'fegeg'
            },
            {
                "name": "日化",
                "link": 'fegeg'
            },
            {
                "name": "干果",
                "link": 'fegeg'
            },
            {
                "name": "茶叶",
                "link": 'fegeg'
            },
            {
                "name": "玩具",
                "link": 'fegeg'
            },
            {
                "name": "美妆",
                "link": 'fegeg'
            },
        ]

        const ShopList = [
            {
                name: "Iphone 8",
                price: 5800,
                image: "https://bosnaufal.github.io/vue-mini-shop/assets/img/mobile.jpg",
                stock: 344,
                good_id: 1,
            },
            {
                name: "Iphone SE",
                price: 2300,
                image: "https://bosnaufal.github.io/vue-mini-shop/assets/img/mobile.jpg",
                stock: 344,
                good_id: 2
            },
        ]
        //var is_login = user_info.login_status;
        //console.log("sattus", is_login)
        return (
            <div className='main_block'>
                <div className='index_search_block'>
                    <Affix style={{ position: 'fixed', top: 0,  backgroundColor: "rgb(225, 48, 50)"}}>
                        <div className='affix_search'>
                            <div style={{width: 53, height: 35, float: "left"}}>
                                {display_mode == 'search_mode'?
                                        <Button icon='close' style={{marginTop: 3, marginLeft: 10}}
                                                onClick={()=>this.on_fcourse(display_mode)}
                                        >
                                        </Button>
                                    :<ShopCart></ShopCart>
                                }
                            </div>
                            <div style={{float: "left"}}>
                            <Search
                                placeholder="热门商品"
                                style={{ width: 500, marginLeft: 0, height: 35}}
                                onSearch={value => console.log(value)}
                                onFocus={() => this.on_fcourse(display_mode)}
                            />
                            </div>
                        </div>
                        <UserInfoComponent/>
                     </Affix>

                    <div style={{ clear:'both' }}></div>
                </div>

                {display_mode == 'show_mode'?
                    <div>
                        <ShopShow shop_type={shop_type}/>
                        <ShopListDetail />
                    </div>
                    :
                    <div style={{height: 900}}>
                        <p>你想搜索什么？</p>
                    </div>
                    }

            </div>
        )
    }
}

App.propTypes = {
    //商品
    goods: PropTypes.arrayOf(PropTypes.shape({
        name: PropTypes.string.isRequired,
        description: PropTypes.string.isRequired,
        price: PropTypes.number.isRequired,
        imgurl: PropTypes.string.isRequired,
        follows: PropTypes.number.isRequired,
        tag: PropTypes.string.isRequired,
        stock: PropTypes.number.isRequired
    }).isRequired).isRequired,
    //购物车数量
    shopcart_count: PropTypes.number.isRequired,
    //当前翻页数量
    current_page: PropTypes.number.isRequired,
    display_mode: PropTypes.string.isRequired,
    hot_courses: PropTypes.arrayOf(PropTypes.string)
};

function select(state) {
    return state.index_manager
}

// 包装 component ，注入 dispatch 和 state 到其默认的 connect(select)(App) 中；
export default connect(select)(App)
