
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label="标题" prop="title"><el-input v-model="queryParams.title" placeholder="请输入标题" clearable
                                                            size="small" @keyup.enter.native="handleQuery"/>
          </el-form-item>
          <el-form-item label="币种" prop="relatedCoinId"><el-select v-model="queryParams.relatedCoinId"
                                                                     placeholder="请选择" clearable size="small" >
            <el-option
              v-for="dict in relatedCoinIdOptions"
              :key="dict.key"
              :label="dict.value"
              :value="dict.key"
            />
          </el-select>
          </el-form-item>
          <el-form-item label="点位" prop="relatedTradingPointId"><el-select v-model="queryParams.relatedTradingPointId"
                                                                             placeholder="请选择" clearable size="small" >
            <el-option
              v-for="dict in relatedTradingPointIdOptions"
              :key="dict.key"
              :label="dict.value"
              :value="dict.key"
            />
          </el-select>
          </el-form-item>
          <el-form-item label="级别" prop="level"><el-select v-model="queryParams.level"
                                                             placeholder="资讯级别:普通/重要/紧急" clearable size="small">
            <el-option
              v-for="dict in levelOptions"
              :key="dict.value"
              :label="dict.label"
              :value="dict.value"
            />
          </el-select>
          </el-form-item>
          <el-form-item label="情绪" prop="sentiment"><el-select v-model="queryParams.sentiment"
                                                                 placeholder="资讯情绪:中性/利好/利空" clearable size="small">
            <el-option
              v-for="dict in sentimentOptions"
              :key="dict.value"
              :label="dict.label"
              :value="dict.value"
            />
          </el-select>
          </el-form-item>
<!--          <el-form-item label="是否头条" prop="isFeatured"><el-select v-model="queryParams.isFeatured"
                                                                      placeholder="资讯是否头条:是/否" clearable size="small">
            <el-option
              v-for="dict in isFeaturedOptions"
              :key="dict.value"
              :label="dict.label"
              :value="dict.value"
            />
          </el-select>
          </el-form-item>-->
          <el-form-item label="状态" prop="status"><el-select v-model="queryParams.status"
                                                              placeholder="资讯状态:草稿/已发布/已下架" clearable size="small">
            <el-option
              v-for="dict in statusOptions"
              :key="dict.value"
              :label="dict.label"
              :value="dict.value"
            />
          </el-select>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:informationNews:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:informationNews:edit']"
              type="success"
              icon="el-icon-edit"
              size="mini"
              :disabled="single"
              @click="handleUpdate"
            >修改
            </el-button>
          </el-col>
          <!--                    <el-col :span="1.5">
                                  <el-button
                                          v-permisaction="['admin:informationNews:remove']"
                                          type="danger"
                                          icon="el-icon-delete"
                                          size="mini"
                                          :disabled="multiple"
                                          @click="handleDelete"
                                  >删除
                                  </el-button>
                              </el-col>-->
        </el-row>

        <el-table v-loading="loading" :data="informationNewsList" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center"/>
          <el-table-column label="标题" align="center" prop="title">
            <template slot-scope="scope">
              <el-tooltip :content="scope.row.title" effect="light" popper-class="custom-tooltip-width" placement="top">
                <span class="overflow-text">{{ scope.row.title }}</span>
              </el-tooltip>
            </template>
          </el-table-column>

          <el-table-column label="内容" align="center" prop="content">
            <template slot-scope="scope">
              <el-tooltip :content="scope.row.content" effect="light" popper-class="custom-tooltip-width" placement="top">
                <span class="overflow-text">{{ scope.row.content }}</span>
              </el-tooltip>
            </template>
          </el-table-column>
          <el-table-column label="来源" align="center" prop="source"
                           :show-overflow-tooltip="true"/>
          <el-table-column label="相关币种" align="center" prop="relatedCoinId" :formatter="relatedCoinIdFormat" width="100">
            <template slot-scope="scope">
              {{ relatedCoinIdFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column label="相关点位" align="center" prop="relatedTradingPointId" :formatter="relatedTradingPointIdFormat" width="100">
            <template slot-scope="scope">
              {{ relatedTradingPointIdFormat(scope.row) }}
            </template>
          </el-table-column><el-table-column label="级别" align="center" prop="level"
                                             :formatter="levelFormat" width="100">
          <template slot-scope="scope">
            {{ levelFormat(scope.row) }}
          </template>
        </el-table-column><el-table-column label="情绪" align="center" prop="sentiment"
                                           :formatter="sentimentFormat" width="100">
          <template slot-scope="scope">
            {{ sentimentFormat(scope.row) }}
          </template>
        </el-table-column><!--<el-table-column label="是否头条" align="center" prop="isFeatured"
                                                 :formatter="isFeaturedFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ isFeaturedFormat(scope.row) }}
                                    </template>
                                </el-table-column>-->
          <el-table-column label="浏览次数" align="center" prop="viewCount"
                          :show-overflow-tooltip="true"/><el-table-column label="点赞次数" align="center" prop="likeCount"
                          :show-overflow-tooltip="true"/><el-table-column label="分享次数" align="center" prop="shareCount"
                          :show-overflow-tooltip="true"/><el-table-column label="状态" align="center" prop="status"
                           :formatter="statusFormat" width="100">
          <template slot-scope="scope">
            {{ statusFormat(scope.row) }}
          </template>
        </el-table-column><el-table-column label="创建时间" align="center" prop="createdAt"
                                           :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.createdAt) }}</span>
          </template>
        </el-table-column>
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-popconfirm
                class="delete-popconfirm"
                title="确认要修改吗?"
                confirm-button-text="修改"
                @confirm="handleUpdate(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:informationNews:edit']"
                  size="mini"
                  type="text"
                  icon="el-icon-edit"
                >修改
                </el-button>
              </el-popconfirm>
              <el-popconfirm
                class="delete-popconfirm"
                title="确认要删除吗?"
                confirm-button-text="删除"
                @confirm="handleDelete(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:informationNews:remove']"
                  size="mini"
                  type="text"
                  icon="el-icon-delete"
                >删除
                </el-button>
              </el-popconfirm>
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total>0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">

            <el-form-item label="标题" prop="title">
              <el-input v-model="form.title" placeholder="标题"
              />
            </el-form-item>
            <el-form-item label="英文标题" prop="enTitle">
              <el-input v-model="form.enTitle" placeholder="英文的标题"
              />
            </el-form-item>
            <el-form-item label="内容" prop="content">
              <el-input v-model="form.content" placeholder="内容"
              />
            </el-form-item>
            <el-form-item label="英文内容" prop="enContent">
              <el-input v-model="form.enContent" placeholder="英文的内容"
              />
            </el-form-item>
            <!--                                    <el-form-item label="摘要" prop="summary">
                                                    <el-input v-model="form.summary" placeholder="摘要"
                                                                  />
                                                </el-form-item>-->
            <!--                                    <el-form-item label="封面图片" prop="coverImage">
                                                    <el-input v-model="form.coverImage" placeholder="封面图片"
                                                                  />
                                                </el-form-item>-->
            <!--                                    <el-form-item label="来源" prop="source">
                                                    <el-input v-model="form.source" placeholder="来源"
                                                                  />
                                                </el-form-item>-->
            <!--                                    <el-form-item label="logo图标" prop="logUrl">
                                                    <el-input v-model="form.logUrl" placeholder="logo图标"
                                                                  />
                                                </el-form-item>-->
            <el-form-item label="相关币种" prop="relatedCoinId">
              <el-select v-model="form.relatedCoinId"
                         placeholder="请选择" >
                <el-option
                  v-for="dict in relatedCoinIdOptions"
                  :key="dict.key"
                  :label="dict.value"
                  :value="dict.key"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="相关点位" prop="relatedTradingPointId">
              <el-select v-model="form.relatedTradingPointId"
                         placeholder="请选择" >
                <el-option
                  v-for="dict in relatedTradingPointIdOptions"
                  :key="dict.key"
                  :label="dict.value"
                  :value="dict.key"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="级别" prop="level">
              <el-select v-model="form.level"
                         placeholder="请选择" >
                <el-option
                  v-for="dict in levelOptions"
                  :key="dict.value"
                  :label="dict.label"
                  :value="dict.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="情绪" prop="sentiment">
              <el-select v-model="form.sentiment"
                         placeholder="请选择" >
                <el-option
                  v-for="dict in sentimentOptions"
                  :key="dict.value"
                  :label="dict.label"
                  :value="dict.value"
                />
              </el-select>
            </el-form-item>
            <!--                                    <el-form-item label="是否头条" prop="isFeatured">
                                                    <el-select v-model="form.isFeatured"
                                                                   placeholder="请选择" >
                                                            <el-option
                                                                    v-for="dict in isFeaturedOptions"
                                                                    :key="dict.value"
                                                                    :label="dict.label"
                                                                    :value="dict.value"
                                                            />
                                                        </el-select>
                                                </el-form-item>-->
            <el-form-item label="状态" prop="status">
              <el-select v-model="form.status"
                         placeholder="请选择" >
                <el-option
                  v-for="dict in statusOptions"
                  :key="dict.value"
                  :label="dict.label"
                  :value="dict.value"
                />
              </el-select>
            </el-form-item>
            <!--                                    <el-form-item label="发布时间" prop="publishedAt">
                                                    <el-date-picker
                                                                v-model="form.publishedAt"
                                                                type="datetime"
                                                                placeholder="选择日期">
                                                        </el-date-picker>
                                                </el-form-item>-->
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {addInformationNews, delInformationNews, getInformationNews, listInformationNews, updateInformationNews} from '@/api/admin/information-news'

import {listInformationCoin } from '@/api/admin/information-coin'
import {listInformationTradingPoint } from '@/api/admin/information-trading-point'
import item from "@/layout/components/Sidebar/Item.vue";
export default {
  name: 'InformationNews',
  components: {
  },
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 日期范围
      dateRange: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      isEdit: false,
      // 类型数据字典
      typeOptions: [],
      informationNewsList: [],
      levelOptions: [],sentimentOptions: [],isFeaturedOptions: [],statusOptions: [],
      // 关系表类型
      relatedCoinIdOptions :[],
      relatedTradingPointIdOptions :[],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        title:undefined,
        relatedCoinId:undefined,
        relatedTradingPointId:undefined,
        level:undefined,
        sentiment:undefined,
        isFeatured:undefined,
        status:undefined,

      },
      // 表单参数
      form: {
      },
      // 表单校验
      rules: {title:  [ {required: true, message: '标题不能为空', trigger: 'blur'} ],
        relatedCoinId:  [ {required: true, message: '相关币种id不能为空', trigger: 'blur'} ],
        relatedTradingPointId:  [ {required: true, message: '相关点位id不能为空', trigger: 'blur'} ],
        level:  [ {required: true, message: '级别：1-普通，2-重要，3-紧急不能为空', trigger: 'blur'} ],
        sentiment:  [ {required: true, message: '情绪：1-中性，2-利好，3-利空不能为空', trigger: 'blur'} ],
        isFeatured:  [ {required: true, message: '是否头条：2-否，1-是不能为空', trigger: 'blur'} ],
        status:  [ {required: true, message: '状态：1-草稿，2-已发布，3-已下架不能为空', trigger: 'blur'} ],
        content:  [ {required: true, message: '内容不能为空', trigger: 'blur'} ],
        enTitle:  [ {required: true, message: '英文标题不能为空', trigger: 'blur'} ],
        enContent:  [ {required: true, message: '英文内容不能为空', trigger: 'blur'} ],
      }
    }
  },
  created() {
    this.getList()
    this.getInformationCoinItems()
    this.getInformationTradingPointItems()
    this.getDicts('information_news_level').then(response => {
      this.levelOptions = response.data
    })
    this.getDicts('information_news_sentiment').then(response => {
      this.sentimentOptions = response.data
    })
    this.getDicts('information_news_is_featured').then(response => {
      this.isFeaturedOptions = response.data
    })
    this.getDicts('information_news_status').then(response => {
      this.statusOptions = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listInformationNews(this.queryParams, this.dateRange).then(response => {
          this.informationNewsList = response.data.list
          this.total = response.data.count
          this.loading = false
        }
      )
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {

        id: undefined,
        title: undefined,
        content: undefined,
        enTitle: undefined,
        enContent: undefined,
        summary: undefined,
        coverImage: undefined,
        source: undefined,
        logoUrl: undefined,
        relatedCoinId: undefined,
        relatedTradingPointId: undefined,
        level: undefined,
        sentiment: undefined,
        isFeatured: undefined,
        status: undefined,
        publishedAt: undefined,
      }
      this.resetForm('form')
    },
    getImgList: function() {
      this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
    },
    fileClose: function() {
      this.fileOpen = false
    },
    relatedCoinIdFormat(row) {
      return this.selectItemsLabel(this.relatedCoinIdOptions, row.relatedCoinId)
    },
    relatedTradingPointIdFormat(row) {
      return this.selectItemsLabel(this.relatedTradingPointIdOptions, row.relatedTradingPointId)
    },
    levelFormat(row) {
      return this.selectDictLabel(this.levelOptions, row.level)
    },
    sentimentFormat(row) {
      return this.selectDictLabel(this.sentimentOptions, row.sentiment)
    },
    isFeaturedFormat(row) {
      return this.selectDictLabel(this.isFeaturedOptions, row.isFeatured)
    },
    statusFormat(row) {
      return this.selectDictLabel(this.statusOptions, row.status)
    },
    // 关系
    getInformationCoinItems() {
      this.getItems(listInformationCoin, undefined).then(res => {
        this.relatedCoinIdOptions = this.setItems(res, 'id', 'symbol')
      })
    },
    getInformationTradingPointItems() {
      this.getItems(listInformationTradingPoint, undefined).then(res => {
        this.relatedTradingPointIdOptions = this.setItems(res, 'id', 'name')
      })
    },
    // 文件
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加资讯'
      this.isEdit = false
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const id =
        row.id || this.ids
      getInformationNews(id).then(response => {
        this.form = response.data

        if (this.form.relatedCoinId !== undefined && this.form.relatedCoinId !== null) {
          this.form.relatedCoinId = String(this.form.relatedCoinId)
        }
        if (this.form.relatedTradingPointId !== undefined && this.form.relatedTradingPointId !== null) {
          this.form.relatedTradingPointId = String(this.form.relatedTradingPointId)
        }

        if (this.form.level !== undefined && this.form.level !== null) {
          this.form.level = String(this.form.level)
        }
        if (this.form.sentiment !== undefined && this.form.sentiment !== null) {
          this.form.sentiment = String(this.form.sentiment)
        }
        if (this.form.isFeatured !== undefined && this.form.isFeatured !== null) {
          this.form.isFeatured = String(this.form.isFeatured)
        }
        if (this.form.status !== undefined && this.form.status !== null) {
          this.form.status = String(this.form.status)
        }
        this.open = true
        this.title = '修改资讯'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateInformationNews(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addInformationNews(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      var Ids = (row.id && [row.id]) || this.ids

      this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function () {
        return delInformationNews( { 'ids': Ids })
      }).then((response) => {
        if (response.code === 200) {
          this.msgSuccess(response.msg)
          this.open = false
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      }).catch(function () {
      })
    }
  }
}
</script>

<style>
.custom-tooltip-width {
  max-width: 800px;
  word-break: break-word;
  white-space: normal;
}
.overflow-text {
  display: inline-block;
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
