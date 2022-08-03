import { computed, ref, Ref } from 'vue'
import { isEmpty, isNullOrUnDef } from '@/utils/is'
import { FormInst } from 'naive-ui'

/**
 * 所有 modal 内容的存储点
 */
const modalMap = {}

/**
 * 所有 modal 是否被激活的存储点
 */
const modalMapStates = {}

export const useModal = <T>(
  registerApi: Function,
  updateApi: Function,
  afterApi: Function,
  modelObj: T,
  modalStyle: Object,
  key: string
): useModalType<T> => {
  /**
   * 创建或者获取一个 modal 需要的相关字段
   */
  const createModal = (): ModalType<T> => {
    let isAdd, style, showModal, form, model
    // 是否创建过该弹窗
    if (isNullOrUnDef(modalMapStates[key])) {
      showModal = ref<boolean>(false)
      style = computed(() => {
        if (isEmpty(modalStyle)) {
          return { width: '600px' }
        }
        return modalStyle
      })
      isAdd = ref<boolean>(true)
      form = ref<FormInst | null>(null)
      model = ref<T>(Object.assign({}, modelObj))
      modalMapStates[key] = true
      modalMap[key] = { isAdd, style, showModal, form, model }
    } else {
      ;({ isAdd, style, showModal, form, model } = modalMap[key])
    }
    return {
      isAdd,
      style,
      showModal,
      form,
      model,
    }
  }

  /**
   * 新增数据
   */
  const handleRegister = async (): Promise<void> => {
    isAdd.value = true
    clearModel()
    openModal()
  }

  /**
   * 提交回调
   * @param e 鼠标点击事件
   */
  const submitCallback = async (e: MouseEvent): Promise<void> => {
    e.preventDefault()
    form.value
      ?.validate(async (errors) => {
        if (errors) {
          return
        }
        if (isAdd.value) {
          await registerApi(model.value, { isMessage: true })
        } else {
          await updateApi(model.value, { isMessage: true })
        }
        await afterApi()
        showModal.value = false
        clearModel()
      })
      .catch((e) => {
        console.log(e)
      })
  }

  /**
   * 清空弹窗
   */
  const clearModel = (): void => {
    model.value = Object.assign({}, modelObj)
  }

  /**
   * 打开弹窗
   */
  const openModal = (): void => {
    showModal.value = true
  }

  /**
   * 取消回调
   */
  const cancelCallback = (): void => {
    clearModel()
    showModal.value = false
    isAdd.value = false
  }

  /**
   * 弹窗标题
   */
  const modalTitle = (): string => {
    return isAdd.value ? '新增' : '编辑'
  }

  const { isAdd, showModal, form, model, style } = createModal()

  return [
    isAdd,
    showModal,
    form,
    model,
    style,
    handleRegister,
    submitCallback,
    clearModel,
    openModal,
    cancelCallback,
    modalTitle,
  ]
}

/**
 * 数据新增
 */
type HandlerRegisterFunc = () => void

/**
 * 点击提交
 */
type HandlerSubmitFunc = (e: MouseEvent) => void

/**
 * 点击取消
 */
type HandlerCancelFunc = () => void

/**
 * 打开弹窗
 */
type HandlerOpenFunc = () => void

/**
 * 清空弹窗
 */
type HandlerClearFunc = () => void

/**
 * 弹窗标题
 */
type TitleFunc = () => string

/**
 * 一个 modal 所需要的相关字段
 */
type ModalType<T> = {
  isAdd: Ref<boolean>
  showModal: Ref<boolean>
  form: Ref<FormInst | null>
  model: Ref<T>
  style: Object
}

/**
 * useModal 方法向外暴漏的问题
 */
type useModalType<T> = [
  Ref<boolean>,
  Ref<boolean>,
  Ref<FormInst | null>,
  Ref<T>,
  Object,
  HandlerRegisterFunc,
  HandlerSubmitFunc,
  HandlerClearFunc,
  HandlerOpenFunc,
  HandlerCancelFunc,
  TitleFunc
]
