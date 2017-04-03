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
.new {
  display: block;
  text-align: right;
  line-height: 50px;
  padding-right: 15px;
}
</style>
