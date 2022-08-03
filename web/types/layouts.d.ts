export interface Menu {
    id: string
    label: string
    icon?: string
    name?: string
    params?: { [key: string]: string }
    children?: Menu[]
}