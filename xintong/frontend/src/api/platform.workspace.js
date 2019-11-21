import request from '@/plugin/axios'
import env from '@/api/env'

export function SaveWorkspace (data) {
  return request({
    url: `${env.url()}/workspace/save`,
    method: 'post',
    data
  })
}

export function UpdateWorkspace (data) {
  return request({
    url: `${env.url()}/workspace/update`,
    method: 'post',
    data
  })
}
export function QueryWorkspaceList (data) {
  return request({
    url: `${env.url()}/workspace/queryPageList`,
    method: 'post',
    data
  })
}

export function DeleteWorkspaceList (data) {
  return request({
    url: `${env.url()}/workspace/deletes`,
    method: 'post',
    data
  })
}

export function DetailWorkspace (data) {
  return request({
    url: `${env.url()}/workspace/detail?id=${data.id}`,
    method: 'get',
    data
  })
}
