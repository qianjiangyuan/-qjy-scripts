<template>
    <d2-container>
        <el-form v-if="!showDetail" ref="form" :inline="true" class="demo-form-inline">
            <div >
                <el-form-item style="">
                    <el-input v-model="name" placeholder="请输入名称"></el-input>
                </el-form-item>
                <el-form-item style="">
                    <el-button type="primary" size="medium" @click="queryList">查询</el-button>
                </el-form-item>
            </div>
            <div>
                <div class="inptH32 mg-b10 pd-b5 border-b1">
                    <el-form-item style="">
                        <el-button size="small" type="primary" @click="addTrain_task">新增</el-button>
                    </el-form-item>
                    <el-form-item style="">
                        <el-button size="small"  @click="deleteTrain_task" :disabled = hideDelete>删除</el-button>
                    </el-form-item>
                </div>
                <div>
                    <el-table :data="dataList" :stripe=true style="width: 100%" v-loading="loading" element-loading-text="加载中" @selection-change="handleSelectionChange" @row-dblclick="detail" highlight-current-row>
                        <el-table-column type="selection" width="55">
                        </el-table-column>
                        <el-table-column
                                prop="name"
                                label="名称">
                        </el-table-column>
                        <el-table-column
                                prop="dataset.name"
                                label="数据集">
                        </el-table-column>
                        <el-table-column
                                prop="codehouse.name"
                                label="代码">
                        </el-table-column>
                        <el-table-column
                                prop="pod_num"
                                label="pod数量">
                            <template slot-scope="scope">
                                {{scope.row.cpu_num}} 核心
                            </template>
                        </el-table-column>
                        <el-table-column
                                prop="cpu_num"
                                label="cpu数量">
                            <template slot-scope="scope">
                                {{scope.row.cpu_num}} 核心
                            </template>
                        </el-table-column>
                        <el-table-column
                                prop="gpu_num"
                                label="GPU数量">
                            <template slot-scope="scope">
                                {{scope.row.gpu_num}} 卡
                            </template>
                        </el-table-column>
                        <el-table-column
                                prop="memory_num"
                                label="内存数量">
                            <template slot-scope="scope">
                                {{scope.row.memory_num}} Gi
                            </template>
                        </el-table-column>
                        <el-table-column fixed="right" label="操作"  >
                            <template slot-scope="scope">

                                   <i class="el-icon-edit" @click="edit(scope.$index)" title="编辑"></i>
                                    <!--<el-button @click="attachPod(scope.$index)" size="small" type="success" icon="el-icon-check" round>SSH</el-button>
                                    <el-button size="small" type="info" icon="el-icon-magic-stick" round>释放</el-button>
                                    -->
                            </template>
                        </el-table-column>
                    </el-table>
                    <el-dialog
                            width="1020px"
                            title="ssh to server"
                            :visible.sync="showSSH"
                            :fullscreen="false"
                            :close-on-click-modal="false"
                            :center="true"
                            @close="close_ssh">
                        <div class="line"></div>
                        <el-tabs v-model="activePod" @tab-click="podTabClick">
                            <el-tab-pane
                                    :name="`pod_${n-1}`"
                                    :label="`pod_${n-1}`"
                                    v-for="n in item.pod_num"
                                    :key="n">
                                <div :id="`xterm-${n-1}`"></div>
                            </el-tab-pane>
                        </el-tabs>
                    </el-dialog>
                </div>
                <div class="block" style="margin-top: 10px;">
                    <!--<br>-->
                    <div class="r_page">
                        <el-pagination
                                @size-change="handleSizeChange"
                                @current-change="handleCurrentChange"
                                :current-page="pageNum"
                                :page-sizes="[10, 50, 200, 500]"
                                :page-size="pageSize"
                                layout="total, sizes, prev, pager, next, jumper"
                                :total="total">
                        </el-pagination>
                    </div>
                </div>
            </div>
        </el-form>
        <template v-if="showDetail">
            <train-task-form :data="item"  @oncancel="cancelTrain_task" @onsuccess="saveTrain_task" :readOnly="readOnly" ></train-task-form>
        </template>
    </d2-container>
</template>

<script>
import { QueryTrainTaskList, DeleteTrainTaskList } from '@/api/platform.trainTask'
import TrainTaskForm from './TrainTaskForm'

import 'xterm/src/xterm.css'
import SockJS from 'sockjs-client'
import { Terminal } from 'xterm'
import * as fit from 'xterm/lib/addons/fit/fit'
import * as attach from 'xterm/lib/addons/attach/attach'

Terminal.applyAddon(fit)
Terminal.applyAddon(attach)

export default {
  name: 'Train_taskList',
  components: {
    TrainTaskForm
  },
  props: {
    datas: {
      type: Array,
      default: function () {
        return []
      }
    },
    type: {
      type: String,
      require: true,
      default: 'parent'
    }
  },
  data () {
    return {
      name: null,
      value: '',
      dataList: this.datas,
      item: {},
      multipleSelection: [],
      pageSize: 10,
      pageNum: 1,
      total: null,
      loading: false,
      showDetail: false,
      hideDelete: false,
      readOnly: false,
      showSSH: false,
      activePod: 'pod_0',
      activeTask: {}
    }
  },
  mounted () {
    this.queryList()
  },
  created () {
  },
  methods: {
    podTabClick (tab, event) {
      console.log(tab.index)
      this.open_ssh(tab.index)
    },
    close_ssh () {
      console.log('close ...')
    },
    open_ssh (index) {
      this.__sock && this.__sock.close()
      this.__term && this.__term.destroy()
      console.log(index)

      let term = new Terminal()
      term.open(document.getElementById('xterm-0'))
      console.log(term)
      console.log(document.getElementById('xterm-0'))
      console.log(document.getElementById('#xterm-0'))
      term.fit()
      term.clear()

      let sock = new SockJS('http://127.0.0.1:8080/train/xterm?index=' + index + '&task_name=' + this.item.name)
      sock.onopen = () => {
        sock.send(JSON.stringify({
          Op: 'resize',
          Cols: term.cols,
          Rows: term.rows
        }))
      }

      sock.onmessage = (e) => {
        console.log('sock message: ', e.data)
        const msg = JSON.parse(e.data)
        console.log(msg)
        term.write(msg.Data)
      }

      sock.onclose = () => {
        console.log('sock close')
        this.__sock.close()
      }

      term.onData(e => {
        this.__sock && this.__sock.send(JSON.stringify({
          Op: 'stdin',
          Data: e
        }))
      })

      term.onResize(() => {
        sock.send(JSON.stringify({
          Op: 'resize',
          Cols: term.cols,
          Rows: term.rows
        }))
      })

      window.addEventListener('resize', () => {
        term.fit()
      })

      this.__term = term
      this.__sock = sock
    },
    queryList () {
      this.loading = true
      let that = this
      let queryparams = {
        params: {
          name: this.name
        },
        pageNum: this.pageNum,
        pageSize: this.pageSize
      }
      QueryTrainTaskList(queryparams).then(res => {
        this.loading = false
        that.dataList = res
      }).catch(res => {
        this.loading = false
      })
    },
    addTrain_task () {
      this.item = {}
      this.showDetail = true
    },
    deleteTrain_task () {
      let ids = []
      this.multipleSelection.forEach(item => {
        ids.push(item.id)
      })
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        DeleteTrainTaskList(ids).then(res => {
          this.$message({
            message: '删除成功',
            type: 'success'
          })
          this.queryList()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    detail (row, event) {
      this.item = row
      this.readOnly = true
      this.showDetail = true
    },
    edit (inx) {
      this.item = this.dataList[inx]
      this.readOnly = false
      this.showDetail = true
    },
    saveTrain_task (item) {
      this.queryList()
      this.showDetail = false
    },
    attachPod: function (inx) {
      this.item = this.dataList[inx]
      this.showSSH = true
    },
    cancelTrain_task () {
      this.showDetail = false
      this.readOnly = false
    },
    handleSelectionChange (val) {
      this.multipleSelection = val
      if (val.length > 0) {
        this.hideDelete = false
      } else {
        this.hideDelete = true
      }
    },
    handleCurrentChange (val) {
      this.pageNum = val
      this.queryList()
    },
    handleSizeChange (val) {
      this.pageSize = val
      this.queryList()
    }

  }
}
</script>
<style>
.xterm {
    min-height: 600px;
}
</style>
