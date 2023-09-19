import Vue from 'vue'
import Vuex from 'vuex'
import Identicon from "identicon.js"

Vue.use(Vuex)

export default new Vuex.Store({
  namespaced: true,
  state: {
    IP:"10.39.5.194",
    // IP:"localhost",
    //  市场合约地址
    MarketContractAddress: "0x7a47aFbD41f9872790Fb094476bFEf1E91a72Bb4",
    // local
    // MarketContractAddress : "0x8d15a407A7B7590a75b3CFbfDee9B456da9Cc631",
    isSearch:false,
    isconnect: false,
    isDark: false,
    currentAddress: "",
    Web3: null,
    ipfs: null,
    //=========
    ownerNFTList: [],
    marketNFTInf:null,
    // ========= 用户信息
    avatar: require("@/assets/webAssets/MetaMask.png")
  },
  getters: {
  },
  mutations: {
    changeAvatar(state, value) {
      state.avatar =
        "data:image/png;base64," +
        new Identicon(value, 120).toString();
    },
    connection(state, value) {
      state.isconnect = value
     
    },
    setIsSearch(state, value){
        state.isSearch = value
    },
    setcurrentAddress(state, value) {
      state.currentAddress = value

    },
    setWEB3(state, value) {
      state.Web3 = value

    },
    setIPFS(state, value) {
      state.ipfs = value;

    },
    setIsDark(state, value) {
      state.isDark = value;
    },
    setOwnerNFTList(state, value) {
      state.ownerNFTList = value;
    },
    setMarketNFTInf(state, value) {
      state.marketNFTInf= value;
    }
  },
  actions: {


  },
  modules: {
  }
})
