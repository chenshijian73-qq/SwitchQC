<template>
  <a-drawer
      title="Add AI Link"
      title-align="'center'"
      :visible="addLinkVisible"
      :closable="false"
      :placement="'right'"
      width="50%"
      @ok="handleAddLinkSubmit"
      @cancel="closeModal"
  >
    <a-form :model="form">
      <a-form-item label="Name:"  prop="name">
        <a-input v-model="form.name" :style="{width:'30%'}" allow-clear />
      </a-form-item>
      <a-form-item label="Link:"  prop="name">
        <a-input v-model="form.link" :style="{width:'100%'}" allow-clear />
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script setup>
import '../../assets/add-file-drawer.css'
import {ref} from "vue";
import {message} from "ant-design-vue";
import {SaveAiLink} from "../../../wailsjs/go/logic/AiLinkLogic";

const props = defineProps({
  getLinks: {
    type: Function,
    required: true,
  },
});

const addLinkVisible = ref(false);
const showModal = () => {
  addLinkVisible.value = true
}
const closeModal = () => {
  addLinkVisible.value = false
}

const form = ref({
  name: '',
  link: '',
});

function handleAddLinkSubmit() {
  const {name, link} = form.value;
  SaveAiLink({
    Name: name,
    Link: link,
  }).then(err => {
    if (err) {
      message.error(`添加文件 ${name} 失败: ${err}`)
    } else {
      message.info(`添加文件 ${name} 成功`)
      form.value.name = ''
      form.value.link = ''
      props.getLinks()
      closeModal()
    }
  })
}

defineExpose({
  showModal,
})

</script>

<style scoped>

</style>