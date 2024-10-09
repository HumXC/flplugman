<script lang="ts" setup>
import { onMounted, reactive, ref } from "vue";
import Greetings from "./components/Greetings/Greetings.vue";
import { GetConfig } from "../wailsjs/go/main/App";
import { LogInfo } from "../wailsjs/runtime/runtime";

const IsGreeted = ref(true)
// 阻止右键菜单
onMounted(() => {
  window.addEventListener("contextmenu", function (e) {
    e.preventDefault();
  });
});

const isGreeted = ref(false);
onMounted(async () => {
  isGreeted.value = (await GetConfig()).is_greeted;
  LogInfo("isGreeted: " + isGreeted);
});
</script>

<template>
  <<<<<<< HEAD <Greetings v-if="IsGreeted" />
  =======
  <div>
    <Greetings v-if="!isGreeted" />
  </div>
  >>>>>>> e104659dc69116a953d877ff1fce92d8590bea3b
</template>

<style></style>
