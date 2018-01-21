<template>
  <span class="back-to-top" v-show="isscroll" @click="BackToTop()"></span>
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
    height: 40px;
    width: 40px;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    background: rgba(86, 85, 85, 0.7);
  }
  .back-icon {
    font-size: 30px;
    color: #fff;
  }
  .back-to-top:before {
    content: '\2191';
    color: #fff;
  }
</style>
