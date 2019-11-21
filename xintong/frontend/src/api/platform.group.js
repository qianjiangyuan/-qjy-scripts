import request from '@/plugin/axios'
import env from '@/api/env'

export function SaveGroup (data) {
  return request({
    url: `${env.url()}/group/save`,
    method: 'post',
    data
  })
}
export function UpdateGroup (data) {
  return request({
    url: `${env.url()}/group/update`,
    method: 'post',
    data
  })
}
export function QueryGroupList (data) {
  return request({
    url: `${env.url()}/group/queryPageList`,
    method: 'post',
    data
  })
}

export function DeleteGroupList (data) {
  return request({
    url: `${env.url()}/group/deletes`,
    method: 'post',
    data
  })
}

export function DetailGroup (data) {
  return request({
    url: `${env.url()}/group/detail?id=${data.id}`,
    method: 'get',
    data
  })
}
