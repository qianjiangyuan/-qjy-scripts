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
        <Workspace-List></Workspace-List>
    </d2-container>
</template>
<script>
import WorkspaceList from '../components/WorkspaceList'
export default {
  name: 'workspace-add',
  components: {
    WorkspaceList
  },
  data () {
    return {
      pageHeader: {
        path: ['工作空间管理', '工作空间列表']
      }
    }
  },
  methods: {
  }
}
</script>

<style lang="scss">
</style>
