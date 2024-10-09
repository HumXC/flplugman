<script lang="ts" setup>
import { onMounted } from "vue";
import Greetings from "./components/Greetings/Greetings.vue";
import { GetConfig } from "../wailsjs/go/main/App";

// 阻止右键菜单
onMounted(() => {
  window.addEventListener("contextmenu", function (e) {
    e.preventDefault();
  });
});
const is_greeted =( await GetConfig()).is_greeted;
</script>

<template>
  <div v-if="!is_greeted">
      <Greetings />
  </div>
</template>

<style></style>
