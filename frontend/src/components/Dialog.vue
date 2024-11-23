<template>
    <main>
        <div class="Container">
            <div class="Header">
                <p>
                    {{ title }}
                </p>
                <div class="Close">
                    <svg width="15" height="14" viewBox="0 0 15 14" fill="#007AFF" xmlns="http://www.w3.org/2000/svg">
                        <path
                            d="M1.32812 13.4922C0.875 13.0469 0.890625 12.2578 1.3125 11.8359L5.78125 7.36719L1.3125 2.91406C0.890625 2.48438 0.875 1.70312 1.32812 1.25C1.78125 0.789062 2.57031 0.804688 2.99219 1.23438L7.45312 5.69531L11.9141 1.23438C12.3516 0.796875 13.1172 0.796875 13.5703 1.25C14.0312 1.70312 14.0312 2.46875 13.5859 2.91406L9.13281 7.36719L13.5859 11.8281C14.0312 12.2734 14.0234 13.0312 13.5703 13.4922C13.125 13.9453 12.3516 13.9453 11.9141 13.5078L7.45312 9.04688L2.99219 13.5078C2.57031 13.9375 1.78906 13.9453 1.32812 13.4922Z">
                        </path>
                    </svg>
                </div>
            </div>
            <p>
                {{ message }}
            </p>
        </div>
    </main>
</template>

<script lang="ts" setup>
defineProps({
    title: String,
    message: String
})
</script>

<style scoped>
main {
    position: absolute;
    top: -260px;
}

.Container {
    width: 240px;
    height: auto;
    min-height: 100px;
    border-radius: 10px;
    border: 2px solid rgba(241, 247, 251, 0.10);
    text-align: left;
}

.Container .Header {
    display: flex;
    justify-content: space-between;
    height: auto;
    width: 100%;
    align-items: center;
}

.Container .Header p {
    margin: 20px auto 0 20px;
    font-size: 16px;
    font-weight: 600;
}

.Container .Header .Close {
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 20px 20px 0 auto;
    height: 20px;
    width: 20px;
    fill: #5121ff;
}

.Container p {
    opacity: 0.8;
    margin: 0 20px 20px 20px;
    font-size: 14px;
}
</style>