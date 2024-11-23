<script lang="ts" setup>
import { onMounted } from "vue";
import { RouterView } from "vue-router";
onMounted(() => {
  window.addEventListener("contextmenu", function (e) {
    e.preventDefault();
  });
});
</script>

<template>
  <router-view v-slot="{ Component }">
    <transition name="scale" mode="out-in">
      <component :is="Component" :key="$route.path" />
    </transition>
  </router-view>
</template>

<style>
.scale-enter-active,
.scale-leave-active {
  opacity: 1;
  transform: scale(1);
  transition: all 0.4s cubic-bezier(0.82, 0, 0.58, 1);
}

.scale-enter-from,
.scale-leave-to {
  opacity: 0;
  transform: scale(0.9);
  /* 缩小效果 */
}
</style>