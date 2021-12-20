type kvListState = {
  isOpenUpdate: boolean
  updateKey: string
  updateDefaultValue: string
}

type kvListAction =
  | {
      type: 'close'
    }
  | {
      type: 'open'
      updateKey: string
      updateDefaultValue: string
    }

export const initialKvListState: kvListState = {
  isOpenUpdate: false,
  updateKey: '',
  updateDefaultValue: '',
}

export const kvListReducer = (state: kvListState, action: kvListAction) => {
  switch (action.type) {
    case 'close':
      return { ...state, isOpenUpdate: false }
    case 'open':
      return {
        ...state,
        isOpenUpdate: true,
        updateKey: action.updateKey,
        updateDefaultValue: action.updateDefaultValue,
      }
  }
}
