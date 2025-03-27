import Basic from "@/views/Basic.vue";
import SD from "@/views/SD.vue";
import { createRouter, createWebHashHistory } from "vue-router";

export default createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: '/',
            component: Basic,
        },
        {
            path: '/sd',
            component: SD
        }
    ]
})