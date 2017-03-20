<template>
  <header>
    <router-link to="/" class="to-home">
      <img src="../img/logo.jpg" class="logo" alt="logo" />
    </router-link>
    <ul class="nav clearfix">
      <li class="nav-item"><a href="mailto:buaaxhm@gmail.com">联系我</a></li>
    </ul>
  </header>
</template>
<style scoped>
header {
  height: 60px;
}
.to-home {
  float: left;
}
.logo {
  width: 60px;
  height: 60px;
  border-radius: 50%;
}
.nav {
  float: right;
}
.nav-item {
  display: inline-block;
  padding: 20px;
}
.nav-item > a {
  font-size: 16px;
  color: #333;
  line-height: 1.5em;
}
@media screen and (max-width: 700px) {
  header {
    padding: 20px 0;
  }
  .to-home,
  .nav {
    float: none;
    display: block;
    text-align: center;
  }
}
@media screen and (min-width: 700px) {
  header {
    padding: 20px 50px;
  }
}
</style>
