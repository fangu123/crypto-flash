
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="用户" prop="userId"><el-select v-model="queryParams.userId"
                                           placeholder="请选择" clearable size="small" >
                                    <el-option
                                            v-for="dict in userIdOptions"
                                            :key="dict.key"
                                            :label="dict.value"
                                            :value="dict.key"
                                    />
                                </el-select>
                            </el-form-item>
                        <el-form-item label="消息标题" prop="notificationId"><el-select v-model="queryParams.notificationId"
                                           placeholder="请选择" clearable size="small" >
                                    <el-option
                                            v-for="dict in notificationIdOptions"
                                            :key="dict.key"
                                            :label="dict.value"
                                            :value="dict.key"
                                    />
                                </el-select>
                            </el-form-item>
                        <el-form-item label="是否已读" prop="isRead"><el-select v-model="queryParams.isRead"
                                               placeholder="消息通知用户是否已读：1-已读，2-未读" clearable size="small">
                                        <el-option
                                                v-for="dict in isReadOptions"
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

<!--                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationNotificationUser:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                :disabled="single"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationNotificationUser:edit']"
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
                                v-permisaction="['admin:informationNotificationUser:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>-->

                <el-table v-loading="loading" :data="informationNotificationUserList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/>
                  <el-table-column label="用户" align="center" prop="userId" :formatter="userIdFormat" width="200">
                                <template slot-scope="scope">
                                    {{ userIdFormat(scope.row) }}
                                </template>
                            </el-table-column>
                  <el-table-column label="消息标题" align="center" prop="notificationId" :formatter="notificationIdFormat" width="500">
                                <template slot-scope="scope">
                                    {{ notificationIdFormat(scope.row) }}
                                </template>
                            </el-table-column><el-table-column label="是否已读" align="center" prop="isRead"
                                                 :formatter="isReadFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ isReadFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="阅读时间" align="center" prop="readAt"
                                                 :show-overflow-tooltip="true">
                                    <template slot-scope="scope">
                                    <span>{{ parseTime(scope.row.readAt) }}</span>
                                    </template>
                                </el-table-column>
<!--                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                         <el-popconfirm
                           class="delete-popconfirm"
                           title="确认要修改吗?"
                           confirm-button-text="修改"
                           @confirm="handleUpdate(scope.row)"
                         >
                           <el-button
                             slot="reference"
                             v-permisaction="['admin:informationNotificationUser:edit']"
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
                              v-permisaction="['admin:informationNotificationUser:remove']"
                              size="mini"
                              type="text"
                              icon="el-icon-delete"
                            >删除
                            </el-button>
                         </el-popconfirm>
                        </template>
                    </el-table-column>-->
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
    import {addInformationNotificationUser, delInformationNotificationUser, getInformationNotificationUser, listInformationNotificationUser, updateInformationNotificationUser} from '@/api/admin/information-notification-user'

    import {listInformationUser } from '@/api/admin/information-user'
    import {listInformationNotification} from "@/api/admin/information-notification";
    export default {
        name: 'InformationNotificationUser',
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
                informationNotificationUserList: [],
                isReadOptions: [],
                // 关系表类型
                userIdOptions :[],
                notificationIdOptions :[],

                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    userId:undefined,
                    notificationId:undefined,
                    isRead:undefined,

                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {userId:  [ {required: true, message: '用户不能为空', trigger: 'blur'} ],
                notificationId:  [ {required: true, message: '消息标题不能为空', trigger: 'blur'} ],
                isRead:  [ {required: true, message: '是否已读：1-已读，2-未读不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getInformationUserItems()
            this.getInformationNewsItems()
            this.getDicts('information_notification_is_read').then(response => {
                this.isReadOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listInformationNotificationUser(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.informationNotificationUserList = response.data.list
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
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            userIdFormat(row) {
                return this.selectItemsLabel(this.userIdOptions, row.userId)
            },
            notificationIdFormat(row) {
                return this.selectItemsLabel(this.notificationIdOptions, row.notificationId)
            },
            isReadFormat(row) {
                return this.selectDictLabel(this.isReadOptions, row.isRead)
            },
            // 关系
            getInformationUserItems() {
               this.getItems(listInformationUser, undefined).then(res => {
                   this.userIdOptions = this.setItems(res, 'id', 'username')
               })
            },
            getInformationNewsItems() {
               this.getItems(listInformationNotification, undefined).then(res => {
                   this.notificationIdOptions = this.setItems(res, 'id', 'title')
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
                this.title = '添加消息通知用户'
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
                getInformationNotificationUser(id).then(response => {
                    this.form = response.data
                  if (this.form.isRead !== undefined && this.form.isRead !== null) {
                    this.form.isRead = String(this.form.isRead)
                  }
                    this.open = true
                    this.title = '修改消息通知用户'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateInformationNotificationUser(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addInformationNotificationUser(this.form).then(response => {
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
                      return delInformationNotificationUser( { 'ids': Ids })
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
