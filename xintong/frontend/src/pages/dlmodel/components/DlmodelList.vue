<template>
  <div>
    <template  v-if="showDetail==false">
      <el-form ref="form" :inline="true" class="demo-form-inline">
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
              <el-button size="small" type="primary" @click="addDlmodel">新增</el-button>
            </el-form-item>
            <el-form-item style="">
              <el-button size="small"  @click="deleteDlmodel" :disabled = hideDelete>删除</el-button>
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
                 prop="version"
                 label="版本">
               </el-table-column>
               <el-table-column
                 prop="trainTaskIdTrain_task.name"
                 label="训练任务">
               </el-table-column>
               <el-table-column
                 prop="label"
                 label="标签">
               </el-table-column>
               <el-table-column
                 prop="platform"
                 label="框架">
               </el-table-column>
               <el-table-column fixed="right" label="操作" width="80">
                 <template slot-scope="scope">
                   <i class="el-icon-edit" @click="edit(scope.$index)" title="编辑"></i>
                 </template>
               </el-table-column>
             </el-table>
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
      </template>
      <template v-if="showDetail">
        <dlmodel-form @oncancel="cancelAddOrUpdateModel" @onsuccess="commitAddOrUpdateModel" :data="item"  :readOnly="readOnly" ></dlmodel-form>
      </template>
   </div>
</template>

<script>
import { QueryDlmodelList, DeleteDlmodelList } from '@/api/platform.dlmodel'
import DlmodelForm from './DlmodelForm'
export default {
  name: 'DlmodelList',
  components: {
    DlmodelForm
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
      readOnly: false
    }
  },
  mounted () {
    this.queryList()
  },
  created () {
  },
  methods: {
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
      QueryDlmodelList(queryparams).then(res => {
        this.loading = false
        that.dataList = res
      }).catch(res => {
        this.loading = false
      })
    },
    addDlmodel () {
      this.item = {}
      this.showDetail = true
    },
    deleteDlmodel () {
      let ids = []
      this.multipleSelection.forEach(item => {
        ids.push(item.id)
      })
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        DeleteDlmodelList(ids).then(res => {
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
    edit: function (inx) {
      this.item = this.dataList[inx]
      this.readOnly = true
      this.showDetail = true
    },
    saveDlmodel (item) {
      this.queryList()
      this.showDetail = false
    },
    cancelDlmodel () {
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
    },

    cancelAddOrUpdateModel () {
      this.showDetail = false
    },
    commitAddOrUpdateModel () {
      this.showDetail = false
      this.queryList()
    }
  }
}
</script>
<style>
</style>
