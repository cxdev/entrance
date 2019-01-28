import { combineReducers } from 'redux'
import pageReducer from 'pages'

const rootReducer = combineReducers({
    'page': pageReducer
})

export default rootReducer