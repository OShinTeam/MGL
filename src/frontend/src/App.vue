<script setup>
import HEADER from 'components/header.vue'
import MAIN from 'components/main.vue'

import {ref,provide } from 'vue'
const LangInfo = ref({});
const Config = ref({});
const GameInfo = ref({});
const Game = ref({});
const UserInfo = ref({});

provide('LangInfo', LangInfo);
provide('Config', Config);
provide('GameInfo', GameInfo);
provide('Game', Game);
provide('UserInfo', UserInfo);
</script>

<template>
  <div id="app" class="window-frame">
    <HEADER />
    <MAIN />
  </div>
</template>

<style>
.body{
  font-family:"宋体","Arial Narrow";
}

</style>