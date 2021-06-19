<template>
	<!-- 卡片视图区 -->
	<el-card>
		<el-row :gutter="25">
			<el-col :span="7">
				<!-- 搜索添加 -->
				<el-input placeholder="请输入搜索供应商的名字" v-model.lazy="queryInfo.searchName" @change="getSupplierList">
				</el-input>
			</el-col>

			<el-col :span="4">
				<el-button type="primary" @click="dialogAddVisible = true">添加供应商</el-button>
			</el-col>
		</el-row>
		<!-- 列表 -->
		<el-table :data="supplierList" border stripe v-loading="isLoading"
			element-loading-background="rgba(255, 255, 255, .5)" element-loading-text="加载中，请稍后..."
			element-loading-spinner="el-icon-loading">
			<el-table-column label="记录号" type="index" fixed="left"></el-table-column>
			<el-table-column label="编号" prop="SupplierId"></el-table-column>
			<el-table-column label="名称" prop="SupplierName"></el-table-column>
			<el-table-column label="简称" prop="SupplierShortname"></el-table-column>
			<el-table-column label="地址" prop="SupplierAddress"></el-table-column>
			<el-table-column label="公司电话" prop="SupplierCtelephone"></el-table-column>
			<el-table-column label="邮箱" prop="SupplierEmail"></el-table-column>
			<el-table-column label="联系人" prop="SupplierLiaisonman"></el-table-column>
			<el-table-column label="联系人电话" prop="SupplierLtelephone"></el-table-column>
			<el-table-column label="备注" prop="SupplierRemarks"></el-table-column>
			<el-table-column label="操作" fixed="right">
				<template #default="scope">
					<!-- 修改 -->
					<el-button type="primary" icon="el-icon-edit" size="mini"
						@click="dialogEditOpen(scope.row.SupplierId, scope.row.SupplierName,  scope.row.SupplierShortname, scope.row.SupplierAddress, scope.row.SupplierCtelephone, scope.row.SupplierEmail, scope.row.SupplierLiaisonman, scope.row.SupplierLtelephone, scope.row.SupplierRemarks)">
					</el-button>
					<!-- 删除 -->
					<el-button type="danger" icon="el-icon-delete" size="mini"
						@click="deleteSupplier(scope.row.SupplierId)"></el-button>
				</template>
			</el-table-column>
		</el-table>
		<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
			:current-page="queryInfo.page" :page-sizes="[1, 2, 10, 100]" :page-size="queryInfo.limit"
			layout="total, sizes, prev, pager, next, jumper" :total="total"></el-pagination>
	</el-card>

	<el-dialog title="添加供应商" v-model="dialogAddVisible" width="50%">
		<el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="100px">
			<el-form-item label="名称" prop="supplierName">
				<el-input v-model="addForm.supplierName"></el-input>
			</el-form-item>
			<el-form-item label="简称" prop="supplierShortname">
				<el-input v-model="addForm.supplierShortname"></el-input>
			</el-form-item>
			<el-form-item label="地址" prop="supplierAddress">
				<el-input v-model="addForm.supplierAddress"></el-input>
			</el-form-item>
			<el-form-item label="公司电话" prop="supplierCtelephone">
				<el-input v-model="addForm.supplierCtelephone"></el-input>
			</el-form-item>
			<el-form-item label="邮箱" prop="supplierEmail">
				<el-input v-model="addForm.supplierEmail"></el-input>
			</el-form-item>
			<el-form-item label="联系人" prop="supplierLiaisonman">
				<el-input v-model="addForm.supplierLiaisonman"></el-input>
			</el-form-item>
			<el-form-item label="联系人电话" prop="supplierLtelephone">
				<el-input v-model="addForm.supplierLtelephone"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="supplierRemarks">
				<el-input v-model="addForm.supplierRemarks"></el-input>
			</el-form-item>
		</el-form>
		<span slot="footer" class="dialog-footer" style="margin-left: 38%;">
			<el-button type="primary" @click="addSupplier">确 定</el-button>
			<el-button @click="dialogAddVisible = false">取 消</el-button>
		</span>
	</el-dialog>

	<!-- 编辑对话框 -->
	<el-dialog title="编辑供应商" v-model="dialogEditVisible" width="50%">
		<el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="100px">
			<el-form-item label="编号" prop="id">
				<el-input v-model="editForm.id" :disabled="true"></el-input>
			</el-form-item>
			<el-form-item label="名称" prop="supplierName">
				<el-input v-model="editForm.supplierName"></el-input>
			</el-form-item>
			<el-form-item label="简称" prop="supplierShortname">
				<el-input v-model="editForm.supplierShortname"></el-input>
			</el-form-item>
			<el-form-item label="地址" prop="supplierAddress">
				<el-input v-model="editForm.supplierAddress"></el-input>
			</el-form-item>
			<el-form-item label="公司电话" prop="supplierCtelephone">
				<el-input v-model="editForm.supplierCtelephone"></el-input>
			</el-form-item>
			<el-form-item label="邮箱" prop="supplierEmail">
				<el-input v-model="editForm.supplierEmail"></el-input>
			</el-form-item>
			<el-form-item label="联系人" prop="supplierLiaisonman">
				<el-input v-model="editForm.supplierLiaisonman"></el-input>
			</el-form-item>
			<el-form-item label="联系人电话" prop="supplierLtelephone">
				<el-input v-model="editForm.supplierLtelephone"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="supplierRemarks">
				<el-input v-model="editForm.supplierRemarks"></el-input>
			</el-form-item>
		</el-form>
		<span slot="footer" class="dialog-footer" style="margin-left: 38%;">
			<el-button type="primary" @click="updateSupplier">确 定</el-button>
			<el-button @click="dialogEditVisible = false">取 消</el-button>
		</span>
	</el-dialog>
</template>

<script>
	export default {
		data() {
			var checkEmail = (rule, value, callback) => {
				const regEmail =
					/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
				if (regEmail.test(value)) {
					return callback();
				}
				callback(new Error("请输入合法的邮箱"));
			};
			return {
				isLoading:false,
				dialogAddVisible: false,
				dialogEditVisible: false,
				queryInfo: {
					searchName: '',
					page: 1,
					limit: 10
				},
				supplierList: [], // 供应商列表
				total: 0, // 最大数据记录
				addForm: {
					supplierName: '',
					supplierShortname: '',
					supplierAddress: '',
					supplierCtelephone: '',
					supplierEmail: '',
					supplierLiaisonman: '',
					supplierLtelephone: '',
					supplierRemarks: '',
				},
				editForm: {
					id: '',
					supplierName: '',
					supplierShortname: '',
					supplierAddress: '',
					supplierCtelephone: '',
					supplierEmail: '',
					supplierLiaisonman: '',
					supplierLtelephone: '',
					supplierRemarks: '',
				},
				addFormRules: {
					supplierName: [{
						required: true,
						message: "请输入供应商名称",
						trigger: "blur"
					}],
					supplierShortname: [{
						required: true,
						message: "请输入供应商简称",
						trigger: "blur"
					}],
					supplierAddress: [{
						required: true,
						message: "请输入供应商地址",
						trigger: "blur"
					}],
					supplierCtelephone: [{
						required: true,
						message: "请输入供应商公司电话",
						trigger: "blur"
					}],
					supplierLtelephone: [{
						required: true,
						message: "请输入供应商联系人电话",
						trigger: "blur"
					}],
					supplierEmail: [{
						required: true,
						trigger: "blur",
						validator: checkEmail,
					}],
					supplierLiaisonman: [{
						required: true,
						message: "请输入供应商联系人",
						trigger: "blur"
					}],
				},
				editFormRules: {
					supplierName: [{
						required: true,
						message: "请输入供应商名称",
						trigger: "blur"
					}],
					supplierShortname: [{
						required: true,
						message: "请输入供应商简称",
						trigger: "blur"
					}],
					supplierAddress: [{
						required: true,
						message: "请输入供应商地址",
						trigger: "blur"
					}],
					supplierCtelephone: [{
						required: true,
						message: "请输入供应商公司电话",
						trigger: "blur"
					}],
					supplierLtelephone: [{
						required: true,
						message: "请输入供应商联系人电话",
						trigger: "blur"
					}],
					supplierEmail: [{
						required: true,
						trigger: "blur",
						validator: checkEmail,
					}],
					supplierLiaisonman: [{
						required: true,
						message: "请输入供应商联系人",
						trigger: "blur"
					}],
				},
			}
		},
		created() {
			this.getSupplierList();
		},
		methods: {
			async getSupplierList() {
				this.isLoading = true;
				// 调用post请求
				const {
					data: res
				} = await this.$http.get("system/apis/supplier", {
					params: this.queryInfo
				});
				if (res.code != 20000) {
					this.$message.error("加载供应商列表失败");
					this.isLoading = false;
				}
				this.supplierList = res.data.value; // 将返回数据赋值
				this.total = res.data.total; // 总个数
				this.isLoading = false;
			},
			// 监听pageSize改变的事件
			handleSizeChange(newLimit) {
				this.queryInfo.limit = newLimit;
				this.getSupplierList(); // 数据发生改变重新申请数据
			},
			// 监听pageNum改变的事件
			handleCurrentChange(newPage) {
				this.queryInfo.page = newPage;
				this.getSupplierList(); // 数据发生改变重新申请数据
			},
			async deleteSupplier(SupplierId) {
				// 弹框
				const confirmResult = await this.$confirm('此操作将永久删除该供应商, 是否继续?', '提示', {
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
				} = await this.$http.delete("system/apis/supplier?id=" + SupplierId);
				if (res.code != 20000) {
					return this.$message.error("删除失败");
				}
				this.$message.success("删除成功");
				this.getSupplierList();
			},
			// 添加员工
			addSupplier() {
				this.$refs.addFormRef.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.post("/system/apis/supplier", this.addForm);
					if (res.code == 20000) {
						this.getSupplierList();
						this.dialogAddVisible = false;
						return this.$message.success("添加成功");
					}
					this.$message.error("添加失败");
				})
			},
			// 编辑供应商
			updateSupplier() {
				this.$refs.editFormRef.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.put("/system/apis/supplier", this.editForm);
					if (res.code == 20000) {
						this.getSupplierList();
						this.dialogEditVisible = false;
						return this.$message.success("编辑成功");
					}
					this.$message.error("编辑失败");
				})
			},
			dialogEditOpen(id, supplierName,  supplierShortname, supplierAddress, supplierCtelephone, supplierEmail, supplierLiaisonman, supplierLtelephone, supplierRemarks) {
				this.editForm.id = String(id);
				this.editForm.supplierName = String(supplierName);
				this.editForm.supplierShortname = String(supplierShortname);
				this.editForm.supplierAddress = String(supplierAddress);
				this.editForm.supplierCtelephone = String(supplierCtelephone);
				this.editForm.supplierEmail = String(supplierEmail);
				this.editForm.supplierLiaisonman = String(supplierLiaisonman);
				this.editForm.supplierLtelephone = String(supplierLtelephone);
				this.editForm.supplierRemarks = String(supplierRemarks);
				this.dialogEditVisible = true;
			},
		},
	}
</script>

<style>
</style>
