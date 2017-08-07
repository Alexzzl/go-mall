

import React from 'react'
import ReactDOM from 'react-dom';
import { createStore, applyMiddleware } from 'redux'
import { Provider } from 'react-redux'

import App from '../app/GoodDetail/Index'
import DetailApp from '../app/GoodDetail/reducers'
import thunk from 'redux-thunk'

let store = createStore(DetailApp, applyMiddleware(thunk))

let rootElement = document.getElementById('root')
ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>,
    rootElement
)
