<script lang="ts" setup>
import { onMounted } from "vue";
import { RouterView } from "vue-router";
import router from "./router";
import { logger } from "./log";
import { config } from "./app";

onMounted(() => {
    window.addEventListener("contextmenu", function (e) {
        e.preventDefault();
    });
});
onMounted(async () => {
    let isGreeted = (await config.Get()).is_greeted;
    logger.Info("is greeted: " + isGreeted);
    if (isGreeted) router.push("/home");
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
    transition: filter 0s ease, opacity 0.5s ease;
}

.scale-enter-from,
.scale-leave-to {
    opacity: 0;
    filter: blur(10px);
}
</style>
