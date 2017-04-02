<template>
  <section class="articles-wrapper clearfix" v-title :data-title='"后台管理"'>
    <article class="edit-article clearfix">
      <input class="title" v-model="title" placeholder="请输入文章标题">
      <textarea class="raw-article pull-left" v-model="rawarticle" debounce=300></textarea>
      <p class="parsed-article markdown-body pull-right" v-html="parseMarkdown"></p>
      <button @click="completeArticle" class="complete-article">提交</button>
    </article>
  </section>
</template>
<script>
  import marked from 'marked';

  export default {
    data() {
      return {
        title: '',
        rawarticle: '',
      };
    },
    computed: {
      parseMarkdown() {
        return marked(this.rawarticle, { sanitize: true });
      },
    },
    methods: {
      completeArticle() {
        if (!this.rawarticle || !this.title) {
          return;
        }
        this.$http.post('/edit', {
          title: this.title,
          content: this.rawarticle,
        })
        .then(() => {
          this.title = '';
          this.rawarticle = '';
        }, (error) => {
          console.log(error);
        });
      },
    },
  };
</script>
<style scoped>
.edit-article {
  position: relative;
  margin: 0 auto;
  padding-bottom: 50px;
}
.parsed-article,
.raw-article {
  width: 48.5%;
  border: 1px solid #ddd;
  padding: 20px;
  box-sizing: border-box;
}
.title {
  display: block;
  margin-bottom: 20px;
  padding: 6px 12px;
}
.raw-article {
  min-height: 500px;
}
.complete-article {
  padding: 3px 12px;
  position: absolute;
  left: 25%;
  margin-left: -28px;
  bottom: 0px;
}
</style>
