<template>
	<!-- 卡片视图区 -->
	<el-card>
		<el-row :gutter="25">
			<el-col :span="7">
				<!-- 搜索添加 -->
				<el-input placeholder="请输入搜索员工的姓名" v-model.lazy="queryInfo.searchName" @change="getStaffList">
				</el-input>
			</el-col>

			<el-col :span="4">
				<el-button type="primary" @click="dialogAddVisible = true">添加员工</el-button>
			</el-col>
		</el-row>
		<!-- 列表 -->
		<el-table :data="staffList" border stripe v-loading="isLoading"
			element-loading-background="rgba(255, 255, 255, .5)" element-loading-text="加载中，请稍后..."
			element-loading-spinner="el-icon-loading">
			<el-table-column label="记录号" type="index" fixed="left"></el-table-column>
			<el-table-column label="员工编号" prop="StaffId"></el-table-column>
			<el-table-column label="姓名" prop="StaffName"></el-table-column>
			<el-table-column label="密码" prop="StaffPassword"></el-table-column>
			<el-table-column label="等级" prop="StaffLevel" width="50"></el-table-column>
			<el-table-column label="电话" prop="StaffTelephone"></el-table-column>
			<el-table-column label="薪水(分)" prop="StaffSalary"></el-table-column>
			<el-table-column label="备注" prop="StaffRemarks"></el-table-column>
			<el-table-column label="操作" fixed="right">
				<template #default="scope">
					<!-- 修改 -->
					<el-button type="primary" icon="el-icon-edit" size="mini"
						@click="dialogEditOpen(scope.row.StaffId, scope.row.StaffName,  scope.row.StaffPassword, scope.row.StaffLevel, scope.row.StaffTelephone, scope.row.StaffSalary, scope.row.StaffRemarks)">
					</el-button>
					<!-- 删除 -->
					<el-button type="danger" icon="el-icon-delete" size="mini"
						@click="deleteStaff(scope.row.StaffId,scope.row.StaffLevel)"></el-button>
				</template>
			</el-table-column>
		</el-table>
		<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
			:current-page="queryInfo.page" :page-sizes="[1, 2, 10, 100]" :page-size="queryInfo.limit"
			layout="total, sizes, prev, pager, next, jumper" :total="total"></el-pagination>
	</el-card>

	<!-- 创建对话框 -->
	<el-dialog title="添加员工" v-model="dialogAddVisible" width="50%">
		<el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="100px">
			<el-form-item label="姓名" prop="staffName">
				<el-input v-model="addForm.staffName"></el-input>
			</el-form-item>
			<el-form-item label="密码" prop="staffPassword">
				<el-input v-model="addForm.staffPassword"></el-input>
			</el-form-item>
			<el-form-item label="等级" prop="staffLevel">
				<el-input v-model="addForm.staffLevel" :disabled="true"></el-input>
			</el-form-item>
			<el-form-item label="电话" prop="staffTelephone">
				<el-input v-model="addForm.staffTelephone"></el-input>
			</el-form-item>
			<el-form-item label="薪水(分)" prop="staffSalary">
				<el-input v-model="addForm.staffSalary"></el-input>
			</el-form-item>
			<el-form-item label="备注" prop="staffRemarks">
				<el-input v-model="addForm.staffRemarks"></el-input>
			</el-form-item>
		</el-form>
		<span slot="footer" class="dialog-footer" style="margin-left: 38%;">
			<el-button type="primary" @click="addStaff">确 定</el-button>
			<el-button @click="dialogAddVisible = false">取 消</el-button>
		</span>
	</el-dialog>

	<!-- 编辑对话框 -->
	<el-dialog title="编辑员工" v-model="dialogEditVisible" width="50%">
		<el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="100px">
			<el-form-item label="员工编号" prop="id">
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
	</el-dialog>
</template>

<script>
	export default {
		data() {
			return {
				isLoading:false,
				dialogAddVisible: false,
				dialogEditVisible: false,
				queryInfo: {
					searchName: '',
					page: 1,
					limit: 10
				},
				staffList: [], // 用户列表
				total: 0, // 最大数据记录
				addForm: {
					staffName: '',
					staffPassword: '',
					staffLevel: '1',
					staffTelephone: '',
					staffSalary: '',
					staffRemarks: '',
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
			this.getStaffList();
		},
		methods: {
			async getStaffList() {
				this.isLoading = true;
				// 调用post请求
				const {
					data: res
				} = await this.$http.get("system/apis/staff", {
					params: this.queryInfo
				});
				if (res.code != 20000) {
					this.$message.error("加载用户列表失败");
					this.isLoading = false;
				}
				this.staffList = res.data.value; // 将返回数据赋值
				this.total = res.data.total; // 总个数
				this.isLoading = false;
			},
			// 监听pageSize改变的事件
			handleSizeChange(newLimit) {
				this.queryInfo.limit = newLimit;
				this.getStaffList(); // 数据发生改变重新申请数据
			},
			// 监听pageNum改变的事件
			handleCurrentChange(newPage) {
				this.queryInfo.page = newPage;
				this.getStaffList(); // 数据发生改变重新申请数据
			},
			async deleteStaff(StaffId, StaffLevel) {
				if (StaffLevel == 2) {
					return this.$message.info("管理员用户不能被删除");
				}
				// 弹框
				const confirmResult = await this.$confirm('此操作将永久删除该员工, 是否继续?', '提示', {
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
				} = await this.$http.delete("system/apis/staff?id=" + StaffId);
				if (res.code != 20000) {
					return this.$message.error("删除失败");
				}
				this.$message.success("删除成功");
				this.getStaffList();
			},
			// 添加员工
			addStaff() {
				this.$refs.addFormRef.validate(async valid => {
					if (!valid) return;
					// 发起请求
					const {
						data: res
					} = await this.$http.post("/system/apis/staff", this.addForm);
					if (res.code == 20000) {
						this.getStaffList();
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
				this.editForm.id =  String(id);
				this.editForm.staffName = String(staffName);
				this.editForm.staffPassword = String(staffPassword);
				this.editForm.staffLevel = String(staffLevel);
				this.editForm.staffTelephone = String(staffTelephone);
				this.editForm.staffSalary = String(staffSalary);
				this.editForm.staffRemarks = String(staffRemarks);
				this.dialogEditVisible = true;
			},
		},
	}
</script>

<style>
</style>
