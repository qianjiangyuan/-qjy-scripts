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
        <Group-Form></Group-Form>
    </d2-container>
</template>

<script>
import GroupForm from '../components/GroupForm'
export default {
  name: 'group-form',
  components: {
    GroupForm
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
