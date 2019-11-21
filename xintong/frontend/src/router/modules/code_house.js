import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/code_house',
  name: 'code_house',
  meta,
  component: layoutHeaderAside,
  children: [
    {
      path: 'add',
      name: 'code-house-add',
      component: () => import('@/pages/code_house/add'),
      meta: { ...meta, title: '代码库上传' }
    },
    {
      path: 'list',
      name: 'code-house-list',
      component: () => import('@/pages/code_house/list'),
      meta: { ...meta, title: '代码库列表' }
    }
  ]
}
