<template>
  <div>
    <Editor :init="init"></Editor>
  </div>
</template>
<script>

import Editor from '@tinymce/tinymce-vue'
import tinymce from './tinymce.min.js'
import './icons/default/icons.min.js'
import './themes/silver/theme.min.js'
import './langs/zh_CN'
import './plugins/preview/'
export default {
  // props: [
  //   value: {

  //   }
  // ],
  components: { Editor },
  data() {
    return {
      init: {
        language: 'zh_CN'
      }
    }
  },

}
</script>
<style scope>
@import url("./skins/ui/oxide/skin.min.css");
</style>