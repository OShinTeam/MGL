<script setup>
import { Load_logger } from 'wailsjs/go/service/App';
import { useRouter } from 'vue-router'
import { inject } from 'vue';
const LangInfo = inject('LangInfo');

const router = useRouter()

function goToUser() {
  router.push('/User')
}

function goToGame() {
  router.push('/Gameview')
}
</script>

<template>
    <el-container class="main-container ">
      <el-main class="scrollable-container">
        <!-- 用户设置卡片 -->
        <el-card class="setting-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>{{ LangInfo.textmap.setting_user }}</span>
            </div>
          </template>
          <div class="setting-content">
            <p>{{ LangInfo.textmap.setting_user_desc }}</p>
          </div>
        </el-card>
        <!-- 基础设置卡片 -->
        <el-card class="setting-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>{{ LangInfo.textmap.setting_normal }}</span>
            </div>
          </template>
          <div class="setting-content">
            <p>{{ LangInfo.textmap.setting_normal_desc }}</p>
          </div>
        </el-card>
        <!-- 高级设置卡片 -->
        <el-card class="setting-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>{{ LangInfo.textmap.setting_adv }}</span>
            </div>
          </template>
          <el-button :disabled=isDisabled type="primary" @click="Load_logger(1)">Debug模式</el-button>
          <el-button :disabled=isDisabled type="primary" @click="goToUser">进入用户管理页</el-button>
          <el-button :disabled=isDisabled type="primary" @click="goToGame">进入游戏详情页</el-button>
          <div class="setting-content">
            <p>{{ LangInfo.textmap.setting_adv_desc }}</p>
          </div>
        </el-card>
        <br />
      </el-main>
    </el-container>
</template>

<style scoped>
.main-container {
  background-color: #f5f7fa;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.setting-card {
  margin-bottom: 24px;
  border-radius: 8px;
}
 
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: bold;
}
 
.setting-content {
  font-size: 14px;
  color: #666;
}
 
.header-btn {
  padding: 0;
  font-size: 12px;
}

.scrollable-container::-webkit-scrollbar {
  display: none;
}
</style>