<template>
    <div>
        <ElImage :src="src" :preview-src-list="list">
            <template #error>
                <ElIcon>
                    <Picture></Picture>
                </ElIcon>
            </template>
        </ElImage>
    </div>
    <div>
        <ElButton v-show="src != ''" @click="onDownload">下载</ElButton>
    </div>
</template>

<script setup>
import download from 'downloadjs';
import { ElButton, ElImage } from 'element-plus';
import { Picture } from '@element-plus/icons-vue';
import { getBlob } from '@/request/http';

const props = defineProps(["src", "list"])

function onDownload() {
    getBlob(props.src, blob => {
        download(blob, "output.jpg")
    })
}

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