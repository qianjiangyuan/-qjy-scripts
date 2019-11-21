<template>
  <el-form ref="groupform" :inline="true"  :model="item"   :rules="rules" class="demo-form-inline" label-width="120px">
    <div>
      <el-form-item prop="name" label="名称">
        <el-input  v-model="item.name" :disabled="readOnly" placeholder="请输入名称"></el-input>
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
import { SaveGroup, UpdateGroup, DetailGroup } from '@/api/platform.group'
export default {
  name: 'GroupForm',
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
      rules: {
      }
    }
  },
  created: async function () {

  },
  mounted () {
    console.log('data=' + this.data)
    if (this.data.id != null) {
      this.detail(this.data.id)
    }
  },
  methods: {
    detail (id) {
      let that = this
      DetailGroup({ id: id }).then(res => {
        that.item = res
      })
    },
    save () {
      this.$refs['groupform'].validate((valid) => {
        if (valid) {
          let that = this
          if (this.child === false) {
            if (this.item.id == null) {
              SaveGroup(this.item).then(res => {
                this.$message({ message: '保存成功', type: 'success' })
                that.item = res
                this.$emit('onsuccess', this.item)
              })
            } else {
              UpdateGroup(this.item).then(res => {
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
    }
  }
}
</script>
<style scoped>
</style>
