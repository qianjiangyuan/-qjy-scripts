import request from '@/plugin/axios'
import env from '@/api/env'

export function SaveUser (data) {
  return request({
    url: `${env.url()}/user/save`,
    method: 'post',
    data
  })
}
export function UpdateUser (data) {
  return request({
    url: `${env.url()}/user/update`,
    method: 'post',
    data
  })
}

export function QueryUserList (data) {
  return request({
    url: `${env.url()}/user/queryPageList`,
    method: 'post',
    data
  })
}

export function DeleteUserList (data) {
  return request({
    url: `${env.url()}/user/deletes`,
    method: 'post',
    data
  })
}

export function CommitUserPassWd (data) {
  return request({
    url: `${env.url()}/user/commitPassWd`,
    method: 'post',
    data
  })
}

export function DetailUser (data) {
  return request({
    url: `${env.url()}/user/detail?id=${data.id}`,
    method: 'get',
    data
  })
}
