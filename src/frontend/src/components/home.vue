<template>
    <el-container class="main-container ">
      <el-main class="scrollable-container">
        <!-- 顶部横幅卡片 -->
        <el-card class="info-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>一些软件信息</span>
            </div>
          </template>
          <div class="info-content">
            <p>信息</p>
          </div>
        </el-card>

          <!-- 中部信息卡片 -->
          <el-row :gutter="20" style="height: 425px;">
            <el-col :span="12" style="display: flex; flex-direction: column; gap: 8px;">
              <!-- 软件版本与更新信息 -->
              <el-card class="info-card" shadow="hover">
                <template #header>
                  <div class="card-header">
                    <span>一些软件信息</span>
                  </div>
                </template>
                <div class="info-content">
                  <!-- 示例内容 -->
                  <p>当前版本：v1.0.0</p>
                  <p>更新时间：2025-07-29</p>
                </div>
              </el-card>

              <!-- 登录信息 -->
              <el-card class="info-card" shadow="hover">
                <template #header>
                  <div class="card-header">
                    <span>登录信息</span>
                  </div>
                </template>
                <div class="info-content">
                  <!-- 示例内容 -->
                  <p>用户: MoGuQAQ</p>
                  <p>一些信息</p>
                </div>
              </el-card>
            </el-col>

            <!-- 最近启动 -->
            <el-col :span="12">
              <el-card class="info-card" shadow="hover" style="height: 400px;">
                <template #header>
                  <div class="card-header">
                    <span>最近启动</span>
                  </div>
                </template>
                <div style="display: flex; flex-direction: column;">
                  <div v-for="n in 3" :key="n" class="last-card">
                    <!-- 左侧图片 -->
                    <div class="last-img"></div>
                
                    <!-- 右侧信息 -->
                    <div class="info-content">
                      <p style="font-weight: bold; font-size: 16px;">游戏名称 {{ n }}</p>
                      <p>累计时长：{{ n * 10 }} 小时</p>
                      <p>上次游玩：2024-10-0{{ n }}</p>
                    </div>
                  </div>
                </div>
              </el-card>
            </el-col>
          </el-row>

        <!-- 系统信息 -->
        <el-card class="info-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>一些系统信息</span>
            </div>
          </template>
          <div class="info-content">
            <p>信息</p>
          </div>
        </el-card>
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

.info-card {
  margin-bottom: 24px;
  border-radius: 8px;
}

.info-content {
  font-size: 14px;
  color: #666;
}
 

.el-card .el-card__header {
  background-color: #f5f7fa;
  font-weight: bold;
}

.scrollable-container::-webkit-scrollbar {
  display: none;
}

.last-card:hover {
  transform: scale(1.03);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.last-card {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 4px;
  line-height: 0.6;
  height: 86px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.1);
  backdrop-filter: blur(2px);
  -webkit-backdrop-filter: blur(10px);
  color: black;
  transition: all 0.3s ease;
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-bottom: 16px;
  padding: 0 10px;
}
 
.last-img {
  width: 59px;
  height: 79px;
  background-color: black;
  border-radius: 4px;
  margin-right: 12px;
  flex-shrink: 0;
}
</style>