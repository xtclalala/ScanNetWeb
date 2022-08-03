/**
 * @Description: listHelper
 * @author: y1t
 * @Date: 2022/7/25
 **/

/**
 * 排序
 * @param list 需要排序的列表
 * @param target 根据哪个字段排序
 * @param defaultNumber  默认值
 */
export const listSort = <T>(list: T[], target: string, defaultNumber: number) => {
  list.sort((a, b) => (a[target] || defaultNumber) - (b[target] || defaultNumber))
}
