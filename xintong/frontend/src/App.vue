<template>
  <div id="app">
    <router-view/>
    <global-uploader></global-uploader>
  </div>
</template>

<script>
import globalUploader from '@/components/vue-simple-uploader/globalUploader'
export default {
  name: 'app',
  components: {
    globalUploader
  }
}
</script>

<style lang="scss">
@import '~@/assets/style/public-class.scss';
</style>
