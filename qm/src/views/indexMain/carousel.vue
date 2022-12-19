<template>
  <div class="carousel">
    <el-carousel :interval="5000" arrow="always" height="470px">
      <el-carousel-item v-for="item in imgList" :key="item">
        <img :src="item.imgUrl" />
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "carousel",
  data() {
    return {
      imgList: [
        {
          imgUrl:
            "https://img14.360buyimg.com/pop/s590x470_jfs/t1/107362/20/24177/90602/6233decdE5627d944/b7194e538e3c8c7a.jpg",
        },
        {
          imgUrl:
            "https://img12.360buyimg.com/pop/s590x470_jfs/t1/204028/36/25480/101429/62f507a7Ec62b0cd3/db5984ae24ce0212.jpg",
        },
        {
          imgUrl:
            "https://imgcps.jd.com/ling4/100022552927/5Lqs6YCJ5aW96LSn/5L2g5YC85b6X5oul5pyJ/p-5f3a47329785549f6bc7a6e6/075268d0/cr/s/q.jpg",
        },
        {
          imgUrl:
            "https://imgcps.jd.com/ling4/100026667910/5Lqs6YCJ5aW96LSn/5L2g5YC85b6X5oul5pyJ/p-5f3a47329785549f6bc7a6e3/d7b3695b/cr/s/q.jpg",
        },
        {
          imgUrl:
            "https://img30.360buyimg.com/pop/s590x470_jfs/t1/201288/25/23710/64955/62f070fdE95cb4f47/22b916d938a43bb5.jpg",
        },
      ],
    };
  },
};
</script>

<style>
.carousel {
    margin: 0 25%;
  /* width: 800px; */
  height: 500px;
  left: 50%;
}
</style>