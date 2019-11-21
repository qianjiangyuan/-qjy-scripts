import request from '@/plugin/axios'
import env from '@/api/env'

export function SaveDataSet (data) {
  return request({
    url: `${env.url()}/dataset/save`,
    method: 'post',
    data
  })
}
export function UpdateDataSet (data) {
  return request({
    url: `${env.url()}/dataset/update`,
    method: 'post',
    data
  })
}
export function QueryDataSetList (data) {
  return request({
    url: `${env.url()}/dataset/queryPageList`,
    method: 'post',
    data
  })
}

export function DeleteDataSetList (data) {
  return request({
    url: `${env.url()}/dataset/deletes`,
    method: 'post',
    data
  })
}

export function QueryDataSetFiles (data) {
  return request({
    url: `${env.url()}/dataset/queryPageFiles`,
    method: 'post',
    data
  })
}

export function DetailDataSet (data) {
  return request({
    url: `${env.url()}/dataset/detail?id=${data.id}`,
    method: 'get',
    data
  })
}
