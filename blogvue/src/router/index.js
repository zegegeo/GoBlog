import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import Register from "../views/Register.vue";
import About from '../views/About.vue';
import ArticleList from '../views/ArticleList.vue';

const routes = [
{ 
    path: "/", 
    name: "Home", 
    component: Home 
},
{ 
    path: "/login", 
    name: "Login", 
    component: Login 
},
{   
    path: "/register", 
    name: "Register", 
    component: Register 
},
{
    path: '/about',
    name: 'About',
    component: About
},
{
    path: '/blog',
    name: 'Blog',
    component: ArticleList,
    meta: {
        requiresAuth: true
    }
},
{
    path: '/articles',
    name: 'ArticleList',
    component: ArticleList,
    meta: { requiresAuth: true }
},
{
    path: '/articles/create',
    name: 'CreateArticle',
    component: () => import('../views/CreateArticle.vue'),
    meta: { requiresAuth: true }
},
{
    path: '/articles/:id(\\d+)',
    name: 'ArticleDetail',
    component: () => import('../views/ArticleDetail.vue'),
    meta: { requiresAuth: true },
    props: true,
    beforeEnter: (to, from, next) => {
        const id = parseInt(to.params.id)
        if (isNaN(id) || id <= 0) {
            next('/articles')
        } else {
            next()
        }
    }
},
{
    path: '/articles/:id(\\d+)/edit',
    name: 'EditArticle',
    component: () => import('../views/EditArticle.vue'),
    meta: { requiresAuth: true },
    props: true,
    beforeEnter: (to, from, next) => {
        const id = parseInt(to.params.id)
        if (isNaN(id) || id <= 0) {
            next('/articles')
        } else {
            next()
        }
    }
}
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    const token = localStorage.getItem('token');
    if (!token) {
      next('/login');
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;