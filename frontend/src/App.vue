<template>
  <div class="app">
    <div class="header">
      <a-button type="primary" v-if="showNav" @click="showNav = false">关闭导航栏</a-button>
      <a-button type="primary" v-else @click="showNav = true">打开导航栏</a-button>
      <a-button type="primary" @click="addFile">添加文件</a-button>
    </div>
    <div class="main">
      <div class="nav" v-if="showNav">
        <ul>
          <li v-for="(file, index) in qcFiles" :key="index" @click="selectFile(file, $event)" @contextmenu.prevent>
            <span class="name">{{ file.name }}</span>
            <a-switch v-model="file.enabled" />
          </li>
        </ul>
      </div>
      <div class="content">
        <a-textarea v-model="selectedFile.content" :disabled="!selectedFile.enabled" @input="saveFile" />
      </div>
    </div>
    <a-drawer
        title="添加文件"
        :visible="addFileVisible"
        :closable="false"
        :placement="placement"
        width="400px"
        @close="addFileVisible = false"
    >
      <a-form ref="form" :model="form" :rules="rules">
        <a-form-item label="文件名" :label-col="{ span: 6 }" :wrapper-col="{ span: 16 }" prop="name">
          <a-input v-model="form.name" />
        </a-form-item>
        <a-form-item label="内容" :label-col="{ span: 16 }" :wrapper-col="{ span: 16 }" prop="content">
          <a-input v-model="form.content" />
        </a-form-item>
      </a-form>
      <div slot="footer">
        <a-button @click="addFileVisible = false">取消</a-button>
        <a-button type="primary" @click="handleSubmit">确定</a-button>
      </div>
    </a-drawer>
  </div>
</template>

<script setup>
import {createApp, onMounted, ref} from 'vue';
import { GetFiles, LogInfo } from '../wailsjs';
import {Form, Input, Modal} from "ant-design-vue";

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

const selectedFile = ref({});
const contextMenuVisible = ref(false);

const addFileVisible = ref(false);
const placement = ref('right');

const form = ref({
  name: '',
  content: '',
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

function addFile() {
  addFileVisible.value = true;
}

function handleSubmit() {
  const { name, content } = form.value;
  qcFiles.value.push({ name, content, enabled: true });
  LogInfo(`添加文件 ${name} 成功`);
  addFileVisible.value = false;
  form.value.name = '';
  form.value.content = '';
}

function selectFile(file, event) {
  selectedFile.value = file;
  if (event.button === 2) {
    showContextMenu(file, event);
  }
}

function editFile(file) {
  // 编辑文件信息的逻辑
  console.log('编辑文件信息', file);
}

function removeFile(file) {
  // 移动文件到回收站的逻辑
  console.log('移动到回收站', file);
}

function saveFile() {
  // 保存文件的逻辑
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
}
</style>