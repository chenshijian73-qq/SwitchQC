<template>
  <div class="app">
    <div class="header">
      <a-button type="primary" v-if="showNav" @click="showNav = false">关闭导航栏</a-button>
      <a-button type="primary" v-else @click="showNav = true">打开导航栏</a-button>
      <a-button type="primary" @click="showAddFile">添加文件</a-button>
    </div>
    <div class="main">
      <div class="nav" v-if="showNav">
        <a-menu mode="pop">
          <a-menu-item v-for="(file, index) in qcFiles" :key="index" :style="{ backgroundColor: selectedFile.name === file.name ? '#e6f7ff' : 'transparent' }" @click="selectFile(file)">
            <div class="file-info">
            <a-switch v-model="file.enabled" size="small" @change="changeStatus(file)"/>
              <span class="name">{{ file.name.slice(0, file.name.lastIndexOf('.')) }}</span>
              <span class="menu-trigger">
              <a-trigger
                  :trigger="['click', 'hover']"
                  clickToClose
                  position="right"
                  v-model:popupVisible="fileMenuVisible[file.name]"
              >
                <div :class="`button-trigger ${fileMenuVisible[file.name] ? 'button-trigger-active' : ''}`">
                  <IconClose v-if="fileMenuVisible[file.name]" />
                  <IconDoubleRight v-else />
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
      </div>
      <div class="content">
        <codemirror v-model="selectedFile.content" :options="cmOptions" :style="{ height: '100%' }" placeholder="Code here..." @blur="saveContent(selectedFile)" :disabled="disableCode" />
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
          <codemirror v-model="form.content" :style="{ height: '100%', width: '100%'}" placeholder="Please enter something" />
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
  IconDoubleRight,
  IconDelete,
  IconClose,
} from '@arco-design/web-vue/es/icon';
import {Codemirror} from "vue-codemirror";
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/neo.css'
import 'codemirror/mode/shell/shell.js'
import "codemirror/addon/hint/anyword-hint.js";
// require active-line.js
import 'codemirror/addon/selection/active-line.js';
// closebrackets
import 'codemirror/addon/edit/closebrackets.js';
// keyMap
import 'codemirror/mode/clike/clike.js';
import 'codemirror/addon/edit/matchbrackets.js';
import 'codemirror/addon/comment/comment.js';
import 'codemirror/addon/dialog/dialog.js';
import 'codemirror/addon/dialog/dialog.css';
import 'codemirror/addon/search/searchcursor.js';
import 'codemirror/addon/search/search.js';
import 'codemirror/keymap/emacs.js';

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
const cmOptions = {
  theme: 'neo',  // 主题
  mode: 'shell' , // 代码格式
  tabSize: 4,  // tab的空格个数
  indentUnit: 2,  // 一个块（编辑语言中的含义）应缩进多少个空格
  autocorrect: true,  // 自动更正
  spellcheck: true,  // 拼写检查
  lint: true,  // 检查格式
  lineNumbers: true,  //是否显示行数
  lineWrapping: true, //是否自动换行
  styleActiveLine: true,  //line选择是是否高亮
  keyMap: 'sublime',  // sublime编辑器效果
  matchBrackets: true,  //括号匹配
  autoCloseBrackets: true,  // 在键入时将自动关闭括号和引号
  matchTags: { bothTags: true },  // 将突出显示光标周围的标签
  foldGutter: true,  // 可将对象折叠，与下面的gutters一起使用
  highlightSelectionMatches: {
    minChars: 2,
    style: "matchhighlight",
    showToken: true
  },
}

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
function selectFile(file) {
  qcFiles.value.forEach(f => {
    if (f.name === file.name) {
      f.fileMenuVisible = true;
      selectedFile.value = f;
      disableCode = false;
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
  console.log(file)
  console.log(selectedFile.value)
  EditFile(selectedFile.value).then(res => {
    if (res == false){
      message.error(`修改文件 ${selectedFile.value.name} 内容失败`)
    } else {
      message.info(`修改文件 ${selectedFile.value.name} 内容成功`)
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
  background-color: #f4fbff;
}

.code-mirror{
  font-size : 13px;
  line-height : 150%;
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
  margin-left: 10px;
  margin-right: 10px;
}
.menu-trigger {
  display: flex;
  justify-content: flex-end;
}

.content {
  flex: 1;
  padding: 10px;
  width: 100%;
  height: 85vh;
}

.content a-textarea {
  width: 100%;
  height: 80vh;
}

.red-icon {
  color: red;
}

.file-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.button-trigger {
  display: inline-block;
  width: 20px;
  height: 20px;
  line-height: 20px;
  text-align: center;
  border-radius: 50%;
  color: #fff;
  cursor: pointer;
  /*margin-right: 10px;*/
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

.CodeMirror-focused .cm-matchhighlight {
  background-position: bottom;
  background-repeat: repeat-x;
}
.CodeMirror-selection-highlight-scrollbar {
  background-color: green;
}
</style>