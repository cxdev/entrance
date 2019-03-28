import React from 'react'
import { Table } from 'semantic-ui-react'


// {
// 	"Data": [
// 		{
// 			"ID": 1,
// 			"CreatedAt": "2018-12-06T18:53:43.562827422+09:00",
// 			"UpdatedAt": "2018-12-06T18:53:43.562827422+09:00",
// 			"Name": "test",
// 			"CommandType": 0,
// 			"CommandSegments": [
// 				{
// 					"Key": "ls",
// 					"Base": "/bin/ls",
// 					"IsRequired": false,
// 					"IsValuable": false
// 				}
// 			]
// 		}
// 	]
// }


// FieldsConfig is a list of following object
// key (must) - field of json obj
// name - display on the header of table, use key if name is null
// format - function for format data to display in the table


const DataRow = (props) => {
    const { handleClick, fieldsConfig } = props

    return (
        <Table.Row onClick={handleClick}>
            {fieldsConfig.map((field, i) => {
                let data = props[field.key]
                if (field.format) {
                    data = field.format(data)
                }
                return (<Table.Cell collapsing key={i}>{data}</Table.Cell>)
            })}
        </Table.Row>
    )
}

const DataTableHeader = (fieldsConfig) => (
    <Table.Header>
        <Table.Row>
            {fieldsConfig.map((field, i) => <Table.HeaderCell key={i}>{field.name || field.key}</Table.HeaderCell>)}
        </Table.Row>
    </Table.Header>
)

const DataTable = (props) => {
    const { data = [], handleClick, fieldsConfig } = props
    const tableHeader = DataTableHeader(fieldsConfig)
    const tableBody = data.map(row => (
        <DataRow {...row} fieldsConfig={fieldsConfig} key={row.ID} handleClick={() => handleClick(row.ID)} />
    ))

    return (
        <Table celled striped>
            {tableHeader}
            <Table.Body>
                {tableBody}
            </Table.Body>
        </Table>
    )
}

export default DataTable
export { DataRow }

