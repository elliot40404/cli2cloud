import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import { useIdStore } from '../stores/main';

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      beforeEnter: (to, from, next) => {
        const store = useIdStore();
        const id = store.id;
        if (id && id.length === 8) {
          next('/cli')
        } else {
          next()
        }
      }
    },
    {
      path: '/cli',
      name: 'cli',
      component: () => import('../views/LoginView.vue'),
      beforeEnter: (to, from, next) => {
        const store = useIdStore();
        const id = to.params.id || store.id;
        if (id && id.length === 8) {
          next()
        } else {
          next('/')
        }
      }
    }
  ]
})

export default router
