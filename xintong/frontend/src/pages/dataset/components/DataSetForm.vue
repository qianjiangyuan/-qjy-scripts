<template>
  <el-form ref="dataSetform" label-width="100px" :inline="true"  :model="item"   :rules="rules" class="demo-form-inline" >
    <div>
      <el-form-item prop="name" label="名称">
        <el-input  v-model="item.name"   :disabled="readOnly"   placeholder="请输入名称"></el-input>
      </el-form-item>
      <br/>
    <el-form-item prop="description" label="描述">
      <el-input  v-model="item.description"   :disabled="readOnly"   type="textarea" :rows="5" placeholder="请输入描述"></el-input>
    </el-form-item>
      <br/>
      <el-form-item prop="visibility"  label="类型">
        <el-radio v-model="item.visibility"  :disabled="readOnly"  label="public">公开</el-radio>
        <el-radio v-model="item.visibility"  :disabled="readOnly"  label="private">私有</el-radio>
      </el-form-item>
     <br/>
      <el-form-item label="框架" prop="platform">
        <el-select  v-model="item.platform"  :disabled="readOnly"
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
      <el-form-item label="标签" prop="label">
        <el-select  v-model="item.label" :disabled="readOnly"  allow-create filterable
                    placeholder="选择标签">
          <el-option
                  v-for="item2 in labelDicts"
                  :key="item2.value"
                  :label="item2.label"
                  :value="item2.value">
          </el-option>
        </el-select>
      </el-form-item>
      <div></div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item  >
                <el-button type="text" size="small"  @click="goPath(-1)">目录</el-button>
            </el-breadcrumb-item>
            <el-breadcrumb-item v-for="(path,index) in pathLevels" >
                <el-button type="text" size="small"  @click="goPath(index)">{{path}}</el-button>
            </el-breadcrumb-item>

        </el-breadcrumb>
        <el-form-item label="文件列表:" prop="fileList" v-if="uploadShow == false">
            <el-table
                    :data="item.Files"
                    style="width: 100%"
                    v-loading="loading"
                    @row-dblclick="nextLevel"
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

      <div></div>
        <el-form-item label="文件上传" prop="platform"  v-if="uploadShow" >
          <uploader :options="options"  @file-success="fileSuccess" @file-error="onFileError" @upload-start="fileUploadStart" class="uploader-example">
            <uploader-unsupport></uploader-unsupport>
            <uploader-drop>
              <p>Drop files here to upload or</p>
              <uploader-btn >选择文件</uploader-btn>
              <uploader-btn  :directory="true">选择文件夹</uploader-btn>
            </uploader-drop>
            <uploader-list></uploader-list>
          </uploader>
        </el-form-item>
        <div></div>
        <el-form-item  >
          <template v-if="!readOnly">
            <el-button  v-if="buttonVisible" type="primary" @click="save">保存</el-button>
          </template>
          <el-button type="primary"  @click="cancel" class="btnReturn" plain>返回</el-button>
        </el-form-item>
    </div>
    <el-tabs type="card" style="margin-top: 15px;">
    </el-tabs>
  </el-form>

</template>
<script>
import { SaveDataSet, QueryDataSetFiles, UpdateDataSet, DetailDataSet } from '@/api/platform.dataset'
import { ListFiles } from '@/api/platform.upload'
import env from '@/api/env'
import Bus from '@/components/vue-simple-uploader/js/bus'
import File from './File'
import util from '@/libs/util'
export default {
  name: 'DataSetForm',
  components: {
    File
  },
  props: {
    data: {
      type: Object,
      default: function () {
        return {}
      }
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
      item: {
        visibility: 'private'
      },
      pathLevels: [],
      cuserUser: [],
      uuid: null,
      userList: [],
      buttonVisible: true,
      platformDicts: [],
      labelDicts: [],
      fileList: [],
      file_total: 0,
      uploadShow: true,
      pageSize: 100,
      pageNum: 1,
      loading: false,
      options: {
        // 可通过 https://github.com/simple-uploader/Uploader/tree/develop/samples/Node.js 示例启动服务
        target: null,
        testChunks: false
      },
      rules: {
        name: [
          { required: true, message: '请输入数据集名称', trigger: 'blur' }
        ]
      }
    }
  },
  created: async function () {
    // this.platformDicts = await this.$dataDict(['platform'])
    if (this.data.id != null) {
      this.detail(this.data.id)
      this.uploadShow = false
    } else {
      this.initUuid(util.uuid())
      this.uploadShow = true
    }
  },
  mounted () {
    this.platformDicts = [{ label: 'OneFlow', value: 'OneFlow' }, { label: 'tensorflow', value: 'TensorFlow' }, { label: 'PyTorch', value: 'PyTorch' }, { label: 'MxNet', value: 'MxNet' }]
    this.labelDicts = [{ label: '图像识别', value: '图像识别' }, { label: '人脸识别', value: '人脸识别' },
      { label: '语音识别', value: '语音识别' }, { label: '文本识别', value: '文本识别' },
      { label: '机器翻译', value: '机器翻译' }, { label: '垃圾邮件', value: '垃圾邮件' } ]
  },
  methods: {
    initUuid (uuid) {
      this.uuid = uuid
      this.options.target = env.url() + '/upload/upload?path=' + uuid
    },
    fileUploadStart () {
      this.buttonVisible = false
    },
    fileSuccess () {
      this.buttonVisible = true
    },
    onFileError () {
      this.buttonVisible = true
    },
    goPath: function (index) {
      this.loading = true
      let path = ''
      for (var i = 0; i <= index; i++) {
        path += '/' + this.pathLevels[i]
      }
      let queryparams = {
        params: {
          url: this.item.url + path
        },
        pageNum: this.pageNum,
        pageSize: this.pageSize
      }
      ListFiles(queryparams).then(res => {
        this.item.Files = res.Files
        this.file_total = res.FileTotal
        this.loading = false
        let len = this.pathLevels.length
        this.pathLevels.splice(index+1, (len - index) - 1)
      }).catch(e => {
        this.loading = false
      })
    },
    queryFileList (dirName) {
      this.loading = true
      let path = ''
      this.pathLevels.forEach(v => {
        path += '/' + v
      })
      if (dirName.length > 0) {
        path += '/' + dirName
      }
      let queryparams = {
        params: {
          url: this.item.url + path
        },
        pageNum: this.pageNum,
        pageSize: this.pageSize
      }
      ListFiles(queryparams).then(res => {
        this.item.Files = res.Files
        this.file_total = res.FileTotal
        if (dirName.length > 0) {
          this.pathLevels.push(dirName)
        }
        this.loading = false
      }).catch(e => {
        this.loading = false
      })
    },
    handleCurrentChange (val) {
      this.pageNum = val
      this.queryList()
    },
    handleSizeChange (val) {
      this.pageSize = val
      this.queryList()
    },
    detail (id) {
      let that = this
      DetailDataSet({ id: id }).then(res => {
        that.item = res
        this.initUuid(that.item.url)
        this.item.Files = res.Files
        this.file_total = res.FileTotal
        this.queryFileList('')
      })
    },
    handleUploadSuccess (res, file) {
      console.log(res)
    },
    nextLevel: function (row, event) {
      if (row.IsDir === true) {
        this.queryFileList(row.name)
      }
    },
    save () {
      this.$refs['dataSetform'].validate((valid) => {
        if (valid) {
          let that = this
          this.item.url = this.uuid
          if (this.child === false) {
            if (this.item.id == null) {
              SaveDataSet(this.item).then(res => {
                this.$message({ message: '保存成功', type: 'success' })
                that.item = res
                this.$emit('onsuccess', this.item)
              })
            } else {
              UpdateDataSet(this.item).then(res => {
                this.$message({ message: '修改成功', type: 'success' })
                that.item = res
                this.$emit('onsuccess', this.item)
              })
            }
          } else {
            that.$emit('onsuccess', that.item)
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    cancel () {
      this.$emit('oncancel', this.item)
    }
  },
  destroyed () {
    Bus.$off('fileAdded')
    Bus.$off('fileSuccess')
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
