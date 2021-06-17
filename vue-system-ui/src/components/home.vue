<template>
	<el-container class="home-container">
		<!--头部布局-->
		<el-header>
			<div>
				<img src="../assets/logo.png" alt class="logo_img" />
				<span>超市进销存管理系统</span>
			</div>
			<el-button type="info" @click="dialogVisible = true">{{username}}</el-button>
		</el-header>
		<el-container>
			<!--侧边布局-->
			<el-aside :width="isCollapse ?'64px':'200px'">
				<!--伸缩按钮-->
				<div class="toggle-button" @click="toggleCollapase">|||</div>
				<el-menu background-color="#545c64" text-color="#fff" active-text-color="#409eff" unique-opened
					:collapse="isCollapse" :collapse-transition="false" :router="true" :default-active="activePath">
					<el-menu-item index="2">
						<i class="el-icon-menu"></i>
						<template #title>导航二</template>
					</el-menu-item>
				</el-menu>
			</el-aside>
			<!--主体布局-->
			<el-main>
				<!--路由占位符-->
				<router-view></router-view>
			</el-main>
		</el-container>
	</el-container>

	<el-dialog title="当前用户信息" v-model="dialogVisible" width="30%">
		<el-descriptions title="">
			<el-descriptions-item label="用户名">
				<el-tag size="small">{{username}}</el-tag>
			</el-descriptions-item>
			<el-descriptions-item label="权限等级">
				<el-tag size="small">{{level}}</el-tag>
			</el-descriptions-item>
		</el-descriptions>
		<template #footer>
			<span class="dialog-footer">
				<el-button type="danger" @click="logout">退出登录</el-button>
				<el-button @click="dialogVisible = false">返回</el-button>
			</span>
		</template>
	</el-dialog>
</template>

<script>
	export default {
		data() {
			return {
				// 左侧菜单
				menuList: [],
				isCollapse: false,
				dialogVisible: false,
				username: "",
				level: 0,
			}
		},
		created() {
			this.username = window.sessionStorage.getItem('username');
			this.level = window.sessionStorage.getItem('level');
			// this.getMenuList();
		},
		methods: {
			logout() {
				window.sessionStorage.clear();
				this.$router.push("/login");
			},
			// 获取所有的导航菜单
			async getMenuList() {
				// const {
				// 	data: res
				// } = await this.$http.get("menus");
				// console.log(res.data);
				// if (res.status != 200) return this.$message.error("操作失败！！！");
				// this.menuList = res.data;
			},
			// 切换菜单折叠与展开
			toggleCollapase() {
				this.isCollapse = !this.isCollapse;
			},
			// // 保存二级菜单的路径
			// saveNavState(activePath) {
			// 	window.sessionStorage.setItem('activePath', activePath); // 存放点击的二级菜单
			// 	this.activePath = activePath; // 给点击的菜单添加高亮
			// },
		}
	};
</script>

<style lang="less" scoped>
	.el-header {
		background-color: #373d41;
		display: flex;
		justify-content: space-between; // 左右贴边
		padding-left: 0%; // 左边界
		align-items: center; // 水平
		color: #fff;
		font-size: 20px;

		>div {
			//左侧div加布局
			display: flex;
			align-items: center;

			span {
				margin-left: 15px;
			}
		}
	}

	.el-aside {
		background-color: #333744;

		.el-menu {
			border-right: none; // 对其右边框
		}
	}

	.el-main {
		background-color: #eaedf1;
	}

	.home-container {
		height: 100%;
	}

	.logo_img {
		width: 20%;
		height: 100%;
	}

	.iconfont {
		margin-right: 10px;
	}

	.toggle-button {
		background-color: #4A5064;
		font-size: 10px;
		line-height: 24px;
		color: #fff;
		text-align: center;
		letter-spacing: 0.2em;
		cursor: pointer; // 显示鼠标指针为：小手
	}
</style>
