<template>
  <section class="tp2" id='poster'>
    <img
      :src="imgUrl"
      alt="posterbg"
      class="image"
      :style='imageStyle'
    >
    <div
      :class="[[position], 'contentWrapper']"
    >
      <div class="poemVertical">
        <div class="poem">
          <div
            v-for="(item, i) in content"
            :key="item"
            class="text"
            :style="{
              textIndent: position !== 'center' ? `${i * 1.2}em` : 0,
            }"
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
    props: ['content', 'author', 'image', 'blur', 'position'],
    data() {
      return {
      };
    },
    computed: {
      imgUrl() {
        console.log('process.env.BASE_URL: ', process.env.BASE_URL)
        // const url = `${process.env.BASE_URL}/tpl/${this.image}.png`
        const url = 'https://poem.fengxiaomao.com/2f472ce78b9294569de51ef263e25711?e=4734147487&token=C4V1LUiqHJXmvGKUjGWOH8oOmZ1KCeFOWWr-C9CX:Jxconpnpx9m2RxpIVhHY3ge7yws='
        return url
      },
      imageStyle() {
        return {
          filter: `blur(${this.blur || 0}px)`,
          '-webkit-filter': `blur(${this.blur || 0}px)`,
        }
      }
    },
    mounted() {
      console.log('this.image: ', this.image)
      console.log('this.blur: ', this.blur)
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
  overflow: hidden;
  background-color: #ffffff;
}
.image {
  width: 100%;
  height: auto;
  vertical-align: middle;
}
.contentWrapper {
  position: absolute;
  padding: 20px;
  box-sizing: border-box;
}

.leftTop,
.leftBottom {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
}

.rightTop,
.rightBottom {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  align-items: center;
}

.leftTop {
  left: 0;
  top: 0;
}

.leftBottom {
  left: 0;
  bottom: 0;
}

.rightTop {
  right: 0;
  top: 0;
}

.rightBottom {
  right: 0;
  bottom: 0;
}

.center {
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}

.poemVertical {
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
