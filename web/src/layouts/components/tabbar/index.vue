<script lang="ts">
export default {
  name: 'YTabBar',
}
</script>
<script setup lang="ts">
import { useViewStore } from '@/store/module/views'
import { useRouter, useRoute } from 'vue-router'
import { useMessage } from 'naive-ui'
import { nextTick, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { YIcon } from '@/components'
import { AppRouteRecordRaw } from '@r/types'
import { router2menu } from '@/utils/yMenu'

const viewStore = useViewStore()
const message = useMessage()
const { viewList, currentView } = storeToRefs(viewStore)
const router = useRouter()
const route = useRoute()
const scrollbar: any = ref(null)
const leftArrowDisabled = ref<boolean>(false)
const rightArrowDisabled = ref<boolean>(false)

/**
 * 监听路由跳转时将视图添加进视图列表
 */
watch(
  route,
  (n) => {
    const r = router2menu(n as AppRouteRecordRaw)
    router.push({ name: n.name as string })
    viewStore.routerPush(r)
  },
  { immediate: true }
)

/**
 * 监听视图列表的长度是否超过显示长度，
 * 超出则显示右移或左移图标
 */
watch(
  viewList,
  async () => {
    await nextTick(() => {
      const { clientWidth, scrollWidth } = scrollbar.value.$el.nextElementSibling.firstChild
      rightArrowDisabled.value = scrollWidth > clientWidth
    })
  },
  { deep: true }
)

/**
 * 点击关闭视图图标
 * @param name
 */
const iconClick = (name: string) => {
  if (viewList.value.length === 1) {
    message.warning('最后一页不能删除')
    return
  }
  closeTab(name)
}

/**
 * 点击视图，并跳转
 * @param name
 * @param e
 */
const itemClick = (name: string, e: MouseEvent) => {
  currentView.value = name
  tabClick(e.target as HTMLElement, name)
}

/**
 * 点击视图，并跳转
 * @param name
 * @param e
 */
const itemChildClick = (name: string, e: MouseEvent) => {
  currentView.value = name
  tabClick(e.target as HTMLElement, name)
}

/**
 * 点击视图，并跳转
 * @param el
 * @param name
 */
const tabClick = (el: HTMLElement, name: string) => {
  scrollbar.value.scrollTo(
    {
      left: el.offsetLeft,
      debounce: true,
      behavior: 'auto',
    } as any,
    0
  )
  router.push({ name })
}

/**
 * 点击左移图标
 */
const leftArrowClick = () => {
  const scrollX = scrollbar.value.$el.nextElementSibling.firstChild.scrollLeft || 0
  scrollbar.value.scrollTo({
    left: Math.max(0, scrollX - 200),
    debounce: true,
    behavior: 'smooth',
  } as any)
  isDisabledArrow()
}

/**
 * 点击右移图标
 */
const rightArrowClick = () => {
  const scrollX = scrollbar.value.$el.nextElementSibling.firstElementChild.scrollLeft || 0
  scrollbar.value.scrollTo(
    {
      left: scrollX + 200,
      debounce: true,
      behavior: 'smooth',
    } as any,
    0
  )
  isDisabledArrow()
}

/**
 * 移动tab栏
 */
const isDisabledArrow = () => {
  setTimeout(() => {
    nextTick(() => {
      const { scrollLeft, scrollWidth, clientWidth } =
        scrollbar.value.$el.nextElementSibling.firstChild
      leftArrowDisabled.value = scrollLeft !== 0
      rightArrowDisabled.value = scrollWidth > scrollLeft + clientWidth + 10
    })
  }, 100)
}

/**
 * 关闭视图，如果关闭当前页，将自动跳转到最后一页
 * @param name
 */
const closeTab = async (name: string) => {
  const index = viewStore.findTabIndexByName(name)
  if (index !== -1) {
    await viewStore.removeTab(index)
    if (currentView.value === name) {
      currentView.value = viewStore.viewListLast.name
      await router.push({ name: currentView.value })
    }
  }
}
</script>

<template>
  <div class="tab-bar">
    <y-icon
      v-if="leftArrowDisabled"
      icon-type="ChevronBack"
      :size="18"
      class="icon-location"
      @click="leftArrowClick"
    />
    <div class="tabs">
      <n-scrollbar ref="scrollbar" x-scrollable>
        <n-button
          v-for="item of viewList"
          :key="item.name"
          :type="currentView === item.name ? 'primary' : 'default'"
          class="tab-item"
          style="--n-height: 24px; --n-font-weight: 200"
          @click.self="itemClick(item.name, $event)"
        >
          <template #icon>
            <y-icon class="show-icon-item" icon-type="Aperture" size="15" />
          </template>
          <span
            style="font-size: 12px; margin-top: 2px"
            class="text-item"
            @click.self="itemChildClick(item.name, $event)"
          >
            {{ item.title }}
          </span>
          <y-icon
            v-if="true"
            class="close-icon-item"
            icon-type="CloseCircleOutline"
            @click="iconClick(item.name)"
          />
        </n-button>
      </n-scrollbar>
    </div>
    <y-icon
      v-if="rightArrowDisabled"
      icon-type="ChevronForward"
      class="icon-location"
      :size="18"
      @click="rightArrowClick"
    />
  </div>
</template>

<style lang="sass" scoped>
.tab-bar
  position: relative
  //box-sizing: border-box
  //white-space: nowrap
  display: flex
  height: 40px
  box-shadow: 10px 5px 10px rgb(0 0 0 / 10%)

  .icon-location
    position: static
    //display: inline-block
    cursor: pointer
    font-size: 20px
    margin: 8px 8px

  .tabs
    //justify-content: flex-start
    //display: inline-block
    width: 100%
    margin-left: 0
    white-space: nowrap
    overflow: hidden

    .tab-item
      padding: 20px 15px
      border-radius: 0
      //margin: 5px 5px
      cursor: pointer
      .text-item
        //color: red
      // todo 字体颜色
      .show-icon-item
        margin-top: 2px
        margin-left: 2px

      .close-icon-item
        margin-left: 0
        width: 0
        height: 0
        transition: all 0.2s ease-in-out
        overflow: hidden

      & :hover
        .close-icon-item
          display: inline
          width: 14px
          height: 14px
          margin-left: 5px
          font-size: 12px
          background-color: rgba(0, 0, 0, 0)
          border-radius: 50%
          padding: 1px
          transition: all 0.2s ease-in-out
</style>
