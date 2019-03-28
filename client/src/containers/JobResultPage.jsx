import React from 'react'
import { Divider, Header, Icon, List } from 'semantic-ui-react'
import FetchComponent from '../components/FetchComponent'
import { job } from '../constants/api'

const JobResultItem = (props) => {
    const { data = [] } = props
    const dataList = data.map(
        (item, i) => <List.Item key={i}>{item}</List.Item>
    )
    return (<List>{dataList}</List>)
}

class JobResultPage extends FetchComponent {

    getFetchRequest = () => {
        const { jobId } = this.props
        return job.endpoint.result(jobId)
    }

    handleSubmit = (event) => {
        console.info(event)
    }

    render = () => {
        const { Data = {} } = this.state.response || {}
        const { OutputData = [], ErrorData = [] } = Data
        return (
            <React.Fragment>
                <Divider horizontal>
                    <Header as='h4'>
                        <Icon name='tag' />
                        Output
                </Header>
                </Divider>
                <JobResultItem data={OutputData} />

                <Divider horizontal>
                    <Header as='h4'>
                        <Icon name='tag' />
                        Error
                    </Header>
                </Divider>
                <JobResultItem data={ErrorData} />
            </React.Fragment>
        )
    }
}

export default JobResultPage