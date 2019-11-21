<template>
    <el-form ref="train_taskform" :inline="true"  :model="item"   :rules="rules" class="demo-form-inline" label-width="150px">
        <div  v-if="!imgetShow && !dataSetShow">
        <div>
            <el-form-item prop="name" label="名称">
                <el-input  v-model="item.name" :disabled="readOnly" placeholder="请输入名称"></el-input>
            </el-form-item>
            <div></div>
            <el-form-item prop="status" label="状态">
                {{item.status}}
            </el-form-item>
            <div></div>
            <el-form-item prop="desc" label="描述">
                <el-input  v-model="item.desc" :disabled="readOnly" type="textarea" :rows="2" placeholder="请输入描述"></el-input>
            </el-form-item>
            <div></div>
            <el-form-item prop="type" label="类型">
                <el-radio v-model="item.type" label="code">代码</el-radio>
                <el-radio v-model="item.type" label="image">镜像</el-radio>
            </el-form-item>
            <div></div>
            <el-form-item v-if="item.type =='code'" label="代码库" prop="codehouse" >
                <el-select  v-model="item.codehouse" :disabled="readOnly" filterable value-key="id"
                            placeholder="选择代码库">
                    <el-option
                            v-for="item2 in codeHouseList"
                            :key="item2.id"
                            :label="item2.name"
                            :value="item2">
                    </el-option>
                </el-select>
                <el-button type="primary" size="small" @click="addDataSet">添加</el-button>
            </el-form-item>

            <el-form-item v-else label="镜像集" prop="imageId" >
                <el-select  v-model="item.image_url" :disabled="readOnly" filterable value-key="id"
                            placeholder="选择镜像集">
                    <el-option
                            v-for="item2 in imageList"
                            :key="item2.image_url"
                            :label="item2.image_url"
                            :value="item2.image_url">
                    </el-option>
                </el-select>
                <el-button type="primary" size="small" @click="addImage">添加</el-button>
            </el-form-item>

            <el-form-item label="数据集" prop="datasetId" >
                <el-select  v-model="item.dataset" :disabled="readOnly" filterable value-key="id"
                            placeholder="选择数据集">
                    <el-option
                            v-for="item2 in dataSetList"
                            :key="item2.id"
                            :label="item2.name"
                            :value="item2">
                    </el-option>
                </el-select>
                <el-button type="primary" size="small" @click="addDataSet">添加</el-button>
            </el-form-item>

            <!--<el-form-item prop="podNum" label="pod数量">-->
                <!--<el-input-number v-model="item.pod_num" :disabled="readOnly" placeholder="请输入pod数量"></el-input-number>-->
            <!--</el-form-item>-->
            <el-form-item prop="cpuNum" label="CPU数量">
                <el-input-number v-model="item.cpu_num" :disabled="readOnly" placeholder="请输入cpu数量"></el-input-number>
            </el-form-item>
            <el-form-item prop="gpuNum" label="GPU数量">
                <el-input-number  v-model="item.gpu_num" :disabled="readOnly" placeholder="请输入GPU数量"></el-input-number>
            </el-form-item>
            <el-form-item prop="memoryNum" label="内存数量">
                <el-input-number  v-model="item.memory_num" :disabled="readOnly" placeholder="请输入内存数量"></el-input-number>
            </el-form-item>
            <div></div>
            <el-form-item prop="cmd" label="命令行">
                <el-input  v-model="item.cmd"  type="textarea" :rows="5" :disabled="readOnly" placeholder="命令行"></el-input>
            </el-form-item>
            <el-form-item label="文件列表:" prop="fileList" v-if="modelFileShow == true&&readOnly==true">
                <el-table
                        :data="item.Files"
                        style="width: 100%"
                        :show-header= false
                        height="300px">
                    <el-table-column  width="80" align="right">
                        <template slot-scope="scope" >
                            <img width="15" v-if="scope.row.IsDir" src="./image/dir.png">
                            <img  width="15" v-if="!scope.row.IsDir" src="./image/file.png">
                        </template>
                    </el-table-column>
                    <el-table-column
                            prop="name"
                            label="文件名"
                            width="580">
                    </el-table-column>

                </el-table>
                <div class="block" style="margin-top: 10px;">
                    <!--<br>-->
                    <div class="r_page">
                        <el-pagination
                                @size-change="handleSizeChange"
                                @current-change="handleCurrentChange"
                                :current-page="pageNum"
                                :page-sizes="[100, 200, 500]"
                                :page-size="pageSize"
                                layout="total, sizes, prev, pager, next, jumper"
                                :total="file_total">
                        </el-pagination>
                    </div>
                </div>
            </el-form-item>
        </div>
        <el-tabs type="card" style="margin-top: 15px;">
        </el-tabs>
        <div>
            <div class="fr" style="">
                <el-form-item>
                    <template >
                        <el-button type="primary" size="small" v-if='item.status != "active"' @click="save">保存</el-button>
                    </template>
                    <el-button type="primary" size="small" v-if='item.status === "active"' @click="stopTask">停止</el-button>
                    <el-button type="primary" size="small" @click="cancel" class="btnReturn" plain>返回</el-button>
                </el-form-item>
            </div>
        </div>
        </div>
        <template v-if="imgetShow">
            <image-form :data="image"  @oncancel="cancelAddImage" @onsuccess="commitAddImage"  :readOnly="readOnly" ></image-form>
        </template>
        <template v-if="dataSetShow">
            <dataSet-form :data="dataset"  @oncancel="cancelAddDataSet" @onsuccess="commitAddDataSet"  :readOnly="readOnly" ></dataSet-form>
        </template>

        <el-dialog :visible.sync="dlModelShow" width="1000px" v-if="dlModelShow">
            <dl-model-form :trainTask="item"  @oncancel="cancelAddDlModel" @onsuccess="commitAddDlMode"  :readOnly="false" ></dl-model-form>
        </el-dialog>

    </el-form>
</template>
<script>
import { SaveTrainTask, UpdateTrainTask, QueryModelFiles, DetailTrainTask, QueryModelFileUrl, StopTrainTask } from '@/api/platform.trainTask'
import { QueryDataSetList } from '@/api/platform.dataset'
import { QueryImageList } from '@/api/platform.image'
import { QueryCodeHouseList } from '@/api/platform.codeHouse'
import ImageForm from '../../image/components/ImageForm'
import DataSetForm from '../../dataset/components/DataSetForm'
import DlModelForm from '../../dlmodel/components/DlmodelForm'
export default {
  name: 'Train_taskForm',
  components: {
    ImageForm,
    DataSetForm,
    DlModelForm
  },
  props: {
    data: {
      type: Object,
      default: function () {
        return {}
      }
    },
    child: {
      type: Boolean,
      default: false
    },
    readOnly: {
      type: Boolean,
      default: false
    }
  },
  data: function () {
    return {
      item: {
        type: 'code'
      },
      dataset_id: null,
      datasetIdDataSet: [],
      dataSetList: [],
      codeHouseList: [],
      imageList: [],
      rules: {},
      dlmodel: {},
      dataset: {},
      image: {},
      dataSetShow: false,
      imgetShow: false,
      dlModelShow: false,
      modelFileShow: true,
      file_total: 0,
      pageSize: 100,
      pageNum: 1
    }
  },
  created: async function () {
    this.selectDataSet()
    this.selectImage()
    this.selectCodeHouse()
  },
  mounted () {
    if (this.data.id != null) {
      this.detail(this.data.id)
      this.modelFileShow = true
    }
  },
  methods: {
    detail: function (id) {
      let that = this
      DetailTrainTask({ id: id }).then(res => {
        that.item = res
        that.dataset_id = res.dataset_id
        that.modelFileShow = true
        that.queryFileList()
        if (that.item.status === 'active') {
          this.readOnly = 'readOnly'
        }

        // SelectByIdDataSet({ id: this.item.datasetId }).then(res=> {
        //  this.datasetIdDataSet = res.data
        // })
        // SelectByIdImage({ id: this.item.imageId }).then(res=> {
        //  this.imageIdImage = res.data
        // })
      })
    },
    queryFileList () {
      let queryparams = {
        params: {
          name: this.item.name
        },
        pageNum: this.pageNum,
        pageSize: this.pageSize
      }
      QueryModelFileUrl(queryparams).then(res => {
        this.item.url = res
        this.queryFileListByUrl()
      })
    },
    queryFileListByUrl () {
      let queryparams = {
        params: {
          url: this.item.url
        },
        pageNum: this.pageNum,
        pageSize: this.pageSize
      }
      QueryModelFiles(queryparams).then(res => {
        this.item.Files = res.Files
        this.file_total = res.FileTotal
      })
    },

    save: function () {
      this.$refs['train_taskform'].validate((valid) => {
        if (valid) {
          let that = this
          this.item.dataset_id = this.item.dataset.id
          if (this.item.codehouse != null) {
            this.item.codehouse_id = this.item.codehouse.id
          }
          if (this.item.id == null) {
            SaveTrainTask(this.item).then(res => {
              this.$message({ message: '保存成功', type: 'success' })
              that.item = res
              this.$emit('onsuccess', this.item)
            })
          } else {
            UpdateTrainTask(this.item).then(res => {
              this.$message({ message: '保存成功', type: 'success' })
              that.item = res
              this.$emit('onsuccess', this.item)
            })
          }
        }
      })
    },
    cancel () {
      this.$emit('oncancel', this.item)
    },
    selectDataSet () {
      let that = this
      QueryDataSetList().then(res => {
        that.dataSetList = res.datas
      })
    },
    selectImage () {
      let that = this
      QueryImageList().then(res => {
        that.imageList = res
      })
    },
    selectCodeHouse () {
      let that = this
      QueryCodeHouseList().then(res => {
        that.codeHouseList = res.datas
      })
    },
    addImage () {
      this.image = {}
      this.imgetShow = true
    },
    addDataSet () {
      this.dataset = {}
      this.dataSetShow = true
    },
    cancelAddImage () {
      this.imgetShow = false
    },
    commitAddImage (model) {
      this.imgetShow = false
      this.item.imageIdImage = model
      this.imageList.unshift(model)
    },
    cancelAddDataSet () {
      this.dataSetShow = false
    },
    commitAddDataSet (dataset) {
      this.dataSetShow = false
      this.item.datasetIdDataSet = dataset
      this.dataSetList.unshift(dataset)
    },
    cancelAddDlModel () {
      this.dlModelShow = false
    },
    commitAddDlMode (dlmode) {
      this.dlModelShow = false
    },
    publish () {
      this.dlModelShow = true
    },
    handleCurrentChange (val) {
      this.pageNum = val
      this.queryFileList()
    },
    handleSizeChange (val) {
      this.pageSize = val
      this.queryFileList()
    },
    stopTask () {
      let that = this
      StopTrainTask({ id: this.item.id }).then(res => {
        that.item = res
        that.dataset_id = res.dataset_id
        that.modelFileShow = true
        that.queryFileList()
      })
    }
  }
}
</script>
<style scoped>
    .el-input {
        width: 500px;
    }
    .el-textarea {
        width: 500px;
    }
    .el-select  {
        width: 500px;
    }
    .upload-demo {
        width: 400px;
    }
    .input-with-select .el-input-group__prepend {
        background-color: #fff;
    }
</style>
