type projectListState = {
  isOpenDelete: boolean
  targetKey: string
  deleteId: string
}

type projectListAction =
  | {
      type: 'closeDelete'
    }
  | {
      type: 'openDelete'
      targetKey: string
      deleteId: string
    }

export const initialProjectListState: projectListState = {
  isOpenDelete: false,
  targetKey: '',
  deleteId: '',
}

export const projectListReducer = (
  state: projectListState,
  action: projectListAction
) => {
  switch (action.type) {
    case 'closeDelete':
      return { ...state, isOpenDelete: false }
    case 'openDelete':
      return {
        ...state,
        isOpenDelete: true,
        targetKey: action.targetKey,
        deleteId: action.deleteId,
      }
  }
}
