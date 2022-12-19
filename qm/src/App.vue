<template>
 <div id="app">
      <router-view/>
  </div>
</template>

<script>




export default {
  name: 'App',

  data: () => ({
    query:[],
    products:[],
    total:[],
  }),
  created(){
    

  },
  methods:{
    
  }
};
</script>
<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0px;
}
body {
  background-color: #2D3A4B;
}
</style>
