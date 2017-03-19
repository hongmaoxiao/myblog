<script>
  import 'github-markdown-css';
  import marked from 'marked';

  export default {
    data() {
      return {
        rawarticle: '""',
        title: '',
        created: '',
        prev: '',
        next: '',
      };
    },
    created() {
      this.$http.get(`/article/${this.$route.params.id}`)
        .then((data) => {
          console.log(data);
          this.title = data.body.article.Title;
          this.rawarticle = data.body.article.Content;
          this.created = data.body.article.CreatedAt;
          this.prev = data.body.prev === 0 ? '' : data.body.prev;
          this.next = data.body.next === 0 ? '' : data.body.next;
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
    <footer>
      <router-link v-if="prev" :to="{name: 'article', params: {id: prev}}">前一篇</router-link>
      <router-link v-if="next" :to="{name: 'article', params: {id: next}}">后一篇</router-link>
    </footer>
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
