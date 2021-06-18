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
		<!-- 用户列表 -->
		<el-table :data="supplierList" border stripe>
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
					<!-- <el-button type="primary" icon="el-icon-edit" size="mini"
						@click="dialogEditOpen(scope.row.StaffId, scope.row.StaffName,  scope.row.StaffPassword, scope.row.StaffLevel, scope.row.StaffTelephone, scope.row.StaffSalary, scope.row.StaffRemarks)">
					</el-button> -->
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
	<!-- 	<el-dialog title="编辑员工" v-model="dialogEditVisible" width="50%">
		<el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="70px">
			<el-form-item label="编号" prop="id">
				<el-input v-model="editForm.id" :disabled="true"></el-input>
			</el-form-item>
			<el-form-item label="姓名" prop="staffName">
				<el-input v-model="editForm.staffName"></el-input>
			</el-form-item>
			<el-form-item label="密码" prop="staffPassword">
				<el-input v-model="editForm.staffPassword"></el-input>
			</el-form-item>
			<el-form-item label="等级" prop="staffLevel">
				<el-input v-model="editForm.staffLevel" :disabled="true"></el-input>
			</el-form-item>
			<el-form-item label="电话" prop="staffTelephone">
				<el-input v-model="editForm.staffTelephone"></el-input>
			</el-form-item>
			<el-form-item label="薪水(分)" prop="staffSalary">
				<el-input v-model="editForm.staffSalary"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="staffRemarks">
				<el-input v-model="editForm.staffRemarks"></el-input>
			</el-form-item>
		</el-form>
		<span slot="footer" class="dialog-footer" style="margin-left: 38%;">
			<el-button type="primary" @click="updateStaff">确 定</el-button>
			<el-button @click="dialogEditVisible = false">取 消</el-button>
		</span>
	</el-dialog> -->
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
				// 请求数据
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
					staffName: '',
					staffPassword: '',
					staffLevel: '1',
					staffTelephone: '',
					staffSalary: '',
					staffRemarks: '',
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
						message: "请输入供应商邮箱",
						validator: checkEmail,
					}],
					supplierLiaisonman: [{
						required: true,
						message: "请输入供应商联系人",
						trigger: "blur"
					}],
				},
				editFormRules: {
					staffName: [{
							required: true,
							message: "请输入用户名",
							trigger: "blur"
						},
						{
							min: 5,
							max: 12,
							message: "长度在 5 到 12 个字符",
							trigger: "blur"
						}
					],
					staffPassword: [{
							required: true,
							message: "请输入密码",
							trigger: "blur"
						},
						{
							min: 5,
							max: 12,
							message: "长度在 5 到 12 个字符",
							trigger: "blur"
						}
					],
					staffLevel: [{
						required: true,
						trigger: "blur"
					}],
					staffSalary: [{
						required: true,
						message: "请输入薪水",
						trigger: "blur"
					}],
					staffTelephone: [{
						required: true,
						message: "请输入电话",
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
				// 调用post请求
				console.log(this.queryInfo)
				const {
					data: res
				} = await this.$http.get("system/apis/supplier", {
					params: this.queryInfo
				});
				console.log(res.data.value)
				if (res.code != 20000) {
					this.$message.error("加载用户列表失败");
				}
				this.supplierList = res.data.value; // 将返回数据赋值
				this.total = res.data.total; // 总个数
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
			// 编辑员工
			updateStaff() {
				this.$refs.editFormRef.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.put("/system/apis/staff", this.editForm);
					if (res.code == 20000) {
						this.getStaffList();
						this.dialogEditVisible = false;
						return this.$message.success("编辑成功");
					}
					this.$message.error("编辑失败");
				})
			},
			dialogEditOpen(id, staffName, staffPassword, staffLevel, staffTelephone, staffSalary, staffRemarks) {
				this.editForm.id = String(id);
				this.editForm.staffName = staffName;
				this.editForm.staffPassword = staffPassword;
				this.editForm.staffLevel = String(staffLevel);
				this.editForm.staffTelephone = staffTelephone;
				this.editForm.staffSalary = staffSalary;
				this.editForm.staffRemarks = staffRemarks;
				this.dialogEditVisible = true;
			},
		},
	}
</script>

<style>
</style>
