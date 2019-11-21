<template>
    <div
            class="d2-layout-header-aside-group"
            :style="styleLayoutMainGroup"
            :class="{grayMode: grayActive}">
        <!-- 半透明遮罩 -->
        <div class="d2-layout-header-aside-mask"></div>
        <!-- 主体内容 -->
        <div class="d2-layout-header-aside-content" flex="dir:top">
            <!-- 顶栏 -->
            <div
                    class="d2-theme-header"
                    :style="{
          opacity: this.searchActive ? 0.5 : 1
        }"
                    flex-box="0"
                    flex>
                <d2-menu-header flex-box="1"/>
                <!-- 顶栏右侧 -->
                <div class="d2-header-right" flex-box="0">
                    <!-- 如果你只想在开发环境显示这个按钮请添加 v-if="$env === 'development'" -->
                    <d2-header-search @click="handleSearchClick"/>
                    <d2-header-error-log/>
                    <d2-header-fullscreen/>
                    <d2-header-theme/>
                    <d2-header-size/>
                    <d2-header-user/>
                </div>
            </div>
            <div style="background-color: white;height: 100%;padding-top: 100px">
                <el-row  :gutter="24">
                    <el-col :span="5" :offset="5">
                        <el-card :body-style="{ padding: '0px' }" >
                            <el-button @click="handleTrainSystem">
                                <img  class="image" src="./image/time3.png">
                                  <div style="padding: 14px;">
                                      <span>监测子系统</span>
                                  </div>
                            </el-button>
                        </el-card>
                    </el-col>
                    <el-col :span="5" >
                        <el-card :body-style="{ padding: '0px' }" >
                           <el-button @click="handleValidateSystem">
                            <img  class="image" src="./image/timg.jpg" >
                               <div style="padding: 14px;">
                                   <span>推断测试子系统</span>
                               </div>
                           </el-button>
                        </el-card>
                    </el-col>
                    <el-col :span="5" >
                        <el-card :body-style="{ padding: '0px' }">
                            <el-button @click="handleManagerSystem">
                                <img  class="image" src="./image/time4.png">
                                <div style="padding: 14px;">
                                    <span>系统运维和管理子系统</span>
                                </div>
                            </el-button>
                        </el-card>
                    </el-col>
                </el-row>
            </div>
        </div>
    </div>
</template>

<script>
import d2MenuSide from '@/layout/header-aside/components/menu-side'
import d2MenuHeader from '@/layout/header-aside/components/menu-header'
import d2Tabs from '@/layout/header-aside/components/tabs'
import d2HeaderFullscreen from '@/layout/header-aside/components/header-fullscreen'
import d2HeaderSearch from '@/layout/header-aside/components/header-search'
import d2HeaderSize from '@/layout/header-aside/components/header-size'
import d2HeaderTheme from '@/layout/header-aside/components/header-theme'
import d2HeaderUser from '@/layout/header-aside/components/header-user'
import d2HeaderErrorLog from '@/layout/header-aside/components/header-error-log'
import { mapState, mapGetters, mapActions } from 'vuex'
import mixinSearch from '@/layout/header-aside/mixins/search'
export default {
  name: 'd2-layout-header-aside',
  mixins: [
    mixinSearch
  ],
  components: {
    d2MenuSide,
    d2MenuHeader,
    d2Tabs,
    d2HeaderFullscreen,
    d2HeaderSearch,
    d2HeaderSize,
    d2HeaderTheme,
    d2HeaderUser,
    d2HeaderErrorLog
  },
  data () {
    return {
      // [侧边栏宽度] 正常状态
      asideWidth: '200px',
      // [侧边栏宽度] 折叠状态
      asideWidthCollapse: '65px'
    }
  },
  computed: {
    ...mapState('d2admin', {
      keepAlive: state => state.page.keepAlive,
      grayActive: state => state.gray.active,
      transitionActive: state => state.transition.active,
      asideCollapse: state => state.menu.asideCollapse
    }),
    ...mapGetters('d2admin', {
      themeActiveSetting: 'theme/activeSetting'
    }),
    /**
       * @description 最外层容器的背景图片样式
       */
    styleLayoutMainGroup () {
      return {
        ...this.themeActiveSetting.backgroundImage ? {
          backgroundImage: `url('${this.$baseUrl}${this.themeActiveSetting.backgroundImage}')`
        } : {}
      }
    }
  },
  methods: {
    ...mapActions('d2admin/menu', [
      'asideCollapseToggle'
    ]),
    /**
       * 接收点击切换侧边栏的按钮
       */
    handleToggleAside () {
      this.asideCollapseToggle()
    },
    handleTrainSystem () {
      this.$router.push('/index?systemType=train')
    },
    handleManagerSystem () {
      this.$router.push('/index?systemType=manager')
    },
    handleValidateSystem () {
      this.$router.push('/index?systemType=validate')
    }
  }
}
</script>

<style lang="scss">
    // 注册主题
    @import '~@/assets/style/theme/register.scss';
    .time {
        font-size: 13px;
        color: #999;
    }

    .bottom {
        margin-top: 13px;
        line-height: 12px;
    }

    .button {
        padding: 0;
        float: right;
    }

    .image {
        width: 100%;
        display: block;
    }

    .clearfix:before,
    .clearfix:after {
        display: table;
        content: "";
    }

    .clearfix:after {
        clear: both
    }
</style>
