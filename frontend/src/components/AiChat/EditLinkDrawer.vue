<template>
  <a-drawer
      title="Edit Link"
      title-align="'center'"
      :visible="addLinkVisible"
      :closable="false"
      :placement="'right'"
      width="50%"
      @ok="handleEditLinkSubmit"
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
import {UpdateAiLink} from "../../../wailsjs/go/logic/AiLinkLogic";

const form = ref({
  id: 0,
  name: '',
  link: '',
});
const props = defineProps({
  getLinks: {
    type: Function,
    required: true,
  },
});

const addLinkVisible = ref(false);
const showEditModal = (input) => {
  addLinkVisible.value = true
  form.value.id = input.ID
  form.value.name = input.Name
  form.value.link = input.Link
}

const closeModal = () => {
  addLinkVisible.value = false
}

function handleEditLinkSubmit() {
  const {id, name, link} = form.value;
  UpdateAiLink({
    ID: id,
    Name: name,
    Link: link,
  }).then(err => {
    if (err) {
      message.error(`更新 ${name} 失败: ${err}`)
    } else {
      message.info(`更新 ${name} 成功`)
      props.getLinks()
      closeModal()
    }
  })
}

defineExpose({
  showEditModal,
})

</script>

<style scoped>

</style>