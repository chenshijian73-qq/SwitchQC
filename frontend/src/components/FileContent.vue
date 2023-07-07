<template>
  <div class="code-editor">
    <codemirror v-model="codeContent"
                :style="{ height: '100%' }"
                :autofocus="true"
                :tabSize="2"
                :extensions="extensions"
                placeholder="Please select file to display content"
                @blur="handleBlur"
                :disabled="disableCode" />
  </div>
</template>

<script setup>
import { ref, watchEffect, emit } from 'vue';
import { Codemirror } from "vue-codemirror";
import { javascript } from "@codemirror/lang-javascript";
import { EditorView } from "@codemirror/view"

const codeContent = ref('');
let myTheme = EditorView.theme({
  // 输入的字体颜色
  "&": {
    color: "#0052D9",
    backgroundColor: "#FFFFFF"
  },
  ".cm-content": {
    caretColor: "#0052D9",
  },
  // 激活背景色
  ".cm-activeLine": {
    backgroundColor: "#FAFAFA"
  },
  // 激活序列的背景色
  ".cm-activeLineGutter": {
    backgroundColor: "#FAFAFA"
  },
  //光标的颜色
  "&.cm-focused .cm-cursor": {
    borderLeftColor: "#0052D9"
  },
  // 选中的状态
  "&.cm-focused .cm-selectionBackground, ::selection": {
    backgroundColor: "#95bdfd",
    color:'#6c9820',
  },
  // 左侧侧边栏的颜色
  ".cm-gutters": {
    backgroundColor: "#FFFFFF",
    color: "#ffcc98", //侧边栏文字颜色
    border: "none"
  }
}, { dark: false })
const extensions = [javascript(), myTheme];const disableCode = false;

function handleBlur() {
  emit('save', codeContent.value);
}

watchEffect(() => {
  codeContent.value = props.codeContent;
});
</script>