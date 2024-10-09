<script lang="ts" setup>
import { onMounted, reactive, ref } from "vue";
import Greetings from "./components/Greetings/Greetings.vue";
import { GetConfig } from "../wailsjs/go/main/App";
import { LogInfo } from "../wailsjs/runtime/runtime";

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
    <div>
        <Greetings v-if="!isGreeted" />
    </div>
</template>

<style></style>
