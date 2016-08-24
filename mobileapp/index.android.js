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
import FCM from 'react-native-fcm';

class Example extends React.Component {
  componentDidMount() {
    FCM.requestPermissions(); // for iOS
    FCM.getFCMToken().then(token => {
      console.log(token)
      // store fcm token in your server
    });
    this.notificationUnsubscribe = FCM.on('notification', (notif) => {
      // there are two parts of notif. notif.notification contains the notification payload, notif.data contains data payload
    });
    this.refreshUnsubscribe = FCM.on('refreshToken', (token) => {
      console.log(token)
      // fcm token may not be available on first load, catch it here
    });

    FCM.subscribeToTopic('/topics/foo-bar');
    FCM.unsubscribeFromTopic('/topics/foo-bar');
  }

  componentWillUnmount() {
    // prevent leaking
    this.refreshUnsubscribe();
    this.notificationUnsubscribe();
  }

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
