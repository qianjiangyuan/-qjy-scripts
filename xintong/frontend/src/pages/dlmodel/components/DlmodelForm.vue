<template>
  <el-form ref="dlmodelform" :inline="true"  :model="item"   :rules="rules" class="demo-form-inline" label-width="120px">
    <div>
      <el-form-item prop="name" label="名称">
      <el-input  v-model="item.name" :disabled="readOnly" placeholder="请输入名称"></el-input>
    </el-form-item>
      <el-form-item prop="model_name" label="模型名称">
        <el-input  v-model="item.model_name" :disabled="readOnly" placeholder="请输入名称"></el-input>
      </el-form-item>
      <el-form-item prop="version" label="版本">
        <el-input  v-model="item.version" :disabled="readOnly" placeholder="请输入版本"></el-input>
      </el-form-item>
      <el-form-item prop="description" label="描述">
        <el-input  v-model="item.description" :disabled="readOnly" type="textarea" :rows="2" placeholder="请输入描述"></el-input>
      </el-form-item>
        <div></div>
        <el-form-item label="标签" prop="label">
        <el-select  v-model="item.label" :disabled="readOnly"
          placeholder="选择标签">
          <el-option
            v-for="item2 in labelDicts"
            :key="item2.value"
            :label="item2.label"
            :value="item2.value">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="框架" prop="platform">
        <el-select  v-model="item.platform" :disabled="readOnly"
          placeholder="选择框架">
          <el-option
            v-for="item2 in platformDicts"
            :key="item2.value"
            :label="item2.label"
            :value="item2.value">
          </el-option>
        </el-select>
      </el-form-item>
      <div></div>
      <el-form-item label="文件列表:" prop="fileList" v-if="uploadShow == false">
        <el-table
                :data="item.Files"
                style="width: 100%"
                :show-header= false
                height="300px">
          <el-table-column  width="80" align="right">
            <template slot-scope="scope" >
              <img width="15" v-if="scope.row.IsDir" src="./image/dir.png">
              <img  width="15" v-if="!scope.row.IsDir" src="./image/file.png">
            </template>
          </el-table-column>
          <el-table-column
                  prop="name"
                  label="文件名"
                  width="580">
          </el-table-column>

        </el-table>
        <div class="block" style="margin-top: 10px;">
          <!--<br>-->
          <div class="r_page">
            <el-pagination
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
                    :current-page="pageNum"
                    :page-sizes="[100, 200, 500]"
                    :page-size="pageSize"
                    layout="total, sizes, prev, pager, next, jumper"
                    :total="file_total">
            </el-pagination>
          </div>
        </div>
      </el-form-item>
          <el-form-item label="文件上传" prop="platform"  v-if="uploadShow">
          <uploader :options="options"  class="uploader-example">
            <uploader-unsupport></uploader-unsupport>
            <uploader-drop>
              <p>Drop files here to upload or</p>
              <uploader-btn >选择文件</uploader-btn>
              <uploader-btn  :directory="true">选择文件夹</uploader-btn>
            </uploader-drop>
            <uploader-list></uploader-list>
          </uploader>
        </el-form-item>
    </div>
    <el-tabs type="card" style="margin-top: 15px;">
    </el-tabs>
    <div>
      <div class="fr" style="">
        <el-form-item>
          <template v-if="!readOnly">
            <el-button type="primary" size="small" @click="save">保存</el-button>
          </template>
          <el-button type="primary" size="small" @click="cancel" class="btnReturn" plain>返回</el-button>
        </el-form-item>
      </div>
    </div>
  </el-form>
</template>
<script>
import { SaveDlmodel, UpdateDlmodel, DetailDlmodel, QueryDlModeFiles } from '@/api/platform.dlmodel'
import env from '@/api/env'
import util from '@/libs/util'
export default {
  name: 'DlmodelForm',
  components: {
  },
  props: {
    data: {
      type: Object,
      default: function () {
        return {}
      }
    },
    trainTask: {
      type: Object
    },
    child: {
      type: Boolean,
      default: false
    },
    readOnly: {
      type: Boolean,
      default: false
    }
  },
  data: function () {
    return {
      item: JSON.parse(JSON.stringify(this.data)),
      trainTaskIdTrain_task: [],
      train_taskList: [],
      labelDicts: [],
      platformDicts: [],
      fileList: [],
      file_total: 0,
      pageSize: 100,
      pageNum: 1,
      uploadShow: true,
      options: {
        // 可通过 https://github.com/simple-uploader/Uploader/tree/develop/samples/Node.js 示例启动服务
        target: null,
        testChunks: false
      },
      rules: {
        name: [{ required: true, message: '名称', trigger: 'blur' }],
        model_name: [{ required: true, message: '请输入模型', trigger: 'blur' }],
        platform: [{ required: true, message: '请选择框架' }]
      }
    }
  },
  created: async function () {
    this.labelDicts = util.labelDicts
    this.platformDicts = util.platformDicts
    if (this.data.id != null) {
      this.detail(this.data.id)
      this.uploadShow = false
    } else {
      this.initUuid(util.uuid())
      this.uploadShow = true
    }
  },
  mounted () {
    if (this.data.id != null) {
      this.detail(this.data.id)
    }
    if (this.trainTask != null) {
      this.item.name = this.trainTask.name
      this.uuid = this.trainTask.modelPath
      this.item.url = this.trainTask.modelPath
      this.uploadShow = false
      this.queryFileList()
    }
  },
  methods: {
    queryFileList () {
      let queryparams = {
        params: {
          url: this.item.url
        },
        pageNum: this.pageNum,
        pageSize: this.pageSize
      }
      QueryDlModeFiles(queryparams).then(res => {
        this.item.Files = res.Files
        this.file_total = res.FileTotal
      })
    },
    cancel () {
      this.$emit('oncancel', this.item)
    },
    initUuid (uuid) {
      this.uuid = uuid
      this.options.target = env.url() + '/upload/upload?path=' + uuid
    },
    detail (id) {
      let that = this
      DetailDlmodel({ id: id }).then(res => {
        that.item = res
        this.initUuid(that.item.url)
      })
    },
    save: function () {
      this.$refs['dlmodelform'].validate((valid) => {
        if (valid) {
          let that = this
          if (this.child !== false) {
            that.$emit('onsuccess', that.item)
          } else {
            if (this.item.id == null) {
              this.item.url = this.uuid
              SaveDlmodel(this.item).then(res => {
                this.$message({ message: '保存成功', type: 'success' })
                that.item = res
                this.$emit('onsuccess', this.item)
              })
            } else {
              UpdateDlmodel(this.item).then(res => {
                this.$message({ message: '保存成功', type: 'success' })
                that.item = res
                this.$emit('onsuccess', this.item)
              })
            }
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>
<style scoped>
  .el-input {
    width: 400px;
  }
  .el-input {
    width: 400px;
  }
  .el-textarea {
    width: 400px;
  }
  .el-select  {
    width: 200px;
  }
  .upload-demo {
    width: 400px;
  }
  .input-with-select .el-input-group__prepend {
    background-color: #fff;
  }
  .uploader-example {
      width: 880px;
      padding: 15px;
      margin: 0px auto 0;
      font-size: 12px;
  }
  .uploader-example .uploader-btn {
      margin-right: 4px;
  }
  .uploader-example .uploader-list {
      max-height: 440px;
      overflow: auto;
      overflow-x: hidden;
      overflow-y: auto;
  }
</style>
