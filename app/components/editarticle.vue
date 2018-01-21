<template>
  <section class="articles-wrapper edit-wrapper clearfix" v-title :data-title='"后台管理"'>
    <p class="manages-wrapper">
      <router-link to="/manages" class="manages">返回管理主页</router-link>
    </p>
    <article class="edit-article">
      <div class="edit-content-fixed">
        <input class="title" v-model="title" placeholder="请输入文章标题" @keyup="saveLocalStorage">
        <textarea class="raw-article pull-left" v-model="rawarticle" debounce=300 @keyup="saveLocalStorage"></textarea>
        <button @click="completeArticle" class="complete-article">提交</button>
      </div>
      <p class="parsed-article markdown-body pull-right" v-html="parseMarkdown"></p>
    </article>
  </section>
</template>
<script>
  import marked from 'marked';
  import debounce from 'lodash.debounce';

  export default {
    data() {
      return {
        title: '',
        rawarticle: '',
        id: '',
        lastModifyTime: 0,
      };
    },
    computed: {
      parseMarkdown() {
        return marked(this.rawarticle, { sanitize: false });
      },
    },
    watch: {
      "$route": "fetchData",
    },
    created() {
      window.that = this;
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
          if (data.ok) {
            alert("提交成功");
          }
        }, (error) => {
          alert("提交失败", error);
        });
      },
      fetchData() {
        let id = this.$route.params.id;
        if (id) {
          this.id = id;
          this.$http.get(`/article/${id}`)
            .then((data) => {
              let article = data.body.article;
              this.title = article.Title;
              this.lastModifyTime = new Date(article.UpdatedAt).getTime();
              const local = this.getLocalStorage('article');
              if (local && local.title && local.title === article.Title && this.lastModifyTime <= local.time) {
                this.rawarticle = local.content;
              } else {
                this.rawarticle = article.Content;
              }
            }, (error) => {
              console.log(error);
            });
        } else {
          const local = this.getLocalStorage("article");
          if (local && local.time) {
            this.rawarticle = local.content;
            this.title = local.title;
          }
        }
      },
      saveLocalStorage: debounce(() => {
        const local = window.that.getLocalStorage('article');
        const has_local = !!local;
        const if_title_or_content_modify = has_local ?
        local.content !== window.that.rawarticle || local.title !== window.that.title :
        false;
        if (!has_local || if_title_or_content_modify) {
          const save_obj = {
            title: window.that.title,
            content: window.that.rawarticle,
            time: new Date().getTime()
          }
          localStorage.setItem("article", JSON.stringify(save_obj));
        }
      }, 2000),
      getLocalStorage(key) {
        const local = localStorage.getItem(key);
        if (!local) {
          return null;
        }
        return JSON.parse(local);
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
.edit-content-fixed {
  position: fixed;
  width: 44%;
  left: 5%;
  text-align: center;
}

.parsed-article,
.raw-article {
  border: 1px solid #ddd;
  padding: 20px;
  box-sizing: border-box;
}

.parsed-article {
  width: 68%;
  margin-right: -20%;
}

.raw-article {
  width: 100%;
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
  margin-top: 20px;
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
