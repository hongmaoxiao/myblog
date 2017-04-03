<template>
  <section>
    <section class="articles-wrapper items" v-title>
      <p class="new-wrapper">
        <router-link to="/edit" class="new">新增</router-link>
      </p>
      <ul class="articles clearfix">
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
.new-wrapper {
  margin-bottom: 15px;
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
.new {
  padding: 3px 15px;
  border-radius: 15px;
  background: #969696;
  color: #fff;
}
</style>
