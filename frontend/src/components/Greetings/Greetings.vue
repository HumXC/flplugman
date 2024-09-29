<script lang="ts" setup>
import { onMounted, reactive } from 'vue'
import * as backend from '../../../wailsjs/go/main/App'
import Button from './Button.vue';
const Logo = {
  Src: "/src/assets/images/logo.jpg",
  Alt: "Logo"
}

const Greetings = {
  Welcome: "欢迎使用",
  Title: "FLPluginMan"
}

// const data = reactive({
//   name: "",
//   resultText: "Please enter your name below 👇",
// })

// function greet() {
//   backend.Greet(data.name).then(result => {
//     data.resultText = result
//   })
// }
function RollingUp() {
  const btn = document.getElementById('btn') as HTMLElement | null;
  const text = document.querySelector('.TextContainer') as HTMLElement | null;
  const ctn = document.querySelector('.Container') as HTMLElement | null;

  if (!btn || !text || !ctn) return;
  const MainHeight = btn.clientHeight + text.clientHeight;
  const SectionHeight = MainHeight / 2;

  // Init
  ctn.style.transform = `translateY(${SectionHeight}px)`;

  ctn.animate([
    { transform: `translateY(${SectionHeight}px)` },
    { transform: `translateY(0)` }
  ], {
    duration: 1000,
    easing: 'cubic-bezier(0.7, -0.01, 0.15, 1.03)',
    fill: 'forwards',
    delay: 1000
  });
}
onMounted(() => {
  RollingUp();
});


</script>

<template>
  <section class="Container">
    <div class="LogoContainer">
      <img :src="Logo.Src" :alt="Logo.Alt">
    </div>
    <div ref="text" class="TextContainer">
      <h2 id="welcome">{{ Greetings.Welcome }}</h2>
      <h1 id="title">{{ Greetings.Title }}</h1>
    </div>
    <!-- <button id="btn" class="animated-gradient">
      →
    </button> -->
    <Button id="btn" />
  </section>
</template>

<style scoped>
@keyframes StartUp {
  0% {
    transform: scale(0);
  }

  100% {
    transform: scale(1);
  }
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

@keyframes FadeIn {
  0% {
    opacity: 0;
  }

  100% {
    opacity: 1;
  }
}

.Container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.LogoContainer {
  display: flex;
  justify-content: center;
  align-items: center;
  height: auto;
  animation: StartUp 1s cubic-bezier(0.7, -0.01, 0.15, 1.03);
}

.LogoContainer img {
  height: 200px;
  width: 200px;
  background-color: transparent;
  border-radius: 100%;
}

.TextContainer {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin: 30px 0 30px 0;
}

.TextContainer h1 {
  opacity: 0;
  margin: 0;
  font-size: 64px;
  font-weight: 400;
  letter-spacing: -4px;
  animation: FadeUp 1s cubic-bezier(0.7, -0.01, 0.15, 1.03);
  animation-delay: 1.1s;
  animation-fill-mode: forwards;
  overflow: hidden;
}

@keyframes Scan {
  0% {
    left: -20%;
    background-size: 0 2px;
  }

  50% {
    left: 0;
    background-size: 100% 2px;
  }

  100% {
    left: 100%;
    background-size: 0 2px;
  }
}

.TextContainer h1::after {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(to right, #FFFFFF, #FFFFFF) no-repeat left bottom;
  background-size: 0 2px;
  animation: Scan 1s cubic-bezier(0.7, -0.01, 0.15, 1.03);
  animation-delay: 2.1s;
  transition: background-size 0.5s;
}

.TextContainer h2 {
  opacity: 0;
  margin: 0;
  font-size: 24px;
  font-weight: 600;
  animation: FadeUp 1s cubic-bezier(0.7, -0.01, 0.15, 1.03);
  animation-delay: 1s;
  animation-fill-mode: forwards;
  overflow: hidden;
}
</style>
