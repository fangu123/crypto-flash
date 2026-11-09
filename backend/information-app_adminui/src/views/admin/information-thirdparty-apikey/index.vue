
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="类型" prop="type"><el-select v-model="queryParams.type"
                                               placeholder="第三方apikey类型:aiCoin/bybit/odaily/blockBeats" clearable size="small">
                                        <el-option
                                                v-for="dict in typeOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="状态" prop="status"><el-select v-model="queryParams.status"
                                               placeholder="apikey状态:启用/停用" clearable size="small">
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
                                v-permisaction="['admin:informationThirdpartyApikey:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationThirdpartyApikey:edit']"
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
                                v-permisaction="['admin:informationThirdpartyApikey:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="informationThirdpartyApikeyList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="类型" align="center" prop="type"
                                                 :formatter="typeFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ typeFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="url地址" align="center" prop="baseUrl"
                                                 :show-overflow-tooltip="true"/><el-table-column label="状态" align="center" prop="status"
                                                 :formatter="statusFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ statusFormat(scope.row) }}
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
                             v-permisaction="['admin:informationThirdpartyApikey:edit']"
                             size="mini"
                             type="text"
                             icon="el-icon-edit"
                           >修改
                           </el-button>
                         </el-popconfirm>
<!--                         <el-popconfirm
                            class="delete-popconfirm"
                            title="确认要删除吗?"
                            confirm-button-text="删除"
                            @confirm="handleDelete(scope.row)"
                         >
                            <el-button
                              slot="reference"
                              v-permisaction="['admin:informationThirdpartyApikey:remove']"
                              size="mini"
                              type="text"
                              icon="el-icon-delete"
                            >删除
                            </el-button>
                         </el-popconfirm>-->
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

                                    <el-form-item label="类型" prop="type">
                                        <el-select v-model="form.type"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in typeOptions"
                                                        :key="dict.value"
                                                        :label="dict.label"
                                                        :value="dict.value"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="url地址" prop="baseUrl">
                                        <el-input v-model="form.baseUrl" placeholder="url地址"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="apikey" prop="apikey">
                                        <el-input v-model="form.apikey" placeholder="apikey"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="secretkey" prop="secretkey">
                                        <el-input v-model="form.secretkey" placeholder="secretkey"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="password" prop="passphrase">
                                        <el-input v-model="form.passphrase" placeholder="passphrase"
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
    import {addInformationThirdpartyApikey, delInformationThirdpartyApikey, getInformationThirdpartyApikey, listInformationThirdpartyApikey, updateInformationThirdpartyApikey} from '@/api/admin/information-thirdparty-apikey'

    export default {
        name: 'InformationThirdpartyApikey',
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
                informationThirdpartyApikeyList: [],
                typeOptions: [],statusOptions: [],
                // 关系表类型

                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    type:undefined,
                    status:undefined,

                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {type:  [ {required: true, message: '类型：1-aiCoin，2-bybit，3-odaily，4-blockBeats不能为空', trigger: 'blur'} ],
                status:  [ {required: true, message: '状态 1-启用 2-停用不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getDicts('information_thirdparty_apikey_type').then(response => {
                this.typeOptions = response.data
            })
            this.getDicts('information_thirdparty_apikey_status').then(response => {
                this.statusOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listInformationThirdpartyApikey(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.informationThirdpartyApikeyList = response.data.list
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
                type: undefined,
                baseUrl: undefined,
                apikey: undefined,
                secretkey: undefined,
                passphrase: undefined,
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
            typeFormat(row) {
                return this.selectDictLabel(this.typeOptions, row.type)
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
                this.title = '添加第三方apikey'
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
                getInformationThirdpartyApikey(id).then(response => {
                    this.form = response.data
                  if (this.form.type !== undefined && this.form.type !== null) {
                    this.form.type = String(this.form.type)
                  }
                  if (this.form.status !== undefined && this.form.status !== null) {
                    this.form.status = String(this.form.status)
                  }
                    this.open = true
                    this.title = '修改第三方apikey'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateInformationThirdpartyApikey(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addInformationThirdpartyApikey(this.form).then(response => {
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
                      return delInformationThirdpartyApikey( { 'ids': Ids })
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
