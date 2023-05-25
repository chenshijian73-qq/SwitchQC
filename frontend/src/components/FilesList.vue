<template>
  <div style="display: flex; flex-direction: column; height: 100%;">
    <div style="height: 95vh;">
      <button @click="createFile" style="align-self: flex-start; margin-top: auto; color: #488d0d">+</button>
      <ul style="text-align: left;padding: 5px;">
        <li v-for="(file, index) in files" :key="index" style="display: flex; align-items: center;">
          <button @click="selectFile(file)" style="text-align: left;margin-right: 5px;width: 90vh;">
            {{ file.name }}
          </button>
          <Button v-if="switchState" @click="toggleSwitch" style="color: green;">ON</Button>
          <Button v-else @click="toggleSwitch" style="color: #ff7271;">OFF</Button>
<!--          <button @click="deleteFile(file)" style="text-align: right;margin-left: 5px; color: #ff7271"> x </button>-->
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import {
  GetFiles,
  LogInfo,
} from '../../wailsjs';
import Button from 'ant-design-vue/lib/button';
import 'ant-design-vue/lib/button/style'

export default {
  components: {
    Button,
  },
  setup() {
    const files = ref([]);

    const switchState = ref(true)
    const toggleSwitch = () => {
      LogInfo("switchState")
      console.log(switchState.value)
      switchState.value = !switchState.value;
    };

    const selectFile = (file) => {
      this.$emit('file-selected', file);
    };

    const createFile = () => {
      // 调用后端方法以创建新文件
      LogInfo("createFile")
    };

    const deleteFile = (file) => {
      // 调用后端方法以删除当前选定的文件
    };

    return {
      files,
      switchState,
      selectFile,
      createFile,
      deleteFile,
      toggleSwitch,
    };
  },
  mounted() {
    this.loadFiles();
  },
  methods: {
    async loadFiles() {
      GetFiles().then((res) => {
        this.files = res;
      });
    },
  },
};
</script>