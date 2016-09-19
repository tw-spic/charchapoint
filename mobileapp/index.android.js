import React from 'react';
import {
  Platform,
  StyleSheet,
  Text,
  View,
  AppRegistry,
  ToastAndroid,
  TouchableWithoutFeedback
} from 'react-native';
import { NativeModules } from 'react-native';
// import FCM from 'react-native-fcm';
import * as firebase from 'firebase'
import { GiftedChat } from 'react-native-gifted-chat';
import geodist from 'geodist';

const firebaseConfig = {
  apiKey: "masked",
  authDomain: "1028630304114-nec2l2fmac4p32dke638rofrv9u1q9tt.apps.googleusercontent.com",
  databaseURL: "charcha-point.firebaseio.com",
  storageBucket: "charcha-point.appspot.com"
};

const zones = [{ // get zones from server
  Id:123,
  Name:"Thoughtworks",
  Description:"TW Bangalore",
  Lat:12.928843,
  Long:77.628614,
  Radius:500
},
{
  Id:124,
  Name:"Home",
  Description:"My Home",
  Lat:12.930197,
  Long:77.634091,
  Radius:500
},
{
  Id:125,
  Name:"TW Pune",
  Description:"Thoughtworks",
  Lat:18.555910,
  Long:73.891793,
  Radius:500
},
{
  Id:125,
  Name:"Sony signal",
  Description:"Sony signal",
  Lat:12.937292,
  Long:77.626935,
  Radius:500
}]
 
const firebaseApp = firebase.initializeApp(firebaseConfig);

class Example extends React.Component {
  constructor(props) {
    super(props);
    this.state = {messages: []};
    this.onSend = this.onSend.bind(this);
    this.findCurrentZone = this.findCurrentZone.bind(this);
    this.subscribeMessages = this.subscribeMessages.bind(this);
  }

  componentWillMount() {
    this.deviceId = Math.trunc(Math.random() * 100); // Some random id for now
    this.setState({
      messages: [],
    });
  }

  componentDidMount() {

    //register location watcher
    navigator.geolocation.getCurrentPosition(
      (position) => {},
      (error) => alert(error),
      {enableHighAccuracy: true, timeout: 20000, maximumAge: 1000}
    );
    this.watchID = navigator.geolocation.watchPosition((position) => {
      var lastPosition = JSON.stringify(position);
      var lat = position.coords.latitude;
      var long = position.coords.longitude;
      var currZone = this.findCurrentZone(lat,long);
      ToastAndroid.show("Current zone is " + currZone.Name, ToastAndroid.LONG);
      this.setState((previousState) => {
        return {
          ...previousState,
          zone:currZone
        };
      });
      this.subscribeMessages();
    });
  }

  subscribeMessages() {
    var messagesRef = firebase.database().ref("messages/" + this.state.zone.Id);
    messagesRef.on('child_added', (msg) => {
      this.setState((previousState) => {
        return {
          messages: GiftedChat.append(previousState.messages, msg.val().message),
        };
      });
    });
  }

  findCurrentZone(lat, long) {
    return zones.find((zone) => { return geodist({lat:lat,lon:long}, { lat:zone.Lat, lon:zone.Long }, {unit: 'meters'}) < zone.Radius})
  }

  onSend(messages = []) { // now we are directly writing to the FCM database, we may need to route it through our server to enable FCM notifications
    firebase.database().ref("messages/" + this.state.zone.Id).push({
        message: messages
    });
    
  }

  componentWillUnmount() {
    navigator.geolocation.clearWatch(this.watchID);
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
