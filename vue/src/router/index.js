import Vue from "vue";
import VueRouter from "vue-router";
import Buy from "../views/Buy.vue";
import About from "../views/About.vue";

Vue.use(VueRouter);

const routes = [
  {
    name: 'buy',
    path: "/",
    component: Buy,
  },
  {
    name: 'about',
    path: "/about",
    component: About,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
