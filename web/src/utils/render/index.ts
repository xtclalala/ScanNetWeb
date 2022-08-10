/**
 * @Description: src\utils\render\index.ts
 * @author: y1t
 * @date 2022/7/26
 **/
import { h, VNodeChild } from 'vue'
import { YIcon } from '@/components'
import { TaskState, TaskStateMap } from '@/enums/bizEnum'
import { NDescriptions, NDescriptionsItem, NEmpty, NTag} from "naive-ui";

/**
 * 渲染图标
 * @param icon
 */
export const renderIcon = (icon: string): (() => VNodeChild) => {
  return () => h(YIcon, { iconType: icon })
}

/**
 * 渲染任务状态
 * @param taskState
 */
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

/**
 * 渲染构建任务
 * @param taskState
 */
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

/**
 * 渲染准备任务
 * @param taskState
 */
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

/**
 * 渲染运行任务
 * @param taskState
 */
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

/**
 * 渲染完成任务
 * @param taskState
 */
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

/**
 * 渲染异常任务
 * @param taskState
 */
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

export const renderTaskResult = (
  result: Array<string>,
  target: 'left' | 'top' = 'left'
): VNodeChild => {
  const renderList: Array<VNodeChild> = []
  const l = result.length
  for (let i = 0; i < l; i = i + 2) {
    renderList.push(renderTaskResultItem([result[i], result[i + 1]]))
  }
  if (renderList.length === 0) {
    return h(NEmpty, { description: '暂无扫描结果' })
  }
  return h(NDescriptions, { labelPlacement: target, column: 1 }, { default: () => renderList })
}

export const renderTaskResultItem = (result: Array<string>): VNodeChild => {
  return h(NDescriptionsItem, {}, { label: () => result[0], default: () => result[1] })
}
