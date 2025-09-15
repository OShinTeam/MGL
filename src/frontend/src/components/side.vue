<script setup>
import { inject,computed } from 'vue';
import { useRouter } from 'vue-router'
import { HomeFilled, Avatar, Tools, Menu, MoreFilled } from '@element-plus/icons-vue';
import { useMenuStore } from '../stores/menuStore'
const menuStore = useMenuStore()
const activeIndex = computed(() => menuStore.activeMenu)
const LangInfo = inject('LangInfo');
const router = useRouter()

function goToHome() {
  router.push('/')
}

function goToSetting() {
  router.push('/Seeting')
}

function goToScript() {
  router.push('/Scriptstore')
}

function goToAll() {
  router.push('/Allgame')
}
</script>

<template>
  <div class="side-bar" v-if="LangInfo && LangInfo.textmap">
    <el-aside style="width: 200px;" >
      <div class="user-info">
        <div class="avatar">
          <el-icon><Avatar /></el-icon>
        </div>
        <div class="username">{{ LangInfo.translator }}</div>
      </div>
  
      <el-menu
        :default-active="activeIndex"
        class="sidebar-menu"
        background-color="#fff"
        text-color="#1d1f21"
        active-text-color="#4da6ff"
        :collapse="false"
      >
        <el-menu-item index="1" @click="goToHome">
          <el-icon><HomeFilled /></el-icon>
          <span>{{ LangInfo.textmap.menu_main }}</span>
        </el-menu-item>
        <el-menu-item index="2" @click="goToAll">
          <el-icon><MoreFilled /></el-icon>
          <span>{{ LangInfo.textmap.menu_allgame }}</span>
        </el-menu-item>
        <el-menu-item index="3" @click="goToScript">
          <el-icon><Menu /></el-icon>
          <span>{{ LangInfo.textmap.menu_script }}</span>
        </el-menu-item>
        <el-menu-item index="4" @click="goToSetting">
          <el-icon><Tools /></el-icon>
          <span>{{ LangInfo.textmap.menu_setting }}</span>
        </el-menu-item>
      </el-menu>
      <!-- 状态区域 -->
      <div class="status-area">
        <p>状态显示</p>
      </div>
    </el-aside>
  </div>
</template>

<style scoped>
.status-area {
  width: 200px;
  position: absolute;
	left: 0;
	bottom: 0;
  padding: 10px 0;
  font-size: 12px;
  color: #666;
  text-align: center;
  border-top: 1px solid #e4e7ed;
  background-color: #fff;
}

.side-bar {
  background-color: #fff;
  height: 100vh;
  width: 200px;
  display: flex;
  flex-direction: column;
}
 
.user-info {
  padding: 20px;
  display: flex;
  align-items: center;
  background-color: #fff;
  color: black;
}
 
.avatar {
  font-size: 24px;
  margin-right: 10px;
}
 
.username {
  font-size: 15px;
  font-weight: 500;
}
 
.sidebar-menu {
  border-right: none;
  flex: 1;
}
</style>
