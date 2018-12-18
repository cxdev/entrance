export const OPEN_COMMANDS_PAGE = 'OPEN_COMMAND_PAGE';
export const OPEN_COMMAND_DETAIL_PAGE = 'OPEN_COMMAND_DETAIL_PAGE';

// TODO: the page change should be handled by react-router
export const actions = {
    openCommandsPage: () => ({ 'type': OPEN_COMMANDS_PAGE }),
    openCommandDetailPage: (id) => ({ 'type': OPEN_COMMAND_DETAIL_PAGE, 'id': id })
}