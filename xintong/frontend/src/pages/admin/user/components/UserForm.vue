<template>
  <el-form ref="userform" :inline="true"  :model="item"   :rules="rules" class="demo-form-inline" label-width="250px">
    <div>

      <el-form-item prop="name" label="名称">
        <el-input  v-model="item.name" :disabled="readOnly" placeholder="请输入名称"></el-input>
      </el-form-item>
      <div></div>
      <div v-if="!updateFlag">
      <el-form-item prop="pass" label="密码">
        <el-input  v-model="item.passwd" type="password"  :disabled="readOnly" placeholder="请输入密码"></el-input>
      </el-form-item>
      <div></div>
      <el-form-item prop="checkPass" label="再输一次密码">
        <el-input  v-model="password" type="password" :disabled="readOnly" placeholder="请输入密码"></el-input>
      </el-form-item>
      <div></div>
      </div>
      <el-form-item prop="administrator" label="类型">
        <el-radio v-model="item.administrator" label="admin">管理员</el-radio>
        <el-radio v-model="item.administrator" label="normal">普通用户</el-radio>
      </el-form-item>
      <br/>
      <el-form-item label="工作空间" prop="workspace" >
      <el-select  v-model="item.workspace" :disabled="readOnly" filterable value-key="id"
                  placeholder="选择工作空间">
        <el-option
                v-for="item2 in workspaceList"
                :key="item2.id"
                :label="item2.name"
                :value="item2">
        </el-option>
      </el-select>
    </el-form-item>
      <div></div>
      <el-form-item label="所属组" prop="group" >
        <el-select  v-model="item.group" :disabled="readOnly" filterable value-key="id"
                    placeholder="所属组">
          <el-option
                  v-for="item2 in groupList"
                  :key="item2.id"
                  :label="item2.name"
                  :value="item2">
          </el-option>
        </el-select>
      </el-form-item>
      <div></div>
      <el-form-item prop="nickname" label="姓名">
        <el-input  v-model="item.nickname"   :disabled="readOnly" placeholder="请输入姓名"></el-input>
      </el-form-item>

      <el-form-item prop="phone" label="手机">
        <el-input  v-model="item.phone" :disabled="readOnly" placeholder="请输入手机"></el-input>
      </el-form-item>
      <el-form-item prop="email" label="邮箱">
        <el-input  v-model="item.email" :disabled="readOnly" placeholder="请输入邮箱"></el-input>
      </el-form-item>
    </div>
    <div>
      <div class="fr" style="">
        <el-form-item>
          <template v-if="!readOnly">
            <el-button type="primary" size="small" @click="save">保存</el-button>
          </template>
          <el-button type="primary" size="small" @click="cancel" class="btnReturn" plain>返回</el-button>
            <el-button  v-if="updateFlag"  ype="primary"  size="small" @click="updatePasswd"  class="btnReturn" >修改密码</el-button>
        </el-form-item>
      </div>
    </div>
    <el-tabs type="card" style="margin-top: 15px;">
    </el-tabs>

    <el-dialog :visible="updatePasswdVisible" :show-close=false >
        <el-form-item prop="pass" label="密码">
            <el-input  v-model="item.passwd" type="password"  :disabled="readOnly" placeholder="请输入密码"></el-input>
        </el-form-item>
        <div></div>
        <el-form-item prop="checkPass" label="再输一次密码">
            <el-input  v-model="password" type="password" :disabled="readOnly" placeholder="请输入密码"></el-input>
        </el-form-item>
        <div></div>
        <div class="fr" style="">
            <el-form-item>
                <template v-if="!readOnly">
                    <el-button type="primary" size="small" @click="commitPasswd">保存</el-button>
                </template>
                <el-button type="primary" size="small" @click="cancelCommitPasswd" class="btnReturn" plain>返回</el-button>
            </el-form-item>
        </div>
    </el-dialog>
  </el-form>
</template>
<script>
import { SaveUser, UpdateUser, CommitUserPassWd, DetailUser } from '@/api/platform.user'
import { QueryWorkspaceList } from '@/api/platform.workspace'
import { QueryGroupList } from '@/api/platform.group'
export default {
  name: 'UserForm',
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
  data: function () {
    return {
      item: {},
      password: null,
      workspaceList: [],
      updateFlag: false,
      updatePasswdVisible: false,
      groupList: [],
      rules: {
        name: [{ required: true, message: '请输入登陆用户名', trigger: 'blur' }],
        nickname: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
        pass: [{ required: true, validator: this.validatePass }],
        checkPass: [{ required: true, validator: this.validatePass2 }],
        group: [{ required: true, message: '请输选择组' }]
      }
    }
  },
  methods: {
    detail (id) {
      let that = this
      DetailUser({ id: id }).then(res => {
        that.item = res
        that.password = that.item.passwd
      })
    },
    cancel () {
      this.$emit('oncancel', this.item)
    },
    validatePass (rule, value, callback) {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else {
        if (this.password !== '') {
          this.$refs.userform.validateField('checkPass')
        }
        callback()
      }
    },
    validatePass2  (rule, value, callback) {
      if (this.password === '') {
        callback(new Error('请再次输入密码'))
      } else if (this.password !== this.item.passwd) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    },

    save () {
      console.log('aaaaaa')
      this.$refs['userform'].validate((valid) => {
        if (valid) {
          let that = this
          if (this.child === false) {
            this.item.workspace_id = this.item.workspace.id
            this.item.group_id = this.item.group.id
            if (this.item.id == null) {
              SaveUser(this.item).then(res => {
                this.$message({ message: '保存成功', type: 'success' })
                that.item = res
                this.$emit('onsuccess', this.item)
              })
            } else {
              UpdateUser(this.item).then(res => {
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
    updatePasswd () {
      this.updatePasswdVisible = true
    },
    queryWorkspace () {
      let that = this
      QueryWorkspaceList().then(res => {
        that.workspaceList = res
      })
    },
    commitPasswd () {
      let that = this
      CommitUserPassWd(this.item).then(res => {
        this.updatePasswdVisible = false
        that.item = res
      })
    },
    cancelCommitPasswd () {
      this.updatePasswdVisible = false
    },
    queryGroup () {
      let that = this
      QueryGroupList().then(res => {
        that.groupList = res
      })
    }
  },
  mounted: function () {
    if (this.data.id != null) {
      this.detail(this.data.id)
      this.updateFlag = true
    } else {
      this.item.administrator = 'normal'
      this.updateFlag = false
    }
  },
  created: async function () {
    this.queryWorkspace()
    this.queryGroup()
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
