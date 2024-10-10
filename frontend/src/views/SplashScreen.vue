<script lang="ts" setup>
import { onMounted, reactive, ref } from "vue";
import Greetings from "./Greet.vue";
import { GetConfig } from "../../wailsjs/go/main/App";
import { LogInfo } from "../../wailsjs/runtime/runtime";
import Button from "../components/Button.vue";
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
    <section class="Container">
        <Greetings v-if="!isGreeted" />
        <RouterLink to="/home"><Button id="btn" /></RouterLink>
    </section>
</template>

<style scoped>
.Container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
}
</style>
