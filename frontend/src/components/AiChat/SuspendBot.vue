<template>
  <div class="suspended-ball"
       @mousedown="dragStart"
       @mousemove="drag"
       @mouseup="dragEnd"
       @mouseleave="dragEnd"
       @click="toggleDialog"
       :style="{ transform: `translate3d(${x}px, ${y}px, 0)` }">
    <a-popover class="dialog" :class="{ 'dialog-left': isLeft, 'dialog-right': !isLeft }">
      <img src="../../assets/cat.gif" style="width: 100px; height: 100px;" alt="icon" />
      <template #content>
        <p>Hello, nice to chat with you!<a-button type="text" @click="handleClick"> Click here to chat more!</a-button> </p>
      </template>
    </a-popover>
    <AiChatWeb ref="aiChatWebComRef"/>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import '../../assets/suspend-bot.css'
import AiChatWeb from "@/components/AiChat/AiChatWeb.vue";

const aiChatWebComRef = ref(null);
const handleClick = () => {
  if (aiChatWebComRef.value) {
    aiChatWebComRef.value.openAiChatWeb()
  }
}

const x = ref(0);
const y = ref(0);
let isDragging = false;
let initialX = 0;
let initialY = 0;
let currentX = 0;
let currentY = 0;
let xOffset = 0;
let yOffset = 0;

const dragStart = (e) => {
  isDragging = true;
  initialX = e.clientX - xOffset;
  initialY = e.clientY - yOffset;
};

const drag = (e) => {
  if (isDragging) {
    e.preventDefault();

    currentX = e.clientX - initialX;
    currentY = e.clientY - initialY;

    xOffset = currentX;
    yOffset = currentY;

    x.value = currentX;
    y.value = currentY;

    // 自动滚动页面
    const scrollThreshold = 50;
    const scrollSpeed = 10;
    const windowHeight = window.innerHeight;
    const windowWidth = window.innerWidth;
    const scrollX = window.scrollX;
    const scrollY = window.scrollY;

    if (e.clientY < scrollThreshold) {
      window.scrollTo(scrollX, scrollY - scrollSpeed);
    } else if (e.clientY > windowHeight - scrollThreshold) {
      window.scrollTo(scrollX, scrollY + scrollSpeed);
    }

    if (e.clientX < scrollThreshold) {
      window.scrollTo(scrollX - scrollSpeed, scrollY);
    } else if (e.clientX > windowWidth - scrollThreshold) {
      window.scrollTo(scrollX + scrollSpeed, scrollY);
    }
  }
};

const dragEnd = () => {
  isDragging = false;
};

const showDialog = ref(false);
const isLeft = ref(false);

const toggleDialog = () => {
  showDialog.value = !showDialog.value;
  isLeft.value = Math.abs(x.value) > window.innerWidth / 2;
};
</script>

<style>
</style>