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
import {GetRecycleList, DeleteFromBin, LogInfo, CleanBin} from "../../wailsjs";
import {ref} from "vue";
import {message} from "ant-design-vue";

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
  console.log("hello")
}

const removeFromBin = (file) => {
  LogInfo("hello")
  DeleteFromBin(file).then(err => {
    if (err !== ""){
      message.error(`删除文件 ${file.Name} 失败: ${err}`)
    } else {
      message.info("jlj")
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