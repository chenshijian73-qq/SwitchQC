<template>
  <div class="app">
    <div class="main">
      <div class="nav" >
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
        <CodeEditor :selectedFile="selectedFile"/>
      </div>
    </div>
    <AddFileDrawer :qcFiles="qcFiles" :addFileVisible="addFileVisible" :getFiles="getFiles" :closeAddFile="closeAddFile" />
  </div>
</template>

<script setup>
import './app.css'
import {onBeforeMount, onMounted, ref} from 'vue';
import { GetFiles, AddFile, EditFile, RemoveFile, LogInfo } from '../wailsjs';
import {message} from "ant-design-vue";
import {
  IconDoubleRight,
  IconArchive,
  IconDelete,
  IconPlus,
} from '@arco-design/web-vue/es/icon';
import CodeEditor from "@/components/CodeEditor.vue";
import AddFileDrawer from "@/components/AddFileDrawer.vue";

const fileMenuVisible = ref(false);
let disableCode = ref(true);
let qcFiles = ref([
  { id: 0, name: 'file1', content: 'This is file1 content.', enabled: true },
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

function showAddFile() {
  addFileVisible.value = true;
}
function closeAddFile() {
  addFileVisible.value = false;
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

function removeFile(file) {
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

onBeforeMount(
    getFiles
);

</script>

<style scoped>
</style>