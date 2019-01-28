import React, { Component } from 'react'
import { Form, Input, Select, Item, Icon, Button } from 'semantic-ui-react'

const options = [
    { key: '1', text: 'BATCH', value: 1 },
    { key: '2', text: 'DAEMON', value: 2 },
    { key: '3', text: 'INTERACTIVE', value: 3 },
]

class CommandAddForm extends Component {
    constructor() {
        super();
        this.state = {
            name: "",
            type: null,
            segments: []
        }
    }

    handleChange = (_, data) => {
        const { name, value } = data || {}
        this.setState({ [name]: value })
    }

    handleSegmentChange = (_, data) => {
        const { name, value, index, checked, type } = data || {}
        let segments = this.state.segments
        switch (type) {
            case "checkbox":
                segments[index][name] = checked
                break
            default:
                segments[index][name] = value
        }
        this.setState({ segments: segments })
    }

    addSegment = () => {
        this.setState((previousState) => ({
            "segments": [
                ...previousState.segments,
                {
                    Key: "",
                    Base: "",
                    IsRequired: false,
                    IsValuable: true
                }
            ]
        }))
    }

    // TODO: bug
    removeSegment = (i) => {
        console.info(i)
        console.info(this.state)
        this.setState((previousState) => {
            let segments = [...previousState.segments]
            console.info(segments)
            return {
                "segments": segments.splice(i, 1)
            }
        })
        console.info(this.state)
    }

    render = () => {
        const segmentItems = this.state.segments.map((segment, i) => {
            return (
                <Item key={i} id={i} >
                    <Form.Field>
                        <label>Name</label>
                        <Input
                            name="Key"
                            placeholder='Key'
                            index={i}
                            value={segment.Key}
                            onChange={this.handleSegmentChange}
                        />
                        <Input
                            name="Base"
                            placeholder='Base'
                            index={i}
                            value={segment.Base}
                            onChange={this.handleSegmentChange}
                        />
                        <Form.Checkbox
                            inline
                            label="Is required"
                            name="IsRequired"
                            index={i}
                            checked={segment.IsRequired}
                            onChange={this.handleSegmentChange}
                        />
                        <Form.Checkbox
                            label="Is valuable"
                            name="IsValuable"
                            index={i}
                            checked={segment.IsValuable}
                            onChange={this.handleSegmentChange}
                        />
                    </Form.Field>
                    <Icon link name='close' onClick={() => this.removeSegment(i)} />
                </Item>
            )
        })

        return (
            <Form onSubmit={() => this.props.handleSubmit(this.state)}>
                <Form.Field>
                    <label>Name</label>
                    <Input
                        name='name'
                        placeholder='Name'
                        value={this.state.name}
                        onChange={this.handleChange}
                    />
                </Form.Field>
                <Form.Field
                    name='type'
                    control={Select}
                    label='Type'
                    options={options}
                    value={this.state.type}
                    placeholder='type'
                    onChange={this.handleChange}
                />
                <div onClick={this.addSegment}>
                    <Icon name='add' /> Add Segment
                </div>
                <Item.Group>
                    {segmentItems}
                </Item.Group>
                <Button type='submit'>Submit</Button>
            </Form>
        )
    }
}

export default CommandAddForm