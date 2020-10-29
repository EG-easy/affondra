import Vue from 'vue'
import App from '@firebase/app';
import '@firebase/auth';
import '@firebase/storage'

// import firebase from 'firebase/app'
// import 'firebase/auth'
// import "firebase/firestore";

// Vue.use(firestorePlugin);

// Your web app's Firebase configuration
var firebaseConfig = {
  apiKey: "AIzaSyB_ATyUz0dG5ip1cBIdz0NkEMcqtJ66J_o",
  authDomain: "affondra.firebaseapp.com",
  databaseURL: "https://affondra.firebaseio.com",
  projectId: "affondra",
  storageBucket: "affondra.appspot.com",
  messagingSenderId: "159065126497",
  appId: "1:159065126497:web:30552b208ddc1d29855c20"
};

var firebase = App.initializeApp(firebaseConfig)

console.log('firebase.storage()', firebase.storage());

const FirebasePlugin = {
  install(Vue) {
    Vue.prototype.$_firebase = firebase
    // Vue.$_firestore = firebase.firestore()
    Vue.prototype.$_firebaseStorage = firebase.storage()
  }
}

Vue.use(FirebasePlugin)