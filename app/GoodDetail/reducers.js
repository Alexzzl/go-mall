
import { combineReducers } from 'redux'
const init_state = {
    comments: [],
    info: {}
}
function comments_redux(state=init_state, action){
    switch (action.type) {
        case "ADD_COMMENT":
            return Object.assign({}, state, {
                comments:[...state.comments, action.text]
            })
        case "GET_COMMENT":
            return Object.assign({}, state, {
                comments: action.comments
            })
        case "GOOD_DETAIL":
            console.log( action.info)
            return Object.assign({}, state, {
                info: action.info
            })
        default:
            return state
    }



}

const DetailApp = combineReducers({
    comments_redux,
})

export default DetailApp
