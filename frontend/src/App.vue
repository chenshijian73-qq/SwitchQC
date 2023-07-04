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
            <a-trigger
                :trigger="['click', 'hover']"
                clickToClose
                position="right"
                v-model:popupVisible="fileMenuVisible[file.name]"
            >
              <div :class="`button-trigger ${fileMenuVisible[file.name] ? 'button-trigger-active' : ''}`">
                <IconClose v-if="fileMenuVisible[file.name]" />
                <IconMenu v-else />
              </div>
              <template #content>
                <a-menu
                    :style="{ marginBottom: '-4px' }"
                    mode="popButton"
                    :tooltipProps="{ position: 'left' }"
                    showCollapseButton
                >
                  <a-menu-item key="1" @click="removeFile(file)">
                    <template #icon><IconEdit /></template>
                    编辑
                  </a-menu-item>
                  <a-menu-item key="1" @click="removeFile(file)">
                    <template #icon><IconDelete /></template>
                    删除
                  </a-menu-item>
                </a-menu>
              </template>
            </a-trigger>
            <span class="name">{{ file.name }}</span>
            <a-switch v-model="file.enabled" @change="changeStatus(file)"/>
          </li>
        </ul>
      </div>
      <div class="content">
        <a-textarea v-model="selectedFile.content" :auto-size="{minRows:32,maxRows:32}" @change="saveContent(selectedFile)" />
      </div>
    </div>
    <a-drawer
        title="添加文件"
        :visible="addFileVisible"
        :closable="false"
        :placement="right"
        width="40%"
        @ok="handleAddFileSubmit"
        @cancel="addFileVisible = false"
    >
      <a-form ref="target" :model="form">
        <a-form-item label="文件名" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }" prop="name" :rules="checkNameUnique(form.name)">
          <a-input v-model="form.name" />
          <div v-if="nameUnique === false" style="color: red">文件名已存在</div>
        </a-form-item>
        <a-form-item label="内容" :label-col="{ span: 16 }" :wrapper-col="{ span: 16 }" prop="content">
          <a-textarea v-model="form.content" placeholder="Please enter something" :auto-size="{minRows:20,maxRows:20}" />
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
import { GetFiles, AddFile, EditFile, RemoveFile, LogInfo } from '../wailsjs';
import {message} from "ant-design-vue";
import {
  IconEdit,
  IconDelete,
  IconClose,
  IconMenu,
} from '@arco-design/web-vue/es/icon';

const showNav = ref(true);
let qcFiles = ref([
  { name: 'file1', content: 'This is file1 content.', enabled: true },
  { name: 'file2', content: 'This is file2 content.', enabled: false },
].map(file => ({ ...file, fileMenuVisible: false })));

function getFiles() {
  GetFiles().then(response => {
    qcFiles.value = response.map(file => ({ ...file, fileMenuVisible: false }));
  });
}

const addFileVisible = ref(false);
const placement = ref('right');
const target = ref(null);
const form = ref({
  name: '',
  content: '',
  enabled: true,
});

const nameUnique = ref(null);
async function checkNameUnique(value) {
  if (!value) {
    nameUnique.value = null;
  } else {
    const isNameUnique = qcFiles.value.every(file => file.name !== value);
    nameUnique.value = isNameUnique;
  }
}

function showAddFile() {
  addFileVisible.value = true;
}
async function handleAddFileSubmit() {
  const isNameUnique = qcFiles.value.every(file => file.name !== form.value.name);
  if (!isNameUnique) {
    message.error('文件名已存在');
    return false;
  }
  const {name, content, enabled} = form.value;
  AddFile({
    Name: name,
    Content: content,
    Enabled: enabled,
  }).then(response => {
    if (response == false) {
      message.error(`添加文件 ${name} 失败`)
    } else {
      qcFiles.value.push({name, content, enabled, fileMenuVisible: false});
      LogInfo(`添加文件 ${name} 成功`);
      message.info(`添加文件 ${name} 成功`)
      addFileVisible.value = false;
      form.value.name = '';
      form.value.content = '';
      form.value.enabled = true;
      getFiles()
    }
  })
}

const selectedFile = ref({});
const fileMenuVisible = ref(false);

function selectFile(file, event) {
  event.stopPropagation();
  qcFiles.value.forEach(f => {
    if (f === file) {
      f.fileMenuVisible = true;
      this.selectedFile = f;
    } else {
      f.fileMenuVisible = false;
    }
  });
}

function changeStatus(file) {
  // 编辑文件信息的逻辑
  console.log('修改文件信息', file);
  EditFile(file).then(res => {
    if (res == false){
      message.error(`修改文件 ${file.name} 状态失败`)
    } else {
      message.info(`修改文件 ${file.name} 状态成功`)
      getFiles()
    }
  })
}
function saveContent(file) {
  EditFile(file).then(res => {
    if (res == false){
      message.error(`修改文件 ${file.name} 内容失败`)
    } else {
      message.info(`修改文件 ${file.name} 内容成功`)
      getFiles()
    }
  })
}

function removeFile(file) {
  // 移动文件到回收站的逻辑
  RemoveFile(file).then(res => {
    if (res == false){
      message.error(`删除文件 ${file.name} 失败`)
    } else {
      message.info(`删除文件 ${file.name} 成功`)
      getFiles()
    }
  })

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
  margin-left: 1px;
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

.button-trigger {
  display: inline-block;
  width: 20px;
  height: 20px;
  line-height: 20px;
  text-align: center;
  border-radius: 50%;
  background-color: #1890ff;
  color: #fff;
  cursor: pointer;
  margin-right: 10px;
}

.button-trigger-active {
  background-color: #ff4d4f;
}

.ant-switch-checked {
  background-color: #1890ff;
}

.ant-switch-checked .ant-switch-handle::before {
  background-color: #fff;
}

.ant-switch:focus {
  box-shadow: none;
}
</style>