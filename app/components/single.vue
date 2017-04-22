<script>
  import 'github-markdown-css';
  import marked from 'marked';
  import Common from './header';
  import Foot from './footer';

  export default {
    components: {
      Common,
      Foot,
    },
    data() {
      return {
        rawarticle: '',
        title: '',
        created: '',
        prev: '',
        next: '',
      };
    },
    computed: {
      parseMarkdown() {
        return marked(this.rawarticle, { sanitize: false });
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
  <section>
    <Common />
    <section class="articles-wrapper" v-title :data-title=title>
      <article class="article">
        <h1 class="article-title">{{title}}</h1>
        <p class="created">{{created}}</p>
        <p class="article-con markdown-body" v-html="parseMarkdown"></p>
      </article>
      <div class="pagination clearfix">
        <router-link class="prev pull-left" v-if="prev" :to="{name: 'article', params: {id: prev}}">前一篇</router-link>
        <router-link class="next pull-right" v-if="next" :to="{name: 'article', params: {id: next}}">后一篇</router-link>
      </div>
    </section>
    <Foot />
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
.pagination {
  margin-top: 40px;
}
.pagination > a {
  padding: 3px 15px;
  background: #969595;
  border-radius: 15px;
  line-height: 1.5em;
  color: #fff;
}
.pagination > a:hover{
  background: #696666;
}
.prev {
  margin-left: 50px;
}
.next {
  margin-right: 50px;
}
</style>
