import React, { Component } from 'react'
import { fetchRequest } from '../components/FetchComponent';
import CommandAddForm from '../components/CommandAddForm'
import { command } from '../constants/api'

class CommandAddPage extends Component {
    handleSubmit = (formData) => {
        console.info(formData)
        const callBack = (data) => {
            console.info(data)
            if (data.isOk) {
                this.props.history.push(`/commands/`);
            }
        }
        const request = new Request(command.endpoint.add,
            {
                method: 'POST',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(formData)
            })
        fetchRequest(request, callBack)
    }

    render = () => (
        <CommandAddForm handleSubmit={this.handleSubmit} />
    )

}

export default CommandAddPage