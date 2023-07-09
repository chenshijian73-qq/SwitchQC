<template>
  <div class="app">
    <div class="main">
      <div class="nav" >
        <FileNav :selectedFile="selectedFile" :selectFile="selectFile" :qcFiles="qcFiles" :getFiles="getFiles"/>
        <div class="options">
          <a-button type="primary" shape="circle" size="mini" @click="showAddFile">
            <icon-plus />
          </a-button>
          <RecycleBin :getFiles="getFiles"/>
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
import './assets/app.css'
import {onBeforeMount, ref} from 'vue';
import {GetFiles, LogInfo} from '../wailsjs';
import {
  IconPlus,
} from '@arco-design/web-vue/es/icon';
import CodeEditor from "@/components/CodeEditor.vue";
import AddFileDrawer from "@/components/AddFileDrawer.vue";
import FileNav from "@/components/FileNav.vue";
import RecycleBin from "@/components/RecycleBin.vue";

const fileMenuVisible = ref(false);
const addFileVisible = ref(false);

let qcFiles = ref([]);

function getFiles() {
  GetFiles().then(response => {
    qcFiles.value = response.map(file => ({ ...file, fileMenuVisible: false }));
    // response.forEach((file, index) => {
    //   console.log(`File ${index}:`, file.Name);
    // });
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
      return;
    }
  }
}

onBeforeMount(
    getFiles
);

</script>

<style scoped>
</style>