<template>
  <a-modal v-model:visible="showAiChatWeb" @ok="handlerClose" :footer="false" fullscreen>
    <template #title>
       AI Chat
    </template>
    <iframe
        :src="externalUrl"
        style="width: 100%; height: 90vh;"
        frameborder="0"
        allow="microphone">
    </iframe>
  </a-modal>
</template>

<script setup>
import {ref} from "vue";

const externalUrl = ref("https://xxx")

const isConfig = ref(false)

const showAiChatWeb = ref(false)
const openModal = () => {
  showAiChatWeb.value = true;
}

function handlerClose () {
  showAiChatWeb.value = false;
}

defineExpose({
  openModal,
})
</script>

<style scoped>

</style>