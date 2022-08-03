/**
 * @Description: objectHelper
 * @author: y1t
 * @Date: 2022/7/25
 **/

/**
 * 合并传入对象的所有属性
 * @param sources Array<object>
 */
export function completeMerger<T>(...sources): T {
  const target = {}
  sources.forEach((item) => {
    Object.keys(item).reduce((prev, current) => {
      if (item[current] !== undefined) {
        prev[current] = item[current]
      }
      return prev
    }, target)
  })
  return target as T
}
