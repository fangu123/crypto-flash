
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="币种" prop="coinId"><el-select v-model="queryParams.coinId"
                                           placeholder="请选择" clearable size="small" >
                                    <el-option
                                            v-for="dict in coinIdOptions"
                                            :key="dict.key"
                                            :label="dict.value"
                                            :value="dict.key"
                                    />
                                </el-select>
                            </el-form-item>
                        <el-form-item label="类型" prop="type"><el-select v-model="queryParams.type"
                                               placeholder="点位类型:做多/做空" clearable size="small">
                                        <el-option
                                                v-for="dict in typeOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="状态" prop="status"><el-select v-model="queryParams.status"
                                               placeholder="点位状态:进行中/已失效/已撤单" clearable size="small">
                                        <el-option
                                                v-for="dict in statusOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="是否通知" prop="isNotified"><el-select v-model="queryParams.isNotified"
                                               placeholder="点位是否已通知:未通知/已通知" clearable size="small">
                                        <el-option
                                                v-for="dict in isNotifiedOptions"
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
                                v-permisaction="['admin:informationTradingPoint:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationTradingPoint:edit']"
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
                                v-permisaction="['admin:informationTradingPoint:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="informationTradingPointList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/>
                  <el-table-column label="币种" align="center" prop="coinId" :formatter="coinIdFormat" width="100">
                                <template slot-scope="scope">
                                    {{ coinIdFormat(scope.row) }}
                                </template>
                            </el-table-column><el-table-column label="类型" align="center" prop="type"
                                                 :formatter="typeFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ typeFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="入场价格" align="center" prop="entryPrice"
                                                 :show-overflow-tooltip="true"/>
                  <el-table-column label="止盈价格" align="center" prop="takeProfits"
                                                 :show-overflow-tooltip="true"/>
                  <el-table-column label="止损价格" align="center" prop="stopLoss"
                                                 :show-overflow-tooltip="true"/>
                  <el-table-column label="有效时长" align="center" prop="validDuration"
                                                 :show-overflow-tooltip="true"/>
                  <el-table-column label="到期时间" align="center" prop="validUntil"
                                                 :show-overflow-tooltip="true">
                                    <template slot-scope="scope">
                                    <span>{{ parseTime(scope.row.validUntil) }}</span>
                                    </template>
                                </el-table-column><el-table-column label="状态" align="center" prop="status"
                                                 :formatter="statusFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ statusFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="是否通知" align="center" prop="isNotified"
                                                 :formatter="isNotifiedFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ isNotifiedFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="备注" align="center" prop="notes"
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
                             v-permisaction="['admin:informationTradingPoint:edit']"
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
                              v-permisaction="['admin:informationTradingPoint:remove']"
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

                                    <el-form-item label="币种" prop="coinId">
                                        <el-select v-model="form.coinId"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in coinIdOptions"
                                                        :key="dict.key"
                                                        :label="dict.value"
                                                        :value="dict.key"
                                                />
                                            </el-select>
                                    </el-form-item>
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
                                    <el-form-item label="入场价格" prop="entryPrice">
                                        <el-input v-model="form.entryPrice" placeholder="入场价格"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="止盈价格" prop="takeProfits">
                                        <el-input v-model="form.takeProfits" placeholder="止盈价格数组"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="止损价格" prop="stopLoss">
                                        <el-input v-model="form.stopLoss" placeholder="止损价格"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="有效时长" prop="validDuration">
                                        <el-input v-model="form.validDuration"  type="number"  placeholder="有效时长：6小时，12小时，24小时，48小时"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="到期时间" prop="validUntil">
                                        <el-date-picker
                                                    v-model="form.validUntil"
                                                    type="datetime"
                                                    placeholder="选择日期">
                                            </el-date-picker>
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
                                    <el-form-item label="是否通知" prop="isNotified">
                                        <el-select v-model="form.isNotified"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in isNotifiedOptions"
                                                        :key="dict.value"
                                                        :label="dict.label"
                                                        :value="dict.value"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="备注" prop="notes">
                                        <el-input v-model="form.notes" placeholder="备注"
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
    import {addInformationTradingPoint, delInformationTradingPoint, getInformationTradingPoint, listInformationTradingPoint, updateInformationTradingPoint} from '@/api/admin/information-trading-point'

    import {listInformationCoin } from '@/api/admin/information-coin'
    export default {
        name: 'InformationTradingPoint',
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
                informationTradingPointList: [],
                typeOptions: [],statusOptions: [],isNotifiedOptions: [],
                // 关系表类型
                coinIdOptions :[],

                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    coinId:undefined,
                    type:undefined,
                    status:undefined,
                    isNotified:undefined,

                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {coinId:  [ {required: true, message: '币种ID不能为空', trigger: 'blur'} ],
                type:  [ {required: true, message: '类型：1-做多，2-做空不能为空', trigger: 'blur'} ],
                status:  [ {required: true, message: '状态：1-进行中，2-已失效，3-已撤单不能为空', trigger: 'blur'} ],
                isNotified:  [ {required: true, message: '是否已通知：1-未通知，2-已通知不能为空', trigger: 'blur'} ],
                entryPrice:  [ {required: true, message: '入场价格不能为空', trigger: 'blur'} ],
                takeProfits:  [ {required: true, message: '止盈不能为空', trigger: 'blur'} ],
                stopLoss:  [ {required: true, message: '止损不能为空', trigger: 'blur'} ],
                validDuration:  [ {required: true, message: '有效时长不能为空', trigger: 'blur'} ],
                validUntil:  [ {required: true, message: '到期时间不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getInformationCoinItems()
            this.getDicts('information_trading_point_type').then(response => {
                this.typeOptions = response.data
            })
            this.getDicts('information_trading_point_status').then(response => {
                this.statusOptions = response.data
            })
            this.getDicts('information_trading_point_is_notified').then(response => {
                this.isNotifiedOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listInformationTradingPoint(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.informationTradingPointList = response.data.list
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
                coinId: undefined,
                type: undefined,
                entryPrice: undefined,
                takeProfits: undefined,
                stopLoss: undefined,
                validDuration: undefined,
                validUntil: undefined,
                status: undefined,
                isNotified: undefined,
                notes: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            coinIdFormat(row) {
                return this.selectItemsLabel(this.coinIdOptions, row.coinId)
            },
            typeFormat(row) {
                return this.selectDictLabel(this.typeOptions, row.type)
            },
            statusFormat(row) {
                return this.selectDictLabel(this.statusOptions, row.status)
            },
            isNotifiedFormat(row) {
                return this.selectDictLabel(this.isNotifiedOptions, row.isNotified)
            },
            // 关系
            getInformationCoinItems() {
               this.getItems(listInformationCoin, undefined).then(res => {
                   this.coinIdOptions = this.setItems(res, 'id', 'symbol')
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
                this.title = '添加点位'
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
                getInformationTradingPoint(id).then(response => {
                    this.form = response.data
                  if (this.form.type !== undefined && this.form.type !== null) {
                    this.form.type = String(this.form.type)
                  }
                  if (this.form.status !== undefined && this.form.status !== null) {
                    this.form.status = String(this.form.status)
                  }
                  if (this.form.isNotified !== undefined && this.form.isNotified !== null) {
                    this.form.isNotified = String(this.form.isNotified)
                  }
                  if (this.form.coinId !== undefined && this.form.coinId !== null) {
                    this.form.coinId = String(this.form.coinId)
                  }
                  // 确保validDuration是数字类型
                  if (this.form.validDuration !== undefined && this.form.validDuration !== null) {
                    this.form.validDuration = Number(this.form.validDuration)
                  }
                    this.open = true
                    this.title = '修改点位'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                      // 确保validDuration是数字类型
                      if (this.form.validDuration !== undefined && this.form.validDuration !== null) {
                        this.form.validDuration = Number(this.form.validDuration)
                      }
                        if (this.form.id !== undefined) {
                            updateInformationTradingPoint(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addInformationTradingPoint(this.form).then(response => {
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
                      return delInformationTradingPoint( { 'ids': Ids })
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
