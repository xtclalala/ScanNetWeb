type ClientViewType = {
  width: number
  height: number
}

export const useClientView = (): ClientViewType => {
  return {
    width: window.screen.availWidth,
    height: window.screen.availHeight,
  }
}
