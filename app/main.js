import 'github-markdown-css';
import Vue from 'vue';
import vueResource from 'vue-resource';
import VueRouter from 'vue-router';
import app from './App.vue';
import routers from './routers';


Vue.use(vueResource);
Vue.use(VueRouter);

Vue.http.options.emulateJSON = true;
// Vue.http.options.emulateHTTP = true;

const App = Vue.extend(app);

const router = new VueRouter({
  routes: routers,
  saveScrollPosition: true,
});

// router.beforeEach((to, from, next) => {
//   if (to.path !== '/') {
//     const cookie = $.cookie('email');
//     console.log(cookie);
//     if (cookie) {
//       next();
//     } else {
//       next('/');
//     }
//   } else {
//     next();
//   }
// });
function runRouter() {
  return new Vue({ el: '#app', router, render: h => h(App) });
}
runRouter();
