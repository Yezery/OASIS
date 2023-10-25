import Vue from 'vue'
import Vuex from 'vuex'
import Identicon from "identicon.js"

Vue.use(Vuex)

export default new Vuex.Store({
  namespaced: true,
  state: {
    ipfsIP: "10.39.5.194", 
    // serverIP:"10.22.62.16",
    serverIP:"localhost",
    //  市场合约地址
    MarketContractAddress: "0x950EA6251a3EF72768A2cb701b3c5eCd0cE3A603",
    // local
    // MarketContractAddress : "0xe3A96a34639C7bf2907B4693EfD2BD7a94479661",
    
    avatar: require("@/assets/webAssets/MetaMask.png"),
    currentAddress: "",
    Web3: null,
    ownerNFTList: [],
    NFTInf:null,
    Theme: false,
    isEmpower: false,
    isGetToken:false,
    isconnect: false,
    gptSocket: null,
    userSocket:null,

    isSearch: false,
    
  },
  getters: {
  },
  mutations: {
    // this.$store.commit("setGetToken", handleAccountsChanged[0]);
    setWEB3(state, value) {
      state.Web3 = value
    },
    setAvatar(state, value) {
      state.avatar =
        "data:image/png;base64," +
        new Identicon(value, 120).toString();
    },
    setOwnerNFTList(state, value) {
      state.ownerNFTList = value;
    },
    setNFTInf(state, value) {
      state.NFTInf= value;
    },
    setTheme(state, value) {
      state.Theme = value;
    },
    setEmpower(state, value) {
      state.isEmpower = value
    },
    setGetToken(state, value) {
      state.isGetToken = value
    },
    setConnection(state, value) {
      state.isconnect = value
    },
    setcurrentAddress(state, value) {
      state.currentAddress = value
    },
    
    setGptSocket(state, value){
      state.gptSocket = value
  },
  setUserSocket(state, value){
        state.userSocket = value
    },
setIsSearch(state, value){
      state.isSearch = value
    },
    


  },
  actions: {


  },
  modules: {
  }
})
