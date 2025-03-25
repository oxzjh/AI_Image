<template>
    <ElPageHeader @back="onBack">
        <template #title>返回</template>
        <template #content>文生图、图生图</template>
    </ElPageHeader>
    <ElDivider />
    <ElContainer>
        <ElMain class="main">
            <ElSelect class="input" placeholder="请选择模型">
            </ElSelect>
            <ElInput class="input" v-model="prompt" :rows="2" type="textarea" placeholder="请输入提示词" />
            <ElInput class="input" v-model="negativePrompt" :rows="2" type="textarea" placeholder="反向提示词" />
        </ElMain>
        <ElAside width="156px">
            <ImageCard />
        </ElAside>
    </ElContainer>
    <ElContainer>
        <ElMain class="main">
            <div class="flex">
                <ElFormItem class="flex-item" label="宽：">
                    <ElInputNumber v-model="width" :min="256" :max="1024" :step="256" controls-position="right" />
                </ElFormItem>
                <ElFormItem class="flex-item" label="高：">
                    <ElInputNumber v-model="height" :min="256" :max="1024" :step="256" controls-position="right" />
                </ElFormItem>
            </div>
            <ElFormItem label="采样方法：">
                <ElSelect v-model="sampleMethod">
                    <ElOption v-for="method, index in samplingMethodList" :key="method" :value="index"
                        :label="method" />
                </ElSelect>
            </ElFormItem>
            <ElFormItem label="采样步数：">
                <ElSlider v-model="sampleSteps" :min="1" :max="40" show-input />
            </ElFormItem>
            <ElFormItem label="CFG Scale:">
                <ElSlider v-model="cfgScale" :min="0.1" :max="20" :step="0.1" show-input />
            </ElFormItem>
            <div class="flex">
                <ElFormItem class="flex-item" label="生成数量：">
                    <ElInputNumber v-model="batchCount" :min="1" :max="4" />
                </ElFormItem>
                <ElFormItem class="flex-item" label="随机种子：">
                    <div class="flex">
                        <ElInputNumber v-model="seed" :min="-1" :max="0x7FFFFFFF" />
                        <ElButton class="btn-refresh" :icon="Refresh" circle @click="onRefreshSeed" />
                    </div>
                </ElFormItem>
            </div>
            <ElFormItem >
                <ElButton class="flex-item" type="primary">生成</ElButton>
            </ElFormItem>
        </ElMain>
        <ElAside class="output" width="300px">aside</ElAside>
    </ElContainer>
</template>

<script setup>
import router from '@/router';
import { ElAside, ElButton, ElContainer, ElDivider, ElFormItem, ElInput, ElInputNumber, ElMain, ElOption, ElPageHeader, ElSelect, ElSlider } from 'element-plus';
import { ref } from 'vue';
import ImageCard from '@/components/ImageCard.vue';
import { Refresh } from '@element-plus/icons-vue';

const prompt = ref("")
const negativePrompt = ref("")
const width = ref(512)
const height = ref(512)
const sampleMethod = ref(0)
const sampleSteps = ref(1)
const cfgScale = ref(1)
const batchCount = ref(1)
const seed = ref(-1)

const samplingMethodList = ["euler_a", "euler", "heun", "dpm2", "dpm++2s_a", "dpm++2m", "dpm++2mv2", "lcm"]

function onBack() {
    router.back()
}

function onRefreshSeed() {
    seed.value = Math.floor(Math.random() * 0x7FFFFFFF)
}
</script>

<style scoped>
.main {
    padding: 0 10px 0 0;
}

.input {
    padding-bottom: 5px;
}

.flex {
    display: flex;
}

.flex-item {
    width: 100%;
}

.btn-refresh {
    margin-left: 5px;
}

.output {
    border: 1px dashed gray;
}
</style>