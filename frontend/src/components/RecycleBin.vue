<template>
  <a-button type="dashed" shape="circle" size="small" @click="showBin">
    <icon-archive />
  </a-button>
  <a-modal v-model:visible="isShowBin" @ok="handleOk" @cancel="handleCancel" width="auto">
    <template #title>
      RecycleBin
    </template>
    <a-list ::style="{ width: '40vh'}" titleAlign="center">
      <a-list-item v-for="(file, index) in recycleList" :index="index">
        <div class="list">
          <a-button type="dashed" status="success" shape="circle" size="mini" @click="restoreFile(file)">
            <icon-refresh />
          </a-button>
          <span class="name">
              {{ file.Name.slice(0, file.Name.lastIndexOf('.')) }}
          </span>
          <span class="deleted-date-font">DeletedAt: {{ file.DeleteAt }}</span>
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
} from '@arco-design/web-vue/es/icon';
import {GetRecycleList, CleanBin} from "../../wailsjs";
import {ref} from "vue";

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