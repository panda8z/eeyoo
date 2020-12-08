<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model
        layout="horizontal"
        ref="loginFormRef"
        class="loginForm"
        :model="formdata"
        :rules="rules"
      >
        <a-form-model-item prop="username">
          <a-input
            v-model="formdata.username"
            placeholder="请输入用户名"
          >
            <a-icon
              slot="prefix"
              type="user"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>

        <a-form-model-item prop="password">
          <a-input
            v-model="formdata.password"
            placeholder="请输入密码"
            type="password"
          >
            <a-icon
              slot="prefix"
              type="lock"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>

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
  </div>
</template>

<script>
export default {

  data() {
    return {
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
      console.log(this)
      this.$refs.loginFormRef.resetFields()
    },
    login() {
      this.$refs.loginFormRef.validate(
        async valid => {
          if (!valid) {
            return this.$message.error('输入非法，请重新输入')
          }
          const { data: res } = await this.$http.post('login', this.formdata)
          console.log(res)
          if (res.status === 200) {
            window.sessionStorage.setItem('token', res.data.token)
            this.$router.push('/admin')
          }
        }
      )
    }
  }

}
</script>

<style scoped>
.container {
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
