const components = {};

components.articles = require('./components/articles.vue');
components.single = require('./components/single.vue');
components.editarticle = require('./components/editarticle.vue');
components.login = require('./components/login.vue');
components.manages = require('./components/manages.vue');
components.PageNotFound = require('./components/PageNotFound.vue');
components.about = require('./components/about.vue');
components.pay = require('./components/pay.vue');
components.share = require('./components/share.vue');

const routers = [
  {
    path: '/',
    name: 'index',
    component: components.articles,
  }, {
    path: '/article/:id',
    name: 'article',
    component: components.single,
  }, {
    path: '/edit',
    name: 'new',
    component: components.editarticle,
  }, {
    path: '/edit/:id',
    name: 'edit',
    component: components.editarticle,
  }, {
    path: '/manages',
    name: 'manages',
    component: components.manages,
  }, {
    path: '/about',
    name: 'about',
    component: components.about,
  }, {
    path: '/login',
    name: 'login',
    component: components.login,
  }, {
    path: '/api/order/payresult',
    name: 'pay',
    component: components.pay,
  }, {
    path: '/share',
    name: 'share',
    component: components.share,
    props: route => ({
      bgc: route.query.bgc,
      content: route.query.content,
      title: route.query.title,
      author: route.query.author,
      color: route.query.color,
    }),
  }, {
    path: '*',
    component: components.PageNotFound,
  },
];
export default routers;
