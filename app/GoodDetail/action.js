

import fetch from 'isomorphic-fetch'

//获取所有的评论
function GetAllCommentsAction(comments) {
    return {
        type: "GET_COMMENT",
        comments
    }
}

export function GetAllComments(good_id) {
    return dispatch =>{
        return dispatch(dispatch_=> {
            fetch("/api/good/" + good_id + '/comments').then(res=>res.json()).then(data=>{
                console.log("fetch process + "+ data.data)
                let comments = data.data.map(item=>{
                    return {
                        comment: item.Comment,
                        user_name: item.Nickname,
                        avator: item.Image,
                        time: item.CreatedTime
                    }
                })
                dispatch_(GetAllCommentsAction(comments))
            })
        })
    }
}


//添加评论
function AddCommentAction(text) {
    return {
        type: "ADD_COMMENT",
        text
    }
}

export function AddComment(good_id, text) {
    return dispatch =>{
        return dispatch(dispatch_=>{
            let data = {Comment: text}
            fetch("/api/good/" + good_id + "/comments", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json, text/plain, */*'
                },
                body: JSON.stringify(data),
                credentials: 'same-origin'
            }).then(res=>res.json()).then(data=>{

                if(data.ok == true){
                    let date = new Date()
                    let currentdate = date.getHours() + ":" + date.getMinutes()
                        + ":" + date.getSeconds();
                    let info = {
                        comment: text,
                        user_name: data.data.Username,
                        avator: data.data.Image,
                        time: currentdate
                    }
                    dispatch_(AddCommentAction(info))
                } else {
                    return
                }
            })
        })
    }
}

function GetDetailAction(info){
    return {
        type: "GOOD_DETAIL",
        info
    }
}
export function GetDetail(good_id) {
    return dispatch=>{
        return dispatch(dispatch_=>{
            fetch("/api/good/"+ good_id + '/info').then(
            res=>res.json()).then(d=>{
                let j = d.data
                let info = {
                    name: j.Good.Name,
                    description: j.Good.Description,
                    price: j.Good.Price,
                    follows: j.Follows,
                    stock: j.Stock,
                    image: j.Image
                }
                dispatch_(GetDetailAction(info))
            })        
        })
    }
}
