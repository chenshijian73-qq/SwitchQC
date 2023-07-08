<template>
  <a-dropdown position="tr">
    <a-button type="dashed" shape="circle" size="small" @click="">
      <icon-archive />
    </a-button>
    <template #content>
      <a-doption>Option 1</a-doption>
      <a-doption>Option 2</a-doption>
      <a-doption>Option 3</a-doption>
    </template>
  </a-dropdown>
</template>

<script setup>
import '../assets/recycle-bin.css'
import {
  IconArchive,
} from '@arco-design/web-vue/es/icon';
import {GetRecycleList, CleanBin} from "../../wailsjs";

</script>

<style scoped>
</style>