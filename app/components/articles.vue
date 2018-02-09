<template>
  <section>
    <Common />
    <Loading v-if="loading" />
    <section class="articles-wrapper items" v-title v-show="!loading">
      <ul class="articles">
        <li v-for="article in articles" :key="article.ID">
          <router-link :to="{ name: 'article', params: {id: article.ID} }">{{article.Title}}</router-link>
        </li>
      </ul>
    </section>
    <Foot v-show="!loading" />
  </section>
</template>
<script>
  import Common from './header';
  import Foot from './footer';
  import Loading from './loading';

  export default {
    components: {
      Common,
      Foot,
      Loading,
    },
    data() {
      return {
        articles: [],
        loading: false,
      };
    },
    mounted() {
      this.loading = true;
      this.$http.get('/articles')
        .then((data) => {
          this.loading = false;
          this.articles = data.body.articles;
        }, (error) => {
          console.log(error);
        });
    },
  };
</script>
<style>
.articles-wrapper {
  margin: 2em 10px 0;
  position: relative;
}
.items {
  margin: 3em 10px 0;
}
.articles {
  text-align: left;
  border: 1px solid #ddd;
  box-shadow: 0 3px 15px #ccc;
  list-style: none;
}

.articles > li {
  position: relative;
  text-align: left;
  border-bottom: 1px solid #ddd;
  cursor: pointer;
  padding: 10px 0;
}
.articles > li:last-child {
  border-bottom: 0;
}
.articles > li > a {
  padding: 0 15px;
  line-height: 1.5em;
  display: block;
}
.articles > li:hover {
  background: #f7f7f7;
}
@media screen and (max-width: 700px) {
  .articles-wrapper {
    margin: 3em 10px 0;
  }
  .articles > li > a {
    padding: 0 10px;
  }
}
</style>
