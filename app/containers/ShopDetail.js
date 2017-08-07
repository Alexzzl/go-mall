

import React, {Component, PropTypes} from 'react';
import '../../app/style/shop_detail.css'
import {Button, Icon} from "antd";
import {connect} from 'react-redux'
import {add_shopping_cart, modal_show, shop_filter} from '../../app/action'
import fetch from 'isomorphic-fetch'


class ShopListDetail extends  Component{
    constructor(props) {
        super(props)
    }
    componentWillMount() {
        console.log("\nproblem has occer here")
        let {dispatch} = this.props
        dispatch(shop_filter("all"))
    }

    add_shopping_cart(good_id){
        let {dispatch} = this.props
        dispatch(add_shopping_cart(good_id, null))
        /* 每一个商品都有不用的类型，比如规格颜色等不一样通过version来区分
         */
        let request_data = {
            GoodId: good_id,
            Version: 1
        }
        fetch("/api/good/joincart", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json, text/plain, */*'
                },
                body: JSON.stringify(request_data),
                credentials: 'same-origin'
            }
        )

    }
    compose_shop_detail(){
        let {dispatch, GoodInfo} = this.props
        let shop_list = GoodInfo.goods
        console.log("shoopppp list: "+ shop_list.length)
        if (shop_list.length > 0) {
            return shop_list.map((item, index) => {
                return <ShopDetail key={index} name={item.Good.Name} image={item.Image} price={item.Price}
                                   good_id={item.Id} stock={item.Stock} likeNum={item.Follows}
                                   JoinCart={() => this.add_shopping_cart(item.Id)}
                ></ShopDetail>
            })
        } else {
            return <p>该分类下暂时还没有商品哦</p>
        }
    }

    render (){
        console.log("lllllllllllllllllllist: ")
        let forms = this.compose_shop_detail()
        return (
          <div>
              {forms}
          </div>
        );
    }
}
function good_like_handle(good_id, like) {
    let url = '/api/good/' + good_id + '/like?status=' + like
    fetch(url, {
        credentials: 'same-origin',
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json, text/plain, */*'
        },
        body: JSON.stringify({Status: like}),
    })
}
class ShopDetail extends Component {
    constructor(props) {
        super(props)
        this.state = {
            like_status: 'like',
            like_nums: props.likeNum
        }
        this.like_status = this.like_status.bind(this)
    }

    like_status(old_status){
        let {good_id} = this.props
        if(old_status == 'like') {
            this.setState({
                like_status: "unlike",
                like_nums: this.state.like_nums + 1
            })
            good_like_handle(good_id, "like")
        } else {
            this.setState({
                like_status: "like",
                like_nums: this.state.like_nums - 1
            })
            good_like_handle(good_id, "unlike")
        }
    }
    render() {
        const {name, price, image, stock, JoinCart, good_id, likeNum} = this.props
        //const name = "Iphone6s", image="https://bosnaufal.github.io/vue-mini-shop/assets/img/mobile.jpg"
        console.log("imggggggg+"+ image)
        return (
                <div className="shop-detail-container">
                    <div className="image">
                        <img src={image}/>
                    </div>
                    <div className="shop-detail-desc">
                        <span className="name">{name}</span>

                        <div className="description">
                            <span className="price">￥{price}</span>
                            <span className="stock"> Stocks: {stock}</span>
                        </div>
                    </div>
                    <div className="action">
                        <div className="action-like">
                            <Button style={{backgroundColor: 'red', height: "38px"}}
                                    size='large' onClick={()=>this.like_status(this.state.like_status)}>
                                {this.state.like_status == 'like'?
                                    < Icon type="like-o" style={{color: "#ffffff"}}></Icon>
                                    :
                                    < Icon type="dislike-o" style={{color: "#ffffff"}}></Icon>
                                }
                                <span style={{color: "#ffffff"}}>{this.state.like_nums}</span></Button>
                        </div>
                        <div className="action-share">
                            <Button style={{backgroundColor: 'green', height: "38px", width: "74px"}}
                                    size='large'>
                                <Icon type="share-alt" style={{color: "#ffffff"}}></Icon>
                                <span style={{color: "#ffffff"}}>分享</span>
                            </Button>
                        </div>
                        <div className="action-cart">
                            <Button style={{backgroundColor: 'blue', height: "38px"}}
                                    onClick={ ()=> JoinCart()} size='large'>
                                <Icon type="shopping-cart" style={{color: "#ffffff"}}></Icon>
                                <span style={{color: "#ffffff"}}>加入购物车</span>
                            </Button>
                        </div>
                    </div>
                </div>
        );
    }
}
function select(state) {
    return {
        //GoodNum: state.shop_cart_manager
        GoodInfo:  state.shop_filter_redux
    }
}
// 包装 component ，注入 dispatch 和 state 到其默认的 connect(select)(App) 中；
export default connect(select)(ShopListDetail)
