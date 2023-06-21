<template>
  <div class="app">
    <div class="header">
      <a-button type="primary" v-if="showNav" @click="showNav = false">关闭导航栏</a-button>
      <a-button type="primary" v-else @click="showNav = true">打开导航栏</a-button>
      <a-button type="primary" @click="addFile">添加文件</a-button>
    </div>
    <div class="main">
      <div class="nav" v-if="showNav">
        <ul>
          <li v-for="(file, index) in qcFiles" :key="index" @click="selectFile(file, $event)" @contextmenu.prevent>
            <span class="name">{{ file.name }}</span>
            <a-switch v-model="file.enabled" />
          </li>
        </ul>
      </div>
      <div class="content">
        <a-textarea v-model="selectedFile.content" :disabled="!selectedFile.enabled" @input="saveFile" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showNav: true,
      qcFiles: [
        { name: 'file1', content: 'This is file1 content.', enabled: true },
        { name: 'file2', content: 'This is file2 content.', enabled: false },
        { name: 'file3', content: 'This is file3 content.', enabled: true },
      ],
      selectedFile: {},
      contextMenuVisible: false,
    };
  },
  methods: {
    addFile() {
      // 添加文件的逻辑
    },
    selectFile(file, event) {
      this.selectedFile = file;
      if (event.button === 2) {
        this.showContextMenu(file, event);
      }
    },
    editFile(file) {
      // 编辑文件信息的逻辑
      console.log('编辑文件信息', file);
    },
    removeFile(file) {
      // 移动文件到回收站的逻辑
      console.log('移动到回收站', file);
    },
    saveFile() {
      // 保存文件的逻辑
    },
  },
};
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
  margin-left: 8px;
}

.content {
  flex: 1;
  padding: 16px;
}
</style>