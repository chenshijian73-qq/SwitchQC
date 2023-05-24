<template>
  <div id="app" style="display: flex; height: 100%;">
    <FileList @file-selected="selectFile" style="width: 20%; height: 100%; overflow-y: auto; border-right: 1px solid #ccc;" />
    <FileContent :selected-file="selectedFile" style="width: 80%; height: 100%; padding-left: 10px; overflow-y: auto;" />
  </div>
</template>

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