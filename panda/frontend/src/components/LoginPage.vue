<script setup>
import { reactive } from 'vue';
import {Login} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  password: "",
})

function login() {
  Login(data.name, data.password).then(result => {
    console.log(result)
  })
}
</script>

<template>
  <div class="container">
    <form action="#" class="login-form">
      <h2>登陆</h2>
      <input type="text" v-model="data.name" name="username" placeholder="用户名">
      <input type="password" v-model="data.password" name="password" placeholder="密码">
      <button  @click="login">登陆</button>
    </form>
  </div>
</template>

<style>
html,body{
  margin: 0;
  font-family: "PingFang SC","Microsoft Yahei",sans-serif;
}

.container{
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: url("/src/assets/images/bg.jpg") fixed no-repeat;  
  background-size: cover;
}

.login-form{
  width: 240px;
  height: 220px;
  display: flex;
  flex-direction: column;
  padding: 40px;
  text-align: center;
  position:relative;
  z-index: 100;
  background: inherit;
  border-radius: 18px;
  overflow: hidden;
}

.login-form::before{
  content: "";
  width: calc(100% + 20px);
  height: calc(100% + 20px);
  position: absolute;
  top: -10px;
  left: -10px;
  overflow: hidden;
  background: inherit;
  box-shadow: inset 0 0 0 200px rgba(255, 255, 255, 0.25);
  filter: blur(5px);
  z-index: -1;
}

.login-form h2{
  font-size: 18px;
  font-weight: 400;
  color: #3d5245;
}

.login-form input,
.login-form button{
  margin: 6px 0;
  height: 36px;
  border: none;
  background-color: rgba(255, 255, 255, 0.3);
  border-radius: 4px;
  padding: 0 14px;
  color: #3d5245;
}
.login-form input::placeholder{
  color: #3d5245;
}
.login-form button:focus,
.login-form input:focus{
  outline: 0;
}

.login-form button{
  margin-top:24px;
  background-color: rgba(209, 179, 9, 0.4);
  color: white;
  position: relative;
  overflow: hidden;
  cursor: pointer;
  transition: 0.4s;
}

.login-form button:hover{
  background-color: rgba(209, 179, 9 , 0.7);
}

.login-form button::before,
.login-form button::after{
  content: "";
  display: block;
  width: 80px;
  height: 100%;
  background: rgba(179, 255, 210, 0.5);
  opacity: 0.5;
  position: absolute;
  top: 0;
  left: 0;
  transform: skewX(-15deg);
  filter: blur(30px);
  overflow: hidden;
  transform: translateX(-100px);
}

.login-form button:hover::after{
  width: 40px;
  background: rgba(179, 255, 210, 0.3);
  left: 60px;
  opacity: 0;
  filter: blur(5px);
}
.login-form button:hover::before{
transition: 1s;
transform: translateX(320px);
opacity: 0.7s;
}
.login-form button:hover::after{
  transition: 1s;
  transform: translateX(320px);
  opacity: 1;
}
</style>