<template>
	<!-- 卡片视图区 -->
	<el-card>
		<el-row :gutter="25">
			<el-col :span="7">
				<!-- 搜索添加 -->
				<el-input placeholder="请输入搜索负责采购员工的姓名" v-model.lazy="queryInfo.searchName"
					@change="getPurchaseListList">
				</el-input>
			</el-col>

			<el-col :span="4">
				<el-button type="primary" @click="dialogAddVisible = true">添加采购清单</el-button>
			</el-col>
		</el-row>
		<!-- 列表 -->
		<el-table :data="purchaseListList" border stripe v-loading="isLoading"
			element-loading-background="rgba(255, 255, 255, 0.5)" element-loading-text="加载中，请稍后..."
			element-loading-spinner="el-icon-loading">
			<el-table-column label="记录号" type="index" fixed="left"></el-table-column>
			<el-table-column label="编号" prop="PurchaseListId"></el-table-column>
			<el-table-column label="员工编号" prop="PurchaseListStaff"></el-table-column>
			<el-table-column label="员工姓名" prop="PurchaseListStaffName"></el-table-column>
			<el-table-column label="采购总数量" prop="PurchaseListNumber"></el-table-column>
			<el-table-column label="采购总价(分)" prop="PurchaseListPrice"></el-table-column>
			<el-table-column label="创建时间" prop="PurchaseListTime"></el-table-column>
			<el-table-column label="备注" prop="PurchaseListRemarks"></el-table-column>
			<el-table-column label="操作" fixed="right">
				<template #default="scope">
					<!-- 修改 -->
					<el-button type="primary" icon="el-icon-edit" size="mini"
						@click="dialogEditOpen(scope.row.PurchaseListId, scope.row.PurchaseListStaff,  scope.row.PurchaseListRemarks)">
					</el-button>
					<!-- 删除 -->
					<el-button type="danger" icon="el-icon-delete" size="mini"
						@click="deletePurchaseList(scope.row.PurchaseListId)">
					</el-button>
				</template>
			</el-table-column>
		</el-table>
		<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
			:current-page="queryInfo.page" :page-sizes="[1, 2, 10, 100]" :page-size="queryInfo.limit"
			layout="total, sizes, prev, pager, next, jumper" :total="total"></el-pagination>
	</el-card>

	<el-dialog title="添加采购清单" v-model="dialogAddVisible" width="50%">
		<el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="100px">
			<el-form-item label="员工编号" prop="purchaseListStaff">
				<el-input v-model="addForm.purchaseListStaff"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="purchaseListRemarks">
				<el-input v-model="addForm.purchaseListRemarks"></el-input>
			</el-form-item>
		</el-form>
		<span slot="footer" class="dialog-footer" style="margin-left: 38%;">
			<el-button type="primary" @click="addPurchaseList()">确 定</el-button>
			<el-button @click="dialogAddVisible = false">取 消</el-button>
		</span>
	</el-dialog>

	<!-- 编辑对话框 -->
	<el-dialog title="编辑商品" v-model="dialogEditVisible" width="50%">
		<el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="100px">
			<el-form-item label="编号" prop="id">
				<el-input v-model="editForm.id" :disabled="true"></el-input>
			</el-form-item>
			<el-form-item label="员工编号" prop="purchaseListStaff">
				<el-input v-model="editForm.purchaseListStaff"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="purchaseListRemarks">
				<el-input v-model="editForm.purchaseListRemarks"></el-input>
			</el-form-item>
		</el-form>
		<span slot="footer" class="dialog-footer" style="margin-left: 38%;">
			<el-button type="primary" @click="updatePurchaseList">确 定</el-button>
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
					searchName: '',
					page: 1,
					limit: 10
				},
				purchaseListList: [], // 供应商列表
				total: 0, // 最大数据记录
				addForm: {
					purchaseListStaff: '',
					purchaseListRemarks: '',
				},
				editForm: {
					id: '',
					purchaseListStaff: '',
					purchaseListRemarks: '',
				},
				addFormRules: {
					purchaseListStaff: [{
						required: true,
						message: "请输入员工编号",
						trigger: "blur"
					}],
					purchaseListRemarks: [{
						required: true,
						message: "请输入备注",
						trigger: "blur"
					}],
				},
				editFormRules: {
					purchaseListStaff: [{
						required: true,
						message: "请输入员工编号",
						trigger: "blur"
					}],
					purchaseListRemarks: [{
						required: true,
						message: "请输入备注",
						trigger: "blur"
					}],
				},
			}
		},
		created() {
			this.getPurchaseListList();
		},
		methods: {
			async getPurchaseListList() {
				this.isLoading = true;
				// 调用post请求
				const {
					data: res
				} = await this.$http.get("system/apis/purchaseList", {
					params: this.queryInfo
				});
				if (res.code != 20000) {
					this.$message.error("加载用户列表失败");
					this.isLoading = false;
				}
				this.purchaseListList = res.data.value; // 将返回数据赋值
				this.total = res.data.total; // 总个数
				this.isLoading = false;
			},
			// 监听pageSize改变的事件
			handleSizeChange(newLimit) {
				this.queryInfo.limit = newLimit;
				this.getPurchaseListList(); // 数据发生改变重新申请数据
			},
			// 监听pageNum改变的事件
			handleCurrentChange(newPage) {
				this.queryInfo.page = newPage;
				this.getPurchaseListList(); // 数据发生改变重新申请数据
			},
			async deletePurchaseList(purchaseList) {
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
				} = await this.$http.delete("system/apis/purchaseList?id=" + purchaseList);
				if (res.code != 20000) {
					return this.$message.error("删除失败");
				}
				this.$message.success("删除成功");
				this.getPurchaseListList();
			},
			// 添加商品
			addPurchaseList() {
				this.$refs.addFormRef.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.post("/system/apis/purchaseList", this.addForm);
					if (res.code == 20000) {
						this.getPurchaseListList();
						this.dialogAddVisible = false;
						return this.$message.success("添加成功");
					}
					this.$message.error("添加失败");
				})
			},
			// 编辑
			updatePurchaseList() {
				this.$refs.editFormRef.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.put("/system/apis/purchaseList", this.editForm);
					if (res.code == 20000) {
						this.getPurchaseListList();
						this.dialogEditVisible = false;
						return this.$message.success("编辑成功");
					}
					this.$message.error("编辑失败");
				})
			},
			dialogEditOpen(id, purchaseListStaff, purchaseListRemarks) {
				this.editForm.id = String(id);
				this.editForm.purchaseListStaff = String(purchaseListStaff);
				this.editForm.purchaseListRemarks = String(purchaseListRemarks);
				this.dialogEditVisible = true;
			},
		},
	}
</script>

<style>
</style>
