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
    <DataSetForm></DataSetForm>
  </d2-container>
</template>
<script>
import DataSetForm from '../components/DataSetForm'
export default {
  name: 'dataset-upload',
  components: {
    DataSetForm
  },
  data () {
    return {
      pageHeader: {
        path: ['数据管理', '数据上传']
      },
      item: {},
      readOnly: false,
      settings: {
        // axios headers
        headers: {
          'X-Requested-With': 'XMLHttpRequest',
          Authorization: `Bearer ${window.localStorage.getItem('user-token')}`
        },
        baseUrl: 'http://test.loc/file-manager/', // overwrite base url Axios
        windowsConfig: 2, // overwrite config
        lang: 'de', // set language
        translation: { // add new translation
          name: 'de',
          content: {
            about: 'Über',
            back: 'Zurück'
          }
        }
      }
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
