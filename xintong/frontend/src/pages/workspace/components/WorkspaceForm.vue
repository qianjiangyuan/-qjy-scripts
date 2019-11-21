<template>
  <el-form ref="workspaceform" :inline="true"  :model="item"   :rules="rules" class="demo-form-inline" label-width="120px">
    <div>
      <el-form-item prop="name" label="名称">
        <el-input  v-model="item.name" :disabled="readOnly" placeholder="请输入名称"></el-input>
      </el-form-item>
      <div></div>
      <el-form-item prop="description" label="描述">
        <el-input  v-model="item.description"  type="textarea" :rows="2" placeholder="请输入描述"></el-input>
      </el-form-item>
      <div></div>
      <!--<el-form-item label="GPU类型" prop="gpuType">-->
        <!--<el-select  v-model="item.gpuType" :disabled="readOnly"-->
          <!--placeholder="选择GPU类型">-->
          <!--<el-option-->
            <!--v-for="item2 in gpuTypeDicts['GPU_TYPE']"-->
            <!--:key="item2.value"-->
            <!--:label="item2.label"-->
            <!--:value="item2.value">-->
          <!--</el-option>-->
        <!--</el-select>-->
      <!--</el-form-item>-->
      <!--<div></div>-->
      <el-form-item prop="gpuNum" label="GPU数量">
        <el-input-number  v-model="item.gpu_num"  placeholder="请输入GPU数量"></el-input-number>
      </el-form-item>
      <div></div>
      <el-form-item prop="cpuNum" label=" CPU数量">
        <el-input-number   v-model="item.cpu_num" placeholder="请输入CPU数量"></el-input-number>
      </el-form-item>
      <div></div>
      <el-form-item prop="memoryNum" label="内存数量">
        <el-input-number  v-model="item.memory_num"  placeholder="请输入内存数量"></el-input-number>Gi
      </el-form-item>
      <div></div>
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
import { SaveWorkspace, UpdateWorkspace, DetailWorkspace } from '@/api/platform.workspace'
export default {
  name: 'WorkspaceForm',
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
      gpuTypeDicts: [],
      rules: {
        name: [{ required: true, message: '请输入登陆用户名', trigger: 'blur' }]
      }
    }
  },
  created: async function () {
    // this.gpuTypeDicts = await this.$dataDict(['GPU_TYPE'])
  },
  mounted () {
    if (this.data.id != null) {
      this.detail(this.data.id)
    }
    this.initvalue()
  },
  methods: {
    detail (id) {
      let that = this
      DetailWorkspace({ id: id }).then(res => {
        that.item = res
        that.initvalue()
      })
    },
    initvalue () {
      if (this.item.gpu_num == null) {
        this.item.gpu_num = 1
      }
      if (this.item.cpu_num == null) {
        this.item.cpu_num = 1
      }
      if (this.item.rdma == null) {
        this.item.rdma = 1
      }
      if (this.item.memory_num == null) {
        this.item.memory_num = 1
      }
    },
    save () {
      this.$refs['workspaceform'].validate((valid) => {
        if (valid) {
          let that = this
          if (this.child === false) {
            if (this.item.id == null) {
              SaveWorkspace(this.item).then(res => {
                this.$message({ message: '保存成功', type: 'success' })
                that.item = res.data
                this.$emit('onsuccess', this.item)
              })
            } else {
              UpdateWorkspace(this.item).then(res => {
                this.$message({ message: '保存成功', type: 'success' })
                that.item = res.data
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
