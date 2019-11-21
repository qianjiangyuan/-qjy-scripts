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
   <Image-Form></Image-Form>
  </d2-container>
</template>

<script>
import ImageForm from '../components/ImageForm'
export default {
  name: 'image-form',
  components: {
    ImageForm
  },
  data () {
    return {
      pageHeader: {
        path: ['镜像管理', '镜像上传']
      }
    }
  },
  methods: {}
}
</script>

<style lang="scss">
</style>
