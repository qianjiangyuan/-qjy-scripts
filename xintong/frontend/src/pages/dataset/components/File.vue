<template>
    <div>
    <img  v-if="isDir" src="./image/dir.png">
        <img  v-if="!isDir" src="./image/file.png">
    <div class="name">
        <div>{{ data.name }}</div>
    </div>
    <!-- 删除icon -->
    <div class="del">
        <i @click="handleFileRemove(data.id)" class="el-icon-delete2"></i>
    </div>
    </div>
</template>

<script>
export default {
  name: 'File',
  props: {
    data: {
      type: Object,
      default: function () {
        return {}
      }
    }
  },
  data: function () {
    return {
      isDir: false
    }
  },
  methods: {
    handleFileRemove: function (data) {

    }
  }
}
</script>

<style scoped>

</style>
