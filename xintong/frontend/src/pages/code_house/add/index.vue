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
        <Code-House-Form></Code-House-Form>
    </d2-container>
</template>

<script>
import CodeHouseForm from '../components/CodeHouseForm'
export default {
  name: 'code-house-form',
  components: {
    CodeHouseForm
  },
  data () {
    return {
      pageHeader: {
        path: ['代码库管理', '代码库列表']
      }
    }
  },
  methods: {}
}
</script>

<style lang="scss">
</style>
