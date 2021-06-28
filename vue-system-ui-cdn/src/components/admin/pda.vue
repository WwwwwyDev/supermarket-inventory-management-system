<template>
	<!-- 卡片视图区 -->
	<el-card>
		<el-row :gutter="25">
			<el-col :span="7">
				<!-- 搜索添加 -->
				<el-input placeholder="请输入搜索的清单号" v-model.lazy="queryInfo.searchList" @change="getPurchaseDetailsList">
				</el-input>
			</el-col>
			<el-col :span="7">
				<el-input placeholder="请输入搜索的商品名" v-model.lazy="queryInfo.searchGoodsName"
					@change="getPurchaseDetailsList">
				</el-input>
			</el-col>
			<el-col :span="4">
				<el-button type="primary" @click="dialogAddVisible = true">添加采购明细</el-button>
			</el-col>
		</el-row>
		<!-- 列表 -->
		<el-table :data="purchaseDetailsList" border stripe v-loading="isLoading"
			element-loading-background="rgba(255, 255, 255, 0.5)" element-loading-text="加载中，请稍后..."
			element-loading-spinner="el-icon-loading">
			<el-table-column label="记录号" type="index" fixed="left"></el-table-column>
			<el-table-column label="明细编号" prop="PurchaseDetailsId"></el-table-column>
			<el-table-column label="清单编号" prop="PurchaseDetailsList"></el-table-column>
			<el-table-column label="商品编号" prop="PurchaseDetailsGoods"></el-table-column>
			<el-table-column label="商品名称" prop="PurchaseDetailsGoodsName"></el-table-column>
			<el-table-column label="数量" prop="PurchaseDetailsNumber"></el-table-column>
			<el-table-column label="总价(分)" prop="PurchaseDetailsPrice"></el-table-column>
			<el-table-column label="备注" prop="PurchaseDetailsRemarks"></el-table-column>
			<el-table-column label="操作" fixed="right">
				<template #default="scope">
					<!-- 修改 -->
					<el-button type="primary" icon="el-icon-edit" size="mini"
						@click="dialogEditOpen(scope.row.PurchaseDetailsId, scope.row.PurchaseDetailsList,  scope.row.PurchaseDetailsGoods, scope.row.PurchaseDetailsNumber, scope.row.PurchaseDetailsRemarks)">
					</el-button>
					<!-- 删除 -->
					<el-button type="danger" icon="el-icon-delete" size="mini"
						@click="deletePurchaseList(scope.row.PurchaseDetailsId)">
					</el-button>
				</template>
			</el-table-column>
		</el-table>
		<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
			:current-page="queryInfo.page" :page-sizes="[1, 2, 10, 100]" :page-size="queryInfo.limit"
			layout="total, sizes, prev, pager, next, jumper" :total="total"></el-pagination>
	</el-card>

	<el-dialog title="添加采购明细" v-model="dialogAddVisible" width="50%">
		<el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="100px">
			<el-form-item label="清单编号" prop="purchaseDetailsList">
				<el-input v-model="addForm.purchaseDetailsList"></el-input>
			</el-form-item>
			<el-form-item label="商品编号" prop="purchaseDetailsGoods">
				<el-input v-model="addForm.purchaseDetailsGoods"></el-input>
			</el-form-item>
			<el-form-item label="数量" prop="purchaseDetailsNumber">
				<el-input v-model="addForm.purchaseDetailsNumber"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="purchaseDetailsRemarks">
				<el-input v-model="addForm.purchaseDetailsRemarks"></el-input>
			</el-form-item>
		</el-form>
		<span slot="footer" class="dialog-footer" style="margin-left: 38%;">
			<el-button type="primary" @click="addPurchaseDetails()">确 定</el-button>
			<el-button @click="dialogAddVisible = false">取 消</el-button>
		</span>
	</el-dialog>

	<el-dialog title="编辑采购明细" v-model="dialogEditVisible" width="50%">
		<el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="100px">
			<el-form-item label="编号" prop="id">
				<el-input v-model="editForm.id" :disabled="true"></el-input>
			</el-form-item>
			<el-form-item label="清单编号" prop="purchaseDetailsList">
				<el-input v-model="editForm.purchaseDetailsList"></el-input>
			</el-form-item>
			<el-form-item label="商品编号" prop="purchaseDetailsGoods">
				<el-input v-model="editForm.purchaseDetailsGoods"></el-input>
			</el-form-item>
			<el-form-item label="数量" prop="purchaseDetailsNumber">
				<el-input v-model="editForm.purchaseDetailsNumber"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="purchaseDetailsRemarks">
				<el-input v-model="editForm.purchaseDetailsRemarks"></el-input>
			</el-form-item>
		</el-form>
		<span slot="footer" class="dialog-footer" style="margin-left: 38%;">
			<el-button type="primary" @click="updatePurchaseDetails()">确 定</el-button>
			<el-button @click="dialogEditVisible = false">取 消</el-button>
		</span>
	</el-dialog>
</template>

<script>
	export default {
		data() {
			return {
				isLoading: false,
				dialogAddVisible: false,
				dialogEditVisible: false,
				queryInfo: {
					searchList: '',
					page: 1,
					limit: 10,
					searchGoodsName: '',
				},
				purchaseDetailsList: [], // 供应商列表
				total: 0, // 最大数据记录
				addForm: {
					purchaseDetailsGoods: '',
					purchaseDetailsNumber: '',
					purchaseDetailsList: '',
					purchaseDetailsRemarks: ''
				},
				editForm: {
					id: '',
					purchaseDetailsGoods: '',
					purchaseDetailsNumber: '',
					purchaseDetailsList: '',
					purchaseDetailsRemarks: ''
				},
				addFormRules: {
					purchaseDetailsGoods: [{
						required: true,
						message: "请输入商品编号",
						trigger: "blur"
					}],
					purchaseDetailsNumber: [{
						required: true,
						message: "请输入数量",
						trigger: "blur"
					}],
					purchaseDetailsList: [{
						required: true,
						message: "请输入所属清单编号",
						trigger: "blur"
					}],
				},
				editFormRules: {
					purchaseDetailsGoods: [{
						required: true,
						message: "请输入商品编号",
						trigger: "blur"
					}],
					purchaseDetailsNumber: [{
						required: true,
						message: "请输入数量",
						trigger: "blur"
					}],
					purchaseDetailsList: [{
						required: true,
						message: "请输入所属清单编号",
						trigger: "blur"
					}],
				},
			}
		},
		created() {
			this.getPurchaseDetailsList();
		},
		methods: {
			async getPurchaseDetailsList() {
				this.isLoading = true;
				// 调用post请求
				const {
					data: res
				} = await this.$http.get("system/apis/purchaseDetails", {
					params: this.queryInfo
				});
				if (res.code != 20000) {
					this.$message.error("加载用户列表失败");
					this.isLoading = false;
				}
				this.purchaseDetailsList = res.data.value; // 将返回数据赋值
				this.total = res.data.total; // 总个数
				this.isLoading = false;
			},
			// 监听pageSize改变的事件
			handleSizeChange(newLimit) {
				this.queryInfo.limit = newLimit;
				this.getPurchaseDetailsList(); // 数据发生改变重新申请数据
			},
			// 监听pageNum改变的事件
			handleCurrentChange(newPage) {
				this.queryInfo.page = newPage;
				this.getPurchaseDetailsList(); // 数据发生改变重新申请数据
			},
			async deletePurchaseDetails(purchaseDetailsId) {
				// 弹框
				const confirmResult = await this.$confirm('此操作将永久删除该清单, 是否继续?', '提示', {
					confirmButtonText: '确定',
					cancelButtonText: '取消',
					type: 'warning'
				}).catch(err => err)
				// 成功删除为confirm 取消为 cancel
				if (confirmResult != 'confirm') {
					return this.$message.info("已取消删除");
				}
				const {
					data: res
				} = await this.$http.delete("system/apis/purchaseDetails?id=" + purchaseDetailsId);
				if (res.code != 20000) {
					return this.$message.error("删除失败");
				}
				this.$message.success("删除成功");
				this.getPurchaseDetailsList();
			},
			// 添加明细
			addPurchaseDetails() {
				this.$refs.addFormRef.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.post("/system/apis/purchaseDetails", this.addForm);
					if (res.code == 20000) {
						this.getPurchaseDetailsList();
						this.dialogAddVisible = false;
						return this.$message.success("添加成功");
					}
					this.$message.error("添加失败");
				})
			},
			// 编辑
			updatePurchaseDetails() {
				this.$refs.editFormRef.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.put("/system/apis/purchaseDetails", this.editForm);
					if (res.code == 20000) {
						this.getPurchaseDetailsList();
						this.dialogEditVisible = false;
						return this.$message.success("编辑成功");
					}
					this.$message.error("编辑失败");
				})
			},
			dialogEditOpen(id, purchaseDetailsList, purchaseDetailsGoods, purchaseDetailsNumber, purchaseDetailsRemarks) {
				this.editForm.id = String(id);
				this.editForm.purchaseDetailsGoods = String(purchaseDetailsGoods);
				this.editForm.purchaseDetailsNumber = String(purchaseDetailsNumber);
				this.editForm.purchaseDetailsList = String(purchaseDetailsList);
				this.editForm.purchaseDetailsRemarks = String(purchaseDetailsRemarks);
				this.dialogEditVisible = true;
			},
		},
	}
</script>

<style>
</style>
