<template>
  <div class="app">
    <div class="main">
      <div class="nav" >
        <FileNav :selectedFile="selectedFile" :selectFile="selectFile" :qcFiles="qcFiles" :removeFile="removeFile"/>
        <a-empty v-if="Object.keys(qcFiles).length==0">
          No data, please add!
        </a-empty>
        <div class="options">
          <a-button type="primary" shape="circle" size="mini" @click="showAddFile">
            <icon-plus />
          </a-button>
          <RecycleBin :getFiles="getFiles"/>
        </div>
      </div>
      <div class="content">
        <CodeEditor :selectedFile="selectedFile" :disableCode="disableCode"/>
      </div>
    </div>
    <AddFileDrawer :qcFiles="qcFiles" :addFileVisible="addFileVisible" :getFiles="getFiles" :closeAddFile="closeAddFile" />
    <SuspendBot />
  </div>
</template>

<script setup>
import './assets/app.css'
import {onBeforeMount, ref} from 'vue';
import {GetFiles, LogInfo, RemoveFile} from '../wailsjs';
import {
  IconPlus,
} from '@arco-design/web-vue/es/icon';
import CodeEditor from "@/components/CodeEditor.vue";
import AddFileDrawer from "@/components/AddFileDrawer.vue";
import FileNav from "@/components/FileNav.vue";
import RecycleBin from "@/components/RecycleBin.vue";
import {message} from "ant-design-vue";
import SuspendBot from "@/components/AiChat/SuspendBot.vue";

const fileMenuVisible = ref(false);
const addFileVisible = ref(false);
const disableCode = ref(true);

let qcFiles = ref([]);

function getFiles() {
  GetFiles().then(response => {
    qcFiles.value = response.map(file => ({ ...file, fileMenuVisible: false }));
    // response.forEach((file, index) => {
    //   console.log(`File ${index}:`, file.Name);
    // });
    console.log(qcFiles)
  });
}

function showAddFile() {
  addFileVisible.value = true;
}
function closeAddFile() {
  addFileVisible.value = false;
}

const selectedFile = ref({});
function selectFile(file) {
  for (const f of qcFiles.value) {
    if (f.Name === file.Name) {
      selectedFile.value = f;
      disableCode.value = false;
      return;
    }
  }
}

function removeFile(file) {
  RemoveFile(file).then(err => {
    if (err !== ""){
      message.error(`删除文件 ${file.Name} 失败: ${err}`)
    } else {
      LogInfo(file.Name)
      console.log(selectedFile)
      if ( file.Name === selectedFile.value.Name) {
        LogInfo("lll")
        selectedFile.value = {}
        disableCode.value = true
      }
    }
    getFiles()
  })
}

onBeforeMount(
    getFiles
);

</script>

<style scoped>
</style>