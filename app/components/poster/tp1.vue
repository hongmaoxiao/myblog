<template>
  <section class="tp1">
    <div class="author">
      {{author}}
    </div>
    <div class="content">
      <div v-for="item in content" :key="item" class="contentText">
        {{item}}
      </div>
    </div>
  </section>
</template>
<script>
  export default {
    props: ['content', 'author'],
    data() {
      return {
      };
    },
    mounted() {
      console.log('this.content: ', this.content)
      console.log('this.author: ', this.author)
    }
  };
</script>
<style scoped>
.tp1 {
  width: 100%;
  height: 736px;
  padding: 20px;
  box-sizing: border-box;
  display: flex;
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
  font-size: 22px;
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
