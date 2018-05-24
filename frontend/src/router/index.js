import Vue from 'vue';
import Router from 'vue-router';
import Meta from 'vue-meta';

const HelloWorld = () => import('@/components/HelloWorld');
const About = () => import('@/components/About');

Vue.use(Router);
Vue.use(Meta);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld,
    },
    {
      path: '/about',
      name: 'About',
      component: About,
    },
  ],
});
