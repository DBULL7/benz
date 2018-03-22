import Vue from 'vue';
import App from './components/App';
import router from './router';
import store from './store'

Vue.config.productionTip = false;

new Vue({
  el: '#root',
  router,
  store,
  components: { App },
  template: '<App/>',
});