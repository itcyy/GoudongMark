<template>
<div id="app">
    <div class="Tilte">
        <div class="ImgFrom">
            <p class="titleView">爆品推荐</p>
            <router-link to="/web/v2/goods" style="font-size:20px ;margin-top: 2px;">查看更多>></router-link>
  
           
        </div>
            
        <div class="ImgFrom">
            <div class="ImgOne">
                <img class="Imgstyle" src="../../components/img/phone01.png" />
                <p class="goodsName">Redmi note 12 5G</p>
                <p class="goodsPrice">￥1192起</p>
            </div>
            <div class="ImgOne">
                <img class="Imgstyle" src="../../components/img/phone02.png" />
                <p class="goodsName">Redmi note 12 Pro</p>
                <p class="goodsPrice">￥1699起</p>
            </div>
            <div class="ImgOne">
                <img class="Imgstyle" src="../../components/img/phone03.png" />
                <p class="goodsName">Redmi note 12 Pro+</p>
                <p class="goodsPrice">￥2090起</p>
            </div>
            <div class="ImgOne">
                <img class="Imgstyle" src="../../components/img/phone04.png" />
                <p class="goodsName">HUAWEI Mate 50 Pro</p>
                <p class="goodsPrice">￥6799起</p>
            </div>
            <div class="ImgOne">
                <img class="Imgstyle" src="../../components/img/phone05.png" />
                <p class="goodsName">HUAWEI MatePad Pro</p>
                <p class="goodsPrice">￥4399起</p>
            </div>
            
            
        </div>
       
       
    </div>
    
    
</div>

</template>


<script>

</script>
<style>
.ImgFrom{
    display: flex;
    flex-direction: row;
   
}
.ImgOne{
    background-color: white;
    width: 220px;
    margin-left: 20px;
    padding: 10px 10px 10px 10px;
}
.Tilte{
    display: flex;
    flex-direction: column;
    
    margin-top: 30px;
    
}
.goodsName{
    font-family: 宋体;
    font-size: 22px;
}
.goodsPrice{
    font-size: 22px;
    color: #FFB152;
    font-family: 宋体;
}
#app{
    width: 100%;
    
    display: flex;
    justify-content: center;
}
.titleView{
    font-family: 华文行楷;
    font-size: 40px;
    color: #cc2900;
    margin-right: 1000px;
    
}
.Imgstyle{
    width: 130px;
    height: 200px;
    margin-left: 20px;
    margin-top: 20px;
}
</style>