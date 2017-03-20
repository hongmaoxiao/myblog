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
    computed: {
      parseMarkdown() {
        return marked(this.rawarticle, { sanitize: true });
      },
    },
    watch: {
      "$route": "fetchData",
    },
    created() {
      this.fetchData();
    },
    methods: {
      fetchData() {
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
      }
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
    <footer class="clearfix">
      <router-link class="prev pull-left" v-if="prev" :to="{name: 'article', params: {id: prev}}">前一篇</router-link>
      <router-link class="next pull-right" v-if="next" :to="{name: 'article', params: {id: next}}">后一篇</router-link>
    </footer>
  </section>
</template>
<style scoped>
.article-title {
  text-align: center;
  font-size: 25px;
  line-height: 1.5em;
  word-wrap: break-word;
}
.created {
  font-size: 16px;
  margin: 0.6em 0 1em;
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
  box-shadow: 0 3px 15px #ccc;
}
footer > a {
  padding: 3px 15px;
  background: #969595;
  border-radius: 15px;
  line-height: 1.5em;
  color: #fff;
}
footer > a:hover{
  background: #696666;
}
.prev {
  margin: 30px 0 30px 50px;
}
.next {
  margin: 30px 50px 30px 0;
}
</style>
