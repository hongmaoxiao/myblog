<template>
  <section>
    <section class="articles-wrapper items" v-title>
      <router-link to="/edit" class="new">新增</router-link>
      <ul class="articles">
        <li v-for="article in articles">
          <router-link :to="{ name: 'edit', params: {id: article.ID} }">
            <span>{{article.Title}}</span>
            <span class="pull-right">修改</span>
          </router-link>
        </li>
      </ul>
    </section>
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
    },
  };
</script>
<style>
.articles-wrapper {
  margin: 2em 10px;
  position: relative;
}
.new {
  display: block;
  text-align: right;
  line-height: 50px;
  padding-right: 15px;
}
.items {
  margin: 3em 10px;
}
.articles {
  text-align: left;
  border: 1px solid #ddd;
  box-shadow: 0 3px 15px #ccc;
}

.articles > li {
  line-height: 3em;
  position: relative;
  text-align: left;
  border-bottom: 1px solid #ddd;
  cursor: pointer;
}
.articles > li:last-child {
  border-bottom: 0;
}
.articles > li > a {
  padding: 0 15px;
  display: block;
}
.articles > li:hover {
  background: #f7f7f7;
}
@media screen and (max-width: 700px) {
  .articles-wrapper {
    margin: 3em 10px;
  }
}
</style>
