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

<style scoped>
.scale-enter-active,
.scale-leave-active {
  opacity: 1;
  filter: blur(0px);
  transition: filter 0s ease, opacity .5s ease;
}

.scale-enter-from,
.scale-leave-to {
  opacity: 0;
  filter: blur(10px);
}
</style>