import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/user',
  name: 'user',
  meta,
  component: layoutHeaderAside,
  children: [
    {
      path: 'list',
      name: 'usermanager',
      component: () => import('@/pages/admin/user'),
      meta: { ...meta, title: '用户管理' }
    }
  ]
}
