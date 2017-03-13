const components = {};

components.article = require('./components/article.vue');
components.single = require('./components/single.vue');
components.editarticle = require('./components/editarticle.vue');

const routers = [
  {
    path: '/',
    name: 'index',
    component: components.article,
  }, {
    path: '/article/:id',
    name: 'article',
    component: components.single,
  }, {
    path: '/admin',
    name: 'admin',
    component: components.editarticle,
  }, {
    path: '*',
    redirect: '/',
  },
];
export default routers;
