import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/train',
  name: 'train',
  meta,
  component: layoutHeaderAside,
  children: [
    {
      path: 'add',
      name: 'trainTaskList',
      component: () => import('@/pages/train/trainTask-add'),
      meta: { ...meta, title: '任务列表' }
    },
    {
      path: 'list',
      name: 'trainTask',
      component: () => import('@/pages/train/trainTask-list'),
      meta: { ...meta, title: '任务训练' }
    }
  ]
}
