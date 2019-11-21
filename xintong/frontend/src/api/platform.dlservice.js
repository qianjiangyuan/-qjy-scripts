import request from '@/plugin/axios'
import env from '@/api/env'

export function SaveDlservice (data) {
  return request({
    url: `${env.url()}/dlservice/save`,
    method: 'post',
    data
  })
}
export function UpdateDlservice (data) {
  return request({
    url: `${env.url()}/dlservice/update`,
    method: 'post',
    data
  })
}
export function QueryDlserviceList (data) {
  return request({
    url: `${env.url()}/dlservice/queryPageList`,
    method: 'post',
    data
  })
}

export function DeleteDlserviceList (data) {
  return request({
    url: `${env.url()}/dlservice/deletes`,
    method: 'post',
    data
  })
}

export function DetailDlservice (data) {
  return request({
    url: `${env.url()}/dlservice/detail?id=${data.id}`,
    method: 'get',
    data
  })
}
