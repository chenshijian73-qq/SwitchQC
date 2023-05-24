<template>
  <div style="display: flex; flex-direction: column; height: 100%;">
    <button @click="createFile" style="align-self: flex-start; margin-top: auto;">创建文件</button>
    <div style="height: 95vh;">
      <!-- <h3>文件列表</h3> -->
      <ul>
        <li v-for="(file, index) in files" :key="index" style="display: flex; align-items: center;">
          <span @click="selectFile(file)">{{ file.name }}</span>
          <button @click="deleteFile(file)" style="margin-left: 5px;">x</button>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';

export default {
  setup() {
    const files = ref([]);

    const selectFile = (file) => {
      this.$emit('file-selected', file);
    };

    const createFile = () => {
      // 调用后端方法以创建新文件
    };

    const deleteFile = (file) => {
      // 调用后端方法以删除当前选定的文件
    };

    return {
      files,
      selectFile,
      createFile,
      deleteFile,
    };
  },
  mounted() {
    this.loadFiles();
  },
  methods: {
    async loadFiles() {
      // 从本地文件夹加载 .quickcmd 文件列表
      // 这里可以替换为实际的 API 调用或其他方法
      this.files = [
        { name: 'file1.quickcmd', content: 'File 1 content' },
        { name: 'file2.quickcmd', content: 'File 2 content' },
        { name: 'file3.quickcmd', content: 'File 3 content' },
        { name: 'file3.quickcmd', content: 'File 3 content' },
        { name: 'file3.quickcmd', content: 'File 3 content' },
        { name: 'file3.quickcmd', content: 'File 3 content' },
        { name: 'file3.quickcmd', content: 'File 3 content' },
        { name: 'file3.quickcmd', content: 'File 3 content' },
        { name: 'file3.quickcmd', content: 'File 3 content' },
        { name: 'file3.quickcmd', content: 'File 3 content' },
        { name: 'file3.quickcmd', content: 'File 3 content' },
        { name: 'file3.quickcmd', content: 'File 3 content' },
      ];
    },
  },
};
</script>