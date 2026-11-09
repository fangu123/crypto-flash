
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="资讯标题" prop="newsId"><el-select v-model="queryParams.newsId"
                                           placeholder="请选择" clearable size="small" >
                                    <el-option
                                            v-for="dict in newsIdOptions"
                                            :key="dict.key"
                                            :label="dict.value"
                                            :value="dict.key"
                                    />
                                </el-select>
                            </el-form-item>
                        <el-form-item label="评论用户" prop="userId"><el-select v-model="queryParams.userId"
                                           placeholder="请选择" clearable size="small" >
                                    <el-option
                                            v-for="dict in userIdOptions"
                                            :key="dict.key"
                                            :label="dict.value"
                                            :value="dict.key"
                                    />
                                </el-select>
                            </el-form-item>
<!--                        <el-form-item label="父级评论" prop="parentId"><el-select v-model="queryParams.parentId"
                                           placeholder="请选择" clearable size="small" >
                                    <el-option
                                            v-for="dict in parentIdOptions"
                                            :key="dict.key"
                                            :label="dict.value"
                                            :value="dict.key"
                                    />
                                </el-select>
                            </el-form-item>
                        <el-form-item label="回复用户" prop="replyToUserId"><el-select v-model="queryParams.replyToUserId"
                                           placeholder="请选择" clearable size="small" >
                                    <el-option
                                            v-for="dict in replyToUserIdOptions"
                                            :key="dict.key"
                                            :label="dict.value"
                                            :value="dict.key"
                                    />
                                </el-select>
                            </el-form-item>-->
                        <el-form-item label="是否置顶" prop="isTop"><el-select v-model="queryParams.isTop"
                                               placeholder="评论是否置顶：是/否" clearable size="small">
                                        <el-option
                                                v-for="dict in isTopOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="状态" prop="status"><el-select v-model="queryParams.status"
                                               placeholder="评论状态：待审核/正常/违规/已删除" clearable size="small">
                                        <el-option
                                                v-for="dict in statusOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
<!--                        <el-form-item label="发布时间" prop="createTime"><el-input v-model="queryParams.createTime" placeholder="请输入发布时间" clearable-->
<!--                                              size="small" <el-date-picker-->
<!--                          v-model="form.createTime"-->
<!--                          type="datetime"-->
<!--                          placeholder="选择日期">-->
<!--                        </el-date-picker>-->
<!--                            </el-form-item>-->

                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>

                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationComment:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationComment:edit']"
                                type="success"
                                icon="el-icon-edit"
                                size="mini"
                                :disabled="single"
                                @click="handleUpdate"
                        >修改
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationComment:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="informationCommentList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="资讯标题" align="center" prop="newsId" :formatter="newsIdFormat" width="100">
                                <template slot-scope="scope">
                                    {{ newsIdFormat(scope.row) }}
                                </template>
                            </el-table-column><el-table-column label="评论用户" align="center" prop="userId" :formatter="userIdFormat" width="100">
                                <template slot-scope="scope">
                                    {{ userIdFormat(scope.row) }}
                                </template>
                   </el-table-column>
                  <!--         <el-table-column label="父级评论" align="center" prop="parentId" :formatter="parentIdFormat" width="100">
                                                  <template slot-scope="scope">
                                                      {{ parentIdFormat(scope.row) }}
                                                  </template>
                                              </el-table-column><el-table-column label="回复用户" align="center" prop="replyToUserId" :formatter="replyToUserIdFormat" width="100">
                                                  <template slot-scope="scope">
                                                      {{ replyToUserIdFormat(scope.row) }}
                                                  </template>
                                              </el-table-column>-->
                  <el-table-column label="评论内容" align="center" prop="content"
                                                 :show-overflow-tooltip="true"/><el-table-column label="点赞数" align="center" prop="likeCount"
                                                 :show-overflow-tooltip="true"/><el-table-column label="是否置顶" align="center" prop="isTop"
                                                 :formatter="isTopFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ isTopFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="状态" align="center" prop="status"
                                                 :formatter="statusFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ statusFormat(scope.row) }}
                                    </template>
                                </el-table-column>
<!--                  <el-table-column label="IP归属地" align="center" prop="ipLocation"
                                                 :show-overflow-tooltip="true"/>-->
                  <el-table-column label="发布时间" align="center" prop="createTime"
                                                 :show-overflow-tooltip="true">
                                    <template slot-scope="scope">
                                    <span>{{ parseTime(scope.row.createdAt) }}</span>
                                    </template>
                                </el-table-column>
<!--                  <el-table-column label="更新时间" align="center" prop="updateTime"-->
<!--                                                 :show-overflow-tooltip="true">-->
<!--                                    <template slot-scope="scope">-->
<!--                                    <span>{{ parseTime(scope.row.updatedAt) }}</span>-->
<!--                                    </template>-->
<!--                                </el-table-column>-->
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
                             v-permisaction="['admin:informationComment:edit']"
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
                              v-permisaction="['admin:informationComment:remove']"
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

                                    <el-form-item label="资讯" prop="newsId">
                                        <el-select v-model="form.newsId"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in newsIdOptions"
                                                        :key="dict.key"
                                                        :label="dict.value"
                                                        :value="dict.key"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="评论用户" prop="userId">
                                        <el-select v-model="form.userId"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in userIdOptions"
                                                        :key="dict.key"
                                                        :label="dict.value"
                                                        :value="dict.key"
                                                />
                                            </el-select>
                                    </el-form-item>
<!--                                    <el-form-item label="父级评论" prop="parentId">
                                        <el-select v-model="form.parentId"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in parentIdOptions"
                                                        :key="dict.key"
                                                        :label="dict.value"
                                                        :value="dict.key"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="回复用户" prop="replyToUserId">
                                        <el-select v-model="form.replyToUserId"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in replyToUserIdOptions"
                                                        :key="dict.key"
                                                        :label="dict.value"
                                                        :value="dict.key"
                                                />
                                            </el-select>
                                    </el-form-item>-->
                                    <el-form-item label="评论内容" prop="content">
                                        <el-input v-model="form.content" placeholder="评论内容"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="点赞数" prop="likeCount">
                                        <el-input v-model="form.likeCount" placeholder="点赞数"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="是否置顶" prop="isTop">
                                        <el-select v-model="form.isTop"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in isTopOptions"
                                                        :key="dict.value"
                                                        :label="dict.label"
                                                        :value="dict.value"
                                                />
                                            </el-select>
                                    </el-form-item>
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
<!--                                    <el-form-item label="IP归属地" prop="ipLocation">
                                        <el-input v-model="form.ipLocation" placeholder="IP归属地"
                                                      />
                                    </el-form-item>-->
                                    <el-form-item label="发布时间" prop="createTime">
                                        <el-date-picker
                                                    v-model="form.createdAt"
                                                    type="datetime"
                                                    placeholder="选择日期">
                                            </el-date-picker>
                                    </el-form-item>
<!--                                    <el-form-item label="更新时间" prop="updateTime">-->
<!--                                        <el-date-picker-->
<!--                                                    v-model="form.updateTime"-->
<!--                                                    type="datetime"-->
<!--                                                    placeholder="选择日期">-->
<!--                                            </el-date-picker>-->
<!--                                    </el-form-item>-->
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
    import {addInformationComment, delInformationComment, getInformationComment, listInformationComment, updateInformationComment} from '@/api/admin/information-comment'

    import {listInformationNews } from '@/api/admin/information-news'
    import {listInformationUser } from '@/api/admin/information-user'
    export default {
        name: 'InformationComment',
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
                // 弹出层标题
                title: '',
                // 是否显示弹出层
                open: false,
                isEdit: false,
                // 类型数据字典
                typeOptions: [],
                informationCommentList: [],
                isTopOptions: [],statusOptions: [],
                // 关系表类型
                newsIdOptions :[],
                userIdOptions :[],
                parentIdOptions :[],
                replyToUserIdOptions :[],

                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    newsId:undefined,
                    userId:undefined,
                    parentId:undefined,
                    replyToUserId:undefined,
                    isTop:undefined,
                    status:undefined,
                    createTime:undefined,

                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {newsId:  [ {required: true, message: '资讯标题不能为空', trigger: 'blur'} ],
                userId:  [ {required: true, message: '评论用户不能为空', trigger: 'blur'} ],
                // parentId:  [ {required: true, message: '父级评论不能为空', trigger: 'blur'} ],
                // replyToUserId:  [ {required: true, message: '回复用户不能为空', trigger: 'blur'} ],
                isTop:  [ {required: true, message: '是否置顶不能为空', trigger: 'blur'} ],
                status:  [ {required: true, message: '状态不能为空', trigger: 'blur'} ],
                  createdAt:  [ {required: true, message: '发布时间不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getInformationNewsItems()
            this.getInformationUserItems1()
            this.getInformationUserItems2()
            this.getInformationUserItems3()
            this.getDicts('information_comment_is_top').then(response => {
                this.isTopOptions = response.data
            })
            this.getDicts('information_comment_status').then(response => {
                this.statusOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listInformationComment(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.informationCommentList = response.data.list
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
                newsId: undefined,
                userId: undefined,
                parentId: undefined,
                replyToUserId: undefined,
                content: undefined,
                likeCount: undefined,
                isTop: undefined,
                status: undefined,
                ipLocation: undefined,
                  createdAt: undefined,
                updateTime: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            newsIdFormat(row) {
                return this.selectItemsLabel(this.newsIdOptions, row.newsId)
            },
            userIdFormat(row) {
                return this.selectItemsLabel(this.userIdOptions, row.userId)
            },
            parentIdFormat(row) {
                return this.selectItemsLabel(this.parentIdOptions, row.parentId)
            },
            replyToUserIdFormat(row) {
                return this.selectItemsLabel(this.replyToUserIdOptions, row.replyToUserId)
            },
            isTopFormat(row) {
                return this.selectDictLabel(this.isTopOptions, row.isTop)
            },
            statusFormat(row) {
                return this.selectDictLabel(this.statusOptions, row.status)
            },
            // 关系
            getInformationNewsItems() {
               this.getItems(listInformationNews, undefined).then(res => {
                   this.newsIdOptions = this.setItems(res, 'id', 'title')
               })
            },
          getInformationUserItems1() {
            this.getItems(listInformationUser, undefined).then(res => {
              this.userIdOptions = this.setItems(res, 'id', 'username')
            })
          },
          getInformationUserItems2() {
            this.getItems(listInformationUser, undefined).then(res => {
              this.parentIdOptions = this.setItems(res, 'id', 'username')
            })
          },
            getInformationUserItems3() {
               this.getItems(listInformationUser, undefined).then(res => {
                   this.replyToUserIdOptions = this.setItems(res, 'id', 'username')
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
                this.title = '添加资讯评论'
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
                getInformationComment(id).then(response => {
                    this.form = response.data
                  if (this.form.status !== undefined && this.form.status !== null) {
                    this.form.status = String(this.form.status)
                  }
                  if (this.form.isTop !== undefined && this.form.isTop !== null) {
                    this.form.isTop = String(this.form.isTop)
                  }
                  if (this.form.newsId !== undefined && this.form.newsId !== null) {
                    this.form.newsId = String(this.form.newsId)
                  }
                  if (this.form.userId !== undefined && this.form.userId !== null) {
                    this.form.userId = String(this.form.userId)
                  }
                  if (this.form.parentId !== undefined && this.form.parentId !== null) {
                    this.form.parentId = String(this.form.parentId)
                  }
                  if (this.form.replyToUserId !== undefined && this.form.replyToUserId !== null) {
                    this.form.replyToUserId = String(this.form.replyToUserId)
                  }
                  if (this.form.likeCount !== undefined && this.form.likeCount !== null) {
                    this.form.likeCount = String(this.form.likeCount)
                  }
                    this.open = true
                    this.title = '修改资讯评论'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateInformationComment(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addInformationComment(this.form).then(response => {
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
                      return delInformationComment( { 'ids': Ids })
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
