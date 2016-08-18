import axios from 'axios';

getContent = () => (dispatch) => {
  dispatch({
    type: 'GET_CONTENT_REQUEST'
  });

  const success = (response) => dispatch({
    type: 'GET_CONTENT_OK',
    response
  });

  const failure = (error) => dispatch({
    type: 'GET_CONTENT_ERROR',
    error
  });

  return axios.get('https://randomuser.me/api/')
  .then(success, failure);
};

module.exports = getContent;
