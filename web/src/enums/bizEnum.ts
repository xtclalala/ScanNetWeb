export enum TaskState {
  def,
  Build,
  Ready,
  Doing,
  Finish,
}

export enum TaskStateValue {
  def = '出错啦，',
  Build = '构建中',
  Ready = '准备完成',
  Doing = '运行中',
  Finish = '运行结束',
}

export enum TaskStateIcon {
  def = 'AlertOutline',
  Build = 'BuildOutline',
  Ready = 'Checkmark',
  Doing = 'Aperture',
  Finish = 'CheckmarkDoneSharp',
}

type TaskStateMapType = {
  [prop: string]: TaskStateMapItem
}

type TaskStateMapItem = { icon: string; str: string; value: TaskState }

export const TaskStateMap: TaskStateMapType = {}

Object.entries(TaskState).forEach(([key, value]) => {
  TaskStateMap[value] = {
    icon: TaskStateIcon[key],
    str: TaskStateValue[key],
    value: value as TaskState,
  }
})
