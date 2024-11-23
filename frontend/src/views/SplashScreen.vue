<script lang="ts" setup>
import { onBeforeMount, onMounted, reactive, ref, shallowRef } from "vue";
import Greetings from "./Greet.vue";
import Button from "../components/Button.vue";
import Initiating from "../components/Initiating.vue";
import InitView from "./Init.vue";
onMounted(() => {
    window.addEventListener("contextmenu", function (e) {
        e.preventDefault();
    });
});

const activeComponent = shallowRef(Greetings)

function toInit() {
    activeComponent.value = InitView;
    const FBtn = document.querySelector(".FloatBtn") as HTMLElement | null;
    const FContainer = document.querySelector(".Float") as HTMLElement | null;
    if (FBtn && FContainer) {
        // 获取按钮的位置
        const rect = FBtn.getBoundingClientRect();
        const distanceToBottom = window.innerHeight - rect.bottom;

        // 将计算出的距离应用为样式的 bottom 值
        FContainer.style.position = 'absolute';
        FContainer.style.bottom = `${distanceToBottom}px`;
    }
}
</script>

<template>
    <section class="Container">
        <Transition name="scale">
            <component :is="activeComponent" />
        </Transition>
        <div class="BtnContainer" id="btn" />

        <div ref="FloatContainer" class="Float">
            <Initiating class="toInitBall" />
            <Button class="FloatBtn" @click="toInit()" />
        </div>
        <!-- <RouterLink to="/home"><Button id="btn" /></RouterLink> -->
    </section>
</template>

<style scoped>
.scale-enter-active,
.scale-leave-active {
    opacity: 1;
    filter: blur(0);
    transition: all .7s cubic-bezier(0.82, 0, 0.58, 1);
}

.scale-enter-from,
.scale-leave-to {
    opacity: 0;
    filter: blur(10px);
}

.Container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.BtnContainer {
    display: flex;
    position: absolute;
    height: 80px;
    width: 80px;
    background-color: transparent;
}

.Float {
    display: flex;
}

.FloatBtn {
    position: relative;
    z-index: 1;
}

.toInitBall {
    position: absolute;
    z-index: -1;
}
</style>
