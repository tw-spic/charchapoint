const initialState = {
  user: null,
  loading: false
};

function reducer(state = initialState, action) {
  switch (action.type) {
    case 'GET_CONTENT_REQUEST':
      return { loading: true, user: null };
    case 'GET_CONTENT_OK':
      console.log(action);
      return { loading: false, user: action.response.data.results[0] };
    case 'GET_CONTENT_ERROR':
      return { loading:false, user: null, error: true }
    default:
      return state;
  }
}

module.exports = reducer;
