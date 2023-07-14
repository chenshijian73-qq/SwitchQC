<template>
  <a-modal v-model:visible="showAiChatWeb"
           @ok="closeAiChatWeb"
           :body-style="{padding: 0, height: '100%',}"
           :footer="false"
           fullscreen >
    <template #title>
      <a-space >
        <a-button type="primary" shape="circle" size="mini" @click="openLinkManage">
          <icon-settings />
        </a-button>
        AI Chat
        <a-select ref="selectorRef"
                  v-model="selectedValue"
                  :style="{width:'150px'}"
                  placeholder="Select ..."
                  @change="selectedUrl"
                  @popup-visible-change="getLinks"
                  allow-search show-footer-on-empty>
          <a-option v-for="(link) in links" :value="link.ID">
            <a-tooltip :content="link.Link" background-color="#3491FA" position="tl">
              <a-space>
                <span class="select-text"> {{ link.Name }} </span>
              </a-space>
            </a-tooltip>
          </a-option>
          <template #footer>
            <div style="padding: 6px 0; text-align: center; background-color: #1a62f2">
              <a-button type="primary" shape="round" @click="openAddLinkForm">
                <icon-plus />
              </a-button>
            </div>
          </template>
        </a-select>
      </a-space>
    </template>
    <iframe
        v-if="externalUrl"
        :src="externalUrl"
        style="width: 100%; height: 99%; border: none; padding: 0px"
        frameborder="0"
        scrolling="no"
        allow="microphone">
    </iframe>
    <a-empty v-else>
      <template #image>
        <icon-exclamation-circle-fill />
      </template>
      No Link Selected, please select!
    </a-empty>
  </a-modal>
  <LinkManage :getLinks="getLinks" :links="links" :openAddLinkForm="openAddLinkForm" ref="LinkManageRef"/>
  <AddLinkDrawer :getLinks="getLinks" ref="AddLinkRef"/>
</template>

<script setup>
import '../../assets/ai-chat-web.css'
import {onMounted, ref} from "vue";
import {GetAiLinkList} from "../../../wailsjs/go/logic/AiLinkLogic";
import {IconSettings, IconExclamationCircleFill, IconPlus } from '@arco-design/web-vue/es/icon';
import LinkManage from "@/components/AiChat/LinkManage.vue";
import AddLinkDrawer from "@/components/AiChat/AddLinkDrawer.vue";

const links= ref([])
const getLinks = () => {
  GetAiLinkList().then(response => {
    links.value = response
  })
}

const externalUrl = ref("")

const LinkManageRef = ref(null)
const openLinkManage = () => {
  LinkManageRef.value.showManage()
}

const AddLinkRef = ref('')
const openAddLinkForm = () => {
  AddLinkRef.value.showModal()
}

const selectorRef = ref(null)
const selectedValue = ref()
const showAiChatWeb = ref(false)
const selectedUrl = (value) => {
  links.value.forEach((link, index) => {
    if(link.ID === value){
      console.log(value)
      selectedValue.value = link.ID
      externalUrl.value = link.Link
    }
  });
}

const openAiChatWeb = () => {
  getLinks()
  console.log(links.value.length)
  if(links.value.length === 1){
    selectedValue.value =  links.value[0].ID
    externalUrl.value = links.value[0].Link
  } else {
    selectedValue.value = null
    externalUrl.value = ''
  }
  showAiChatWeb.value = true
}
function closeAiChatWeb () {
  showAiChatWeb.value = false
}

defineExpose({
  openAiChatWeb,
})

onMounted(
    getLinks,
)

</script>

<style scoped>



</style>