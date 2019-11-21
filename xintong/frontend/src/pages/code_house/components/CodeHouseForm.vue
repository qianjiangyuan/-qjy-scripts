<!--suppress ALL -->
<template>
  <el-form ref="codeHouseform" :inline="true"  :model="item"   :rules="rules" class="demo-form-inline" label-width="120px">
    <div>
      <el-form-item prop="name" label="名称">
        <el-input  v-model="item.name" :disabled="readOnly" placeholder="请输入name"></el-input>
      </el-form-item>
      <div></div>
      <el-form-item prop="git" label="git地址">
        <el-input  v-model="item.git" :disabled="readOnly" placeholder="请输入git地址"></el-input>
      </el-form-item>
      <div></div>
      <el-form-item prop="visibility"  label="类型">
        <el-radio v-model="item.visibility"  :disabled="readOnly"  label="public">公开</el-radio>
        <el-radio v-model="item.visibility"  :disabled="readOnly"  label="private">私有</el-radio>
      </el-form-item>
      <div></div>
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
      <div></div>

      <el-form-item prop="cmd" label="启动脚本">
        <el-input  v-model="item.cmd" :disabled="readOnly" type="textarea" :rows="4"  placeholder="请输入启动"></el-input>
      </el-form-item>
      <div></div>
        <el-form-item prop="describe" label="描述">
            <el-input  v-model="item.describe" :disabled="readOnly"   type="textarea" :rows="4" placeholder="请输入描述"></el-input>
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
import util from '@/libs/util'
import { SaveCodeHouse, UpdateCodeHouse,DetailCodeHouse } from '@/api/platform.codeHouse'
export default {
  name: 'CodeHouseForm',
  components: {
  },
  props: {
    data: {
      type: Object,
        labelDicts: [],
        platformDicts: [],

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
      platformDicts: [],
      labelDicts: [],
      rules: {
      }
    }
  },
  created: async function () {
    this.labelDicts = util.labelDicts
    this.platformDicts = util.platformDicts
  },
  mounted () {
    if (this.data.id != null) {
      this.detail(this.data.id)
    }
  },
  methods: {
    detail (id) {
      let that = this
      DetailCodeHouse({ id: id }).then(res => {
        that.item = res
      })
    },
    save () {
      this.$refs['codeHouseform'].validate((valid) => {
        if (valid) {
          let that = this
          if (this.child === false) {
              if (this.item.id == null) {
                  SaveCodeHouse(this.item).then(res => {
                      this.$message({message: '保存成功', type: 'success'})
                      that.item = res.data
                      this.$emit('onsuccess', this.item)
                  })
              }
              else{
                  UpdateCodeHouse(this.item).then(res => {
                      this.$message({message: '保存成功', type: 'success'})
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
</style>
