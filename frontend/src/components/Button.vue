<template>
    <div class="ButtonContainer">
        <div class="button" @click="handleClick">
            <div class="arrow">
                <svg t="1727575500694" viewBox="0 0 1000 1000" version="1.1" xmlns="http://www.w3.org/2000/svg"
                    p-id="1494" xmlns:xlink="http://www.w3.org/1999/xlink">
                    <path
                        d="M227.986 584.688l257.492 257.492c20.11 20.11 52.709 20.11 72.819 0l257.492-257.492c20.11-20.11 20.11-52.709 0-72.819s-52.709-20.11-72.819 0l-169.585 169.585v-493.664c0-28.453-23.046-51.499-51.499-51.499s-51.499 23.046-51.499 51.499v493.664l-169.585-169.585c-10.042-10.043-23.226-15.089-36.41-15.089s-26.367 5.021-36.41 15.089c-20.11 20.11-20.11 52.709 0 72.819z"
                        p-id="1495"></path>
                </svg>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
function handleClick() {
    console.log("Button clicked");
}
import * as backend from '../../wailsjs/go/main/App'
</script>

<style scoped>
@keyframes BlurOut {
    0% {
        filter: blur(30px);
        opacity: 0;
    }

    100% {
        filter: blur(0);
        opacity: 1;
        box-shadow: inset 0 0 10px rgba(133, 133, 133, 0.55);
    }
}

.ButtonContainer {
    display: flex;
    position: relative;
    height: 80px;
    width: 80px;
    background-color: transparent;
}

.button {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    opacity: 0;
    position: absolute;
    height: 80px;
    width: 80px;
    background-color: #3d3d3d;
    border-radius: 100%;
    color: #fff;
    cursor: pointer;
    font-size: 1.5rem;
    animation: BlurOut 2s cubic-bezier(0.7, -0.01, 0.15, 1.03);
    animation-delay: 2s;
    animation-fill-mode: forwards;
}

.button .arrow {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    height: 30px;
    width: 30px;
    z-index: 2;
    pointer-events: none;
}

.button .arrow svg {
    fill: #fff;
    height: 100%;
    width: 100%;
    transform: rotate(-90deg);
}

.button::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 80px;
    height: 80px;
    border-radius: 50px;
    z-index: 1;
    transform: scale(0);
    box-shadow: inset 0 0 10px rgba(255, 255, 255, 0.55);
    transition: transform 0.5s cubic-bezier(0.7, -0.01, 0.15, 1.03);
    background: radial-gradient(circle at center, #2409bb, #5121ff, #8143ff, #3258ff, #6200ff);
    background-size: 660px 660px;
    color: #1d2125;
    --anime-speed: 4s;
    -webkit-animation: FadeGradient var(--anime-speed) linear infinite;
    -moz-animation: FadeGradient var(--anime-speed) linear infinite;
    animation: FadeGradient var(--anime-speed) linear infinite;
}

.button:hover::before {
    transform: scale(1.2);
    transition: transform 0.3s cubic-bezier(0.7, -0.01, 0.15, 1.03);
}

.button:active::before {
    transform: scale(1.1);
    transition: transform 0.1s cubic-bezier(0.7, -0.01, 0.15, 1.03);
}

@keyframes FadeGradient {
    0% {
        background-position: 0% 0px;
    }

    100% {
        background-position: 660px -660px;
    }
}
</style>