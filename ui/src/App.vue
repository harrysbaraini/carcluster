<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: "App"
};
</script>

<style>
#app {
  @import url('https://fonts.googleapis.com/css?family=Comfortaa&display=swap');

  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

html, body {
  padding: 0;
  margin: 0;
}
</style>
