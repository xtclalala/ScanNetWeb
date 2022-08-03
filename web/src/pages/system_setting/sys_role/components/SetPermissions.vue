<script lang="ts">
export default {
  name: 'SetPermissions',
}
</script>

<script setup lang="ts">
import { reactive, ref, toRef } from 'vue'
import { BaseMenu, SelectMenu } from '@/api/system_setting/types/sys_menu'
import { allMenu } from '@/api/system_setting/sys_menu'
import { list2Tree } from '@/utils/helper/treeHelper'
import { roleAddMenu, roleAddPer } from '@/api/system_setting/sys_role'

type PerProps = {
  show: boolean
  currentRole: number
  defaultMenuData: number[]
  defaultPerData: number[]
}
const props = withDefaults(defineProps<PerProps>(), {
  show: false,
  currentRole: 0,
})
const emit = defineEmits(['update:show', 'afterUpdate'])

const currentTab = ref<string>('Menu')
const show = toRef(props, 'show')
const menuCheckedKeys = toRef(props, 'defaultMenuData')
const perCheckedKeys = toRef(props, 'defaultPerData')
const menuData = ref<SelectMenu[]>([])
const perData = ref<SelectMenu[]>([])

let menuUpdateKeys = reactive<number[]>(menuCheckedKeys.value)
let perUpdateKeys = reactive<number[]>(perCheckedKeys.value)

const allMenus = async () => {
  const res = await allMenu<BaseMenu[]>({ isMessage: false })
  menuData.value = list2Tree<BaseMenu>(res)
  perData.value = res.flatMap((item) => {
    item.disabled = true
    return item
  })
}
const closeDrawer = () => {
  emit('update:show', false)
}
const menuUpdateCheckedKeys = (v: Array<number>) => {
  menuUpdateKeys = [...v]
}
const perUpdateCheckedKeys = (v: number[]) => {
  perUpdateKeys = [...v]
}
const selectValue = async (v: string) => {
  currentTab.value = v
}
const handleUpdate = async () => {
  if (currentTab.value === 'Menu') {
    await roleAddMenu({ id: props.currentRole, menuIds: menuUpdateKeys })
  } else {
    await roleAddPer({ id: props.currentRole, permissionIds: perUpdateKeys })
  }
  emit('afterUpdate', props.currentRole)
}

allMenus()
</script>

<template>
  <n-drawer v-model:show="show" :width="700" @mask-click="closeDrawer">
    <n-drawer-content>
      <n-tabs
        type="line"
        animated
        default-value="Menu"
        :value="currentTab"
        @update:value="selectValue"
      >
        <n-tab-pane name="Menu" tab="页面">
          <n-tree
            block-line
            cascade
            checkable
            :data="menuData"
            key-field="id"
            label-field="title"
            :default-checked-keys="menuCheckedKeys"
            @update:checked-keys="menuUpdateCheckedKeys"
          />
        </n-tab-pane>
        <n-tab-pane name="Permission" tab="按钮">
          <n-tree
            block-line
            cascade
            checkable
            default-expand-all
            :data="perData"
            key-field="id"
            label-field="title"
            children-field="permissions"
            :default-checked-keys="perCheckedKeys"
            @update:checked-keys="perUpdateCheckedKeys"
          />
        </n-tab-pane>
      </n-tabs>

      <template #footer>
        <n-button @click="handleUpdate">更新</n-button>
      </template>
    </n-drawer-content>
  </n-drawer>
</template>

<style scoped></style>
