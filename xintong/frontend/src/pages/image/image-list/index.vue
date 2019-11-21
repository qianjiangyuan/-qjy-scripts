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
    <Image-List></Image-List>
  </d2-container>
</template>

<script>
import ImageList from '../components/ImageList'
export default {
  name: 'image-add',
  components: {
    ImageList
  },
  data () {
    return {
      pageHeader: {
        path: ['镜像管理', '镜像列表']
      }
    }
  },
  methods: {}
}
</script>

<style lang="scss">
</style>
