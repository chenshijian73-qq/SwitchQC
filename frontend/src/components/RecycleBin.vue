<template>
  <a-button type="dashed" shape="circle" size="small" @click="showBin">
    <icon-archive />
  </a-button>
  <a-modal v-model:visible="isShowBin" @ok="handleOk" @cancel="handleCancel" width="auto">
    <template #title>
      RecycleBin
    </template>
    <a-list :style="{ width: '50vh'}" titleAlign="center">
      <a-list-item v-for="(file, index) in recycleList" :index="index">
        <div class="list">
          <span class="name">
              {{ file.Name.slice(0, file.Name.lastIndexOf('.')) }}
           </span>
          <a-space>
            <span class="deleted-date-font">Deleted: {{ file.DeleteAt }}</span>
            <a-button type="dashed" status="success" shape="circle" size="mini" @click="restoreFile(file)">
              <icon-refresh />
            </a-button>
            <a-button type="dashed" status="danger" shape="circle" size="mini" @click="removeFromBin(file)">
              <icon-minus />
            </a-button>
          </a-space>
        </div>
      </a-list-item>
    </a-list>
  </a-modal>
</template>

<script setup>
import '../assets/recycle-bin.css'
import {
  IconRefresh,
  IconArchive,
    IconMinus
} from '@arco-design/web-vue/es/icon';
import {GetRecycleList, CleanFileFromBin, RestoreFileFromBin} from "../../wailsjs/go/main/Recycle";
import {ref} from "vue";
import {message} from "ant-design-vue";

const props = defineProps({
  getFiles: {
    type: Function,
    required: true,
  },
});

const recycleList = ref({})
function getRecycleList() {
  GetRecycleList().then(response => {
    recycleList.value = response;
    // response.forEach((file, index) => {
    //   console.log(`File ${index}:`, file);
    // });
  });
}
const restoreFile = (file) => {
  RestoreFileFromBin(file).then(err => {
    if (err !== ""){
      message.error(`恢复文件 ${file.Name} 失败: ${err}`)
    }
    getRecycleList()
    props.getFiles()
  })
}

function removeFromBin(file) {
  CleanFileFromBin(file).then(err => {
    if (err !== ""){
      message.error(`删除文件 ${file.Name} 失败: ${err}`)
    }
    getRecycleList()
  })
}

const isShowBin = ref(false)
const showBin = () => {
  getRecycleList()
  isShowBin.value = true
}
const handleOk = () => {
}
const handleCancel = () => {
  isShowBin.value = false
}

defineExpose({
  showBin,
  handleCancel
})

</script>

<style scoped>
</style>