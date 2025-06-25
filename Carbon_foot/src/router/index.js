import { createRouter, createWebHistory } from 'vue-router'
import addContact from "@/components/addContact.vue"
import contactList from "@/components/contactList.vue"
import person from "@/components/person.vue"
import TodoView from '../views/TodoView.vue'

import HomeView from "@/views/HomeView.vue"
import Show from "@/views/show.vue"
import adovocate from "@/views/adovocate.vue"
import Mylogin from "@/views/MyLogin.vue"
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Mylogin',
      component: Mylogin
    },
    {
      path: '/show',
      name: 'show',
      component: Show
    },
    {
      path: '/adovocate',
      name: 'adovocate',
      component: adovocate
    },
    {
      path: '/todolist',
      component: TodoView
    },
    {
      path: '/addContact',
      name: 'add',
      component: addContact
    },
    {
      path: '/contactList',
      name: 'list',
      component: contactList
    },
    {
      path: '/person',
      name: 'person',
      component: person
    },


  ]
})


// 路由守卫
router.beforeEach((to, from, next) => {
  const publicPages = ['/']; // 公共页面路径数组，例如登录页
  const authRequired = !publicPages.includes(to.path);
  const token = localStorage.getItem('token');

  if (authRequired && !token) {
    // 如果访问的是需要登录的页面且没有 token，则跳转至登录页
    next('/');
  } else {
    next(); // 否则继续导航
  }
});

export default router
