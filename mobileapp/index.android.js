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
// import FCM from 'react-native-fcm';
import * as firebase from 'firebase'
import { GiftedChat } from 'react-native-gifted-chat';

const firebaseConfig = {
  apiKey: "masked",
  authDomain: "1028630304114-nec2l2fmac4p32dke638rofrv9u1q9tt.apps.googleusercontent.com",
  databaseURL: "charcha-point.firebaseio.com",
  storageBucket: "charcha-point.appspot.com"
};

const firebaseApp = firebase.initializeApp(firebaseConfig);

class Example extends React.Component {
  constructor(props) {
    super(props);
    this.state = {messages: []};
    this.onSend = this.onSend.bind(this);
  }
  // componentDidMount() {
  //   FCM.requestPermissions(); // for iOS
  //   FCM.getFCMToken().then(token => {
  //     console.log(token)
  //     // store fcm token in your server
  //   });
  //   this.notificationUnsubscribe = FCM.on('notification', (notif) => {
  //     this.setState({
  //       data: this.state.data + "\n" + notif.message
  //     })
  //     // there are two parts of notif. notif.notification contains the notification payload, notif.data contains data payload
  //   });
  //   this.refreshUnsubscribe = FCM.on('refreshToken', (token) => {
  //     console.log(token)
  //     // fcm token may not be available on first load, catch it here
  //   });

  //   FCM.subscribeToTopic('/topics/foo-bar');
  //   FCM.unsubscribeFromTopic('/topics/foo-bar');
  // }

  componentWillMount() {
    this.deviceId = Math.trunc(Math.random() * 100); // Some random id for now
    this.setState({
      messages: [],
    });
  }

  componentDidMount() {
    firebase.database().ref('messages').set({});
    var messagesRef = firebase.database().ref('messages');
    messagesRef.on('child_added', (msg) => {
      this.setState((previousState) => {
        return {
          messages: GiftedChat.append(previousState.messages, msg.val().message),
        };
      });
    });

  }

  onSend(messages = []) {
    firebase.database().ref('messages').push({
        message: messages
    });
    
  }

  componentWillUnmount() {
    // // prevent leaking
    // this.refreshUnsubscribe();
    // this.notificationUnsubscribe();
  }

  render() {
    console.log(this.deviceId);
    return (
      <GiftedChat
        messages={this.state.messages}
        onSend={this.onSend}
        user={{
          _id: this.deviceId,
        }}
      />
    );
  }
}

AppRegistry.registerComponent('CharchaPoint', () => Example);
