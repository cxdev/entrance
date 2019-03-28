import React, { Component } from 'react'
import { Form, Input, Button } from 'semantic-ui-react'

// {
//     "Key": "ls",
//     "Base": "/bin/ls",
//     "IsRequired": false,
//     "IsValuable": false
// }
const SegmentItem = (props) => {
    const { Key, Base, IsRequired, IsValuable, handleChange } = props

    if (IsValuable) {
        return (
            <Form.Field required={IsRequired}>
                <label>{Key}</label>
                <Input name={Key} placeholder={Key} onChange={handleChange} />
            </Form.Field>
        )
    } else {
        return <Form.Checkbox inline label={`Use ${Key}?`} required={IsRequired} name={Key} onChange={handleChange} />
    }
}

// TODO form validation
class CommandExecForm extends Component {
    constructor() {
        super();
        this.state = {
            data: {}
        }
    }
    handleChange = (_, { name, value }) => {
        let updatedData = Object.assign(this.state.data, { [name]: value })
        this.setState({ data: updatedData })
    }

    render = () => {
        console.info(this.state)
        const { ID, CreatedAt, UpdatedAt, Name, CommandType, CommandSegments, handleSubmit } = this.props
        const segmentItems = CommandSegments.map((commandSegment, i) => <SegmentItem key={i} {...commandSegment} handleChange={this.handleChange} />)
        return (
            <Form onSubmit={() => handleSubmit(this.state.data)}>
                {segmentItems}
                <Button type='submit'>Submit</Button>
            </Form>
        )
    }
}


export default CommandExecForm