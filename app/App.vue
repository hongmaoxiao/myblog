<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>
<style>
*{
  margin: 0;
  padding: 0;
  box-sizing: content-box;
}
html{
  height: 100%;
}

body{
  position: relative;
  zoom: 1;
  height: 100%;
  margin: auto;
}
@media screen and (max-width: 700px) {
  body {
    width: 100%;
  }
}
@media screen and (min-width: 700px) {
  body {
    width: 100%;
    max-width: 1200px;
  }
}
@media screen and (min-width: 700px) and (max-width: 1200px) {
  body {
    width: 95%;
  }
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
  margin: 0;
  padding: 0;
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
