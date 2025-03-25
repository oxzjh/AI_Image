import Basic from "@/views/Basic.vue";
import SD from "@/views/SD.vue";
import { createRouter, createWebHistory } from "vue-router";

export default createRouter({
    history: createWebHistory(),
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