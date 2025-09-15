<script setup>
import { ref, onMounted,inject } from 'vue';
import { ElMessage } from 'element-plus';
import { Minus, FullScreen, Close } from '@element-plus/icons-vue';
import { WindowToggleMaximise,WindowMinimise,Quit} from 'wailsjs/runtime';
import {
  InitConfig,
  GetConfig,
  GetLang
} from 'wailsjs/go/service/App';

const LangInfo = inject('LangInfo');
const Config = inject('Config');
const loading = ref(true);
const initSuccess = ref(false);

onMounted(async () => {
  try {
    await InitConfig();
    initSuccess.value = true;
  } catch (e) {
    console.error('InitConfig() 出错:', e);
    ElMessage.error('初始化失败');
    return;
  } finally {
    loading.value = false;
  }

  try {
    const resultLang = await GetLang();
    LangInfo.value = JSON.parse(resultLang);
  } catch (error) {
    console.error('JSON 解析失败 (语言数据):', error);
    ElMessage.error("语言数据解析失败!");
  }

  try {
    const resultConfig = await GetConfig();
    Config.value = JSON.parse(resultConfig);
  } catch (error) {
    console.error('JSON 解析失败 (配置数据):', error);
    ElMessage.error("配置数据解析失败!");
  }
});

defineExpose({
  Config,
  LangInfo
})
</script>

<template>
    <!-- 上边栏 -->
    <div class="top-bar" @dblclick="WindowToggleMaximise">
      <!-- 左侧 Logo 和标题 -->
      <div class="left-panel" v-if="Config">
        <span class="el-text">MGL·{{Config.Version}}</span>
      </div>
 
      <!-- 右侧窗口控制按钮 -->
      <div class="right-panel">
        <el-button circle size="small" @click="WindowMinimise">
          <el-icon><Minus /></el-icon>
        </el-button>
        <el-button circle size="small" @click="WindowToggleMaximise">
          <el-icon><FullScreen /></el-icon>
        </el-button>
        <el-button circle size="small" @click="Quit">
          <el-icon><Close /></el-icon>
        </el-button>
      </div>
    </div>
</template>

<style scoped>
.window-frame {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #f5f7fa;
  font-family: sans-serif;
  user-select: none;
}
 
.top-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 12px;
  height: 30px;
  background-color: #fff;
  border-bottom: 1px solid #e4e7ed;
  --wails-draggable:drag
}
 
.left-panel {
  display: flex;
  align-items: center;
}
 
.left-panel .logo {
  height: 24px;
  margin-right: 8px;
}
 
.right-panel {
  display: flex;
  align-items: center;
  gap: 8px;
  -webkit-app-region: no-drag;
}
 
.right-panel button {
  width: 24px;
  height: 24px;
  padding: 0;
  border: none;
  background-color: transparent;
  color: #666;
  transition: all 0.2s ease;
}
 
.right-panel button:hover {
  background-color: rgba(0, 0, 0, 0.1);
}
 
.main-content {
  flex: 1;
  overflow: auto;
  padding: 16px;
}
</style>