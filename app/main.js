import Vue from 'vue';
import vueResource from 'vue-resource';
import VueRouter from 'vue-router';
import app from './App.vue';
import routers from './routers';

const baseTitle = '冯小懋的个人博客';

Vue.use(vueResource);
Vue.use(VueRouter);

Vue.http.options.emulateJSON = true;
Vue.http.options.emulateHTTP = true;

Vue.directive('title', {
  inserted: (el) => {
    document.title = el.dataset.title ? `${el.dataset.title} - ${baseTitle}` : baseTitle;
  },
  componentUpdated: (el) => {
    document.title = el.dataset.title ? `${el.dataset.title} - ${baseTitle}` : baseTitle;
  },
});

const App = Vue.extend(app);

const router = new VueRouter({
  routes: routers,
  saveScrollPosition: true,
});

function runRouter() {
  return new Vue({ el: '#app', router, render: h => h(App) });
}
runRouter();
