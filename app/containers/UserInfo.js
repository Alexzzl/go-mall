import React, { Component, PropTypes } from 'react'
import {connect} from 'react-redux'
import {Modal, Form, Icon, Input, Button, Checkbox} from 'antd'
const FormItem = Form.Item;
import fetch from 'isomorphic-fetch'

import { add_shopping_cart, AsyncGetUserInfo, AsyncUserLogin, modal_show} from '../../app/action'

import '../../app/style/user_info.css'

class UserInfoComponent extends Component{
    constructor(props){
        super(props)
        this.LoginHandler = this.LoginHandler.bind(this)
        this.LogionEvent = this.LogionEvent.bind(this)
    }
    // 用户登录
    LoginHandler(){
        // this.setState({modal_status: true})
        let {dispatch,  GoodNum} = this.props
        dispatch(modal_show(GoodNum.modal_status))
        return false;
    }

    LogionEvent(e){
        e.preventDefault();

        let {dispatch, modal_status} = this.props
        dispatch(AsyncUserLogin(this.username.refs.input.value, this.password.refs.input.value))
        return false;
    }
    cancel () {
        let {dispatch, GoodNum} = this.props
        dispatch(modal_show(GoodNum.modal_status))

    }
    componentWillMount() {
        let {dispatch, GoodNum} = this.props
        dispatch(AsyncGetUserInfo(GoodNum.login_status))
    }

    render () {
        const {GoodNum} = this.props
        let user_info_url = ""
        let login_status = GoodNum.login_status
        //console.log("start 222222render", GoodNum.login_status)
        if(login_status == true) {
            user_info_url = '/user/' + GoodNum.UserId
        }
        return (
            <div className='affix_user'>
                {login_status == false ?
                    <a href="javascript:void(0)" onClick={this.LoginHandler} className='login_btn'>登录</a>
                    :
                    <a href={user_info_url}> <span className="login-icon"></span></a>
                }
                <Modal title="Title"
                       visible={GoodNum.modal_status}
                       confirmLoading={false}
                       footer={null}
                       onCancel={()=>this.cancel()}
                >

                    <Form onSubmit={this.LogionEvent} className="login-form">
                        <FormItem>
                                <Input ref={(input)=>this.username = input} prefix={<Icon type="user" style={{ fontSize: 13 }} />} placeholder="Username" />
                        </FormItem>
                        <FormItem>

                                <Input ref={(input)=> {this.password = input}} prefix={<Icon type="lock" style={{ fontSize: 13 }} />} type="password" placeholder="Password" />
                        </FormItem>
                        <FormItem>
                            <Button type="primary" htmlType="submit" className="login-form-button">
                                Login
                            </Button>
                        </FormItem>
                    </Form>
                    {
                        this.props.err_message == true?
                        <p> 登录失败</p>:null
                    }

                </Modal>
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
export default connect(select)(UserInfoComponent)