// 菜单 侧边栏
export default [
  {
    title: '首页',
    icon: 'home',
    path: '/'
  },
  {
    title: '开发空间',
    icon: 'folder',
    path: '/devworkspace'

  },
  {
    title: '数据管理',
    icon: 'folder',
    children: [
      { title: '数据上传', path: '/dataset/upload' },
      { title: '数据列表', path: '/dataset/list' }
    ]
  },
  {
    title: '代码库',
    icon: 'cube',
    children: [
      { title: '代码库管理', path: '/code_house/add' },
      { title: '代码库列表', path: '/code_house/list' }
    ]
  },
  {
    title: '镜像管理',
    icon: 'cube',
    children: [
      { title: '镜像上传', path: '/image/upload' },
      { title: '镜像列表', path: '/image/list' }
    ]
  },
  {
    title: '训练任务',
    icon: 'tasks',
    children: [
      { title: '任务提交', path: '/train/add' },
      { title: '任务列表', path: '/train/list' }
    ]
  },
  {
    title: '模型管理',
    icon: 'tasks',
    children: [
      { title: '模型提交', path: '/dlModel/add' },
      { title: '模型列表', path: '/dlModel/list' }
    ]
  },
  {
    title: '服务管理',
    icon: 'tasks',
    children: [
      { title: '服务发布', path: '/dlService/add' },
      { title: '服务列表', path: '/dlService/list' }
    ]
  },
  {
    title: '日志管理',
    icon: 'cube',
    path: '/log'
  },
  {
    title: '资源监控',
    icon: 'cube',
    path: '/resources'
  },
  {
    title: '工作空间管理',
    icon: 'cube',
    administrator: true,
    children: [
      { title: '创建工作空间', path: '/workspace/add' },
      { title: '工作空间列表', path: '/workspace/list' }
    ]
  },
  {
    title: '组管理',
    icon: 'cube',
    administrator: true,
    children: [
      { title: '组管理', path: '/group/add' },
      { title: '组列表', path: '/group/list' }
    ]
  },

  {
    title: '用户管理',
    administrator: true,
    icon: 'cube',
    children: [
      { title: '用户列表', path: '/user/list' }
    ]
  }
]
