import Vue from 'vue'
import App from './App.vue'
import router from '@/router'
import store from '@/store'
import BootstrapVue from 'bootstrap-vue'
import VueSwal from 'vue-swal'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import 'vue-sidebar-menu/dist/vue-sidebar-menu.css'

Vue.use(BootstrapVue)
Vue.use(VueSwal)

Vue.config.productionTip = false

import VueCompositionApi from '@vue/composition-api';

Vue.use(VueCompositionApi);

let vm = new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')

export default vm