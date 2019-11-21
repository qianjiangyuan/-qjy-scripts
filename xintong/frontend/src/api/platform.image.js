import request from '@/plugin/axios'
import env from '@/api/env'

export function SaveImage (data) {
  return request({
    url: `${env.url()}/image/save`,
    method: 'post',
    data
  })
}
export function UpdateImage (data) {
  return request({
    url: `${env.url()}/image/update`,
    method: 'post',
    data
  })
}
export function QueryImageList (data) {
  return request({
    url: `${env.url()}/image/queryPageList`,
    method: 'post',
    data
  })
}

export function DeleteImageList (data) {
  return request({
    url: `${env.url()}/image/deletes`,
    method: 'post',
    data
  })
}

export function DetailImage (data) {
  return request({
    url: `${env.url()}/image/detail?id=${data.id}`,
    method: 'get',
    data
  })
}
