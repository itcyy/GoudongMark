import Vue from 'vue'
import App from './App.vue'
import SlideVerify from 'vue-monoplasty-slide-verify';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import router from './router'
import axios from 'axios'
import vuetify from './plugins/vuetify'
import iView from 'iview';
import 'iview/dist/styles/iview.css';
import api from './api/index';


import './assets/css/style.css'
import './assets/css/font.css'
// npm install --save font-awesome
import 'font-awesome/css/font-awesome.min.css'
import { encryptDes, decryptDes } from '@/api'




Vue.prototype.encryptDes = encryptDes;
Vue.prototype.decryptDes = decryptDes;

Vue.prototype.$api = api
Vue.use(iView);
Vue.prototype.$axios = axios
Vue.config.productionTip = false
Vue.use(ElementUI);
Vue.use(SlideVerify);
new Vue({
  axios,
  router,
  el:'#app',
  vuetify,
  render: h => h(App)
}).$mount('#app')
