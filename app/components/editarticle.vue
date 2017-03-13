<template>
  <div class="articles-wrapper clearfix">
    <div class="edit-article clearfix">
      <textarea class="raw-article pull-left" v-model="rawarticle" debounce=300></textarea>
      <p class="parsed-article pull-right" v-html="parseMarkdown"></p>
      <button @click="completeArticle" class="complete-article">提交</button>
    </div>
  </div>
</template>
<script>
  import marked from 'marked';

  export default {
    data() {
      return {
        rawarticle: '',
      };
    },
    computed: {
      parseMarkdown: function () {
        return marked(this.rawarticle, { sanitize: true });
      },
    },
    methods: {
      completeArticle() {
        if (!this.rawarticle) {
          return;
        }
        this.$http.post('/api/admin', {
          title: "test article",
          content: this.rawarticle,
        })
        .then((data) => {
          this.rawarticle = '';
        }, (error) => {
          console.log(error);
        });
      },
    },
  };
</script>
<style scoped>
.edit-article {
  position: relative;
  width: 80%;
  margin: 0 auto;
  padding-bottom: 50px;
}
.parsed-article,
.raw-article {
  width: 47.5%;
  border: 1px solid #ddd;
  padding: 20px;
}
.raw-article {
  min-height: 500px;
}
.complete-article {
  padding: 3px 12px;
  position: absolute;
  left: 25%;
  margin-left: -28px;
  bottom: 0px;
}
</style>
