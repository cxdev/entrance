import React from 'react'
import { Divider, Header, Icon } from 'semantic-ui-react'
import DataBoard from '../components/DataBoard'
import CommandExecForm from '../components/CommandExecForm'
import FetchComponent from '../components/FetchComponent'
import { command } from '../constants/api'

class CommandDetailPage extends FetchComponent {

    getFetchAPI = (props) => {
        const { commandId } = props.match.params
        return `/command/${commandId}/info`

    }

    handleSubmit = (event) => {
        console.info(event)
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