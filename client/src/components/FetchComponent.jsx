import { Component } from 'react';

const fetchRequest = (request, successCallBack, errorCallback = successCallBack) => {
    const time = new Date().toLocaleString()
    if (request) {
        fetch(request)
            .then(response => {
                if (response.ok) {
                    return response.json()
                } else {
                    throw new Error('Response with code:', response.ok)
                }
            })
            .then(responseJson => {
                successCallBack({
                    "response": responseJson,
                    "isOk": true,
                    "message": null,
                    "time": time
                })
            })
            .catch(error => {
                errorCallback({
                    "isOk": false,
                    "message": error.message,
                    "time": time
                })
            }
            )
    } else {
        errorCallback({
            "isOk": false,
            "message": "No Request for fetching",
            "time": time
        })
    }
}

class FetchComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            "response": null,
            "isOK": true,
            "message": null
        }
    }

    componentDidMount = () => {
        const request = this.getFetchRequest()
        if (request) {
            this.fetchRequestToState(request)
        }
    }

    getFetchRequest = () => {
        return this.props.request
    }

    fetchRequestToState = (request) => {
        const callBack = (data) => this.setState(data)
        fetchRequest(request, callBack, (errorData) => {
            console.warn(errorData)
            callBack(errorData)
        })
    }
}

export default FetchComponent
export { fetchRequest }
