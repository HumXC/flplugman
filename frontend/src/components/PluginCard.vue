<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { plugin } from "../app";
import { logger } from "../log";
import Bitsize from "./Bitsize.vue";

const props = defineProps<{ plugin: plugin.Plugin }>();
const dist = ref(props.plugin.PresetPath);

async function movePreset() {
    try {
        await props.plugin.MoveTo(dist.value);
    } catch (e) {
        const err = e as string;
        if (err.endsWith("The filename, directory name, or volume label syntax is incorrect.")) {
            logger.Error("文件名不合法");
        } else {
            logger.Error(err);
        }
    }
}
</script>

<template>
    <div class="card">
        <div class="top">
            <span class="name">{{ props.plugin.Name }}</span>
            <div class="category-container">
                <span class="category" v-for="c in props.plugin.Category">{{ c }}</span>
            </div>
            <span class="tip">{{ props.plugin.Nfo.Tip }}</span>
        </div>
        <div class="bottom">
            <div class="bottom-left">
                <Bitsize :bitsize="props.plugin.Bitsize" />
            </div>
            <span class="vendorname">{{ props.plugin.Vendorname }}</span>
        </div>
    </div>
</template>

<style scoped>
/* card 分成 3 部分，垂直排列，中间区域为主要区域 */
.card {
    display: flex;
    flex-direction: column;
    align-items: start;
    justify-content: flex-start;
    height: 130px;
    max-width: 350px;
    border-radius: 10px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    background-color: #ffffff38;
    padding: 10px;
}

.top {
    display: flex;
    flex-direction: column;
    align-items: start;
    flex-grow: 1;
    width: 100%;
}
.tip {
    padding-left: 3px;
    font-size: 12px;
}
.name {
    font-size: 18px;
    margin-bottom: 10px;
}
.category-container {
    display: flex;
    flex-wrap: wrap;
    margin-bottom: 10px;
    flex-direction: row;
    gap: 5px;
    width: 80%;
}
.category {
    background-color: #8383838c;
    padding: 2px 8px;
    font-size: 14px;
    border-radius: 12px;
}
.bottom {
    display: flex;
    justify-content: center;
    align-items: center;
    display: flex;
    flex-direction: row;
    width: 100%;
}
.bottom-left {
    display: flex;
    flex-grow: 1;
    align-items: start;
    justify-content: start;
}

.vendorname {
    font-size: 14px;
    color: #acacac;
}
</style>
