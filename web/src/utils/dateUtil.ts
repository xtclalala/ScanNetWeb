/**
 * Independent time operation tool to facilitate subsequent switch to dayjs
 */

const DATE_TIME_FORMAT = 'YYYY-MM-DD HH:mm:ss'
const DATE_FORMAT = 'YYYY-MM-DD'

export function formatToDateTime(date: Date, format: string = DATE_TIME_FORMAT): string {
  const year: string = date.getFullYear().toString()
  const month: string = (date.getMonth() + 1).toString()
  const day: string = date.getDate().toString()
  const hour: string = date.getHours().toString()
  const minutes: string = date.getMinutes().toString()
  const seconds: string = date.getSeconds().toString()

  return format
    .replace('YYYY', year)
    .replace('MM', month)
    .replace('DD', day)
    .replace('HH', hour)
    .replace('mm', minutes)
    .replace('ss', seconds)
}

export function formatToDate(date: Date, format: string = DATE_FORMAT): string {
  const year: string = date.getFullYear().toString()
  const month: string = (date.getMonth() + 1).toString()
  const day: string = date.getDate().toString()

  return format.replace('YYYY', year).replace('MM', month).replace('DD', day)
}
