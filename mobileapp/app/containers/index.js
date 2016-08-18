import React, { Component } from 'react';
import { Provider } from 'react-redux';
import { createStore, compose, applyMiddleware } from 'redux';
import ReduxThunk from 'redux-thunk';

var reducer = require('../reducers');
var RandomUserContainer = require('./RandomUserContainer');

const store = createStore(
  reducer,
  compose(
    applyMiddleware(ReduxThunk)
  )
);

class App extends Component {
  render() {
    return (
      <Provider store={store}>
          <RandomUserContainer navigator={this.props.navigator}/>
      </Provider>
    );
  }
}

module.exports = App;
