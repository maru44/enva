type kvListState = {
  isOpenUpdate: boolean
  targetKey: string
  updateDefaultValue: string
  deleteId: string
  isOpenDelete: boolean
}

type kvListAction =
  | {
      type: 'closeUpdate'
    }
  | {
      type: 'closeDelete'
    }
  | {
      type: 'openUpdate'
      targetKey: string
      updateDefaultValue: string
    }
  | {
      type: 'openDelete'
      targetKey: string
      deleteId: string
    }

export const initialKvListState: kvListState = {
  isOpenUpdate: false,
  targetKey: '',
  updateDefaultValue: '',
  isOpenDelete: false,
  deleteId: '',
}

export const kvListReducer = (state: kvListState, action: kvListAction) => {
  switch (action.type) {
    case 'closeUpdate':
      return { ...state, isOpenUpdate: false }
    case 'closeDelete':
      return { ...state, isOpenDelete: false }
    case 'openUpdate':
      return {
        ...state,
        isOpenUpdate: true,
        targetKey: action.targetKey,
        updateDefaultValue: action.updateDefaultValue,
      }
    case 'openDelete':
      return {
        ...state,
        isOpenDelete: true,
        targetKey: action.targetKey,
        deleteId: action.deleteId,
      }
  }
}
