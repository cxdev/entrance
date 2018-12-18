import React from 'react';
import { withRouter } from "react-router-dom";
import FetchToDataTable from '../components/FetchToDataTable'
import { job } from '../constants/api'

const { endpoint, fieldsConfig } = job

class JobsPage extends FetchToDataTable {

    handleClick = (jobId) => {
        this.props.history.push(`/job/${jobId}`);
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

export default withRouter(JobsPage)