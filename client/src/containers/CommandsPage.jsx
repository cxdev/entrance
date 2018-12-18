import React from 'react';
import { withRouter } from "react-router-dom";
import FetchToDataTable from '../components/FetchToDataTable'
import { command } from '../constants/api'

const { endpoint, fieldsConfig } = command

class CommandsPage extends FetchToDataTable {

    handleClick = (commandId) => {
        this.props.history.push(`/command/${commandId}`);
    }

    render = () => {
        return (
            <FetchToDataTable
                api={endpoint}
                handleClick={this.handleClick}
                fieldsConfig={fieldsConfig}
            />
        )
    }
}

export default withRouter(CommandsPage)