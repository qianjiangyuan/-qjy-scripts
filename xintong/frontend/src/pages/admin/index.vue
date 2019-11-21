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
        管理员中心
    </d2-container>
</template>

<script>
export default {
  name: 'admin',
  data () {
    return {
      pageHeader: {
        path: ['管理员中心']
      }
    }
  },
  methods: {}
}
</script>

<style lang="scss">
</style>
