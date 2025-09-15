<script setup>
import { Plus } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import { inject } from 'vue';
const LangInfo = inject('LangInfo');
function AddGame(){
  ElMessage.info("添加游戏")
}
</script>

<template>
    <el-container class="main-container">
      <!-- 分类导航 -->
      <div class="category-nav">
        <el-menu
          mode="horizontal"
          default-active="all"
          class="top-menu"
          background-color="#FFFFFF"
          text-color="#333"
          active-text-color="#409EFF"
          :ellipsis="false"
        >
          <el-menu-item index="all">{{ LangInfo.textmap.game_all }}</el-menu-item>
          <el-menu-item index="recent">{{ LangInfo.textmap.game_recent }}</el-menu-item>

          <!-- 分类菜单 -->
          <el-sub-menu
          index="category"
          popper-append-to-body
          :show-timeout="200"
          :hide-timeout="100"
          >
            <template #title>{{ LangInfo.textmap.game_cate }}</template>
            <el-menu-item
              v-for="(_, index) in 5"
              :key="index"
              :index="`category-${index}`"
            >
              分类占位
            </el-menu-item>
          </el-sub-menu>
          <div class="right-menu">
            <el-menu-item @click="AddGame"><el-icon><Plus /></el-icon></el-menu-item>
          </div>
        </el-menu>
      </div>
      <!-- 游戏卡片展示 -->
      <el-row :gutter="20" style="padding: 20px;" class="game-card-scrollable scrollable-container">
        <el-col
          v-for="(_, index) in 50"
          :key="index"
          :xs="12"
          :sm="6"
          :md="4"
          :lg="3"
          :xl="2"
        >
          <div class="game-card">
            <div class="game-card__image"></div>
            <div class="game-card__footer">
              <span>游戏占位</span>
            </div>
          </div>
        </el-col>
      </el-row>
    </el-container>
</template>

<style scoped>
.main-container {
  background-color: #f5f7fa;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.game-card {
  background-color: #f5f5f5;
  border-radius: 4px;
  overflow: hidden;
  transition: all 0.3s ease;
  display: flex;
  flex-direction: column;
  margin-bottom: 16px;
}
 
.game-card:hover {
  transform: scale(1.03);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}
 
.game-card__image {
  width: 100%;
  padding-top: 135%;
  background-color: black;
  border-radius: 0;
}
 
.game-card__footer {
  padding: 10px;
  font-size: 14px;
  font-weight: 500;
  color: #333;
  text-align: center;
}

.right-menu {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-left: auto;
}

.scrollable-container::-webkit-scrollbar {
  display: none;
}

.game-card-scrollable {
  overflow-y: auto;
}
</style>