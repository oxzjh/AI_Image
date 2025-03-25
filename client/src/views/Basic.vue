<template>
    <div class="flex">
        <div class="box">
            <input type="file" accept=".jpg,.jpeg,.png,.webp,.tif,.tiff" @change="onChange">
            <Image ref="source" :src="imageURL" />
            <div class="flex">
                <ElSelect v-model="handler" class="option">
                    <ElOptionGroup v-for="item in data" :label="item.label" :key="item.id">
                        <ElOption v-for="h in item.handlers" :label="h.label" :value="h.value" :key="h.value" />
                    </ElOptionGroup>
                </ElSelect>
                <ElCheckbox v-model="extra" class="option" v-if="handler == 'headseg' || handler == 'line'">{{ handler
                    == 'headseg' ?
                    '填充白色' : '增强模式' }}</ElCheckbox>
                <ElButton :disabled="imageURL == '' || handler == ''" @click="onGenerate">生成</ElButton>
            </div>
        </div>
        <div class="box right">
            <RouterLink to="/sd">文生图，图生图➜</RouterLink>
            <Image :src="generateURL" />
            <!-- <ElButton :disabled="generateURL == ''" @click="onDownload">下载</ElButton> -->
        </div>
    </div>
</template>

<script setup>
import Image from '@/components/Image.vue'
import { getURL, post } from '@/request'
import { ElButton, ElCheckbox, ElOption, ElOptionGroup, ElSelect } from 'element-plus';
import { ref } from 'vue';

const source = ref()
const imageURL = ref("")
const generateURL = ref("")
const handler = ref("")
const extra = ref(false)

const data = [
    { id: "upscaler", label: "", handlers: [{ label: "高清扩图", value: "upscale" }] },
    { id: "headseg", label: "", handlers: [{ label: "人物头部提取", value: "headseg" }] },
    { id: "cartoon", label: "", handlers: [{ label: "头像卡通化", value: "cartoon" }] },
    { id: "line", label: "", handlers: [{ label: "线稿提取", value: "line" }] },
    {
        id: "colorizer", label: "黑白图上色", handlers: [
            { label: "", value: "eccv16" },
            { label: "", value: "siggraph17" }
        ]
    },
    {
        id: "style", label: "风格转换", handlers: [
            { label: "", value: "candy" },
            { label: "", value: "mosaic" },
            { label: "", value: "pointilism" },
            { label: "", value: "princess" },
            { label: "", value: "udnie" }
        ]
    }
]

function onChange(e) {
    const reader = new FileReader()
    reader.readAsDataURL(e.target.files[0])
    reader.onload = function () {
        imageURL.value = reader.result
    }
}

function onGenerate() {
    generateURL.value = ""
    const index = imageURL.value.indexOf(",")
    const data = { image: imageURL.value.substring(index + 1) }
    let route = "/image/" + handler.value
    if (handler.value == "headseg") {
        data.fill_white = extra.value
    } else if (handler.value == "line") {
        data.enhance = extra.value
    } else if (handler.value == "eccv16" || handler.value == "siggraph17") {
        route = "/image/coloring"
        data.method = handler.value
    } else if (handler.value == "candy"
        || handler.value == "mosaic"
        || handler.value == "pointilism"
        || handler.value == "princess"
        || handler.value == "udnie"
    ) {
        route = "/image/transfer"
        data.style = handler.value
    }
    post(route, data, () => {
        generateURL.value = getURL("/image/download?"+Date.now())
        console.log(generateURL.value)
    })
}

</script>

<style scoped>
.flex {
    display: flex;
}

.box {
    width: 100%;
    margin: 0 5px;
}

.option {
    margin-right: 10px;
}

.right {
    text-align: right;
}
</style>