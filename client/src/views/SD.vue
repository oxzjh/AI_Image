<template>
    <ElPageHeader @back="onBack">
        <template #title>返回</template>
        <template #content>文生图、图生图</template>
    </ElPageHeader>
    <ElDivider />
    <ElContainer>
        <ElMain class="main">
            <div class="flex">
                <ElSelect class="input" v-model="model" placeholder="请选择模型" no-data-text="未检索到可用模型">
                    <ElOption v-for="item in models" :key="item" :value="item" />
                </ElSelect>
                <ElSelect class="input" v-model="taesd" placeholder="Taesd模型" no-data-text="未检索到可用模型">
                    <ElOption v-for="item in taesds" :key="item" :value="item" />
                </ElSelect>
            </div>
            <ElInput class="input" v-model="prompt" :rows="2" type="textarea" placeholder="请输入提示词" />
            <ElInput class="input" v-model="negativePrompt" :rows="2" type="textarea" placeholder="反向提示词" />
        </ElMain>
        <ElAside width="156px">
            <ImageCard ref="inputImage">图生图输入图片</ImageCard>
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
            <ElFormItem label="随机种子：">
                <div class="flex">
                    <ElInputNumber v-model="seed" :min="-1" :max="0x7FFFFFFF" />
                    <ElButton class="btn-refresh" :icon="Refresh" circle @click="onRefreshSeed" />
                </div>
            </ElFormItem>
            <ElFormItem>
                <ElButton class="flex-item" type="primary" @click="onGenerateClick">生成</ElButton>
            </ElFormItem>
            <ElLink v-show="hasCUDA && !useCUDA"
                href="https://github.com/oxzjh/AI_Image/releases/download/SD_CUDA/SD_CUDA.zip" type="warning">
                提示：检测到您的设备支持显卡推理，点击下载CUDA版SD进行高效推理</ElLink>
        </ElMain>
        <ElAside class="output" width="300px">
            <ImageView :src="output" :list="[output]" />
        </ElAside>

    </ElContainer>
</template>

<script setup>
import router from '@/router';
import { ElAside, ElButton, ElContainer, ElDivider, ElFormItem, ElImage, ElInput, ElInputNumber, ElLink, ElMain, ElMessage, ElOption, ElPageHeader, ElSelect, ElSlider } from 'element-plus';
import { onMounted, ref } from 'vue';
import ImageCard from '@/components/ImageCard.vue';
import { Refresh } from '@element-plus/icons-vue';
import { getURL, post, stream } from '@/request';
import ImageView from '@/components/ImageView.vue';

const models = ref([])
const taesds = ref([])
const hasCUDA = ref(false)
const useCUDA = ref(false)

const model = ref("")
const taesd = ref("")
const prompt = ref("")
const negativePrompt = ref("")
const inputImage = ref()
const width = ref(512)
const height = ref(512)
const sampleMethod = ref(0)
const sampleSteps = ref(1)
const cfgScale = ref(1)
const seed = ref(-1)

const output = ref("")

const samplingMethodList = ["euler_a", "euler", "heun", "dpm2", "dpm++2s_a", "dpm++2m", "dpm++2mv2", "lcm"]

function onBack() {
    router.back()
}

onMounted(() => {
    post("/sd/models", {}, result => {
        models.value = result.models
        taesds.value = result.taesd
        hasCUDA.value = result.has_cuda
        useCUDA.value = result.use_cuda
    })
})

function onRefreshSeed() {
    seed.value = Math.floor(Math.random() * 0x7FFFFFFF)
}

function onGenerateClick() {
    if (model.value == "") {
        ElMessage({ message: "请选择模型", type: "error" })
        return
    }
    if (prompt.value == "") {
        ElMessage({ message: "请输入提示词", type: "error" })
        return
    }
    getInputImageData(inputImageData => {
        stream("/sd/generate", {
            model: model.value,
            taesd: taesd.value,
            prompt: prompt.value,
            negative_prompt: negativePrompt.value,
            input_image: inputImageData,
            width: width.value,
            height: height.value,
            sample_method: sampleMethod.value,
            sample_steps: sampleSteps.value,
            cfg_scale: cfgScale.value,
            seed: seed.value
        }, (content, loading) => {
            try {
                const result = JSON.parse(content)
                if (Array.isArray(result)) {
                    output.value = getURL("/sd/download?file=" + result[0] + "&" + Date.now())
                } else {
                    loading.text.value = result.step + "/" + result.steps
                }
            } catch { }
        })
    })
}

function getInputImageData(callback) {
    const file = inputImage.value.getFile()
    if (file) {
        const reader = new FileReader()
        reader.readAsDataURL(file.raw)
        reader.onload = function () {
            const index = reader.result.indexOf(",")
            callback(reader.result.substring(index + 1))
        }
    } else {
        callback("")
    }
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

.image-list {
    display: flex;
    justify-content: center;
}
</style>