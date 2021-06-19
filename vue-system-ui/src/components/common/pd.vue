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
		</el-table>
		<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
			:current-page="queryInfo.page" :page-sizes="[1, 2, 10, 100]" :page-size="queryInfo.limit"
			layout="total, sizes, prev, pager, next, jumper" :total="total"></el-pagination>
	</el-card>

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
			}
		},
	}
</script>

<style>
</style>
