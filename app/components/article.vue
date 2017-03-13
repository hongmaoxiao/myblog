<template>
  <div class="articles-wrapper">
    <div class="articles">
      <ul>
        <li v-for="article in articles" class="clearfix">
          <router-link :to="{ name: 'article', params: {id: article.ID} }">{{article.Title}}</router-link>
        </li>
      </ul>
    </div>
  </div>
</template>
<script>
  export default {
    data() {
      return {
        articles: [],
      };
    },
    mounted() {
      this.$http.get('/api/articles')
        .then((data) => {
          console.log(data);
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
  margin-top: 70px;
  width: 100%;
  overflow: hidden;
  float: left;
  position: relative;
}
.articles{
  float: left;
  width: 60%;
  margin-left: 8%;
  text-align: left;
  border: 1px solid #ddd;
}

.articles > ul > li{
  padding: 6px 12px;
  line-height: 30px;
  position: relative;
  text-align: left;
  border-bottom: 1px solid #ddd;
}
.articles > ul > li:last-child{
  border-bottom: 0;
}
</style>
