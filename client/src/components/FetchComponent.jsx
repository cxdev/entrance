import { Component } from 'react';

export default class FetchComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            "response": null,
            "isOK": true,
            "message": null
        }
    }

    getFetchAPI = (props) => {
        return props.api
    }

    componentDidMount = () => {
        const time = new Date().toLocaleString()
        const api = this.getFetchAPI(this.props)
        if (api) {
            fetch(api)
                .then(response => {
                    if (response.ok) {
                        return response.json()
                    } else {
                        throw new Error('Response with code:', response.ok)
                    }
                })
                .then(responseJson => {
                    this.setState({
                        "response": responseJson,
                        "isOk": true,
                        "message": null,
                        "time": time
                    })
                })
                .catch(error => {
                    console.warn(error)
                    this.setState({
                        "isOk": false,
                        "message": error.message,
                        "time": time
                    })
                }
                )
        } else {
            this.setState({
                "isOk": false,
                "message": "No API for fetching",
                "time": time
            })
        }
    }
}
