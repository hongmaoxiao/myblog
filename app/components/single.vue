<script>
  import 'github-markdown-css';
  import marked from 'marked';

  export default {
    data() {
      return {
        rawarticle: '""',
        title: '',
        created: '',
      };
    },
    created() {
      this.$http.get(`/article/${this.$route.params.id}`)
        .then((data) => {
          console.log(data);
          this.title = data.body.article.Title;
          this.rawarticle = data.body.article.Content;
          this.created = data.body.article.CreatedAt;
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
  <section class="articles-wrapper">
    <article class="article">
      <h1 class="article-title">{{title}}</h1>
      <p class="created">{{created}}</p>
      <p class="article-con markdown-body" v-html="parseMarkdown"></p>
    </article>
  </section>
</template>
<style scoped>
.article-title {
  text-align: center;
  font-size: 32px;
  line-height: 1.5em;
}
.created {
  font-size: 16px;
  margin: 1em 0 1.5em;
  line-height: 1.5em;
  text-align: center;
  color: #5f6465;
}
.article {
  margin: 0 auto;
}
.article-con {
  text-align: left;
  border: 1px solid #ddd;
  padding: 20px;
  border-radius: 6px;
}
</style>
