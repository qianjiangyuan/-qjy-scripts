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
    <dataSet-list :readOnly="readOnly" @cancelForm="cancelAddForm" ></dataSet-list>
  </d2-container>
</template>
<script>
import DataSetList from '../components/DataSetList'
export default {
  name: 'dataset-list',
  components: {
    DataSetList
  },
  data () {
    return {
      pageHeader: {
        path: ['数据管理', '数据上传']
      },
      item: {},
      readOnly: false
    }
  },
  methods: {
    cancelAddForm () {
    }
  }
}
</script>

<style lang="scss">
</style>
