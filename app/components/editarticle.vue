<template>
  <section class="articles-wrapper edit-wrapper clearfix" v-title :data-title='"后台管理"'>
    <p class="manages-wrapper">
      <router-link to="/manages" class="manages">返回管理主页</router-link>
    </p>
    <article class="edit-article">
      <input class="title" v-model="title" placeholder="请输入文章标题">
      <textarea class="raw-article pull-left" v-model="rawarticle" debounce=300></textarea>
      <p class="parsed-article markdown-body pull-right" v-html="parseMarkdown"></p>
      <button @click="completeArticle" class="complete-article">提交</button>
    </article>
  </section>
</template>
<script>
  import marked from 'marked';

  export default {
    data() {
      return {
        title: '',
        rawarticle: '',
        id: '',
      };
    },
    computed: {
      parseMarkdown() {
        return marked(this.rawarticle, { sanitize: true });
      },
    },
    watch: {
      "$route": "fetchData",
    },
    created() {
      this.fetchData();
    },
    methods: {
      completeArticle() {
        if (!this.rawarticle || !this.title) {
          return;
        }
        this.$http.post('/edit', {
          title: this.title,
          content: this.rawarticle,
          id: this.id,
        })
        .then((data) => {
          console.log(data);
        }, (error) => {
          console.log(error);
        });
      },
      fetchData() {
        let id = this.$route.params.id;
        if (id) {
          this.id = id;
          this.$http.get(`/article/${id}`)
            .then((data) => {
              this.title = data.body.article.Title;
              this.rawarticle = data.body.article.Content;
            }, (error) => {
              console.log(error);
            });
        }
      }
    },
  };
</script>
<style scoped>
.edit-article {
  position: relative;
  margin: 0 auto;
  padding-bottom: 50px;
}
.parsed-article,
.raw-article {
  width: 48.5%;
  border: 1px solid #ddd;
  padding: 20px;
  box-sizing: border-box;
}
.title {
  display: block;
  margin-bottom: 20px;
  padding: 6px 12px;
}
.raw-article {
  min-height: 500px;
}
.complete-article {
  padding: 3px 12px;
  position: absolute;
  left: 25%;
  margin-left: -28px;
  top: 580px;
}
.manages-wrapper {
  margin-bottom: 15px;
  display: flex;
  justify-content: flex-end;
  align-item: center;
}
.manages {
  padding: 3px 15px;
  border-radius: 15px;
  background: #969696;
  color: #fff;
}
</style>
