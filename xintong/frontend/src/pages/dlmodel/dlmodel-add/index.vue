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
        <Dlmodel-Form></Dlmodel-Form>
    </d2-container>
</template>

<script>
import DlmodelForm from '../components/DlmodelForm'
export default {
  name: 'dlmodel-form',
  components: {
    DlmodelForm
  },
  data () {
    return {
      pageHeader: {
        path: ['模型管理', '模型管理']
      }
    }
  },
  methods: {}
}
</script>

<style lang="scss">
</style>
