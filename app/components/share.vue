<template>
  <section v-bind:style="styleObject" class='wrapper'>
    <Tp1
      v-if="+tid === 1"
      :content='contentList'
      :author='authorInfo'
    />
    <Tp2
      v-if="+tid === 2"
      :content='contentList'
      :author='authorInfo'
      :image='image'
      :blur='blur'
      position='leftTop'
    />
    <Tp2
      v-if="+tid === 3"
      :content='contentList'
      :author='authorInfo'
      :image='image'
      :blur='blur'
      position='rightTop'
    />
    <Tp2
      v-if="+tid === 4"
      :content='contentList'
      :author='authorInfo'
      :image='image'
      :blur='blur'
      position='center'
    />
    <Tp2
      v-if="+tid === 5"
      :content='contentList'
      :author='authorInfo'
      :image='image'
      :blur='blur'
      position='leftBottom'
    />
    <Tp2
      v-if="+tid === 6"
      :content='contentList'
      :author='authorInfo'
      :image='image'
      :blur='blur'
      position='rightBottom'
    />
  </section>
</template>
<script>
  import Tp1 from './poster/tp1'
  import Tp2 from './poster/tp2'
  import compact from 'lodash.compact'
  export default {
    components: {
      Tp1,
      Tp2,
    },
    props: ['bgc', 'color', 'content', 'title', 'author', 'blur', 'image', 'tid', 'font'],
    data() {
      return {
      }
    },
    computed: {
      styleObject() {
        const bgc = this.bgc === 'transparent' ? 'transparent' : `#${this.bgc}`
        return {
          color: `#${this.color}`,
          backgroundColor: bgc,
          fontFamily: this.font
        }
      },
      contentList() {
        return this.content && this.content.split('|') || []
      },
      authorInfo() {
        return compact([this.title, this.author]).join(' • ')
      }
    }
  }
</script>
<style>
.wrapper {
  width: 414px;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
}
/* .defaultFont {
  font-family: TChinese, SimSun, FangSong, STSong, STZhongsong, LiSu, KaiTi, -apple-system, BlinkMacSystemFont, 'Helvetica Neue', Helvetica,
  Segoe UI, Arial, Roboto, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft Yahei',
  sans-serif;
} */
</style>
