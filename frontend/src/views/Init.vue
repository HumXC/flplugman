<template>
    <div>
        <main id="Main">
            <h2 id="Title">{{ Init.title }}</h2>
            <div id="Context">
                <h1 id="Message">{{ Init.choose }}</h1>
                <span style="font-size: 24px; font-weight: 500;">{{ Init.chooseRoute }}</span>
                <div class="button" @click="chooseDir()">
                    <p>{{ Init.chooseDir }}</p>
                </div>
            </div>
        </main>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import Dialog from '../components/Dialog.vue';
import { utils } from '../app'

const Init = ref({
    title: '现在，进行一些初始设置...',
    choose: '选择你的 FL Studio 用户数据目录',
    chooseTip: '你已选择以下的路径，它没问题吗？',
    chooseDir: '选择文件夹',
    chooseRoute: '',
});

async function chooseDir() {
    utils.ChooseDir()
        .then(dir => {
            Init.value.title = ''
            Init.value.choose = Init.value.chooseTip;
            Init.value.chooseRoute = dir;
            Init.value.chooseDir = '重新选择';
        })
        .catch(err => {
            Init.value.chooseRoute = '错误';
            Dialog.message = err;
        });
}

function StartUp() {
    const Main = document.getElementById('Main') as HTMLElement;
    const Title = document.getElementById('Title') as HTMLElement;
    const Context = document.getElementById('Context') as HTMLElement;

    if (!Main || !Context || !Title) return;

    const SectionHeight = Context.offsetHeight - 32;

    // 防止与 Vue Router 的组件过渡动画冲突

    Main.style.transition = 'none';

    // Init 位置
    Main.style.transform = `translateY(${SectionHeight}px)`;

    Title.animate([{

    }, {
        fontSize: '24px',
    }], {
        duration: 1000,
        easing: "cubic-bezier(0.7, -0.01, 0.15, 1.03)",
        fill: "forwards",
        delay: 1000,
    }

    );

    Main.animate(
        [
            { transform: `translateY(${SectionHeight}px)` },
            { transform: `translateY(0)` },
        ],
        {
            duration: 1000,
            easing: "cubic-bezier(0.7, -0.01, 0.15, 1.03)",
            fill: "forwards",
            delay: 1000,
        }
    );

    setTimeout(() => {
        Main.style.transition = '';
    }, 1000);
}

onMounted(() => {
    StartUp();
})
</script>

<style scoped>
main,
div {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    color: #fff;
    gap: 16px;
}

main {
    height: 100vh;
    width: 100vw;
}

* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

h1,
h2 {
    font-size: 42px;
}

@keyframes FadeUp {
    0% {
        opacity: 0;
        transform: translateY(100%);
        filter: blur(10px);
    }

    100% {
        opacity: 1;
        transform: translateY(0);
        filter: blur(0);
    }
}

#Context {
    opacity: 0;
    animation: FadeUp 1s cubic-bezier(0.7, -0.01, 0.15, 1.03) forwards;
    animation-delay: 1s;
}

.button {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    height: 40px;
    width: auto;
    background-color: #3d3d3d;
    border-radius: 80px;
    box-shadow: inset 0 0 6px rgba(255, 255, 255, 0.55);
    color: #fff;
    cursor: pointer;
    font-size: 1.5rem;
    transition: transform 0.5s cubic-bezier(0.7, -0.01, 0.15, 1.03);
    background: radial-gradient(circle at center, #2409bb, #5121ff, #8143ff, #3258ff, #6200ff);
    background-size: 660px 660px;
    --anime-speed: 4s;
    -webkit-animation: FadeGradient var(--anime-speed) linear infinite;
    -moz-animation: FadeGradient var(--anime-speed) linear infinite;
    animation: FadeGradient var(--anime-speed) linear infinite;
}

@keyframes FadeGradient {
    0% {
        background-position: 0% 0px;
    }

    100% {
        background-position: 660px -660px;
    }
}

.button p {
    font-size: 16px;
    margin: 0 20px 0 20px;
    font-weight: 600;
}
</style>