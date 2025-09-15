<script setup>
import {VideoPlay , Operation } from '@element-plus/icons-vue';
</script>

<template>
    <el-container class="main-container ">
      <div class="game-ui-container">
        <!-- 背景图片 -->
        <div class="background-image">
          <img src="https://www.yuanshen.com/images/poster.948ee40e.jpg" alt="Background" />
        </div>

        <!-- 右下角功能区 -->
       <div class="function-container">
          <!-- 功能区 -->
          <el-card class="function-card">
            <div class="card-content">
              <p>信息卡片</p>
            </div>
          </el-card>
          <br />
          <br />
          <!-- 按钮区 -->
          <div class="button-area">
            <el-button style="width: 28vh;height: 8vh;" class="start-btn" @click="startGame">
              <el-icon><VideoPlay /></el-icon>
              <span>开始游戏</span>
            </el-button>
            <el-button style="width: 8vh;height: 8vh;" :icon="Operation" circle />
          </div>
        </div>
      </div>
      <br />
    </el-container>
</template>

<style scoped>
.main-container {

  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.game-ui-container {
  position: relative;
  width: 100%;
  height: 100vh;
}

.background-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
}

.background-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.activity-title {
  position: absolute;
  top: 50px;
  left: 50px;
  color: #fff;
}

.card-header {
  display: flex;
  flex-direction: column;
}

.card-content {
  margin-top: 10px;
}

.function-container {
  position: absolute;
  bottom: 50px;
  right: 20px;
  width: 40vh;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
}
 
.function-card {
  width: 100%;
  height: 30vh;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.1);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(2px);
  -webkit-backdrop-filter: blur(10px);
  border-radius: 16px;
  color: black;
  font-family: sans-serif;
}
 
.card-content {
  padding: 20px;
}
 
.button-area {
  display: flex;
  gap: 10px;
}

</style>