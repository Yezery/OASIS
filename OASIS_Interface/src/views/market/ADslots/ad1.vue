<template>
  <div class="ADMainBox">
    <div class="innerBOX">
      <span style="position: absolute; left: 65px; top: 45px; font-size: 39px">
        The
      </span>
      <span
        style="
          font-family: 'TESLA';
          position: absolute;
          left: 140px;
          top: 90px;
          font-size: 100px;
          z-index: 2;
        "
      >DIABLO</span>
      <span style="position: absolute; left: 430px; bottom: 30%; font-size: 30px">
        is launched
      </span>

      <img
        src="@/assets/AD1Assets/Ellipse 1.png"
        alt=""
        width="160px"
        height="160px"
        style="position: absolute; left: 223px; bottom: 45px"
      >
      <img
        src="@/assets/AD1Assets/image 1.png"
        alt=""
        width="190px"
        height="190px"
        style="position: absolute; left: 200px; bottom: 0; z-index: 3"
      >

      <button class="ad1btn">
        <span>Discover Now</span>
      </button>
    </div>
  </div>
</template>
  
  <script>
  export default {
    name:"AD1",
    data() {
      return {};
    },
    methods: {},
  };
</script>
  
<style lang="scss" scoped>
.ADMainBox {
  font-family: "Transformers_Movie";
  width: 100%;
  height: 100%;
  background-color: black;
  position: relative;
  color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: all 0.6s;
  &:hover {
    transform: scale(1.1);

    transition: all 0.6s;
  }
}
.innerBOX {
  width: 600px;
  height: 100%;
  position: relative;
}
.ad1btn {
  position: absolute;
  left: 68px;
  bottom: 20%;
  width: 130px;
  height: 30px;
  background-color: black;
  border-radius: 8px;
  font-size: 15px;
  font-family: "Transformers_Movie";
  border: 1px solid rgba(255, 255, 255, 0);
  animation: textshine 1s linear 1s infinite alternate;
  z-index: 4;
  color: white;
  border: 1px solid white;
  user-select: none;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  display: flex;
  transition: all 0.4s ease-in-out;
  &:hover {
    color: black;
    border: 1px solid black;
    background-color: white;
    transition: all 0.4s ease-in-out;
  }
  & span {
    transform: translateY(2px);
  }
}
@keyframes textshine {
  100% {
    box-shadow: var(--boxshdow-style--);
  }
}
</style>