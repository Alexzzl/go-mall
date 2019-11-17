// export const ADD_TODO = 'ADD_TODO';
// export const COMPLETE_TODO = 'COMPLETE_TODO';
// export const SET_VISIBILITY_FILTER = 'SET_VISIBILITY_FILTER'

/*
 * 其它的常量
 */

// export const VisibilityFilters = {
//   SHOW_ALL: 'SHOW_ALL',
//   SHOW_COMPLETED: 'SHOW_COMPLETED',
//   SHOW_ACTIVE: 'SHOW_ACTIVE'
// };

/*
 * action 创建函数
 */

// export function addTodo(text) {
//   return { type: ADD_TODO, text }
// }
//
// export function completeTodo(index) {
//   return { type: COMPLETE_TODO, index }
// }
//
import fetch from 'isomorphic-fetch'

export function setVisibilityFilter(filter) {
  return { type: SET_VISIBILITY_FILTER, filter }
}

export const DISPLAY_MODE = 'DISPLAY_MODE'  //type 值
export const LOGIN_STATUS = 'LOGIN_STATUS' //用户状态
export const USER_LOIN = 'USER_LOGIN' //用户进行登录

//进行搜索模式还是显示模式
export function searchAction_(status) {
  return {
        type: DISPLAY_MODE,
        status
  }
};
// 用户信息的获取同时判断登录状态
export function GetUserInfoAction(user_info) {
    return {
        type: LOGIN_STATUS,
        user_info
    }
}
function FetchGetUserInfo() {
    return dispatch => {
        return fetch("/api/user_info", {
                credentials: 'same-origin'
            })
            .then(response => response.json())
            .then(json => {
                    let user_info = {
                        username: json.ok == true ?json.data.UserName: "",
                        image: json.ok == true ?json.data.image: "",
                        Userid: json.ok == true ? json.data.UserId:"",
                        login_status: json.ok
                    }
                    console.log("sttttt" + user_info.login_status)
                    if ( json.ok == true) {
                        user_info.GoodNum = json.data.ShopCartNum
                    }
                    console.log("\npostion", user_info)
                    dispatch(GetUserInfoAction(user_info))
                }
            )
    }
}
export function AsyncGetUserInfo(login_status) {
    return dispatch => {
        const current_stats = login_status
        if (current_stats == false){
            return dispatch(FetchGetUserInfo())
        } else {
            return Promise.resolve()
        }

    }
}

// 用户进行登录操作
export function UserLoginAction(user_info) {
    return {
        type: USER_LOIN,
        user_info
    }
}

export function AsyncUserLogin(username, password) {
    return (dispatch, Getstate) => {
        return dispatch(dispatch2 => {
            let request_data = {
                'UserName': username,
                'Password': password}
            fetch("/api/user/login", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json, text/plain, */*'
                },
                body: JSON.stringify(request_data),
                credentials: 'same-origin'
            })
                .then(res => res.json())
                .then(json => {
                    let user_info;
                    console.log("00000000000")
                    if (json.ok == true){
                        user_info = {
                            username: json.data.UserName,
                            image: json.data.image,
                            login_status: true,
                            GoodNum: json.data.ShopCartNum,
                            modal_status: false,
                            err_message:false,
                            UserId: json.data.UserId

                        }
                    } else {
                        user_info = {
                            username: "",
                            image: "",
                            login_status: false,
                            err_message: true,
                            modal_status: true,
                            GoodNum: -1,
                            UserId: ""
                        }
                    }
                    console.log("\n\nsusesss" + user_info.login_status)
                dispatch2(UserLoginAction(user_info))
                })
        })
    }
}

export function modal_show(status) {
    return {
        type: "SHOW_MODAL",
        status
    }
}

//购物车数量变化
export function add_shopping_cart(GoodId, Num) {
    return {
        type: "SHOP_CART",
        GoodId, Num
    }
}

// 商品筛选
function shop_filter_action(goods) {
    return {
        type: "SHOP_TYPE",
        goods
    }
};

export function shop_filter(type_code) {
    console.log("type_code" + type_code)

    return (dispatch) => {
        return dispatch(dispatch2 => {

            fetch("/api/good/good_type?" + "good_type=" + type_code,
            )
                .then(res => res.json())
                .then(data => {
                    console.log("shop_filter: "+ data.data)
                    dispatch2(shop_filter_action(data.data))
                })

        })
    }
}

//商品秒杀
function sec_shop_action(sec_goods){
    return {
        type: "SEC_SHOP",
        sec_goods
    }
}
export function sec_shop() {
    return (dispatch) => {
        return dispatch(dispatch2 => {
            fetch("/api/good/seckill").then(res => res.json()).then(data => {
                    console.log("api--------------", data.data)
                    dispatch2(sec_shop_action(data.data))
                })

        })
    }
}