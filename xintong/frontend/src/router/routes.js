import image from './modules/image'
import dataset from './modules/dataset'
import workspace from './modules/workspace'
import user from './modules/user'
import train from './modules/train'
import dlmodel from './modules/dlmodel'
import dlservice from './modules/dlservice'
import codeHouse from './modules/code_house'
import group from './modules/group'
import layoutHeaderAside from '@/layout/header-aside'

/**
 * 在主框架内显示
 */
const frameIn = [
  // {
  //   path: '/',
  //   name: 'home',
  //   meta: {
  //     auth: true
  //   },
  //   component: () => import('@/pages/home')
  // },
  {
    path: '/',
    redirect: { name: 'index' },
    component: layoutHeaderAside,
    children: [
      // 首页 必须 name:index
      {
        path: 'index',
        name: 'index',
        meta: {
          auth: true
        },
        component: () => import('@/pages/index')
      },
      {
        path: 'admin',
        name: 'admin',
        meta: {
          auth: true,
          title: '管理员'
        },
        component: () => import('@/pages/admin')
      },
      {
        path: '/devworkspace',
        name: 'devworkspace',
        meta: {
          auth: true,
          title: '开发空间'
        },
        component: () => import('@/pages/devworkspace')
      },
      {
        path: '/resources',
        name: 'resources',
        meta: {
          auth: true,
          title: '资源监控'
        },
        component: () => import('@/pages/resources')
      },
      {
        path: 'profile',
        name: 'profile',
        meta: {
          auth: true,
          title: '个人中心'
        },
        component: () => import('@/pages/profile')
      },
      // 刷新页面 必须保留
      {
        path: 'refresh',
        name: 'refresh',
        hidden: true,
        component: {
          beforeRouteEnter (to, from, next) {
            next(vm => vm.$router.replace(from.fullPath))
          },
          render: h => h()
        }
      },
      // 页面重定向 必须保留
      {
        path: 'redirect/:route*',
        name: 'redirect',
        hidden: true,
        component: {
          beforeRouteEnter (to, from, next) {
            next(vm => vm.$router.replace(JSON.parse(from.params.route)))
          },
          render: h => h()
        }
      }
    ]
  },
  image,
  dataset,
  workspace,
  user,
  train,
  dlservice,
  dlmodel,
  group,
  codeHouse
]

/**
 * 在主框架之外显示
 */
const frameOut = [
  // 登录
  {
    path: '/login',
    name: 'login',
    component: () => import('@/pages/login')
  }
]

/**
 * 错误页面
 */
const errorPage = [
  // 404
  {
    path: '*',
    name: '404',
    component: () => import('@/pages/error-page-404')
  }
]

// 导出需要显示菜单的
export const frameInRoutes = frameIn

// 重新组织后导出
export default [
  ...frameIn,
  ...frameOut,
  ...errorPage
]
