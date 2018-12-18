import React from 'react'
import { Table } from 'semantic-ui-react'

// {
// 	"Data": [
// 		{
// 			"ID": 1,
// 			"CreatedAt": "2018-12-06T18:53:52.683314924+09:00",
// 			"UpdatedAt": "2018-12-06T18:53:52.683314924+09:00",
// 			"Status": 0,
// 			"CommandID": 1,
// 			"Arguments": {
// 				"ls": "/"
// 			},
// 			"SysCmd": "/bin/ls /"
// 		}
// 	]
// }

const headerRow = ['ID', 'CreatedAt', 'UpdatedAt', 'Status', 'CommandID', 'Arguments', 'SysCmd']