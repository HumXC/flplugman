<script lang="ts" setup>
import { onMounted } from 'vue';
import Greetings from './components/Greetings/Greetings.vue';

// 阻止右键菜单
onMounted(() => {
  window.addEventListener('contextmenu', function (e) {
    e.preventDefault();
  });
})
</script>

<template>
  <Greetings />
</template>

<style></style>