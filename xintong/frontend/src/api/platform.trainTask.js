import request from '@/plugin/axios'
import env from '@/api/env'

export function SaveTrainTask (data) {
  return request({
    url: `${env.url()}/train/save`,
    method: 'post',
    data
  })
}

export function QueryTrainTaskList (data) {
  return request({
    url: `${env.url()}/train/queryPageList`,
    method: 'post',
    data
  })
}

export function DeleteTrainTaskList (data) {
  return request({
    url: `${env.url()}/train/deletes`,
    method: 'post',
    data
  })
}

export function DetailTrainTask (data) {
  return request({
    url: `${env.url()}/train/detail?id=${data.id}`,
    method: 'get',
    data
  })
}

export function UpdateTrainTask (data) {
  return request({
    url: `${env.url()}/train/update`,
    method: 'post',
    data
  })
}

export function QueryModelFiles (data) {
  return request({
    url: `${env.url()}/train/queryPageFiles`,
    method: 'post',
    data
  })
}

export function QueryModelFileUrl (data) {
  return request({
    url: `${env.url()}/train/queryModelFileUrl`,
    method: 'post',
    data
  })
}

export function StopTrainTask (data) {
  return request({
    url: `${env.url()}/train/stopTrainTask?id=${data.id}`,
    method: 'get',
    data
  })
}
