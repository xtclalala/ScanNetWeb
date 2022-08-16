import { isNullOrUnDef } from '@/utils/is'
import { reactive, ref } from 'vue'
import { Page } from '@/api/common/types/login'
import { useClientView } from '@/hooks/comHooks/useClientView'

const tableMap = {}
const tableMapStates = {}

export const useTable = <T, D>(
  dataApi: Function,
  page: Page,
  searchObj: D,
  key: string,
  tableHeight = 0.5
) => {
  // 创建分页对象
  const createTable = (page: Page, key: string) => {
    let pagination, loading, data, searchData
    if (isNullOrUnDef(tableMapStates[key])) {
      pagination = reactive({
        page: page.page,
        pageSize: page.pageSize,
        itemCount: 0,
        onChange: (pageNum: number) => {
          loading.value = true
          pagination.page = pageNum
          getData({ page: pageNum, pageSize: pagination.pageSize, desc: page.desc })
        },
      })
      loading = ref<boolean>(true)
      data = ref<Array<T>>([])
      searchData = ref({ ...searchObj })
      tableMapStates[key] = true
      tableMap[key] = { pagination, loading, data, searchData }
    } else {
      ;({ pagination, loading, data, searchData } = tableMap[key])
    }
    return {
      pagination,
      loading,
      data,
      searchData,
    }
  }

  // 获取搜索数据
  const getData = async (page: Page): Promise<void> => {
    const res = await dataApi(page, searchData)
    pagination.itemCount = res.total
    data.value = res.items
    loading.value = false
  }

  // 点击搜索按钮
  const doSearch = (): void => {
    loading.value = true
    pagination.page = 1
    getData({ page: 1, pageSize: pagination.pageSize, desc: page.desc })
  }

  // 重置搜索条件
  const doReset = (): void => {
    searchData.value = { ...searchObj }
  }

  const key2id = (row) => row.id

  const getTableHeight = () => {
    const { height } = useClientView()
    return height * tableHeight
  }

  const { pagination, loading, data, searchData } = createTable(page, key)

  return [
    pagination,
    loading,
    data,
    searchData,
    getData,
    doSearch,
    doReset,
    key2id,
    getTableHeight(),
  ]
}
