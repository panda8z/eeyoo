<template>
  <div class="container">
    <!-- 登录框 -->
    <div class="loginBox">
      <!-- 登录信息表单 -->
      <a-form-model
        layout="horizontal"
        ref="loginFormRef"
        class="loginForm"
        :model="formdata"
        :rules="rules"
      >
        <!-- 用户名 -->
        <a-form-model-item prop="username">
          <a-input v-model="formdata.username" placeholder="请输入用户名">
            <a-icon
              slot="prefix"
              type="user"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>

        <!-- 密码 -->
        <a-form-model-item prop="password">
          <a-input
            v-model="formdata.password"
            placeholder="请输入密码"
            enter-button
            type="password"
            v-on:keyup.enter="login"
          >
            <a-icon
              slot="prefix"
              type="lock"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>
        <!-- 登录，取消按钮 -->
        <a-form-model-item class="loginBoxBtns">
          <a-button type="primary" class="loginBtn" @click="login()"
            >Login</a-button
          >
          <a-button type="info" class="loginBtn" @click="resetForm()"
            >Cancle</a-button
          >
        </a-form-model-item>
      </a-form-model>
    </div>
    <a-spin size="large" :spinning="spinning" />
  </div>
</template>

<script>
export default {

  data() {
    return {
      spinning: false,
      formdata: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名长度4-12字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 20, message: '密码长度6-20字符', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    resetForm() {
      this.$refs.loginFormRef.resetFields()
    },
    login() {
      this.spinning = true
      this.$refs.loginFormRef.validate(
        async valid => {
          if (!valid) {
            return this.$message.error('输入非法，请重新输入')
          }
          const { data: res } = await this.$http.post('login', this.formdata)
          this.spinning = false
          if (res.status !== 200) {
            return this.$message.error(res.msg)
          }
          window.sessionStorage.setItem('token', res.data.token)
          this.$router.push('index')
          this.$message.success('登录成功！')
        }
      )
    }
  }

}
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  background-color: #282c34;
}

.loginBox {
  width: 450px;
  height: 300px;
  background-color: aliceblue;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
}

.loginBoxBtns {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
}

.loginBtn {
  margin: 10px;
}
</style>
