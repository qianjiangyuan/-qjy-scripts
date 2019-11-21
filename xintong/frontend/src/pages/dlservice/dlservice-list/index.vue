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
        <Dlservice-Form></Dlservice-Form>
    </d2-container>
</template>

<script>
import DlserviceForm from '../components/DlserviceList'
export default {
  name: 'dlservice-form',
  components: {
    DlserviceForm
  },
  data () {
    return {
      pageHeader: {
        path: ['服务管理', '服务列表']
      }
    }
  },
  methods: {}
}
</script>

<style lang="scss">
</style>
