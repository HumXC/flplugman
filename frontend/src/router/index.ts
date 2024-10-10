import { createRouter, createWebHistory } from 'vue-router'
import { createMemoryHistory } from 'vue-router';
import SplashScreen from '../views/SplashScreen.vue';
import HomeView from '../views/Home.vue';
const router = createRouter({
    history: createMemoryHistory(),
    routes: [
        {
            path: '/',
            name: '欢迎页',
            component: SplashScreen,
        }, {
            path: '/home',
            name: '首页',
            component: HomeView
        }
    ]
});

router.beforeEach((to, from, next) => {
    document.title = to.meta.title as string;
    next();
});

export default router;
