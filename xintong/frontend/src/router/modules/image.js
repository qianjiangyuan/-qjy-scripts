import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/image',
  name: 'image',
  meta,
  component: layoutHeaderAside,
  children: [
    {
      path: 'upload',
      name: 'image-upload',
      component: () => import('@/pages/image/image-upload'),
      meta: { ...meta, title: '镜像上传' }
    },
    {
      path: 'list',
      name: 'image-list',
      component: () => import('@/pages/image/image-list'),
      meta: { ...meta, title: '镜像列表' }
    }
  ]
}
