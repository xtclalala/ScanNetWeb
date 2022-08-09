/**
 * @Description: src\utils\render\index.ts
 * @author: y1t
 * @date 2022/7/26
 **/
import { h, VNodeChild } from 'vue'
import { YIcon } from '@/components'
import { TaskState, TaskStateMap } from '@/enums/bizEnum'
import { NTag } from 'naive-ui'

/**
 * 渲染图标
 * @param icon
 */
export const renderIcon = (icon: string): (() => VNodeChild) => {
  return () => h(YIcon, { iconType: icon })
}

export const renderTaskState = (taskState: TaskState): VNodeChild => {
  switch (taskState) {
    case TaskState.Build:
      return renderTaskStateBuild()
    case TaskState.Ready:
      return renderTaskStateReady()
    case TaskState.Doing:
      return renderTaskStateDoing()
    case TaskState.Finish:
      return renderTaskStateFinish()
    default:
      return renderTaskStateDefault()
  }
}

export const renderTaskStateBuild = (): VNodeChild =>
  h(
    NTag,
    {
      round: true,
      bordered: false,
      type: 'info',
    },
    {
      default: () => TaskStateMap[TaskState.Build].str,
      icon: renderIcon(TaskStateMap[TaskState.Build].icon),
    }
  )

export const renderTaskStateReady = (): VNodeChild =>
  h(
    NTag,
    {
      round: true,
      bordered: false,
      type: 'success',
    },
    {
      default: () => TaskStateMap[TaskState.Ready].str,
      icon: renderIcon(TaskStateMap[TaskState.Ready].icon),
    }
  )

export const renderTaskStateDoing = (): VNodeChild =>
  h(
    NTag,
    {
      round: true,
      bordered: false,
      type: 'warning',
    },
    {
      default: () => TaskStateMap[TaskState.Doing].str,
      icon: renderIcon(TaskStateMap[TaskState.Doing].icon),
    }
  )

export const renderTaskStateFinish = (): VNodeChild =>
  h(
    NTag,
    {
      round: true,
      bordered: false,
      type: 'success',
    },
    {
      default: () => TaskStateMap[TaskState.Finish].str,
      icon: renderIcon(TaskStateMap[TaskState.Finish].icon),
    }
  )

export const renderTaskStateDefault = (): VNodeChild =>
  h(
    NTag,
    {
      round: true,
      bordered: false,
      type: 'error',
    },
    {
      default: () => TaskStateMap[TaskState.def].str,
      icon: renderIcon(TaskStateMap[TaskState.def].icon),
    }
  )
