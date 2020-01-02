import Vue from 'vue'
import App from './App.vue'
import router from '@/router'
import store from '@/store'
import BootstrapVue from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import 'vue-sidebar-menu/dist/vue-sidebar-menu.css'

Vue.use(BootstrapVue)

import Icon from 'vue-awesome/components/Icon'
Vue.component('icon', Icon)

import 'vue-awesome/icons/times'
import 'vue-awesome/icons/plus'
import 'vue-awesome/icons/chevron-right'

Vue.config.productionTip = false

import VueCompositionApi from '@vue/composition-api';

Vue.use(VueCompositionApi);

import {digitGrouping, dateFormat} from '@/utils/filter'
Vue.filter('digitgrouping', digitGrouping)
Vue.filter('dateformat', dateFormat)

let vm = new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')

export default vm