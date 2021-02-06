<template>
  <section class="tp2" id='poster'>
    <img :src="image" alt="posterbg" class="image">
    <div class="contentWrapper">
      <div class="poemVertical">
        <div class="poem">
          <div
            v-for="(item, i) in content"
            :key="item"
            class="text"
            :style='{
              textIndent: `${i * 1.2}em`
            }'
          >
            {{item}}
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
  export default {
    props: ['content', 'author', 'image', 'blur'],
    data() {
      return {
      };
    },
    mounted() {
      console.log('this.image: ', this.image)
      console.log('this.author: ', this.author)
    }
  };
</script>
<style scoped>
.tp2 {
  width: 100%;
  height: auto;
  padding: 0;
  margin: 0;
  position: relative;
}
.image {
  width: 100%;
  height: auto;
  vertical-align: middle;
}
.contentWrapper {
  position: absolute;
  left: 0;
  top: 0;
  padding: 20px;
  box-sizing: border-box;
}
.poemVertical {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
  margin: 0 auto;
}
.poem {
  writing-mode: vertical-rl;
}
.text {
  letter-spacing: 6px;
  writing-mode: vertical-rl;
  line-height: 1.2em;
  font-weight: normal;
  font-size: 26PX;
}
</style>
