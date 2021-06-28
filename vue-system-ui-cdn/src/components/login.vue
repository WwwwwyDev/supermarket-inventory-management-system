<template>
	<div class="login_container">
		<div class="login_box">
			<div class="avatar_box">
				<!--头像-->
				<img src="../assets/pre.jpg" alt />
			</div>
			<!--添加表单-->
			<el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" class="login_form" label-width="0px">
				<el-form-item prop="username">
					<el-input v-model="loginForm.username" prefix-icon="el-icon-user"></el-input>
				</el-form-item>
				<el-form-item prop="password">
					<el-input v-model="loginForm.password" prefix-icon="el-icon-lock" type="password"></el-input>
				</el-form-item>
				<el-form-item class="btns">
					<el-button @click="login" type="primary">登录</el-button>
					<el-button @click="resetLoginForm" type="info">重置</el-button>
				</el-form-item>
			</el-form>
		</div>
	</div>
</template>

<script>
	export default {
		data() {
			return {
				loginForm: {
					username: "admin",
					password: "admin"
				},
				loginRules: {
					username: [{
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
					password: [{
							required: true,
							message: "请输入密码",
							trigger: "blur"
						},
						{
							min: 5,
							max: 12,
							message: "密码为 5~12 位",
							trigger: "blur"
						}
					]
				}
			};
		},
		methods: {
			resetLoginForm() {
				this.$refs.loginFormRef.resetFields();
			},
			login() {
				this.$refs.loginFormRef.validate(async valid => {
					if (!valid) return;
					const {
						data: res
					} = await this.$http.post("system/apis/login", this.loginForm);
					//console.log(res)
					if (res.code == 30000) {
						this.$message.error("没有此用户");
						return;
					}
					if (res.code == 30001) {
						this.$message.error("密码错误");
						return;
					}
					if (res.code == 20000) {
						this.$message.success("登录成功");
						//console.log(res.data.user.StaffName)
						//console.log(res.data.user.StaffLevel)
						//console.log(res.data.user)
						window.sessionStorage.setItem('username', res.data.user.StaffName);
						window.sessionStorage.setItem('level', res.data.user.StaffLevel);
						window.sessionStorage.setItem('salary', res.data.user.StaffSalary);
						window.sessionStorage.setItem('remarks', res.data.user.StaffRemarks);
						window.sessionStorage.setItem('telephone', res.data.user.StaffTelephone);
						this.$router.push({path: "/home"});
						return;
					}
					this.$message.error("后台错误");
				});
			},
		}
	}
</script>

<style lang="less" scoped>
	// 跟节点样式
	.login_container {
		background-color: #3e6d40;
		height: 100%;
	}

	// 登录框
	.login_box {
		width: 450px;
		height: 300px;
		background-color: #fff;
		border-radius: 3px; // 圆角
		position: absolute; // 绝对定位
		left: 50%;
		top: 50%;
		transform: translate(-50%, -50%); // 根据自己位置 以自己为长度位移

		// 头像框
		.avatar_box {
			width: 130px;
			height: 130px;
			border: 1px solid #eee;
			border-radius: 50%; // 加圆角
			box-shadow: 0 0 10px #ddd; // 盒子阴影
			position: absolute;
			left: 50%;
			transform: translate(-50%, -50%);

			img {
				width: 100%;
				height: 100%;
				border-radius: 50%; // 加圆角
				background-color: #eee;
			}
		}

		.btns {
			display: flex; // 弹性布局
			justify-content: center; // 尾部对齐
		}

		.login_form {
			position: absolute;
			bottom: 0%;
			width: 100%;
			padding: 0 10px;
			box-sizing: border-box; // 设置边框
		}
	}
</style>
