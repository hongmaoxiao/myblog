<template>
  <section class="login">
    <form @submit.prevent="login">
      <p><label>用户名:</label><input v-model="username" name="username" placeholder="请输入用户名" required /></p>
      <p><label>密码:</label><input v-model="password" name="password" type="password" placeholder="请输入密码" required /></p>
      <p class="submit-wrapper">
        <input name="submit" type="submit" class="login-btn" value="登录" />
      </p>
    </form>
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
        .then((res) => {
          if (res.status === 200) {
            console.log("session", res);
            this.$router.push({name: "manages"});
          } else {
            alert("用户名或者密码错误！")
          }
        }, (error) => {
          alert("用户名或者密码错误！")
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
.login p {
  margin-bottom: 20px;
}
.login input {
  padding: 6px 12px;
}
.login p > label {
  width: 70px;
  display: inline-block;
}
.login p.submit-wrapper {
  padding-top: 15px;
  width: 100%;
  text-align: center;
  margin-bottom: 0;
}
.login-btn {
  padding: 3px 15px;
  border-radius: 5px;
  border: 0;
  outline: 0;
  font-size: 16px;
}
</style>
