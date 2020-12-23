const components = {};

components.articles = require('./components/articles.vue');
components.single = require('./components/single.vue');
components.editarticle = require('./components/editarticle.vue');
components.login = require('./components/login.vue');
components.manages = require('./components/manages.vue');
components.PageNotFound = require('./components/PageNotFound.vue');
components.about = require('./components/about.vue');
components.pay = require('./components/pay.vue');

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
    path: '/order/payresult',
    name: 'pay',
    component: components.pay,
  }, {
    path: '*',
    component: components.PageNotFound,
  },
];
export default routers;
