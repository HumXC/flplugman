<script lang="ts" setup>
import Greetings from "./Greet.vue";
import Button from "../components/Button.vue";
</script>

<template>
    <section class="Container">
        <Greetings />
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
