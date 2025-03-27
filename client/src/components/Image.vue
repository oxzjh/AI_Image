<template>
    <ElImage :src="src" :preview-src-list="[src]">
        <template #error>
            <ElIcon>
                <Picture></Picture>
            </ElIcon>
        </template>
    </ElImage>
</template>

<script setup>
import { ElIcon, ElImage } from 'element-plus';
import { Picture } from '@element-plus/icons-vue';

defineProps(["src"])
</script>

<style scoped>
.el-image {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    min-height: 100px;
    background-color: var(--el-fill-color-light);
    color: var(--el-text-color-secondary);
    font-size: 30px;
}
</style>