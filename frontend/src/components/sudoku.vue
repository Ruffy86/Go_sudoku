<template>
    <div id="app-sudoku">
        <div class="container">
            <h1>
                <table class="table-bordered" style="margin-left: auto; margin-right: auto;">
                    <tbody>
                    <tr v-for="(item, itemIndex) in message" v-bind:key="itemIndex" style="border-color: white; border-style: solid;">
                        <td v-for="(article, articleIndex) in item" :key="articleIndex" style="border-color: white; border-style: solid;">
                            <label v-if="article!=0">{{article}}</label>
                            <input type="text" v-if="article == 0" v-model="userInput.number[String(itemIndex) + String(articleIndex)]" style="height: 30px; width: 20px" />
                        </td>
                    </tr>
                    </tbody>
                </table>
            </h1>

            <a @click="getMessage">Empezar!</a>

            <span>{{ userInput }}</span>

        </div>
    </div>
</template>

<script>

export default {
  data() {
    return {
      message: [],
        userInput :{
            number: []
        }
    };
  },
  methods: {
    getMessage: function() {
      var self = this;

      window.backend.MyStruct.DisplayArray().then(mensaje => {
      self.message = mensaje;
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
    align-self: center;
}
a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
.table-bordered{
    border-color: white;
   }
</style>
