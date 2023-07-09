<template>
    <codemirror v-model="props.selectedFile.Content"
                :style="{ height: '100%' }"
                :autofocus="true"
                :tabSize="2"
                :extensions="extensions"
                placeholder="Please select file to display content"
                @blur="handleBlur"
                :disabled="disableCode" />
</template>

<script setup>
import {ref, watch} from 'vue';
import { Codemirror } from "vue-codemirror";
import { javascript } from "@codemirror/lang-javascript";
import { EditorView } from "@codemirror/view"
import {EditFile} from "../../wailsjs";
import {message} from "ant-design-vue";

// codeEditor config
let codeTheme = EditorView.theme({
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
    backgroundColor: "#d3e1f5"
  },
  // 激活序列的背景色
  ".cm-activeLineGutter": {
    backgroundColor: "#c4d9f8"
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
    backgroundColor: "#ffffff",
    color: "#ea8767", //侧边栏文字颜色
    border: "none"
  }
}, { dark: false })
const extensions = [javascript(), codeTheme];

const props = defineProps({
  selectedFile: {
    type: Object,
    required: true,
  },
});

const disableCode = ref(true);

function handleBlur() {
  EditFile(props.selectedFile).then(err => {
    console.log(`${err}`)
    if (err !== ""){
      message.error(`修改 ${props.selectedFile.name} 内容失败: ${err}`)
    }
  })
}

watch(() => props.selectedFile, (value) => {
  if (value) {
    disableCode.value = false;
  } else {
    disableCode.value = true;
  }
});

</script>