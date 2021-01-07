<template>
  <section v-bind:style="styleObject" class='wrapper'>
    <div class="author">
      {{authorInfo}}
    </div>
    <div class="content">
      <div v-for="item in contentList" :key="item" class="contentText">
        {{item}}
      </div>
    </div>
  </section>
</template>
<script>
  import compact from 'lodash.compact'
  export default {
    props: ['bgc', 'color', 'content', 'title', 'author'],
    data() {
      return {
      }
    },
    computed: {
      styleObject() {
        return {
          color: `#${this.color}`,
          backgroundColor: `#${this.bgc}`,
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
  width: 100vw;
  height: 100vh;
  margin: 0;
  padding: 50px;
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: flex-start;
}
.content {
  font-size: 30px;
  flex: 1;
  height: 100%;
  writing-mode: vertical-rl;
  text-align: start;
}

.contentText {
  text-align: start;
  writing-mode: vertical-rl;
  line-height: 1.4em;
}

.author {
  font-size: 24px;
  width: auto;
  height: 100%;
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  align-items: flex-start;
  writing-mode: vertical-rl;
}

.authorText {
  writing-mode: vertical-rl;
  display: inline-block;
}
</style>
