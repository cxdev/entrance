import React, { Component } from 'react';
import { withRouter } from "react-router-dom";
import FetchToDataTable from '../components/FetchToDataTable'
import { job } from '../constants/api'

const { endpoint, fieldsConfig } = job

class JobsPage extends Component {

    handleClick = (jobId) => {
        this.props.history.push(`/job/${jobId}`);
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

export default withRouter(JobsPage)