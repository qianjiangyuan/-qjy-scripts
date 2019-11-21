<template>
  <el-form ref="imageform" :inline="true"  :model="item"   :rules="rules" class="demo-form-inline" label-width="250px">
    <div>
      <el-form-item prop="name" label="名称">
        <el-input  v-model="item.name" :disabled="readOnly" placeholder="请输入名称"></el-input>
      </el-form-item>
      <el-form-item prop="desc" label="版本">
        <el-input  v-model="item.version"  :disabled="readOnly" placeholder="请输入版本"></el-input>
      </el-form-item>
      <el-form-item prop="desc" label="描述">
        <el-input  v-model="item.desc" type="textarea"  :disabled="readOnly" placeholder="请输入描述"></el-input>
      </el-form-item>

      <el-form-item prop="comman" label="命令行">
        <el-input  v-model="item.command" type="textarea"  :disabled="readOnly" placeholder="请输入命令行"></el-input>
      </el-form-item>
      <el-form-item prop="visibility" label="类型">
        <el-radio v-model="item.visibility" label="public">公开</el-radio>
        <el-radio v-model="item.visibility" label="private">私有</el-radio>
      </el-form-item>
      <br/>
      <el-form-item prop="platform" label="框架">
        <el-select  v-model="item.platform" :disabled="readOnly"  allow-create  filterable
                    placeholder="选择标签">
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

      <div></div>
      <el-form-item label="文件上传" prop="platform">
        <el-upload
                class="upload-demo"
                :action=uploadUrl
                multiple
                :limit="3"
                :on-success="handleUploadSuccess"
                :before-upload="handelBeforUpload"
                :file-list="fileList">
          <el-button size="small" type="primary">点击上传</el-button>
        </el-upload>
      </el-form-item>
      <div></div>
      <el-form-item >
        <template v-if="!readOnly">
          <el-button v-if="comitButtonShow" type="primary" size="small" @click="save">保存</el-button>
        </template>
        <el-button type="primary" size="small" @click="cancel" class="btnReturn" plain>返回</el-button>
      </el-form-item>

    </div>
    <el-tabs type="card" style="margin-top: 15px;">
    </el-tabs>
  </el-form>
</template>
<script>
import { SaveImage, UpdateImage, DetailImage } from '@/api/platform.image'
import env from '@/api/env'
import util from '@/libs/util'
export default {
  name: 'ImageForm',
  components: {
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
  data () {
    return {
      item: JSON.parse(JSON.stringify(this.data)),
      labelDicts: [],
      platformDicts: [],
      uuid: null,
      uploadUrl: env.url() + '/upload/upload',
      cuserUser: [],
      userList: [],
      fileList: [],
      comitButtonShow: true,
      rules: {
        name: [
          { required: true, message: '请输入镜像名称', trigger: 'blur' }
        ]
      }
    }
  },
  created: async function () {
  },
  mounted () {
    if (this.data.id != null) {
      this.detail(this.data.id)
    }
    this.platformDicts = [{ label: 'OneFlow', value: 'OneFlow' }, { label: 'tensorflow', value: 'TensorFlow' }, { label: 'PyTorch', value: 'PyTorch' }, { label: 'MxNet', value: 'MxNet' }]
    this.labelDicts = [{ label: '图像识别', value: '图像识别' }, { label: '人脸识别', value: '人脸识别' },
      { label: '语音识别', value: '语音识别' }, { label: '文本识别', value: '文本识别' },
      { label: '机器翻译', value: '机器翻译' }, { label: '垃圾邮件', value: '垃圾邮件' } ]
    this.initUuid(util.uuid())
  },
  methods: {
    initUuid (uuid) {
      this.uuid = uuid
      this.uploadUrl = env.url() + '/upload/upload?path=' + uuid
    },
    detail (id) {
      let that = this
      DetailImage({ id: id }).then(res => {
        that.item = res
        this.initUuid(that.item.url)
        this.fileList = that.item.Files
        this.item.filename = this.item.Files[0].name
      })
    },
    handleUploadSuccess: function (res, file) {
      this.item.filename = file.name
      console.log(res)
      this.comitButtonShow = true
    },
    save () {
      this.$refs['imageform'].validate((valid) => {
        if (valid) {
          let that = this
          if (this.child === false) {
            if (this.item.id == null) {
              this.item.url = this.uuid
              SaveImage(this.item).then(res => {
                this.$message({ message: '保存成功', type: 'success' })
                that.item = res
                this.$emit('onsuccess', this.item)
              })
            } else {
              UpdateImage(this.item).then(res => {
                this.$message({ message: '保存成功', type: 'success' })
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
    handelBeforUpload (res, file) {
      this.comitButtonShow = false
    },
    cancel () {
      this.$emit('oncancel', this.item)
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
</style>
