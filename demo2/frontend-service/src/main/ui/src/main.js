import Vue from 'vue'
// eslint-disable-next-line
import _ from './logger.js'

import App from './App.vue'
import Router from './router.js'

Vue.config.productionTip = false;

new Vue({
  render: h => h(App),
  router: Router,
}).$mount('#app');
