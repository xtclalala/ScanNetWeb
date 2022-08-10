import { isNullOrUnDef } from '@/utils/is'

/**
 * 所有 emit 内容的存储点
 */
const emitMap = {}

/**
 * 所有 emit 是否被激活的存储点
 */
const emitMapStates = {}

type CallBackType = () => void

type EmitType = {
  SaveEmit: (fn: CallBackType, target: string) => void
  GetEmit: (target: string) => CallBackType | null
}

export const useEmit = (): EmitType => {
  const SaveEmit = (fn: CallBackType, target: string): void => {
    // 是否创建过该弹窗
    // todo 待执行函数可以存多个
    if (isNullOrUnDef(emitMapStates[target])) {
      emitMapStates[target] = true
      emitMap[target] = fn
    } else {
      emitMap[target] = fn
    }
  }

  const GetEmit = (target: string): CallBackType | null => {
    if (isNullOrUnDef(emitMapStates[target])) {
      return null
    }
    return emitMap[target]
  }
  return {
    SaveEmit,
    GetEmit,
  }
}
