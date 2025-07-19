import { requestAjax } from '@/api/http.ts'
import { FormDataVO } from '@/utils/interfaces/form.ts'

export function getFormConfig (params?: any) {
  return requestAjax({
    url: '/initialize',
    method: 'get',
    params
  })
}

export function postFormConfig (data: FormDataVO) {
  for (const key in data) {
    if (key.includes('Port') ||
      key.includes('Max') ||
      key.includes('Length') ||
      key.includes('Size') ||
      key.includes('Count') ||
      key.includes('RedisDB')) {
      data[key] = Number(data[key])
    }
  }
  return requestAjax({
    url: '/initialize',
    method: 'post',
    data
  })
}

export function getCopyrightConfig (params?: any) {
  return requestAjax({
    url: '/copyright',
    method: 'get',
    params
  })
}
