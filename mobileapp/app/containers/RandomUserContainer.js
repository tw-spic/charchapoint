import React, { Component } from 'react';
import {
  AppRegistry,
  StyleSheet,
  Text,
  View,
  Navigator,
  TouchableNativeFeedback,
  Image
} from 'react-native';
import { connect } from 'react-redux';

var getRandomUser = require('../actions/getRandomUser');

class RandomUserContainer extends Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    this.props.getRandomUser();
  }

  renderUserInfo(user) {
    return (
      <View>
        <Image
        style={{
            width: 120,
            height: 120,
            backgroundColor: 'transparent',
            marginLeft: 30,
            borderRadius: 20
          }}
          resizeMode={Image.resizeMode.contain}
          source={{uri: user.picture.large}}
        />
        <Text>First Name: { user.name.first }</Text>
        <Text>Last Name: { user.name.last }</Text>
      </View>
    );
  }
  render() {
    return (
      <View>
        <Text>Random User Container</Text>
        {(this.props.loading) && <Text>Loading user data...</Text>}
        {(this.props.user) && this.renderUserInfo(this.props.user)}
      </View>
    );
  }
}

function mapStateToProps(state) {
    var loading = state.loading;
    var user = (state.user) ? state.user : null;
    return { user, loading };
}

const mapDispatchToProps = {
  getRandomUser
}

module.exports = connect(mapStateToProps, mapDispatchToProps)(RandomUserContainer);
