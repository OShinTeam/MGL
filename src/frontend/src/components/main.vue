<script setup lang="ts">
import Side from 'components/side.vue';
</script>

<template>
    <el-container style="height: 100vh;">
      <Side />
      <router-view />
    </el-container>
</template>