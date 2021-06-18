<template>
	<!-- 卡片视图区 -->
	<el-card>
		<el-row :gutter="25">
			<el-col :span="10">
				<!-- 搜索添加 -->
				<el-input placeholder="请输入搜索内容">
					<el-button slot="append" icon="el-icon-search"></el-button>
				</el-input>
			</el-col>

			<el-col :span="4">
				<el-button type="primary">添加员工</el-button>
			</el-col>
		</el-row>
		<!-- 用户列表 -->
		<el-table :data="staffList" border stripe>
			<el-table-column label="记录号" type="index" fixed="left"></el-table-column>
			<el-table-column label="编号" prop="StaffId"></el-table-column>
			<el-table-column label="姓名" prop="StaffName"></el-table-column>
			<el-table-column label="密码" prop="StaffPassword"></el-table-column>
			<el-table-column label="等级" prop="StaffLevel" width="50"></el-table-column>
			<el-table-column label="电话" prop="StaffTelephone"></el-table-column>
			<el-table-column label="薪水" prop="StaffSalary"></el-table-column>
			<el-table-column label="备注" prop="StaffRemarks"></el-table-column>
			<el-table-column label="操作" fixed="right">
				<template #default="scope">
					<!-- 修改 -->
					<el-button type="primary" icon="el-icon-edit" size="mini"></el-button>
					<!-- 删除 -->
					<el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteStaff(scope.row.StaffId,scope.row.StaffLevel)"></el-button>
				</template>
			</el-table-column>
		</el-table>
		<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
			:current-page="queryInfo.page" :page-sizes="[1, 2, 5, 100]" :page-size="queryInfo.limit"
			layout="total, sizes, prev, pager, next, jumper" :total="total"></el-pagination>
	</el-card>
</template>

<script>
	export default {
		data() {
			return {
				// 请求数据
				queryInfo: {
					searchName: "",
					page: 1,
					limit: 5
				},
				staffList: [], // 用户列表

				total: 0, // 最大数据记录
			}
		},
		created() {
			this.getStaffList();
		},
		methods: {
			async getStaffList() {
				// 调用post请求
				const {
					data: res
				} = await this.$http.get("system/apis/staff", {
					params: this.queryInfo
				});
				if (res.code != 20000) {
					this.$message.error("加载用户列表失败");
				}
				this.staffList = res.data.value; // 将返回数据赋值
				this.total = res.data.total; // 总个数
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
			}
		},
	}
</script>

<style>
</style>
