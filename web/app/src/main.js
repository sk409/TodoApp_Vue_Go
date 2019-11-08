import Vue from 'vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import lang from "element-ui/lib/locale/lang/ja"
import locale from "element-ui/lib/locale"
import App from './App.vue'
import router from './router'
import store from './store'

locale.use(lang)
Vue.use(ElementUI);
Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')