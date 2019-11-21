<template>
  <el-form ref="dlserviceform" :inline="true"  :model="item"   :rules="rules" class="demo-form-inline" label-width="120px">
    <div>
      <el-form-item prop="name" label="名称">
        <el-input  v-model="item.name" :disabled="readOnly" placeholder="请输入名称"></el-input>
      </el-form-item>
      <el-form-item prop="version" label="版本">
        <el-input  v-model="item.version" :disabled="readOnly" placeholder="请输入版本"></el-input>
      </el-form-item>
      <el-form-item prop="description" label="描述">
        <el-input  v-model="item.description" :disabled="readOnly" type="textarea" :rows="2" placeholder="请输入描述"></el-input>
      </el-form-item>
      <div></div>
      <el-form-item label="标签" prop="label">
        <el-select  class="short" v-model="item.label" :disabled="readOnly"
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
      <el-form-item label="框架" prop="platform">
        <el-select    class="short" v-model="item.platform" :disabled="readOnly"
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
      <el-form-item prop="path" label="路径">
        <el-input  v-model="item.path" :disabled="readOnly" placeholder="请输入路径"></el-input>
      </el-form-item>
      <div></div>
     <el-form-item label="镜像" prop="imageID" >
       <el-select  v-model="item.image" :disabled="readOnly" filterable value-key="id"
         placeholder="选择镜像">
         <el-option
           v-for="item2 in imageList"
             :key="item2.id"
             :label="item2.name"
             :value="item2">
          </el-option>
        </el-select>
     </el-form-item>

      <el-form-item label="模型" prop="dlmodelID" >
        <el-select  v-model="item.dlmodel" :disabled="readOnly" filterable value-key="id"
                    placeholder="选择镜像">
          <el-option
                  v-for="item2 in dlmodelList"
                  :key="item2.id"
                  :label="item2.name"
                  :value="item2">
          </el-option>
        </el-select>
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
import { SaveDlservice, UpdateDlservice, DetailDlservice } from '@/api/platform.dlservice'
import { QueryImageList } from '@/api/platform.image'
import { QueryDlmodelList } from '@/api/platform.dlmodel'
import util from '@/libs/util'
export default {
  name: 'DlserviceForm',
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
      item: {},
      labelDicts: [],
      platformDicts: [],
      imageIdImage: [],
      imageList: [],
      dlmodelList: [],
      rules: {
      }
    }
  },
  created: async function () {
    this.labelDicts = util.labelDicts
    this.platformDicts = util.platformDicts
    this.selectImage()
    this.selectDlmodel()
  },
  mounted () {
    if (this.data.id != null) {
      this.detail(this.data.id)
    }
    // this.selectImage();
  },
  methods: {
    detail (id) {
      let that = this
      DetailDlservice({ id: id }).then(res => {
        that.item = res
      })
    },
    save () {
      this.$refs['dlserviceform'].validate((valid) => {
        if (valid) {
          let that = this
          if (this.child === false) {
            this.item.image_id = this.item.image.id
            this.item.dlmodel_id = this.item.dlmodel.id
            if (this.item.id == null) {
              SaveDlservice(this.item).then(res => {
                this.$message({ message: '保存成功', type: 'success' })
                that.item = res
                this.$emit('onsuccess', this.item)
              })
            } else {
              UpdateDlservice(this.item).then(res => {
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
    cancel () {
      this.$emit('oncancel', this.item)
    },
    selectImage () {
      let that = this
      QueryImageList().then(res => {
        that.imageList = res
      })
    },
    selectDlmodel () {
      let that = this
      QueryDlmodelList().then(res => {
        that.dlmodelList = res
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
  .short {
    width: 200px;
  }
  .upload-demo {
    width: 400px;
  }
  .el-select  {
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
