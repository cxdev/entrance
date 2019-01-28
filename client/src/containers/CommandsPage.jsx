import React, { Component } from 'react';
import { withRouter } from "react-router-dom";
import FetchToDataTable from '../components/FetchToDataTable'
import { command } from '../constants/api'

const { endpoint, fieldsConfig } = command

class CommandsPage extends Component {

    handleClick = (commandId) => {
        this.props.history.push(`/command/${commandId}`);
    }

    render = () => {
        return (
            <FetchToDataTable
                request={endpoint.find}
                handleClick={this.handleClick}
                fieldsConfig={fieldsConfig}
            />
        )
    }
}

export default withRouter(CommandsPage)