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
        <Dlmodel-List></Dlmodel-List>
    </d2-container>
</template>

<script>
import DlmodelList from '../components/DlmodelList'
export default {
  name: 'dlmodel-form',
  components: {
    DlmodelList
  },
  data () {
    return {
      pageHeader: {
        path: ['模型管理', '模型列表']
      }
    }
  },
  methods: {}
}
</script>

<style lang="scss">
</style>
