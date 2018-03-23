import React from 'react'
import ReactDOM from 'react-dom'
import App from './components/App/AppContainer'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import { configureStore } from './configStore'
import { Provider } from 'react-redux'


const store = configureStore()


ReactDOM.render(
  <Provider store={store}>
    <Router>
      <Route path='/' component={App} />
    </Router>
  </Provider>
  , document.getElementById('root'))