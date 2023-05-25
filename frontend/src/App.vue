<template>
  <div id="app" style="display: flex; height: 100%;">
    <FileList @file-selected="selectFile" style="width: 15%; height: 100%; overflow-y: scroll; border-right: 1px solid #eec5c5;" />
    <FileContent :selected-file="selectedFile" style="width: 80%; height: 100%; padding-left: 10px; overflow-y: scroll;" />
  </div>
</template>
<style>
  @import 'ant-design-vue/dist/reset.css';
</style>
<script>
import { ref } from 'vue';
import FileList from './components/FilesList.vue';
import FileContent from './components/FileContent.vue';

export default {
  components: {
    FileList,
    FileContent,
  },
  setup() {
    const selectedFile = ref(null);

    const selectFile = (file) => {
      selectedFile.value = file;
    };

    return {
      selectedFile,
      selectFile,
    };
  },
};
</script>