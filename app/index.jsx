import React from 'react'
import ReactDOM from 'react-dom';
import { createStore, applyMiddleware } from 'redux'
import { Provider } from 'react-redux'
import App from '../app/containers/Index'
import todoApp from '../app/reducers'
import thunk from 'redux-thunk'


let store = createStore(todoApp, applyMiddleware(thunk))

let rootElement = document.getElementById('root')
ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  rootElement
)
