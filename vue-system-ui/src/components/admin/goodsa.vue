<template>
	<!-- 卡片视图区 -->
	<el-card>
		<el-row :gutter="25">
			<el-col :span="7">
				<!-- 搜索添加 -->
				<el-input placeholder="请输入搜索商品的名字" v-model.lazy="queryInfo.searchName" @change="getGoodsList">
				</el-input>
			</el-col>

			<el-col :span="4">
				<el-button type="primary" @click="dialogAddVisible = true">添加供应商</el-button>
			</el-col>
		</el-row>
		<!-- 列表 -->
		<el-table :data="goodsList" border stripe>
			<el-table-column label="记录号" type="index" fixed="left"></el-table-column>
			<el-table-column label="编号" prop="GoodsId"></el-table-column>
			<el-table-column label="名称" prop="GoodsName"></el-table-column>
			<el-table-column label="单价(分)" prop="GoodsPrice"></el-table-column>
			<el-table-column label="供应商编号" prop="GoodsSupplier"></el-table-column>
			<el-table-column label="供应商名字" prop="GoodsSupplierName"></el-table-column>
			<el-table-column label="简介" prop="GoodsSynopsis"></el-table-column>
			<el-table-column label="备注" prop="GoodsRemarks"></el-table-column>
			<el-table-column label="操作" fixed="right">
				<template #default="scope">
					<!-- 修改 -->
					<el-button type="primary" icon="el-icon-edit" size="mini"
						@click="dialogEditOpen(scope.row.GoodsId, scope.row.GoodsName,  scope.row.GoodsPrice, scope.row.GoodsSupplier,  scope.row.GoodsSynopsis, scope.row.GoodsRemarks)">
					</el-button>
					<!-- 删除 -->
					<el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteGoods(scope.row.GoodsId)">
					</el-button>
				</template>
			</el-table-column>
		</el-table>
		<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
			:current-page="queryInfo.page" :page-sizes="[1, 2, 10, 100]" :page-size="queryInfo.limit"
			layout="total, sizes, prev, pager, next, jumper" :total="total"></el-pagination>
	</el-card>

	<el-dialog title="添加商品" v-model="dialogAddVisible" width="50%">
		<el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="100px">
			<el-form-item label="名称" prop="goodsName">
				<el-input v-model="addForm.goodsName"></el-input>
			</el-form-item>
			<el-form-item label="单价(分)" prop="goodsPrice">
				<el-input v-model="addForm.goodsPrice"></el-input>
			</el-form-item>
			<el-form-item label="供应商编号" prop="goodsSupplier">
				<el-input v-model="addForm.goodsSupplier"></el-input>
			</el-form-item>
			<el-form-item label="简介" prop="goodsSynopsis">
				<el-input v-model="addForm.goodsSynopsis"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="goodsRemarks">
				<el-input v-model="addForm.goodsRemarks"></el-input>
			</el-form-item>
		</el-form>
		<span slot="footer" class="dialog-footer" style="margin-left: 38%;">
			<el-button type="primary" @click="addGoods()">确 定</el-button>
			<el-button @click="dialogAddVisible = false">取 消</el-button>
		</span>
	</el-dialog>

	<!-- 编辑对话框 -->
	<el-dialog title="编辑商品" v-model="dialogEditVisible" width="50%">
		<el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="100px">
			<el-form-item label="编号" prop="id">
				<el-input v-model="editForm.id" :disabled="true"></el-input>
			</el-form-item>
			<el-form-item label="名称" prop="goodsName">
				<el-input v-model="editForm.goodsName"></el-input>
			</el-form-item>
			<el-form-item label="单价(分)" prop="goodsPrice">
				<el-input v-model="editForm.goodsPrice"></el-input>
			</el-form-item>
			<el-form-item label="供应商编号" prop="goodsSupplier">
				<el-input v-model="editForm.goodsSupplier"></el-input>
			</el-form-item>
			<el-form-item label="简介" prop="goodsSynopsis">
				<el-input v-model="editForm.goodsSynopsis"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="goodsRemarks">
				<el-input v-model="editForm.goodsRemarks"></el-input>
			</el-form-item>
		</el-form>
		<span slot="footer" class="dialog-footer" style="margin-left: 38%;">
			<el-button type="primary" @click="updateGoods">确 定</el-button>
			<el-button @click="dialogEditVisible = false">取 消</el-button>
		</span>
	</el-dialog>
</template>

<script>
	export default {
		data() {
			return {
				// 请求数据
				dialogAddVisible: false,
				dialogEditVisible: false,
				queryInfo: {
					searchName: '',
					page: 1,
					limit: 10
				},
				goodsList: [], // 供应商列表
				total: 0, // 最大数据记录
				addForm: {
					goodsName: '',
					goodsPrice: '',
					goodsSupplier: '',
					goodsSynopsis: '',
					goodsRemarks: '',
				},
				editForm: {
					id: '',
					goodsName: '',
					goodsPrice: '',
					goodsSupplier: '',
					goodsSynopsis: '',
					goodsRemarks: '',
				},
				addFormRules: {
					goodsName: [{
						required: true,
						message: "请输入商品名称",
						trigger: "blur"
					}],
					goodsPrice: [{
						required: true,
						message: "请输入商品单价",
						trigger: "blur"
					}],
					goodsSupplier: [{
						required: true,
						message: "请输入供应商编号",
						trigger: "blur"
					}],
					goodsSynopsis: [{
						required: true,
						message: "请输入简介",
						trigger: "blur"
					}],
					goodsRemarks: [{
						required: true,
						message: "请输入备注",
						trigger: "blur"
					}],
				},
				editFormRules: {
					goodsName: [{
						required: true,
						message: "请输入商品名称",
						trigger: "blur"
					}],
					goodsPrice: [{
						required: true,
						message: "请输入商品单价",
						trigger: "blur"
					}],
					goodsSupplier: [{
						required: true,
						message: "请输入供应商编号",
						trigger: "blur"
					}],
					goodsSynopsis: [{
						required: true,
						message: "请输入简介",
						trigger: "blur"
					}],
					goodsRemarks: [{
						required: true,
						message: "请输入备注",
						trigger: "blur"
					}],
				},
			}
		},
		created() {
			this.getGoodsList();
		},
		methods: {
			async getGoodsList() {
				// 调用post请求
				const {
					data: res
				} = await this.$http.get("system/apis/goods", {
					params: this.queryInfo
				});
				if (res.code != 20000) {
					this.$message.error("加载用户列表失败");
				}
				this.goodsList = res.data.value; // 将返回数据赋值
				this.total = res.data.total; // 总个数
			},
			// 监听pageSize改变的事件
			handleSizeChange(newLimit) {
				this.queryInfo.limit = newLimit;
				this.getGoodsList(); // 数据发生改变重新申请数据
			},
			// 监听pageNum改变的事件
			handleCurrentChange(newPage) {
				this.queryInfo.page = newPage;
				this.getGoodsList(); // 数据发生改变重新申请数据
			},
			async deleteGoods(GoodsId) {
				// 弹框
				const confirmResult = await this.$confirm('此操作将永久删除该商品, 是否继续?', '提示', {
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
				} = await this.$http.delete("system/apis/goods?id=" + GoodsId);
				if (res.code != 20000) {
					return this.$message.error("删除失败");
				}
				this.$message.success("删除成功");
				this.getGoodsList();
			},
			// 添加商品
			addGoods() {
				this.$refs.addFormRef.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.post("/system/apis/goods", this.addForm);
					if (res.code == 20000) {
						this.getGoodsList();
						this.dialogAddVisible = false;
						return this.$message.success("添加成功");
					}
					this.$message.error("添加失败");
				})
			},
			// 编辑
			updateGoods() {
				this.$refs.editFormRef.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.put("/system/apis/goods", this.editForm);
					if (res.code == 20000) {
						this.getGoodsList();
						this.dialogEditVisible = false;
						return this.$message.success("编辑成功");
					}
					this.$message.error("编辑失败");
				})
			},
			dialogEditOpen(id, goodsName, goodsPrice, goodsSupplier, goodsSynopsis, goodsRemarks) {
				this.editForm.id = String(id);
				this.editForm.goodsName = goodsName;
				this.editForm.goodsPrice = goodsPrice;
				this.editForm.goodsSupplier = goodsSupplier;
				this.editForm.goodsSynopsis = goodsSynopsis;
				this.editForm.goodsRemarks = goodsRemarks;
				this.dialogEditVisible = true;
			},
		},
	}
</script>

<style>
</style>
