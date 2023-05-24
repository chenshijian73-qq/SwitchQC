<template>
    <div class="container">
        <!-- <h3>文件内容</h3> -->
        <div class="textarea-container">
        <textarea v-model="fileContent" style="width: 100%; height: 100vh;"></textarea>
        </div>
    </div>
</template>

<style>
.container {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.textarea-container {
  flex: 1;
  display: flex;
}

</style>

<script>
import { ref } from 'vue';

export default {
  props: {
      selectedFile: {
      type: Object,
      default: null,
      },
  },
  setup(props) {
      const fileContent = ref('');

      return {
        fileContent,
      };
  },
  watch: {
      selectedFile: {
        immediate: true,
        handler(newFile) {
            this.fileContent = newFile ? newFile.content : '';
        },
      },
  },
};
</script>