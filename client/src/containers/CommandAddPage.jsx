import React, { Component } from 'react'
import FetchComponent from '../components/FetchComponent';
import CommandAddForm from '../components/CommandAddForm'
import { command } from '../constants/api'

class CommandAddPage extends Component {
    handleSubmit = (formData) => {
        console.info(formData)
        const callBack = (data) => {
            console.info(data)
        }
        // this.fetchApiToState(callBack)
    }

    render = () => (
        <CommandAddForm handleSubmit={this.handleSubmit} />
    )

}

export default CommandAddPage