<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import Greetings from './components/Greetings/Greetings.vue';
import * as backend from '../wailsjs/go/main/App'

const IsGreeted = ref(true)
// 阻止右键菜单
onMounted(() => {
  window.addEventListener('contextmenu', function (e) {
    e.preventDefault();
  });
})
</script>

<template>
  <Greetings v-if="IsGreeted" />
</template>

<style></style>