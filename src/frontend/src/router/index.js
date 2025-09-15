import { createRouter, createWebHashHistory } from 'vue-router'
import Home from 'components/home.vue';
import Seeting from 'components/setting.vue';
import Gameview from 'components/gameview.vue';
import Scriptstore from 'components/scriptstore.vue';
import Allgame from 'components/allgame.vue';
import User from 'components/user.vue';

const routes = [
  { path: '/', component: Home, meta: { menuIndex: '1' } },
  { path: '/Seeting', component: Seeting, meta: { menuIndex: '4' }  },
  { path: '/Gameview', component: Gameview, meta: { menuIndex: '2' }  },
  { path: '/Scriptstore', component: Scriptstore, meta: { menuIndex: '3' }  },
  { path: '/Allgame', component: Allgame, meta: { menuIndex: '2' }  },
  { path: '/User', component: User, meta: { menuIndex: '4' } }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

import { useMenuStore } from '../stores/menuStore'
 
router.beforeEach((to, from, next) => {
  const menuStore = useMenuStore()
  const menuIndex = to.meta.menuIndex || '1'
  menuStore.setActiveMenu(menuIndex)
  next()
})
export default router