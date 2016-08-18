import React, { Component } from 'react';
import {
  AppRegistry,
  StyleSheet,
  Text,
  View,
  Navigator,
  TouchableNativeFeedback
} from 'react-native';

class FirstPage extends Component {
  constructor(props) {
    super(props);
  }

  navigateToSecondPage() {
    this.props.navigator.push({id: 'SecondPage'});
  }

  render() {
    return (
      <View>
        <Text>First Page</Text>
        <TouchableNativeFeedback
          onPress={this.navigateToSecondPage.bind(this)}
          background={TouchableNativeFeedback.SelectableBackground()}>
            <View>
              <Text>Go to Second Page</Text>
            </View>
        </TouchableNativeFeedback>
      </View>
    );
  }
}

module.exports = FirstPage;
