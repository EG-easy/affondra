import Vue from "vue";
import VueRouter from "vue-router";
import Buy from "../views/Buy.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Buy,
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
