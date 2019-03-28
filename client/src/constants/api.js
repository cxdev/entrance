const command = {
    'endpoint': {
        'add': '/admin/command/add',
        'find': '/command',
        'fetchOne': (commandId) => `/command/${commandId}/info`,
        'execute': (commandId) => `/command/${commandId}/execute`
    },
    'fieldsConfig': [
        { 'key': 'ID' },
        { 'key': 'CreatedAt' },
        { 'key': 'UpdatedAt' },
        { 'key': 'Name' },
        { 'key': 'CommandType' },
        {
            'key': 'CommandSegments',
            'name': 'Command Segments',
            'format': (value) => JSON.stringify(value)
        }]
}

const job = {
    'endpoint': {
        'find': '/job',
        'fetchOne': (jobId) => `/job/${jobId}/info`,
        'result': (jobId) => `/job/${jobId}/result`
    },
    'fieldsConfig': [
        { 'key': 'ID' },
        { 'key': 'CreatedAt' },
        { 'key': 'UpdatedAt' },
        { 'key': 'Status' },
        { 'key': 'CommandID' },
        {
            'key': 'Arguments',
            'format': (value) => JSON.stringify(value)
        },
        { 'key': 'SysCmd' }]
}

const api = {
    'command': command,
    'job': job
}

export default api
export { command, job }
