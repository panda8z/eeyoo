import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';

Vue.use(Vuetify);

// 多语言：添加中文\English
import zhHans from 'vuetify/es5/locale/zh-Hans'
import en from 'vuetify/es5/locale/en'

export default new Vuetify({

  // 添加多语言支持目前仅支持中文
  lang: {
    locales: { zhHans, en },
    current: 'zhHans',
  },
});
