let action = () => {}

export default () => {
  return action
}

export function SetAction(newAction) {
  action = newAction
}
