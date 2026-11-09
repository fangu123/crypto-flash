
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="用户名" prop="username"><el-input v-model="queryParams.username" placeholder="请输入用户名" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="设备类型" prop="deviceType"><el-select v-model="queryParams.deviceType"
                                               placeholder="用户设备类型:iOS/Android/Web" clearable size="small">
                                        <el-option
                                                v-for="dict in deviceTypeOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="钱包地址" prop="walletAddress">
                          <el-input v-model="queryParams.walletAddress" placeholder="请输入绑定的钱包地址" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="点位通知" prop="pointNotificationSettings">
                          <el-select v-model="queryParams.pointNotificationSettings"
                                               placeholder="用户点位通知:部分/全部" clearable size="small">
                                        <el-option
                                                v-for="dict in pointNotificationSettingsOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="状态" prop="status"><el-select v-model="queryParams.status"
                                               placeholder="用户状态:正常/锁定/禁用" clearable size="small">
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

<!--                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationUser:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationUser:edit']"
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
                                v-permisaction="['admin:informationUser:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>-->

                <el-table v-loading="loading" :data="informationUserList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/>
                  <el-table-column label="用户名" align="center" prop="username"
                                                 :show-overflow-tooltip="true"/>
                  <el-table-column label="邮箱" align="center" prop="email"
                                                 :show-overflow-tooltip="true"/>
                  <el-table-column label="设备类型" align="center" prop="deviceType"
                                                 :formatter="deviceTypeFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ deviceTypeFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="系统版本" align="center" prop="systemVersion"
                                                 :show-overflow-tooltip="true"/>
                  <el-table-column label="App版本号" align="center" prop="appVersion"
                                                 :show-overflow-tooltip="true"/>
                  <el-table-column label="钱包地址" align="center" prop="walletAddress"
                                                 :show-overflow-tooltip="true"/>
                  <el-table-column label="点位查看" align="center" prop="pointNotificationSettings"
                                                 :formatter="pointNotificationSettingsFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ pointNotificationSettingsFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="最后登录时间" align="center" prop="lastLoginAt"
                                                 :show-overflow-tooltip="true">
                                    <template slot-scope="scope">
                                    <span>{{ parseTime(scope.row.lastLoginAt) }}</span>
                                    </template>
                                </el-table-column><el-table-column label="最后登录IP" align="center" prop="lastLoginIp"
                                                 :show-overflow-tooltip="true"/><el-table-column label="状态" align="center" prop="status"
                                                 :formatter="statusFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ statusFormat(scope.row) }}
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
                             v-permisaction="['admin:informationUser:edit']"
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
                              v-permisaction="['admin:informationUser:remove']"
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
    import {addInformationUser, delInformationUser, getInformationUser, listInformationUser, updateInformationUser} from '@/api/admin/information-user'

    export default {
        name: 'InformationUser',
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
                informationUserList: [],
                deviceTypeOptions: [],twoFactorEnabledOptions: [],pointNotificationSettingsOptions: [],statusOptions: [],
                // 关系表类型

                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    username:undefined,
                    deviceType:undefined,
                    walletAddress:undefined,
                    pointNotificationSettings:undefined,
                    status:undefined,

                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {username:  [ {required: true, message: '用户名不能为空', trigger: 'blur'} ],
                deviceType:  [ {required: true, message: '设备类型：1-iOS，2-Android，3-Web不能为空', trigger: 'blur'} ],
                walletAddress:  [ {required: true, message: '绑定的钱包地址不能为空', trigger: 'blur'} ],
                pointNotificationSettings:  [ {required: true, message: '点位通知设置：2-全部，1-部分不能为空', trigger: 'blur'} ],
                status:  [ {required: true, message: '状态：3-禁用，1-正常，2-锁定不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getDicts('information_user_device_type').then(response => {
                this.deviceTypeOptions = response.data
            })
            this.getDicts('information_user_two_factor_enabled').then(response => {
                this.twoFactorEnabledOptions = response.data
            })
            this.getDicts('information_user_point_notification_settings').then(response => {
                this.pointNotificationSettingsOptions = response.data
            })
            this.getDicts('information_user_status').then(response => {
                this.statusOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listInformationUser(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.informationUserList = response.data.list
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
                status: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            deviceTypeFormat(row) {
                return this.selectDictLabel(this.deviceTypeOptions, row.deviceType)
            },
            twoFactorEnabledFormat(row) {
                return this.selectDictLabel(this.twoFactorEnabledOptions, row.twoFactorEnabled)
            },
            pointNotificationSettingsFormat(row) {
                return this.selectDictLabel(this.pointNotificationSettingsOptions, row.pointNotificationSettings)
            },
            statusFormat(row) {
                return this.selectDictLabel(this.statusOptions, row.status)
            },
            // 关系
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
                this.title = '添加用户'
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
                getInformationUser(id).then(response => {
                    this.form = response.data
                  if (this.form.status !== undefined && this.form.status !== null) {
                    this.form.status = String(this.form.status)
                  }
                    this.open = true
                    this.title = '修改用户'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateInformationUser(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addInformationUser(this.form).then(response => {
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
                      return delInformationUser( { 'ids': Ids })
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
