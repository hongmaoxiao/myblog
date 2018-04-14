<template>
  <section>
    <Common />
    <section class="about markdown-body">
      <h2>关于我</h2>
      <ul>
        <li>JavaScript、Go，会点Python、Node。</li>
        <li>喜欢三体，喜欢玩游戏，闲时看点书。</li>
        <li>研究源代码是一件很好玩的事情。</li>
      </ul>
      <h2>联系我</h2>
      <p>Github： <a href='https://github.com/hongmaoxiao'>https://github.com/hongmaoxiao</a></p>
      <p>Email： buaaxhm#gmail.com(#替换为@)</p>
      <p>掘金：<a href='https://juejin.im/user/584a098ab123db00660937e2/posts'>supermao</a></p>
      <p>知乎专栏：<a href='https://zhuanlan.zhihu.com/c_145195250'>凿壁偷点光</a></p>
      <p>segmentfault：<a href='https://segmentfault.com/u/hongmaoxiao'>supermao</a></p>
      <p>微信公众号：大部分都是原创技术文章，欢迎关注。</p>
      <p><img src='https://fengxiaomao.com/upload/qrcode_wx_p.jpg' alt='qrcode' /></p>
      <h2>开源</h2>
      <p>nowjs：now is a time toolkit for javascript.  <a href='https://github.com/hongmaoxiao/now'>https://github.com/hongmaoxiao/now</a></p>
    </section>
    <Foot />
  </section>
</template>
<script>
  // import 'github-markdown-css';
  import Common from './header';
  import Foot from './footer';

  export default {
    components: {
      Common,
      Foot,
    },
  };
</script>
<style>
.about {
  margin: 3em 10px 5em;
  box-shadow: 0 3px 15px #ccc;
  border: 1px solid #ddd;
  padding: 20px;
}
</style>
