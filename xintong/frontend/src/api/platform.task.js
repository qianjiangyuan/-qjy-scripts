import request from '@/plugin/axios'
import env from '@/api/env'

export function TaskCreate (data) {
  return request({
    url: `${env.url()}/task/create`,
    method: 'post',
    data
  })
}

export function TaskList (params) {
  return request({
    url: `${env.url()}/task/list`,
    method: 'get',
    params
  })
}

export function TaskDel (params) {
  return request({
    url: `${env.url()}/task/del/${params.id}`,
    method: 'get'
  })
}

export function TaskLog (params) {
  return request({
    url: `${env.url()}/task/log/${params.id}`,
    method: 'get',
    params: {
      pid: params.pid
    }
  })
}

export function TaskConf (params) {
  return request({
    url: `${env.url()}/task/conf/${params.id}`,
    method: 'get'
  })
}
