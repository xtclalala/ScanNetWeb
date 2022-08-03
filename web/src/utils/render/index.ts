/**
 * @Description: src\utils\render\index.ts
 * @author: y1t
 * @date 2022/7/26
 **/
import { h } from 'vue'
import { YIcon } from '@/components'

/**
 * 渲染图标
 * @param icon
 */
export const renderIcon = (icon: string) => {
  return () => h(YIcon, { iconType: icon })
}
