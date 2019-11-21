import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/dataset',
  name: 'dataset',
  meta,
  // redirect: { name: 'demo-page1' },
  component: layoutHeaderAside,
  // children: (pre => [
  //   { path: 'page1', name: `${pre}page1`, component: () => import('@/pages/demo/page1'), meta: { ...meta, title: '页面 1' } },
  //   { path: 'page2', name: `${pre}page2`, component: () => import('@/pages/demo/page2'), meta: { ...meta, title: '页面 2' } },
  //   { path: 'page3', name: `${pre}page3`, component: () => import('@/pages/demo/page3'), meta: { ...meta, title: '页面 3' } }
  // ])('demo-')
  children: [
    {
      path: 'upload',
      name: 'dataset-upload',
      component: () => import('@/pages/dataset/dataset-upload'),
      meta: { ...meta, title: '数据上传' }
    },
    {
      path: 'list',
      name: 'dataset-list',
      component: () => import('@/pages/dataset/dataset-list'),
      meta: { ...meta, title: '数据列表' }
    }
  ]
}
