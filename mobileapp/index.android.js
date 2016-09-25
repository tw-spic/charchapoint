import React from 'react';
import {
  Platform,
  StyleSheet,
  Text,
  View,
  AppRegistry,
  ToastAndroid,
  ActivityIndicator,
  TouchableWithoutFeedback
} from 'react-native';
import { NativeModules } from 'react-native';
import * as firebase from 'firebase'
import { GiftedChat } from 'react-native-gifted-chat';
import geodist from 'geodist';
import DeviceInfo from 'react-native-device-info';

const firebaseConfig = {
  apiKey: "masked",
  authDomain: "1028630304114-nec2l2fmac4p32dke638rofrv9u1q9tt.apps.googleusercontent.com",
  databaseURL: "charcha-point.firebaseio.com",
  storageBucket: "charcha-point.appspot.com"
};

const styles = StyleSheet.create({
  view:{
    flex:1,
    flexDirection:'column',
    alignItems:'center',
    justifyContent:'center'
  },
  titleText: {
    fontSize: 20,
    fontWeight: 'bold',
    textAlign: 'center',
    paddingBottom:20,
  },  
  normalText: {
    textAlign: 'center',
    paddingBottom:20,
  },
});

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

class CharchaPoint extends React.Component {
  constructor(props) {
    super(props);
    this.deviceId = DeviceInfo.getUniqueID();
    this.onSend = this.onSend.bind(this);
    this.findCurrentZone = this.findCurrentZone.bind(this);
    this.subscribeMessages = this.subscribeMessages.bind(this);
    this.registerLocationWatcher = this.registerLocationWatcher.bind(this);
    this.setCurrentZone = this.setCurrentZone.bind(this);
    this.state = {messages: []};
  }

  componentDidMount() {
    this.registerLocationWatcher();
  }

  registerLocationWatcher() {
    navigator.geolocation.getCurrentPosition(
      (position) => {},
      (error) => {},
      {enableHighAccuracy: true, timeout: 20000, maximumAge: 1000}
    );
    this.watchID = navigator.geolocation.watchPosition((position) => {
      this.lat = position.coords.latitude;
      this.long = position.coords.longitude;
      this.setCurrentZone();
    });
  }

  subscribeMessages() {
    var messagesRef = firebase.database().ref("messages/" + this.state.zone.Id);
    messagesRef.on('child_added', (msg) => {
      this.setState((previousState) => {
        return {
          ...previousState,
          messages: GiftedChat.append(previousState.messages, msg.val().message),
        };
      });
    });
  }

  setCurrentZone() {
    if (!this.lat || !this.long) {
      return;
    }

    var currZone = this.findCurrentZone(this.lat, this.long);
    if (!currZone || currZone === this.state.zone) {
      return;
    }

    ToastAndroid.show("You are in " + currZone.Name + " zone. \n" + currZone.Description, ToastAndroid.LONG);
    this.setState((previousState) => {
      return {
        ...previousState,
        zone:currZone
      };
    });

    this.subscribeMessages();
  }

  findCurrentZone(lat, long) {
    if (!zones) {
      return;
    }
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
    if(!this.state.zone){
       return (
           <View style={styles.view}>
              <Text style={styles.titleText} >Charcha point </Text>
              <Text style={styles.normalText}  >We are finding zone near to you. Please wait...</Text>
              <ActivityIndicator color="#0000ff" size="large"/>
            </View>
        );
     } else {
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
}

AppRegistry.registerComponent('CharchaPoint', () => CharchaPoint);
