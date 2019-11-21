import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/dlservice',
  name: 'dlservice',
  meta,
  component: layoutHeaderAside,
  children: [
    {
      path: 'add',
      name: 'dlservice-add',
      component: () => import('@/pages/dlservice/dlservice-add'),
      meta: { ...meta, title: '服务发布' }
    },
    {
      path: 'list',
      name: 'image-list',
      component: () => import('@/pages/dlservice/dlservice-list'),
      meta: { ...meta, title: '服务列表' }
    }
  ]
}
