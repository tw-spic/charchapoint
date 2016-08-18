import React, { Component } from 'react';
import {
  AppRegistry,
  StyleSheet,
  Text,
  View,
  Navigator
} from 'react-native';

var FirstPage = require('./app/components/FirstPage');
var SecondPage = require('./app/components/SecondPage');

class CharchaPoint extends Component {
  setUpRoute(route) {
    if (route.sceneConfig) {
      return route.sceneConfig;
    }
    return Navigator.SceneConfigs.FloatFromRight;
  }

  renderScene(route, navigator) {
    var routeId = route.id;
    if (routeId === 'FirstPage') {
      return (
        <FirstPage navigator={navigator}/>
      );
    }

    if(routeId === 'SecondPage') {
      return (
        <SecondPage navigator={navigator}/>
      );
    }
  }

  render() {
    return (
      <Navigator
        initialRoute={{id: 'FirstPage', name: 'Index'}}
        renderScene={this.renderScene.bind(this)}
        configureScene={this.setUpRoute.bind(this)}
      />
    );
  }
}

AppRegistry.registerComponent('CharchaPoint', () => CharchaPoint);
