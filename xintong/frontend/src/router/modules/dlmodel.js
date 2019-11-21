import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/dlmodel',
  name: 'dlmodel',
  meta,
  component: layoutHeaderAside,
  children: [
    {
      path: 'add',
      name: 'dlmodel-add',
      component: () => import('@/pages/dlmodel/dlmodel-add'),
      meta: { ...meta, title: '模型提交' }
    },
    {
      path: 'list',
      name: 'image-list',
      component: () => import('@/pages/dlmodel/dlmodel-list'),
      meta: { ...meta, title: '模型列表' }
    }
  ]
}
