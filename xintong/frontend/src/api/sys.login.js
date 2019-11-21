import request from '@/plugin/axios'
import env from '@/api/env'

export function AccountLogin (data) {
  return request({
    url: `${env.url()}/auth/login`,
    method: 'post',
    data
  })
}

export function GithubLogin (data) {
  return request({
    url: `${env.url()}/auth/github-login`,
    method: 'post',
    data
  })
}
