<template>
    <ElUpload :class="{ 'eluploadlimited': limited }" accept=".jpg,.jpeg.png,.webp,.tif,.tiff" list-type="picture-card"
        :auto-upload="false" :on-preview="onPreview" :on-change="onChange" :limit="1" :on-remove="onRemove">
    </ElUpload>
    <ElDialog v-model="previewVisible" width="80%">
        <ElImage :src="previewImageURL" />
    </ElDialog>
</template>

<script setup>
import { ElDialog, ElImage, ElUpload } from 'element-plus';
import { ref } from 'vue';

const limited = ref(false)
const previewImageURL = ref("")
const previewVisible = ref(false)

let _file = null

defineExpose({
    getFile: () => _file
})

function onPreview(file) {
    previewImageURL.value = file.url
    previewVisible.value = true
}

function onChange(file, fileList) {
    _file = file
    limited.value = fileList.length >= 1
}

function onRemove() {
    limited.value = false
}

</script>

<style>
.eluploadlimited .el-upload--picture-card {
    display: none;
}
</style>