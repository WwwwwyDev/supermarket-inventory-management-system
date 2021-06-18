<template>
	<el-backtop :bottom="60"></el-backtop>
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
					:collapse="isCollapse" :collapse-transition="false" :router="true" :default-active="$route.path">
					<el-menu-item index="/welcome" key="0">
						<i class="el-icon-s-home"></i>
						<template #title>主页</template>
					</el-menu-item>
					<el-menu-item :index="it.MenuAddress" v-for="it in menuList" :key="it.MenuId">
						<i class="el-icon-menu"></i>
						<template #title>{{it.MenuName}}</template>
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
			<el-descriptions-item label="薪水(元)">
				<el-tag size="small">{{salary}}</el-tag>
			</el-descriptions-item>
			<el-descriptions-item label="电话">
				<el-tag size="small">{{telephone}}</el-tag>
			</el-descriptions-item>
			<el-descriptions-item label="备注">
				<el-tag size="small">{{remarks}}</el-tag>
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
				salary: 0,
				remarks: "",
				telephone: "",
			}
		},
		created() {
			this.username = window.sessionStorage.getItem('username');
			this.level = window.sessionStorage.getItem('level');
			this.salary = window.sessionStorage.getItem('salary') / 100;
			this.remarks = window.sessionStorage.getItem('remarks');
			this.telephone = window.sessionStorage.getItem('telephone');
			this.getMenuList();
		},
		methods: {
			logout() {
				window.sessionStorage.clear();
				this.$router.push("/login");
			},
			// 获取所有的导航菜单
			async getMenuList() {
				const {
					data: res
				} = await this.$http.get("system/apis/menu?level=" + this.level);
				// console.log(res);
				if (res.code != 20000) return this.$message.error("菜单加载失败");
				this.menuList = res.data.menus;
				//console.log(res.data.menus)
			},
			// 切换菜单折叠与展开
			toggleCollapase() {
				this.isCollapse = !this.isCollapse;
			},
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
