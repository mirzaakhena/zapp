import Vue from 'vue'
import App from './App.vue'
import BootstrapVue from 'bootstrap-vue'
import VueCompositionApi from '@vue/composition-api';

import router from '@/router'
import store from '@/store'
import {digitGrouping, dateFormat} from '@/utils/filter'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import 'vue-sidebar-menu/dist/vue-sidebar-menu.css'
import '@fortawesome/fontawesome-free/css/all.css'
import '@/assets/style.css'

Vue.use(BootstrapVue)
Vue.use(VueCompositionApi);

Vue.filter('digitgrouping', digitGrouping)
Vue.filter('dateformat', dateFormat)

Vue.config.productionTip = false

let vm = new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')

export default vm
