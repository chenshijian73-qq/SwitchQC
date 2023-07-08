<template>
  <a-drawer
      title="Add File"
      :visible="addFileVisible"
      :closable="false"
      :placement="'right'"
      width="50%"
      @ok="handleAddFileSubmit"
      @cancel="props.closeAddFile"
  >
    <a-form :model="form">
      <a-form-item label="Name:"  prop="name" :rules="checkNameUnique(form.name)">
        <a-input v-model="form.name" :style="{width:'30%'}" allow-clear />
        <span v-if="nameUnique === false" style="color: red; margin-left: 10px">文件名已存在</span>
      </a-form-item>
      <a-form-item label="Content:"  prop="content">
        <codemirror v-model="form.content"
                    :style="{ height: '50vh', width: '100%'}"
                    :autofocus="true"
                    :tabSize="2"
                    :extensions="extensions"
                    placeholder="Please enter something" />
      </a-form-item>
      <a-form-item label="Enabled:" :label-col="{ span: 16 }" :wrapper-col="{ span: 16 }" prop="enabled">
        <a-switch v-model="form.enabled" />
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script setup>
import '../assets/add-file-drawer.css'

import {javascript} from "@codemirror/lang-javascript";
import {Codemirror} from "vue-codemirror";
import {ref, watch} from "vue";
import {message} from "ant-design-vue";
import {AddFile} from "../../wailsjs";

// codeEditor config
const extensions = [javascript()];

const props = defineProps({
  qcFiles: {
    type: Array,
    required: true,
  },
  addFileVisible: {
    type: Boolean,
    required: true,
  },
  getFiles: {
    type: Function,
    required: true,
  },
  closeAddFile: {
    type: Function,
    required: true,
  }
});

const addFileVisible = ref(props.addFileVisible);
const nameUnique = ref(null);
const form = ref({
  name: '',
  content: '',
  enabled: true,
});

async function checkNameUnique(value) {
  if (!value) {
    nameUnique.value = null;
  } else {
    const isNameUnique = props.qcFiles.every(file => file.name.slice(0, file.name.lastIndexOf('.')) !== value);
    nameUnique.value = isNameUnique;
  }
}

function handleAddFileSubmit() {
  const isNameUnique = props.qcFiles.every(file => file.name !== form.value.name);
  if (!isNameUnique) {
    message.error('文件名已存在');
    return false;
  }
  const {name, content, enabled} = form.value;
  AddFile({
    Name: name,
    Content: content,
    Enabled: enabled,
  }).then(err => {
    if (err !== "") {
      message.error(`添加文件 ${name} 失败: ${err}`)
    } else {
      message.info(`添加文件 ${name} 成功`)
      addFileVisible.value = false;
      form.value.name = '';
      form.value.content = '';
      form.value.enabled = true;
      props.closeAddFile()
      props.getFiles()
    }
  })
}

watch(() => props.addFileVisible, (value) => {
  addFileVisible.value = value
});

</script>

<style scoped>

</style>