import request from '@/plugin/axios'
import env from '@/api/env'

export function SaveDlmodel (data) {
  return request({
    url: `${env.url()}/dlmodel/save`,
    method: 'post',
    data
  })
}
export function UpdateDlmodel (data) {
  return request({
    url: `${env.url()}/dlmodel/update`,
    method: 'post',
    data
  })
}
export function QueryDlmodelList (data) {
  return request({
    url: `${env.url()}/dlmodel/queryPageList`,
    method: 'post',
    data
  })
}

export function DeleteDlmodeList (data) {
  return request({
    url: `${env.url()}/dlmodel/deletes`,
    method: 'post',
    data
  })
}
export function QueryDlModeFiles (data) {
  return request({
    url: `${env.url()}/dlmodel/queryPageFiles`,
    method: 'post',
    data
  })
}

export function DetailDlmodel (data) {
  return request({
    url: `${env.url()}/dlmodel/detail?id=${data.id}`,
    method: 'get',
    data
  })
}
