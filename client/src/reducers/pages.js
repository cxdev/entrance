import { OPEN_COMMANDS_PAGE, OPEN_COMMAND_DETAIL_PAGE } from '../actions/commands'

const PageCode = Object.freeze({
    HOME: Symbol("HOME"),
    COMMANDS: Symbol("COMMANDS"),
    COMMAND_DETAIL: Symbol("COMMAND_DETAIL")
});

const pageReducer = (state = { 'viewing': PageCode.HOME }, action) => {
    switch (action.type) {
        case OPEN_COMMANDS_PAGE:
            return {
                'viewing': PageCode.COMMANDS,
                'action': action
            }
        case OPEN_COMMAND_DETAIL_PAGE:
            return {
                'viewing': PageCode.COMMAND_DETAIL,
                'action': action
            }
        default:
            return state
    }
}

export default pageReducer