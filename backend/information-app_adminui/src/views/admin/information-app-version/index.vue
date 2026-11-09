
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="平台" prop="platform"><el-select v-model="queryParams.platform"
                                               placeholder="版本管理平台:iOS/Android" clearable size="small">
                                        <el-option
                                                v-for="dict in platformOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="是否更新" prop="forceUpdate"><el-select v-model="queryParams.forceUpdate"
                                               placeholder="版本是否强制更新:是/否" clearable size="small">
                                        <el-option
                                                v-for="dict in forceUpdateOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="类型" prop="updateType"><el-select v-model="queryParams.updateType"
                                               placeholder="版本管理类型:功能更新/问题修复/性能优化" clearable size="small">
                                        <el-option
                                                v-for="dict in updateTypeOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="状态" prop="status"><el-select v-model="queryParams.status"
                                               placeholder="版本管理状态:发布/草稿/撤回" clearable size="small">
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
                                v-permisaction="['admin:informationAppVersion:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationAppVersion:edit']"
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
                                v-permisaction="['admin:informationAppVersion:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="informationAppVersionList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="版本号" align="center" prop="version"
                             :show-overflow-tooltip="true"/><el-table-column label="平台" align="center" prop="platform"
                             :formatter="platformFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ platformFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="下载地址" align="center" prop="downloadUrl"
                             :show-overflow-tooltip="true"/><el-table-column label="是否强制更新" align="center" prop="forceUpdate"
                             :formatter="forceUpdateFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ forceUpdateFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="当前更新内容" align="center" prop="content"
                             :show-overflow-tooltip="true"/><el-table-column label="类型" align="center" prop="updateType"
                             :formatter="updateTypeFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ updateTypeFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="发布渠道" align="center" prop="channel"
                             :show-overflow-tooltip="true"/><el-table-column label="状态" align="center" prop="status"
                             :formatter="statusFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ statusFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="版本历史日志" align="center" prop="versionHistoryLog"
                                                 :show-overflow-tooltip="true"/>
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
                             v-permisaction="['admin:informationAppVersion:edit']"
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
                              v-permisaction="['admin:informationAppVersion:remove']"
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

                                    <el-form-item label="版本号" prop="version">
                                        <el-input v-model="form.version" placeholder="版本号"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="平台" prop="platform">
                                        <el-select v-model="form.platform"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in platformOptions"
                                                        :key="dict.value"
                                                        :label="dict.label"
                                                        :value="dict.value"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="下载地址" prop="downloadUrl">
                                        <el-input v-model="form.downloadUrl" placeholder="下载地址"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="是否更新" prop="forceUpdate">
                                        <el-select v-model="form.forceUpdate"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in forceUpdateOptions"
                                                        :key="dict.value"
                                                        :label="dict.label"
                                                        :value="dict.value"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="当前更新" prop="content">
                                        <el-input v-model="form.content" placeholder="当前更新内容"
                                                      />
                                    </el-form-item>
                      <el-form-item label="英文更新" prop="enContent">
                        <el-input v-model="form.enContent" placeholder="英文的当前更新内容"
                        />
                      </el-form-item>
                                    <el-form-item label="类型" prop="updateType">
                                        <el-select v-model="form.updateType"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in updateTypeOptions"
                                                        :key="dict.value"
                                                        :label="dict.label"
                                                        :value="dict.value"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="发布渠道" prop="channel">
                                        <el-input v-model="form.channel" placeholder="发布渠道"
                                                      />
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
                                    <el-form-item label="文件" prop="file">
                                        <el-input v-model="form.file" placeholder="文件"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="版本日志" prop="versionHistoryLog">
                                        <el-input v-model="form.versionHistoryLog" placeholder="版本历史日志"
                                                      />
                                    </el-form-item>
                      <el-form-item label="英文日志" prop="enVersionHistoryLog">
                        <el-input v-model="form.enVersionHistoryLog" placeholder="英文的版本历史日志"
                        />
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
    import {addInformationAppVersion, delInformationAppVersion, getInformationAppVersion, listInformationAppVersion, updateInformationAppVersion} from '@/api/admin/information-app-version'

    export default {
        name: 'InformationAppVersion',
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
                informationAppVersionList: [],
                platformOptions: [],forceUpdateOptions: [],updateTypeOptions: [],statusOptions: [],
                // 关系表类型

                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    platform:undefined,
                    forceUpdate:undefined,
                    updateType:undefined,
                    status:undefined,

                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {platform:  [ {required: true, message: '平台:1-iOS，2-Android不能为空', trigger: 'blur'} ],
                forceUpdate:  [ {required: true, message: '是否强制更新:2-否，1-是不能为空', trigger: 'blur'} ],
                updateType:  [ {required: true, message: '类型:1-功能更新 2-问题修复 3-性能优化不能为空', trigger: 'blur'} ],
                status:  [ {required: true, message: '状态:1-发布 2-草稿 3-撤回不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getDicts('information_app_version_platform').then(response => {
                this.platformOptions = response.data
            })
            this.getDicts('information_app_version_force_update').then(response => {
                this.forceUpdateOptions = response.data
            })
            this.getDicts('information_app_version_update_type').then(response => {
                this.updateTypeOptions = response.data
            })
            this.getDicts('information_app_version_status').then(response => {
                this.statusOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listInformationAppVersion(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.informationAppVersionList = response.data.list
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
                version: undefined,
                platform: undefined,
                downloadUrl: undefined,
                forceUpdate: undefined,
                content: undefined,
                enContent: undefined,
                updateType: undefined,
                channel: undefined,
                status: undefined,
                file: undefined,
                versionHistoryLog: undefined,
                enVersionHistoryLog: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            platformFormat(row) {
                return this.selectDictLabel(this.platformOptions, row.platform)
            },
            forceUpdateFormat(row) {
                return this.selectDictLabel(this.forceUpdateOptions, row.forceUpdate)
            },
            updateTypeFormat(row) {
                return this.selectDictLabel(this.updateTypeOptions, row.updateType)
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
                this.title = '添加版本管理'
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
                getInformationAppVersion(id).then(response => {
                    this.form = response.data
                  // 将数字转换为字符串，以匹配字典选项的 value 类型，确保下拉框正确回显 label
                  if (this.form.platform !== undefined && this.form.platform !== null) {
                    this.form.platform = String(this.form.platform)
                  }
                  if (this.form.forceUpdate !== undefined && this.form.forceUpdate !== null) {
                    this.form.forceUpdate = String(this.form.forceUpdate)
                  }
                  if (this.form.updateType !== undefined && this.form.updateType !== null) {
                    this.form.updateType = String(this.form.updateType)
                  }
                  if (this.form.status !== undefined && this.form.status !== null) {
                    this.form.status = String(this.form.status)
                  }
                    this.open = true
                    this.title = '修改版本管理'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateInformationAppVersion(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addInformationAppVersion(this.form).then(response => {
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
                      return delInformationAppVersion( { 'ids': Ids })
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
