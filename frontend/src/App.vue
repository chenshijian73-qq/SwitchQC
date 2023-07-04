<template>
  <div class="app">
    <div class="header">
      <a-button type="primary" v-if="showNav" @click="showNav = false">关闭导航栏</a-button>
      <a-button type="primary" v-else @click="showNav = true">打开导航栏</a-button>
      <a-button type="primary" @click="showAddFile">添加文件</a-button>
    </div>
    <div class="main">
      <div class="nav" v-if="showNav">
        <ul>
          <li v-for="(file, index) in qcFiles" :key="index" @click="selectFile(file, $event)" @contextmenu.prevent>
            <span class="name">{{ file.name }}</span>
            <a-switch v-model="file.enabled" @change="editFile(file)"/>
          </li>
        </ul>
      </div>
      <div class="content">
        <a-textarea v-model="selectedFile.content" :disabled="!selectedFile.enabled" auto-size=true @change="saveContent(selectedFile)" />
      </div>
    </div>
    <a-drawer
        title="添加文件"
        :visible="addFileVisible"
        :closable="false"
        :placement="placement"
        width="400px"
        @ok="handleAddFileSubmit"
        @cancel="addFileVisible = false"
    >
      <a-form ref="form" :model="form" :rules="rules">
        <a-form-item label="文件名" :label-col="{ span: 6 }" :wrapper-col="{ span: 16 }" prop="name">
          <a-input v-model="form.name" />
        </a-form-item>
        <a-form-item label="内容" :label-col="{ span: 16 }" :wrapper-col="{ span: 16 }" prop="content">
          <a-textarea v-model="form.content" auto-size=true />
        </a-form-item>
        <a-form-item label="状态" :label-col="{ span: 16 }" :wrapper-col="{ span: 16 }" prop="enabled">
          <a-switch v-model="form.enabled" />
        </a-form-item>
      </a-form>
    </a-drawer>
  </div>
</template>

<script setup>
import {onMounted, ref} from 'vue';
import { GetFiles, AddFile, EditFile, LogInfo } from '../wailsjs';
import {message} from "ant-design-vue";

const showNav = ref(true);
let qcFiles = ref([
  { name: 'file1', content: 'This is file1 content.', enabled: true },
  { name: 'file2', content: 'This is file2 content.', enabled: false },
  { name: 'file3', content: 'This is file3 content.', enabled: true },
]);

function getFiles() {
  GetFiles().then(response => {
    qcFiles.value = response;
  });
}

const addFileVisible = ref(false);
const placement = ref('right');
const form = ref({
  name: '',
  content: '',
  enabled: true,
});
const rules = ref({
  name: [
    { required: true, message: '请输入文件名', trigger: 'blur' },
    { min: 2, max: 20, message: '文件名长度在 2 到 20 个字符之间', trigger: 'blur' },
  ],
  content: [
    { required: true, message: '请输入文件内容', trigger: 'blur' },
  ],
});
function showAddFile() {
  addFileVisible.value = true;
}
function handleAddFileSubmit() {
  const { name, content, enabled } = form.value;
  AddFile({
    Name: name,
    Content: content,
    Enabled: enabled,
  }).then(response => {
    if (response == false) {
      message.error(`添加文件 ${name} 失败`)
    } else {
      qcFiles.value.push({ name, content, enabled });
      LogInfo(`添加文件 ${name} 成功`);
      message.info(`添加文件 ${name} 成功`)
      addFileVisible.value = false;
      form.value.name = '';
      form.value.content = '';
      form.value.enabled = true;
    }
  })
}

const selectedFile = ref({});
const contextMenuVisible = ref(false);
function selectFile(file, event) {
  selectedFile.value = file;
  if (event.button === 2) {
    showContextMenu(file, event);
  }
}

function editFile(file) {
  // 编辑文件信息的逻辑
  console.log('修改文件信息', file);
  EditFile(file).then(res => {
    if (res == false){
      message.error(`修改文件 ${file.name} 信息失败`)
    } else {
      message.info(`修改文件 ${file.name} 信息成功`)
    }
  })
}
function saveContent(file) {
  EditFile(file).then(res => {
    if (res == false){
      message.error(`修改文件 ${file.name} 内容失败`)
    } else {
      message.info(`修改文件 ${file.name} 内容成功`)
    }
  })
}

function removeFile(file) {
  // 移动文件到回收站的逻辑
  console.log('移动到回收站', file);
}


onMounted(
    getFiles
);

</script>

<style scoped>
.app {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.header {
  display: flex;
  justify-content: space-between;
  padding: 16px;
}

.main {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.nav {
  width: 200px;
  border-right: 1px solid #e8e8e8;
  overflow: auto;
}

ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

li {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 16px;
  cursor: pointer;
}

.name {
  margin-left: 8px;
}

.content {
  flex: 1;
  padding: 16px;
  width: 100%;
  height: 80vh;
}

.content a-textarea {
  width: 100%;
  height: 80vh;
}

</style>