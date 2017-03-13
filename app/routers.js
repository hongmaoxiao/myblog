const components = {};

components.article = require('./components/article.vue');
components.single = require('./components/single.vue');

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
    path: '*',
    redirect: '/',
  },
];
export default routers;
