<script>
  import 'github-markdown-css';
  import marked from 'marked';

  export default {
    data() {
      return {
        rawarticle: '',
      };
    },
    created() {
      this.$http.get(`/api/article/${window.location.href.match(/article\/(\d+)$/)[1]}`)
        .then((data) => {
          console.log(data.body.article.Content);
          this.rawarticle = data.body.article.Content;
        }, (error) => {
          console.log(error);
        });
    },
    computed: {
      parseMarkdown: () => marked(this.rawarticle || '', { sanitize: true }),
    },
  };
</script>
<template>
  <div class="articles-wrapper">
    <div class="article">
      <p class="article-con" v-html="parseMarkdown"></p>
    </div>
  </div>
</template>
<style scoped>
.article {
  float: left;
  width: 60%;
  margin-left: 8%;
  text-align: left;
  border: 1px solid #ddd;
  padding: 10px;
}
</style>
