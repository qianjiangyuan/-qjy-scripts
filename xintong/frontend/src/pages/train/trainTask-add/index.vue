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
        <Train-Task-Form></Train-Task-Form>
    </d2-container>
</template>

<script>
import TrainTaskForm from '../components/TrainTaskForm'
export default {
  name: 'trainTaskForm',
  components: {
    TrainTaskForm
  },
  data () {
    return {
      pageHeader: {
        path: ['训练任务', '训练任务提交']
      }
    }
  },
  methods: {}
}
</script>

<style lang="scss">
</style>
