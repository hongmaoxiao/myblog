<script>
  import 'github-markdown-css';
  import marked from 'marked';

  export default {
    data() {
      return {
        rawarticle: '""',
        title: '',
      };
    },
    created() {
      this.$http.get(`/article/${this.$route.params.id}`)
        .then((data) => {
          this.title = data.body.article.Title;
          this.rawarticle = data.body.article.Content;
        }, (error) => {
          console.log(error);
        });
    },
    computed: {
      parseMarkdown() {
        return marked(this.rawarticle, { sanitize: true });
      },
    },
  };
</script>
<template>
  <div class="articles-wrapper">
    <div class="article">
      <h1 class="article-title">{{title}}</h1>
      <p class="article-con markdown-body" v-html="parseMarkdown"></p>
    </div>
  </div>
</template>
<style scoped>
.article-title {
  text-align: center;
  font-size: 32px;
  line-height: 2em;
  margin-bottom: 20px;
}
.article {
  width: 80%;
  margin: 0 auto;
}
.article-con {
  text-align: left;
  border: 1px solid #ddd;
  padding: 20px;
  border-radius: 6px;
}
</style>
