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
        <Code-House-List></Code-House-List>
    </d2-container>
</template>

<script>
import CodeHouseList from '../components/CodeHouseList'
export default {
  name: 'codehouse-list',
  components: {
    CodeHouseList
  },
  data () {
    return {
      pageHeader: {
        path: ['组管理', '组列表']
      }
    }
  },
  methods: {}
}
</script>

<style lang="scss">
</style>
