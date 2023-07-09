<template>
  <a-button type="dashed" shape="circle" size="small" @click="showBin">
    <icon-archive />
  </a-button>
  <a-modal v-model:visible="isShowBin" @ok="handleOk" @cancel="handleCancel" width="auto">
    <template #title>
      RecycleBin
    </template>
    <a-menu mode="pop" :style="{ width: '200px', borderRadius: '4px', backgroundColor: 'transparent'}">
      <a-menu-item v-for="(file, index) in recycleList" :index="index">
        <div class="file-info">
          <span class="name">{{ file.name.slice(0, file.name.lastIndexOf('.')) }}</span>
          <span class="deleted-date">deletedAt: {{ file.deleted }}</span>
        </div>
      </a-menu-item>
    </a-menu>
  </a-modal>
</template>

<script setup>
import '../assets/recycle-bin.css'
import {
  IconArchive,
} from '@arco-design/web-vue/es/icon';
import {GetRecycleList, CleanBin} from "../../wailsjs";
import {ref} from "vue";

const recycleList = ref({})
function getRecycleList() {
  GetRecycleList().then(response => {
    recycleList.value = response;
  });
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