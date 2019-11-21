import request from '@/plugin/axios'
import env from '@/api/env'

export function SaveCodeHouse (data) {
  return request({
    url: `${env.url()}/code-house/save`,
    method: 'post',
    data
  })
}
export function UpdateCodeHouse (data) {
  return request({
    url: `${env.url()}/code-house/update`,
    method: 'post',
    data
  })
}
export function QueryCodeHouseList (data) {
  return request({
    url: `${env.url()}/code-house/queryPageList`,
    method: 'post',
    data
  })
}

export function DeleteCodeHouseList (data) {
  return request({
    url: `${env.url()}/code-house/deletes`,
    method: 'post',
    data
  })
}

export function DetailCodeHouse (data) {
  return request({
    url: `${env.url()}/code-house/detail?id=${data.id}`,
    method: 'get',
    data
  })
}
