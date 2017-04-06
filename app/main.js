import Vue from 'vue';
import vueResource from 'vue-resource';
import VueRouter from 'vue-router';
import VueCookie from 'vue-cookie';
import app from './App.vue';
import routers from './routers';

const baseTitle = '冯小懋的个人博客';

Vue.use(vueResource);
Vue.use(VueRouter);
Vue.use(VueCookie);

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

const checkLogin = ['new', 'edit', 'manages'];

router.beforeEach((to, from, next) => {
  const mysession = VueCookie.get('mysession');
  if (to.name === 'login' && mysession) {
    router.push('manages');
  }
  if (checkLogin.indexOf(to.name) > -1) {
    if (mysession) {
      next();
    } else {
      router.push('login');
    }
  } else {
    next();
  }
});

function runRouter() {
  return new Vue({ el: '#app', router, render: h => h(App) });
}
runRouter();
