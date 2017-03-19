<template>
  <section class="articles-wrapper items">
    <ul class="articles">
      <li v-for="article in articles">
        <router-link :to="{ name: 'article', params: {id: article.ID} }">{{article.Title}}</router-link>
      </li>
    </ul>
  </section>
</template>
<script>
  export default {
    data() {
      return {
        articles: [],
      };
    },
    mounted() {
      this.$http.get('/articles')
        .then((data) => {
          this.articles = data.body.articles;
        }, (error) => {
          console.log(error);
        });
      console.log(this.articles);
    },
  };
</script>
<style>
.articles-wrapper {
  margin: 2em 10px;
  position: relative;
}
.items {
  margin: 3em 10px;
}
.articles{
  text-align: left;
  border: 1px solid #ddd;
}

.articles > li{
  padding: 0 12px;
  line-height: 3em;
  position: relative;
  text-align: left;
  border-bottom: 1px solid #ddd;
}
.articles > li:last-child{
  border-bottom: 0;
}
.articles > li > a {
  display: block;
}
</style>
