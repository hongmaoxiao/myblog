<template>
  <a href="javascript: void(0);" class="back-to-top" v-show="isscroll" @click="BackToTop()">
    <i class="back-icon">&uarr;</i>
  </a href="javascript: void(0);">
</template>
<script>
  export default {
    data() {
      return {
        isscroll: false,
      };
    },
    created() {
      window.addEventListener('scroll', this.handleScroll);
    },
    methods: {
      handleScroll() {
        this.isscroll = window.scrollY > 200;
      },
      BackToTop() {
        window.scrollTo(0, 0);
      }
    },
    destroyed() {
      window.removeEventListener('scroll', this.handleScroll);
    },
  };
</script>
<style>
  .back-to-top {
    position: fixed;
    right: 40px;
    bottom: 40px;
    height: 50px;
    width: 50px;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    background: rgba(86, 85, 85, 0.7);
  }
  .back-icon {
    font-size: 30px;
    color: #fff;
  }
</style>
