<template>
    <!--main container-->
    <d2-container>
        <!--page header-->
        <template slot="header">
            <el-breadcrumb separator-class="el-icon-arrow-right">
                <el-breadcrumb-item :key="p" v-for="p in pageHeader.path">
                    {{p}}
                </el-breadcrumb-item>
            </el-breadcrumb>
        </template>
        <User-List></User-List>
    </d2-container>
</template>
<script>
import UserList from './components/UserList'
export default {
  name: 'user-list',
  components: {
    UserList
  },
  data () {
    return {
      pageHeader: {
        path: ['用户管理', '用户列表']
      }
    }
  },
  methods: {
  }
}
</script>

<style lang="scss">
</style>
