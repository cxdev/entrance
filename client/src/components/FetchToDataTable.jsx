import React, { Component } from 'react';
import FetchComponent from './FetchComponent';
import DataTable from '../components/DataTable';

class FetchToDataTable extends FetchComponent {

    render = () => {
        return (
            <DataTable
                data={this.state.response ? this.state.response.Data : []}
                handleClick={this.props.handleClick}
                fieldsConfig={this.props.fieldsConfig}
            />
        )
    }
}

export default FetchToDataTable