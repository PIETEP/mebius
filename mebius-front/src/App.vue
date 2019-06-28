<template>
  <div>
    <myheader></myheader>
    <p v-if="msg.length > 0">
      {{msg}}
    </p>
    <p v-else>
      no text
    </p>
    <input type="text" v-model="msg">
    <button @click="clear()">clear</button>
  </div>
</template>

<script>
import myheader from './components/myheader'
export default {
  components: {
    myheader
  },
  data(){
    return {
      msg: 'Hello World'
    }
  },
  methods: {
    clear () {
      this.msg=''
    }
  },
  created () {
    fetch('http://localhost:8080/users')
    .then( response => {
      return response.json()
    })
    .then( json => {
      console.log(json)
      this.msg = json[0].firstname
    })
    .catch( () => {
    });
  }
}
</script>
