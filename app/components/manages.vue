<template>
  <section>
    <section class="articles-wrapper items" v-title :data-title='"后台管理"'>
      <p class="new-wrapper">
        <router-link to="/edit" class="btn-like new">新增</router-link>
      </p>
      <ul class="articles clearfix edit-articles">
        <li v-for="article in articles">
          <span class="title">{{article.Title}}</span>
          <div class="edit-handle">
            <router-link class="edit-link" :to="{ name: 'edit', params: {id: article.ID} }">
              <span class="btn-like edit">修改</span>
            </router-link>
            <span class="btn-like delete" @click="del(article)">删除</span>
          </div>
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
    methods: {
      del(article) {
        console.log(article);
        this.$http.post(`/delete/${article.ID}`)
        .then((data) => {
          console.log(data);
        }, (error) => {
          console.log(error);
        });
      }
    }
  };
</script>
<style>
.new-wrapper {
  margin-bottom: 15px;
  text-align: right;
}
.new {
  background: #73e043;
}
.edit {
  background: #969696;
}
.delete {
  background: #f13030;
}
.btn-like {
  padding: 6px 15px;
  border-radius: 15px;
  color: #fff;
  cursor: pointer;
}
.edit-articles > li {
  cursor: default;
}
.title {
  padding-left: 15px;
}
.edit-link {
  margin-right: 10px;
}
.edit-handle {
  display: inline-block;
  vertical-align: middle;
  float: right;
}
</style>
