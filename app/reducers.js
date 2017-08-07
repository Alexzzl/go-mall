import { combineReducers } from 'redux'
// import { ADD_TODO, COMPLETE_TODO, SET_VISIBILITY_FILTER, VisibilityFilters } from '../app/action'
// const { SHOW_ALL } = VisibilityFilters
//
// function visibilityFilter(state = SHOW_ALL, action) {
//   switch (action.type) {
//     case SET_VISIBILITY_FILTER:
//       return action.filter
//     default:
//       return state
//   }
// }
//
// function todos(state = [], action) {
//   switch (action.type) {
//     case ADD_TODO:
//       return [
//         ...state,
//         {
//           text: action.text,
//           completed: false
//         }
//       ]
//     case COMPLETE_TODO:
//       return [
//         ...state.slice(0, action.index),
//         Object.assign({}, state[action.index], {
//           completed: true
//         }),
//         ...state.slice(action.index + 1)
//       ]
//     default:
//       return state
//   }
// }

import {DISPLAY_MODE } from '../app/action'

const instance = {
    display_mode: "show_mode", //展示模式
    hot_courses: ["Python", "Go"], //热门课程
    goods: [
        {
            name: "Iphone 5S",
            description: "这是一款好手机",
            price: 32.5,
            imgurl: "http://www.baidu.com",
            follows: 34,
            tag: "手机",
            stock: 344
        }
    ],
    shopcart_count: 0,
    current_page: 1,
    modal_status: false,
}

function trancs_status(status) {
    if (status == 'show_mode'){
        return "search_mode";
    }
    return "show_mode"
}

function index_manager(state=instance , action) {

  switch (action.type){
      case "DISPLAY_MODE":  //展示模式的切换
        return Object.assign({}, state, {
            hot_courses: ["Python", "Go"],
            display_mode: trancs_status(action.status)
        })
      default:
        return state
  }
}

const good_cart = {
    modal_status: false,
    login_status: false,
    username: "",
    image: "",
    err_message: false,
    UserId: "",
    GoodNum: -1 //购物车中商品数量
}

function shop_cart_manager(state=good_cart, action) {
    switch (action.type) {
        case "SHOP_CART":
            let nums = state.GoodNum
            if (action.GoodId == null){
                nums = action.Num
            } else if (nums != undefined) {
                nums += 1
            } else {
                nums = 1
            }
            return Object.assign({}, state, {
                GoodNum: nums
            })
        case "USER_LOGIN":

            return Object.assign({}, state, {
                GoodNum: action.user_info.GoodNum,
                modal_status: action.modal_status,
                login_status: action.login_status,
                username: action.user_info.username,
                image: action.user_info.image,
                UserId: action.user_info.UserId
            })
        case "LOGIN_STATUS":
            return Object.assign({}, state, {
                GoodNum: action.user_info.GoodNum,
                login_status: action.user_info.login_status,
                username: action.user_info.username,
                image: action.user_info.image,
                UserId: action.user_info.UserId
            })

        case "SHOW_MODAL":
            let new_status = false
            if (action.status == false){
                new_status = true
            }
            return Object.assign({}, state, {
                modal_status: new_status
            })
        default:
            return state
    }
}


const shop_type_initstate = {
    goods:[]
}

// 商品类型筛选
function shop_filter_redux(state=shop_type_initstate, action) {
    switch (action.type){
        case "SHOP_TYPE":
            return Object.assign({}, state, {
                goods: action.goods
            })
        default:
            return state
    }
}

const sec_shops = {
    sec_goods: []
}
function sec_shop_redux(state=sec_shops, action) {
    switch (action.type){
        case "SEC_SHOP":
            return Object.assign({}, state, {
                sec_goods: action.sec_goods
            })
        default:
            return state
    }
}
const todoApp = combineReducers({
    index_manager,
    shop_cart_manager,
    shop_filter_redux,
    sec_shop_redux
})

export default todoApp
