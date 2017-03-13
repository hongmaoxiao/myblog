<script>
export default {
};
</script>
<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>
<style>
*{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
html{
  height: 100%;
}

body{
  position: relative;
  zoom: 1;
  width: 100%;
  height: 100%;
}
a:link,
a:visited,
a:hover,
a:active{
  text-decoration: none;
}

#app{
  font-family: -apple-system,SF UI Text,Arial,PingFang SC,Hiragino Sans GB,Microsoft YaHei,WenQuanYi Micro Hei,sans-serif;
  position: relative;
}
ul{
  list-style: none;
}
.pull-left{
  float: left;
}
.pull-right{
  float: right;
}

.clearfix {
  zoom: 1;
}

.clearfix:after {
  content: '';
  display: block;
  height: 0;
  visibility: hidden;
  clear: both;
}
</style>
