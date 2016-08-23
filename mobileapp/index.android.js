import React from 'react';
import {
  Platform,
  StyleSheet,
  Text,
  View,
  AppRegistry,
  TouchableWithoutFeedback
} from 'react-native';
import { NativeModules } from 'react-native';

class Example extends React.Component {
  render() {
    NativeModules.MyToastAndroid.show('Awesome', 300);
    return (
      <View>
        <TouchableWithoutFeedback
          onPress={() =>
            NativeModules.MyToastAndroid.show('Awesome', 300)}>
          <Text>Click me.</Text>
        </TouchableWithoutFeedback>
      </View>
    );
  }
}

AppRegistry.registerComponent('CharchaPoint', () => Example);
