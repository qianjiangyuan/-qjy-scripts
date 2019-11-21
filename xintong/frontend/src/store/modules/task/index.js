import { TaskCreate, TaskList, TaskDel, TaskLog, TaskConf } from '@/api/platform.task'

export default {
  namespaced: true,
  actions: {
    create ({ dispatch }, data) {
      return new Promise((resolve, reject) => {
        TaskCreate(data).then(async rs => {
          resolve(rs)
        }).catch(reject)
      })
    },

    list ({ dispatch }, param) {
      return new Promise((resolve, reject) => {
        TaskList(param).then(async rs => {
          resolve(rs)
        }).catch(reject)
      })
    },

    delete ({ dispatch }, param) {
      return new Promise((resolve, reject) => {
        TaskDel(param).then(async rs => {
          resolve(rs)
        }).catch(reject)
      })
    },

    log ({ dispatch }, param) {
      return new Promise((resolve, reject) => {
        TaskLog(param).then(async rs => {
          resolve(rs)
        }).catch(reject)
      })
    },

    conf ({ dispatch }, param) {
      return new Promise((resolve, reject) => {
        TaskConf(param).then(async rs => {
          resolve(rs)
        }).catch(reject)
      })
    }
  }
}
