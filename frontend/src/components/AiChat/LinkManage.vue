<template>
  <a-modal v-model:visible="isShowManage" :footer="false" width="auto">
    <template #title>
      Links Manage
    </template>
    <a-space direction="vertical">
      <a-list :style="{ width: '30vh'}" titleAlign="center">
        <a-list-item v-for="(link, index) in props.links" :index="index">
          <div class="list">
            <a-tooltip :content="link.Link" background-color="#3491FA" position="tl">
                <span class="select-text"> {{ link.Name }} </span>
            </a-tooltip>
            <a-space>
              <a-button type="dashed" status="danger" shape="circle" size="mini" @click="removeFromManage(link)">
                <icon-minus />
              </a-button>
            </a-space>
          </div>
        </a-list-item>
      </a-list>
      <div style="padding: 0 0; text-align: center;">
        <a-divider :size="2" style="border-bottom-style: dotted" />
        <a-button type="primary" shape="round" @click="props.openAddLinkForm">
          <icon-plus />
        </a-button>
      </div>
    </a-space>
  </a-modal>
</template>

<script setup>
import '../../assets/ai-chat-web.css'
import {
  IconPlus,
  IconMinus,
} from '@arco-design/web-vue/es/icon';
import {DeleteAiLink} from "../../../wailsjs/go/logic/AiLinkLogic";
import {ref} from "vue";
import {message} from "ant-design-vue";

const props = defineProps({
  links:{
    type: Array,
    required: true,
  },
  getLinks: {
    type: Function,
    required: true,
  },
  openAddLinkForm: {
    type: Function,
    required: true,
  }
});

function removeFromManage(link) {
  DeleteAiLink(link).then(err => {
    if (err){
      message.error(`删除 ${link.Name} 失败: ${err}`)
    }
  })
  props.getLinks()
}

const isShowManage = ref(false)
const showManage = () => {
  props.getLinks()
  isShowManage.value = true
}

defineExpose({
  showManage,
})

</script>

<style scoped>
</style>