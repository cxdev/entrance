import React, { Component } from 'react'
import { Divider, Header, Icon } from 'semantic-ui-react'
import DataBoard from '../components/DataBoard'
import CommandExecForm from '../components/CommandExecForm'
import FetchComponent, { fetchRequest } from '../components/FetchComponent'
import { command } from '../constants/api'

class CommandDetailPage extends FetchComponent {

    getFetchRequest = () => {
        const { commandId } = this.props.match.params
        return command.endpoint.fetchOne(commandId)
    }

    getExecuteRequest = (json) => {
        const { commandId } = this.props.match.params
        const endpoint = command.endpoint.execute(commandId)
        return new Request(
            endpoint,
            {
                method: 'POST',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(json)
            }
        )
    }

    handleSubmit = (formData) => {
        console.info(formData)
        const callBack = (data) => {
            console.info(data)
            if (data.isOk) {
                this.props.history.push(`/job/${data.response.Data}`);
            }
        }

        const request = this.getExecuteRequest(formData)

        fetchRequest(request, callBack)
    }

    render = () => {
        const commandDetail = this.state.response ? <DataBoard {...this.state.response['Data']} fieldsConfig={command.fieldsConfig} /> : "No record"
        const executeForm = this.state.response ? <CommandExecForm {...this.state.response['Data']} handleSubmit={this.handleSubmit} /> : null
        return (
            <React.Fragment>
                <Divider horizontal>
                    <Header as='h4'>
                        <Icon name='tag' />
                        Description
                    </Header>
                </Divider>
                {commandDetail}
                <Divider horizontal>
                    <Header as='h4'>
                        <Icon name='paper plane outline' />
                        Execute
                    </Header>
                </Divider>
                {executeForm}
            </React.Fragment>
        )
    }
}

export default CommandDetailPage