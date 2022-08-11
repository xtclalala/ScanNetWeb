import { isNullOrUnDef } from '@/utils/is'

/**
 * 所有 emit 内容的存储点
 */
const emitMap: { [prop: string]: Array<CallBackType> } = {}

/**
 * 所有 emit 是否被激活的存储点
 */
const emitMapNumber: { [Props: string]: number } = {}

type CallBackType = () => void

type EmitType = {
  SaveEmit: (fn: CallBackType, target: string) => void
  GetEmit: (target: string) => CallBackType | null
}

export const useEmit = (): EmitType => {
  const SaveEmit = (fn: CallBackType, target: string): void => {
    // 是否创建过该弹窗
    if (isNullOrUnDef(emitMapNumber[target])) {
      emitMapNumber[target] = 1
      emitMap[target] = [fn]
    } else {
      emitMapNumber[target]++
      emitMap[target].push(fn)
    }
  }

  const GetEmit = (target: string): CallBackType => {
    if (isNullOrUnDef(emitMapNumber[target])) {
      return () => {}
    }
    return (): void => {
      if (isNullOrUnDef(emitMapNumber[target])) {
        return
      }
      if (emitMapNumber[target] !== emitMap[target].length) {
        // window.$message?.warning("xxxxx")
        return
      }
      for (const fn of emitMap[target]) {
        fn()
      }
    }
  }
  return {
    SaveEmit,
    GetEmit,
  }
}
