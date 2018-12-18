import React from 'react'
import { Table } from 'semantic-ui-react'

const DataBoard = (props) => {
    const { fieldsConfig } = props
    return (

        <Table definition>
            <Table.Body>

                {fieldsConfig.map((field, i) => {
                    const name = field.name || field.key
                    let data = props[field.key]
                    if (field.format) {
                        data = field.format(data)
                    }
                    return (
                        <Table.Row key={i}>
                            <Table.Cell>{name}</Table.Cell>
                            <Table.Cell>{data}</Table.Cell>
                        </Table.Row>
                    )
                })}
            </Table.Body>
        </Table>
    )
}

export default DataBoard