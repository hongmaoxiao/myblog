<template>
  <section>
    <div>{{bgc}}</div>
    <div>{{content}}</div>
    <div>{{title}}</div>
    <div>{{author}}</div>
  </section>
</template>
<script>
  export default {
    data() {
      return {
        bgc: '',
        content: '',
        title: '',
        author: '',
      };
    },
    mounted() {
      this.initParams();
    },
    methods: {
      initParams() {
        const {
          bgc,
          content,
          title,
          author,
        } = this.$route.params

        this.bgc = bgc;
        this.content = content;
        this.title = title;
        this.author = author;
      }
    }
  };
</script>
<style>
.about {
  margin: 3em 10px 5em;
  box-shadow: 0 3px 15px #ccc;
  border: 1px solid #ddd;
  padding: 20px;
}
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
