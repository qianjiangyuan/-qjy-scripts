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
    首页
  </d2-container>
</template>

<script>
var echarts = require('echarts')
export default {
  name: 'admin',
  data () {
    return {
      pageHeader: {
        path: ['首页']
      }
    }
  },
  methods: {},
  mounted: function () {
    // 基于准备好的dom，初始化echarts实例
    var myChart = echarts.init(document.getElementById('echartContainer'))
    // 绘制图表
    myChart.setOption({
      title: { text: 'ECharts 入门示例' },
      tooltip: {},
      xAxis: {
        data: ['衬衫', '羊毛衫', '雪纺衫', '裤子', '高跟鞋', '袜子']
      },
      yAxis: {},
      series: [{
        name: '销量',
        type: 'bar',
        data: [5, 20, 36, 10, 10, 20]
      }]
    })
  }
}
</script>

<style lang="scss">
</style>
