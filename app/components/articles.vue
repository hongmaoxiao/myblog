<template>
  <section>
    <Common />
    <section class="articles-wrapper items" v-title>
      <ul class="articles">
        <li v-for="article in articles">
          <router-link :to="{ name: 'article', params: {id: article.ID} }">{{article.Title}}</router-link>
        </li>
      </ul>
    </section>
    <Foot />
  </section>
</template>
<script>
  import Common from './header';
  import Foot from './footer';

  export default {
    components: {
      Common,
      Foot,
    },
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
    margin: 3em 10px 0;
  }
}
</style>
