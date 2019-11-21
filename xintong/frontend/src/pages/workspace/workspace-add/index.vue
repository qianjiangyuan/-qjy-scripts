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
        <Workspace-Form :readOnly="false"></Workspace-Form>
    </d2-container>
</template>
<script>
import WorkspaceForm from '../components/WorkspaceForm'
export default {
  name: 'workspace-add',
  components: {
    WorkspaceForm
  },
  data () {
    return {
      pageHeader: {
        path: ['工作空间管理', '工作空间创建']
      }
    }
  },
  methods: {
  }
}
</script>

<style lang="scss">
</style>
