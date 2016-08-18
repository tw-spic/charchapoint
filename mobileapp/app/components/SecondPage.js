import React, { Component } from 'react';
import {
  AppRegistry,
  StyleSheet,
  Text,
  View,
  Navigator,
  TouchableNativeFeedback
} from 'react-native';

class SecondPage extends Component {
  constructor(props) {
    super(props);
  }

  navigateToFirstPage() {
    this.props.navigator.push({id: 'FirstPage'});
  }

  render() {
    return (
      <View>
        <Text>Second Page</Text>
        <TouchableNativeFeedback
          onPress={this.navigateToFirstPage.bind(this)}
          background={TouchableNativeFeedback.SelectableBackground()}>
            <View>
              <Text>Go Back to First Page</Text>
            </View>
        </TouchableNativeFeedback>
      </View>
    );
  }
}

module.exports = SecondPage;
