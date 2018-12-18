const command = {
    'endpoint': '/command',
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
    'endpoint': '/job',
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
