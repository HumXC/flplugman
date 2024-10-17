<script lang="ts" setup>
import { onMounted, ref } from "vue";
const props = defineProps<{ bitsize: number }>();
const colors = (bitsize: number) => {
    switch (bitsize) {
        case 32:
            return "rgb(207, 196, 48)";
        case 64:
            return "rgb(48, 133, 207)";
        case 96:
            return "rgb(179, 64, 140)";
        default:
            return "gray";
    }
};
const color = ref(colors(props.bitsize));
const bitsize = ref("?");
onMounted(() => {
    if (props.bitsize === 96) {
        bitsize.value = "32/64";
    } else {
        bitsize.value = props.bitsize.toString();
    }
});
</script>

<template>
    <div class="bitsize">
        <span>{{ bitsize }}</span>
    </div>
</template>

<style scoped>
.bitsize {
    display: inline-block;
    background-color: v-bind(color);
    padding: 2px 6px;
    border-radius: 6px;
}
</style>
