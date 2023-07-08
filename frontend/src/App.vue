<template>
  <div class="app">
    <div class="main">
      <div class="nav" >
        <FileNav :selectedFile="selectedFile" :selectFile="selectFile" :qcFiles="qcFiles" :getFiles="getFiles"/>
        <div class="options">
          <a-button type="primary" shape="circle" size="mini" @click="showAddFile">
            <icon-plus />
          </a-button>
          <RecycleCin/>
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
import { GetFiles } from '../wailsjs';
import {
  IconArchive,
  IconPlus,
} from '@arco-design/web-vue/es/icon';
import CodeEditor from "@/components/CodeEditor.vue";
import AddFileDrawer from "@/components/AddFileDrawer.vue";
import FileNav from "@/components/FileNav.vue";
import RecycleCin from "@/components/RecycleCin.vue";


let qcFiles = ref([
  { id: -1, name: 'file1', content: 'This is file1 content.', enabled: true },
].map(file => ({ ...file, fileMenuVisible: false })));
const fileMenuVisible = ref(false);
const addFileVisible = ref(false);

function getFiles() {
  GetFiles().then(response => {
    qcFiles.value = response.map(file => ({ ...file, fileMenuVisible: false }));
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
  qcFiles.value.forEach(f => {
    if (f.name === file.name) {
      selectedFile.value = f;
    }
  });
}

onBeforeMount(
    getFiles
);

</script>

<style scoped>
</style>