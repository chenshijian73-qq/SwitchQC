<template>
  <div class="app">
    <div class="main">
      <div class="nav" v-if="showNav">
        <a-menu mode="pop" :style="{ width: '200px', borderRadius: '4px', backgroundColor: 'transparent'}">
          <a-menu-item v-for="(file, index) in qcFiles" :key="index" :style="{ backgroundColor: selectedFile.name === file.name ? '#d9e5f6' : 'transparent' }" @click="selectFile(file)">
            <div class="file-info">
            <a-switch v-model="file.enabled" size="small" @change="changeStatus(file)"/>
              <span class="name">{{ file.name.slice(0, file.name.lastIndexOf('.')) }}</span>
              <span class="menu-trigger">
              <a-trigger
                  :trigger="['hover']"
                  clickToClose
                  position="right"
                  @show="file.fileMenuVisible=true"
                  @hide="file.fileMenuVisible=false"
              >
                <div :class="`button-trigger ${file.fileMenuVisible ? 'button-trigger-active' : ''}`">
                  <IconDoubleRight class="icon-double-right " />
                </div>
                <template #content>
                  <a-menu
                      :style="{ marginBottom: '-4px' }"
                      mode="popButton"
                      :tooltipProps="{ position: 'right' }"
                      showCollapseButton
                  >
                    <a-menu-item key="index" class="red-icon" @click="removeFile(file)">
                      <template #icon><IconDelete class="red-icon"/></template>
                      Remove
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-trigger>
            </span>
            </div>
          </a-menu-item>
        </a-menu>
        <div class="options">
          <a-button type="primary" shape="circle" size="mini" @click="showAddFile">
            <icon-plus />
          </a-button>
          <a-button type="dashed" shape="circle" size="small" @click="">
            <icon-archive />
          </a-button>
        </div>
      </div>
      <div class="content">
        <codemirror v-model="selectedFile.content"
                     :style="{ height: '100%' }"
                     :autofocus="true"
                     :tabSize="2"
                     :extensions="extensions"
                     placeholder="Please select file to display content"
                     @blur="saveContent(selectedFile)"
                     :disabled="disableCode" />
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
        <a-form-item label="文件名"  prop="name" :rules="checkNameUnique(form.name)">
          <a-input v-model="form.name" :style="{width:'30%'}" allow-clear />
          <span v-if="nameUnique === false" style="color: red; margin-left: 10px">文件名已存在</span>
        </a-form-item>
        <a-form-item label="内容"  prop="content">
          <codemirror v-model="form.content"
                      :style="{ height: '200px', width: '100%'}"
                      :autofocus="true"
                      :tabSize="2"
                      :extensions="extensions"
                      placeholder="Please enter something" />
        </a-form-item>
        <a-form-item label="状态" :label-col="{ span: 16 }" :wrapper-col="{ span: 16 }" prop="enabled">
          <a-switch v-model="form.enabled" />
        </a-form-item>
      </a-form>
    </a-drawer>
  </div>
</template>

<script setup>
import './app.css'
import {onMounted, ref} from 'vue';
import { GetFiles, AddFile, EditFile, RemoveFile, LogInfo } from '../wailsjs';
import {message} from "ant-design-vue";
import {
  IconDoubleRight,
  IconArchive,
  IconDelete,
  IconPlus,
} from '@arco-design/web-vue/es/icon';
import { Codemirror } from "vue-codemirror";
import { javascript } from "@codemirror/lang-javascript";
import { EditorView } from "@codemirror/view"

let myTheme = EditorView.theme({
  // 输入的字体颜色
  "&": {
    color: "#0052D9",
    backgroundColor: "#FFFFFF"
  },
  ".cm-content": {
    caretColor: "#0052D9",
  },
  // 激活背景色
  ".cm-activeLine": {
    backgroundColor: "#FAFAFA"
  },
  // 激活序列的背景色
  ".cm-activeLineGutter": {
    backgroundColor: "#FAFAFA"
  },
  //光标的颜色
  "&.cm-focused .cm-cursor": {
    borderLeftColor: "#0052D9"
  },
  // 选中的状态
  "&.cm-focused .cm-selectionBackground, ::selection": {
    backgroundColor: "#95bdfd",
    color:'#6c9820',
  },
  // 左侧侧边栏的颜色
  ".cm-gutters": {
    backgroundColor: "#FFFFFF",
    color: "#ffcc98", //侧边栏文字颜色
    border: "none"
  }
}, { dark: false })
const extensions = [javascript(), myTheme];

const showNav = ref(true);
const fileMenuVisible = ref(false);
let disableCode = ref(true);
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
    const isNameUnique = qcFiles.value.every(file => file.name.slice(0, file.name.lastIndexOf('.')) !== value);
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
  }).then(err => {
    if (err !== "") {
      message.error(`添加文件 ${name} 失败: ${err}`)
    } else {
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
function selectFile(file) {
  qcFiles.value.forEach(f => {
    if (f.name === file.name) {
      selectedFile.value = f;
      disableCode = false;
    }
  });
}

function changeStatus(file) {
  EditFile(file).then(err => {
    if (err !== ""){
      message.error(`修改文件 ${file.name} 状态失败: ${err}`)
    }
  })
}
function saveContent() {
  EditFile(selectedFile.value).then(err => {
    if (err !== ""){
      message.error(`修改 ${selectedFile.value.name} 内容失败: ${err}`)
    }
  })
}
function removeFile(file) {
  // 移动文件到回收站的逻辑
  RemoveFile(file).then(err => {
    if (err !== ""){
      message.error(`删除文件 ${file.name} 失败: ${err}`)
    } else {
      if( selectedFile.value.name == file.name){
        selectedFile.value = {}
        disableCode=true
      }
      getFiles()
    }
  })

}

onMounted(
    getFiles
);

</script>

<style scoped>
</style>