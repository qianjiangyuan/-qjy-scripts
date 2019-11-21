import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/workspace',
  name: 'workspace',
  meta,
  component: layoutHeaderAside,
  children: [
    {
      path: 'add',
      name: 'workspace-add',
      component: () => import('@/pages/workspace/workspace-add'),
      meta: { ...meta, title: '工作空间创建' }
    },
    {
      path: 'list',
      name: 'workspace-list',
      component: () => import('@/pages/workspace/workspace-list'),
      meta: { ...meta, title: '工作空间列表' }
    }
  ]
}
