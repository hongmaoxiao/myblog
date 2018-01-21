<template>
  <section class="loading-outer">
    <div class="loading">
      <span></span>
      <span></span>
      <span></span>
      <span></span>
      <span></span>
      <span></span>
      <span></span>
      <span></span>
    </div>
  </section>
</template>
<style scoped>
.loading-outer {
  position: fixed;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
}

.loading {
    width: 100px;
    height: 100px;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
}
.loading span {
    display: inline-block;
    width: 20px;
    height: 20px;
    border-radius: 50%;
    background: lightgreen;
    position: absolute;
    animation: load 1.04s ease infinite;
}
@keyframes load{
    0%{
        transform: scale(1.2);
        opacity: 1;
    }
    100%{
        transform: scale(.3);
        opacity: 0.5;
    }
}
.loading span:nth-child(1){
    left: 0;
    top: 50%;
    margin-top:-10px;
    animation-delay:0.13s;
}
.loading span:nth-child(2){
    left: 14px;
    top: 14px;
    animation-delay:0.26s;
}
.loading span:nth-child(3){
    left: 50%;
    top: 0;
    margin-left: -10px;
    animation-delay:0.39s;
}
.loading span:nth-child(4){
    top: 14px;
    right:14px;
    animation-delay:0.52s;
}
.loading span:nth-child(5){
    right: 0;
    top: 50%;
    margin-top:-10px;
    animation-delay:0.65s;
}
.loading span:nth-child(6){
    right: 14px;
    bottom:14px;
    animation-delay:0.78s;
}
.loading span:nth-child(7){
    bottom: 0;
    left: 50%;
    margin-left: -10px;
    animation-delay:0.91s;
}
.loading span:nth-child(8){
    bottom: 14px;
    left: 14px;
    animation-delay:1.04s;
}
</style>
