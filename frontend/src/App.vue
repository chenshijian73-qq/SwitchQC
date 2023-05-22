<template>
  <div id="app">
    <h1>SwitchHosts</h1>
    <button @click="addHost">添加 Host</button>
    <button @click="removeHost">删除 Host</button>
    <button @click="editHost">编辑 Host</button>
    <button @click="toggleHost">切换 Host</button>
    <button @click="viewHost">查看 Host</button>
  </div>
  <div>
    <h3>Host 文件列表</h3>
    <ul>
      <li v-for="(host, index) in hosts" :key="index">
        {{ host }}
      </li>
    </ul>
  </div>
</template>

<script setup>
import {
  AddHost,
  RemoveHost,
  EditHost,
  ToggleHost,
  ViewHost,
  LogInfo,
} from '../wailsjs';

const addHost = async () => {
  LogInfo("addHost");
  AddHost().then((res) => {
    LogInfo("addHost finished")
  });
};

const removeHost = async () => {
  LogInfo("RemoveHost");
  RemoveHost().then((res) => {
    LogInfo("RemoveHost finished")
  });
};

const editHost = async () => {
  LogInfo("EditHost");
  EditHost().then((res) => {
    LogInfo("EditHost finished")
  });
};

const toggleHost = async () => {
  LogInfo("ToggleHost");
  ToggleHost().then((res) => {
    LogInfo("ToggleHost finished")
  });
};

const viewHost = async () => {
  LogInfo("ViewHost");
  ViewHost().then((res) => {
    LogInfo("ViewHost finished")
  });
};


</script>