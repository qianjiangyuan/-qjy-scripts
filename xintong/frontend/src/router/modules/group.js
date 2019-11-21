import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/group',
  name: 'group',
  meta,
  component: layoutHeaderAside,
  children: [
    {
      path: 'add',
      name: 'group-add',
      component: () => import('@/pages/group/add'),
      meta: { ...meta, title: ' 组创建' }
    },
    {
      path: 'list',
      name: 'group-list',
      component: () => import('@/pages/group/list'),
      meta: { ...meta, title: '组列表' }
    }
  ]
}
