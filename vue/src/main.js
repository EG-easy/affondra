import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import _ from "lodash";
Object.defineProperty(Vue.prototype, "$lodash", { value: _ });

Vue.config.devtools = true;
Vue.config.productionTip = false;

// import all componets here
const ComponentContext = require.context("./", true, /\.vue$/i, "lazy");
ComponentContext.keys().forEach((componentFilePath) => {
  const componentName = componentFilePath.split("/").pop().split(".")[0];
  Vue.component(componentName, () => ComponentContext(componentFilePath));
});

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
