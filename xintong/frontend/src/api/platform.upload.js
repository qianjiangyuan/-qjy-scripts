import request from '@/plugin/axios'
import env from '@/api/env'

export function uploadUrl () {
  return env.url() + '/upload/upload'
}
export function mergeUpload (data) {
  return request({
    url: `${env.url()}/upload/mergeUpload`,
    method: 'post',
    data
  })
}

export function ListFiles (data) {
  return request({
    url: `${env.url()}/upload/list`,
    method: 'post',
    data
  })
}
