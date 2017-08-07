

import React, { Component, PropTypes } from 'react'
import {connect} from 'react-redux'
import '../../app/style/shop_detail.css'
import {GetDetail, GetAllComments, AddComment}  from "../../app/GoodDetail/action"

import {Button} from 'antd'

class App extends  Component {

    constructor(props) {
        // 没有使用route库，直接操作location来进行跳转
        super(props);
        let path = location.pathname
        let re = /good\/(\d+)/
        let result = path.match(re)
        if (result != null) {
            this.good_id = result[1]
        } else {
            location.href = "/"
        }
    }
    componentWillMount() {
        let {dispatch} = this.props
        dispatch(GetDetail(this.good_id))
        dispatch(GetAllComments(this.good_id))
    }
    render_comment_content(){
        let {comments} = this.props
        return comments.comments.map((item, index) => {
            return <Comments key={index} avator={item.avator}  username={item.user_name}
                             comment={item.comment} time={item.time}></Comments>
        })
    }

    render() {
        let obj = {
            image: this.props.comments.info.image, 
            name: this.props.comments.info.name, 
            price: this.props.comments.info.price, 
            stock: this.props.comments.info.stock
        }
        let CommentContent = this.render_comment_content()

        let {image, name, price, stock} = obj
        return (
            <div className="main-block-detail">
                <div className="shop-detail-container-detail">
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
                </div>

                <div className="good-description">
                    <p>
                        This is the details of the Iphone, Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                        sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
                        quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                        Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
                        pariatur.
                        Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit
                        anim
                        id est laborum.
                        <br/>
                        <br/>
                        Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt
                        ut
                        labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco
                        laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in
                        voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat
                        cupidatat
                        non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
                    </p>
                </div>
                <div className="comments">
                    <div className="title">
                        <span>{this.props.comments.comments.length} Comments</span>
                    </div>

                    {CommentContent}
                    <CommentSubmit dispatch={this.props.dispatch}></CommentSubmit>

                </div>

            </div>

        );
    }
}


class Comments extends Component {
    render(){
        return (
            <div className="content">
                <span className="avator"><img src=""></img></span>
                <span className="name">{this.props.username}</span>
                <span className="detail-content">{this.props.comment}</span>
                <span className="update_time">{this.props.time}</span>
            </div>
        );
    }
}

class CommentSubmit extends Component {
    constructor(props){
        super(props)
        this.submit_comment = this.submit_comment.bind(this)
    }

    submit_comment(){
        // textarea很奇怪不用通过refs来获取值，和input有区别
        let {dispatch} = this.props
        let content = this.comment_content.value
        dispatch(AddComment(1, content))
    }

    render(){
        return (
         <div className="submit">
             <div style={{paddingTop: "20px"}}>
                <span className="title">Leave a Comment</span>
             </div>
             <div className="div">
                <textarea  className="content" ref={(input)=>this.comment_content = input} placeholder="Write a comment here..."></textarea>
             </div>

             <Button type="submit" style={{width: "100%", height: "50px", backgroundColor: "rgb(9, 76, 165)"}} onClick={()=>this.submit_comment()}>Submit</Button>
         </div>
        )}
}


function state_select(state){
    return {comments: state.comments_redux}
}
export default connect(state_select)(App);
