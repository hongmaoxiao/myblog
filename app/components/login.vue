<template>
  <section class="login">
    <p><label>用户名:</label><input v-model="username" placeholder="请输入用户名" /></p>
    <p><label>密码:</label><input v-model="password" placeholder="请输入密码" /></p>
    <button @click="login">登录</button>
  </section>
</template>
<script>
  export default {
    data() {
      return {
        username: '',
        password: '',
      };
    },
    methods: {
      login() {
        if (!this.username || !this.password) {
          return;
        }
        this.$http.post('/login', {
          username: this.username,
          password: this.password,
        })
        .then(() => {
          this.username = '';
          this.password = '';
        }, (error) => {
          console.log(error);
        });
      },
    },
  };
</script>
<style scoped>
.login {
  width: 300px;
  margin: 100px auto;
  padding: 50px;
  border: 1px solid #ddd;
  box-shadow: 0 3px 15px #ccc;
}
.login > p {
  margin-bottom: 20px;
}
.login input {
  padding: 6px 12px;
}
.login > p > label {
  width: 70px;
  display: inline-block;
}
.login-btn {
  padding: 3px 12px;
}
</style>
