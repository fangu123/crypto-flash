
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="链" prop="blockchainId"><el-select v-model="queryParams.blockchainId"
                                           placeholder="请选择" clearable size="small" >
                                    <el-option
                                            v-for="dict in blockchainIdOptions"
                                            :key="dict.key"
                                            :label="dict.value"
                                            :value="dict.key"
                                    />
                                </el-select>
                            </el-form-item>
                        <el-form-item label="币种标识" prop="symbol"><el-input v-model="queryParams.symbol" placeholder="请输入币种标识" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="状态" prop="status"><el-select v-model="queryParams.status"
                                               placeholder="状态:启用/禁用" clearable size="small">
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
                                v-permisaction="['admin:informationCoin:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:informationCoin:edit']"
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
                                v-permisaction="['admin:informationCoin:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>-->

                <el-table v-loading="loading" :data="informationCoinList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/>
                  <el-table-column label="链" align="center" prop="blockchainId" :formatter="blockchainIdFormat" width="100">
                                <template slot-scope="scope">
                                    {{ blockchainIdFormat(scope.row) }}
                                </template>
                            </el-table-column><el-table-column label="名称" align="center" prop="name"
                                                 :show-overflow-tooltip="true"/><el-table-column label="标识" align="center" prop="symbol"
                                                 :show-overflow-tooltip="true"/><el-table-column label="币Logo" align="center" prop="logoUrl" width="300">
                  <template slot-scope="scope">
                    <el-image
                      v-if="scope.row.logoUrl"
                      :src="scope.row.logoUrl"
                      style="width: 40px; height: 40px"
                      lazy>
                      <div slot="error" class="image-slot">
                        <i class="el-icon-picture-outline"></i>
                      </div>
                    </el-image>
                  </template>
                </el-table-column>

                  <!--                  <el-table-column label="当前价格" align="center" prop="price"
                                                                   :show-overflow-tooltip="true"/><el-table-column label="24小时价格变化" align="center" prop="priceChange24h"
                                                                   :show-overflow-tooltip="true"/><el-table-column label="市值" align="center" prop="marketCap"
                                                                   :show-overflow-tooltip="true"/><el-table-column label="24小时交易量" align="center" prop="volume24h"
                                                                   :show-overflow-tooltip="true"/><el-table-column label="流通量" align="center" prop="circulatingSupply"
                                                                   :show-overflow-tooltip="true"/><el-table-column label="总量" align="center" prop="totalSupply"
                                                                   :show-overflow-tooltip="true"/><el-table-column label="最大量" align="center" prop="maxSupply"
                                                                   :show-overflow-tooltip="true"/><el-table-column label="排名" align="center" prop="rank"
                                                                   :show-overflow-tooltip="true"/>-->
                  <el-table-column label="状态" align="center" prop="status"
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
                             v-permisaction="['admin:informationCoin:edit']"
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
                              v-permisaction="['admin:informationCoin:remove']"
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

                                    <el-form-item label="链" prop="blockchainId">
                                        <el-select v-model="form.blockchainId"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in blockchainIdOptions"
                                                        :key="dict.key"
                                                        :label="dict.value"
                                                        :value="dict.key"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="币种名称" prop="name">
                                        <el-input v-model="form.name" placeholder="币种名称"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="币种标识" prop="symbol">
                                        <el-input v-model="form.symbol" placeholder="币种标识"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="币Logo" prop="logoUrl">
                                        <el-input v-model="form.logoUrl" placeholder="币Logo URL"
                                                      />
                                    </el-form-item>
<!--                                    <el-form-item label="排名" prop="rank">
                                        <el-input v-model="form.rank" placeholder="排名"
                                                      />
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
    import {addInformationCoin, delInformationCoin, getInformationCoin, listInformationCoin, updateInformationCoin} from '@/api/admin/information-coin'

    import {listInformationBlockchain } from '@/api/admin/information-blockchain'
    export default {
        name: 'InformationCoin',
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
                informationCoinList: [],
                statusOptions: [],
                // 关系表类型
                blockchainIdOptions :[],

                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    blockchainId:undefined,
                    symbol:undefined,
                    status:undefined,

                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {blockchainId:  [ {required: true, message: '所属链不能为空', trigger: 'blur'} ],
                symbol:  [ {required: true, message: '币种标识不能为空', trigger: 'blur'} ],
                status:  [ {required: true, message: '状态:1-启用，2-禁用不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getInformationBlockchainItems()
            this.getDicts('information_coin_status').then(response => {
                this.statusOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listInformationCoin(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.informationCoinList = response.data.list
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
                blockchainId: undefined,
                name: undefined,
                symbol: undefined,
                logoUrl: undefined,
                rank: undefined,
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
            blockchainIdFormat(row) {
                return this.selectItemsLabel(this.blockchainIdOptions, row.blockchainId)
            },
            statusFormat(row) {
                return this.selectDictLabel(this.statusOptions, row.status)
            },
            // 关系
            getInformationBlockchainItems() {
               this.getItems(listInformationBlockchain, undefined).then(res => {
                   this.blockchainIdOptions = this.setItems(res, 'id', 'symbol')
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
                this.title = '添加币种'
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
                getInformationCoin(id).then(response => {
                    this.form = response.data
                  if (this.form.status !== undefined && this.form.status !== null) {
                    this.form.status = String(this.form.status)
                  }
                    this.open = true
                    this.title = '修改币种'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateInformationCoin(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addInformationCoin(this.form).then(response => {
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
                      return delInformationCoin( { 'ids': Ids })
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
