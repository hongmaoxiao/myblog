<template>
  <div class="page-not-found">
    <p class="not-found-content">
      您来到了没有知识的荒原！　
    </p>
  </div>
</template>
<style scoped>
.page-not-found {
  position: fixed;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  background: #e8e8e8;
}

.not-found-content {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  padding: 30px;
  background: #fff;
  min-width: 300px;
  min-height: 200px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}

@media screen and (max-width: 700px) {
  .not-found-content {
    width: 80%;
    max-width: 80%;
    min-width: 80%;
  }
}
</style>
