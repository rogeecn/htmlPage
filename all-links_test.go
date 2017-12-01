package htmlPage

import (
	"testing"
)

func Test_ListUrlFetch(t *testing.T) {

	htmlBody := `<body>
<nav id="w1" class="navbar navbar-static-top theme-nav"><div class="container"><div class="navbar-header"><button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#w1-collapse"><span class="sr-only">Toggle navigation</span>
<span class="icon-bar"></span>
<span class="icon-bar"></span>
<span class="icon-bar"></span></button><a class="navbar-brand" href="/">711XD</a></div><div id="w1-collapse" class="collapse navbar-collapse"><ul id="w2" class="navbar-nav nav"><li><a href="/">首页</a></li>
<li><a href="/tag">标签</a></li>
<li><a href="/about-me">About</a></li>
<form class="navbar-form navbar-left" method="get" action="/search">
    <div class="form-group">
        <input type="text" name="keyword" class="form-control" placeholder="Search">
    </div>
</form></ul><ul id="w3" class="navbar-nav navbar-right nav"><li><a href="/manage/login">登录</a></li></ul></div></div></nav><div class="container">
    <div class="row">

        <div class="col-md-9" id="main">
            <div id="w0" class="list-view"><div data-key="36"><article class="panel panel-default">
    <div class="panel-body post-title">
        <h1>
            <a class="article-title" href="/2017/11/30/muilti-nginx-https-domain">Nginx配置多个HTTPS域名</a>        </h1>

        <div class="article-meta">
            <span class="post-date">
                <span class="glyphicon glyphicon-time"></span>
                <time>2017-11-30</time>
            </span>
                    </div>
    </div>

    <div class="panel-body markdown-body"><p>最近在玩微信小程序，手头有：</p>
<ul>
<li>一台云服务器：CentOS 7</li>
<li>多个一级域名</li>
</ul>
<p>开发测试过程中，因为某些原因，想要让手头的A、B域名同时指向云服务器的443端口，支持HTTPS。</p>
<p>Nginx支持TLS协议的SNI扩展（同一个IP上可以支持多个不同证书的域名），只需要重新安装Nginx，使其支持TLS即可。</p>
<h3>安装Nginx</h3>
<pre><code class="language-source-shell">[root]#  wget http://nginx.org/download/nginx-1.12.0.tar.gz
[root]#  tar zxvf nginx-1.12.0.tar.gz
[root]#  cd nginx-1.12.0
[root]#  ./configure --prefix=/usr/local/nginx --with-http_ssl_module \
--with-openssl=./openssl-1.0.1e \
--with-openssl-opt="enable-tlsext"

</code></pre>
<p>备注：在安装的过程中发现，云服务器的环境中缺少一些库，下载后，重新执行Nginx的<code>./configure</code>指令，具体操作如下：</p>
</div>

    <div class="panel-body">
                    <a class="badge" href="/tag/Nginx">Nginx</a>                    <a class="badge" href="/tag/LetsEncrypt">LetsEncrypt</a>                    <a class="badge" href="/tag/CentOS">CentOS</a>                    <a class="badge" href="/tag/HTTPS">HTTPS</a>                    <a class="badge" href="/tag/CetBot">CetBot</a>

        <div class="pull-right">
            <div class="article-more-link">
                <a href="/2017/11/30/muilti-nginx-https-domain">Read More</a>            </div>
        </div>
    </div>
</article>

</div>
<div data-key="35"><article class="panel panel-default">
    <div class="panel-body post-title">
        <h1>
            <a class="article-title" href="/2017/11/30/point-for-css3-flexbox">聊聊CSS3 Flexbox</a>        </h1>

        <div class="article-meta">
            <span class="post-date">
                <span class="glyphicon glyphicon-time"></span>
                <time>2017-11-30</time>
            </span>
                    </div>
    </div>

    <div class="panel-body markdown-body"><p>本文涉及内容如下： flexbox的基本概念、容器属性学习、项目属性学习、实战演练。 flexbox 堪称布局神器，但属性实在太多让人无从下手，因此将自己所学的知识做个总结。</p>
<h2>基本概念</h2>
<p>flexbox是Flexible Box的缩写，译为弹性布局。flex 布局主要是让容器中的子项目可以根据配置改变自身的宽高及顺序，以最佳的方式填充容器，达到<code>弹性</code>的目的。容器有横轴和纵轴来划分容器，横轴默认为水平方向纵轴为垂直方向。</p>
<p><img src="http://oss.ipaoyun.com/blog/1-1512030664.png" alt="聊聊CSS3 Flexbox"></p>
</div>

    <div class="panel-body">
                    <a class="badge" href="/tag/Flexbox">Flexbox</a>                    <a class="badge" href="/tag/Flex">Flex</a>                    <a class="badge" href="/tag/CSS3">CSS3</a>                    <a class="badge" href="/tag/%E5%B8%83%E5%B1%80">布局</a>

        <div class="pull-right">
            <div class="article-more-link">
                <a href="/2017/11/30/point-for-css3-flexbox">Read More</a>            </div>
        </div>
    </div>
</article>

</div>
<div data-key="34"><article class="panel panel-default">
    <div class="panel-body post-title">
        <h1>
            <a class="article-title" href="/2017/11/30/css3-flexbox-layout-course">理解CSS3 Flexbox 布局</a>        </h1>

        <div class="article-meta">
            <span class="post-date">
                <span class="glyphicon glyphicon-time"></span>
                <time>2017-11-30</time>
            </span>
                    </div>
    </div>

    <div class="panel-body markdown-body"><h2>Flexbox 让人困惑</h2>
<p>有很多谈及 Flexbox 的文章，但依然有不少前端对此感到困惑。一方面，flex 相关的 CSS 属性繁多，影响到的具体效果也包含多个方面；另一方面，CSS 可以使用 <strong>Shorthand properties</strong> 风格的写法（例如最常见的 <code>background: url(images/bg.gif) no-repeat left top;</code>），很容易让新手弄不清具体含义。</p>
<p>这篇文章要讲的 Flexbox 当然还是 CSS3 规范中的弹性盒模型，不过写出前面一段，是因为我希望这篇文章可以解决那些问题——简单说，就是 Flexbox 让人困惑这样的问题。解决的方法，就是<strong>理解它</strong>。</p>
</div>

    <div class="panel-body">
                    <a class="badge" href="/tag/Flexbox">Flexbox</a>                    <a class="badge" href="/tag/Flex">Flex</a>                    <a class="badge" href="/tag/CSS3">CSS3</a>                    <a class="badge" href="/tag/Html5">Html5</a>

        <div class="pull-right">
            <div class="article-more-link">
                <a href="/2017/11/30/css3-flexbox-layout-course">Read More</a>            </div>
        </div>
    </div>
</article>

</div>
<div data-key="33"><article class="panel panel-default">
    <div class="panel-body post-title">
        <h1>
            <a class="article-title" href="/2017/11/30/git-flight-rules">Git飞行规则(Flight Rules)</a>        </h1>

        <div class="article-meta">
            <span class="post-date">
                <span class="glyphicon glyphicon-time"></span>
                <time>2017-11-30</time>
            </span>
                    </div>
    </div>

    <div class="panel-body markdown-body"><h4>前言</h4>
<ul>
<li>英文原版<a href="https://github.com/k88hudson/git-flight-rules/blob/master/README.md">README</a></li>
<li>翻译可能存在错误或不标准的地方，欢迎大家指正和修改，谢谢！</li>
</ul>
<h4>什么是"飞行规则"?</h4>
<p>一个 <a href="http://www.jsc.nasa.gov/news/columbia/fr_generic.pdf">宇航员指南</a> (现在, 程序员们都在使用GIT) 是关于出现问题过后应该怎么操作。</p>
<blockquote><p> <em>飞行规则(Flight Rules)</em> 是记录在手册上的来之不易的一系列知识，记录了某个事情发生的原因，以及怎样一步一步的进行处理。本质上, 它们是特定场景的非常详细的标准处理流程。 [...]</p>
</blockquote>
<blockquote><p>自20世纪60年代初以来，NASA一直在捕捉(capturing)我们的失误，灾难和解决方案, 当时水星时代(Mercury-era)的地面小组首先开始将“经验教训”收集到一个纲要(compendium)中，该纲现在已经有上千个问题情景，从发动机故障到破损的舱口把手到计算机故障，以及它们对应的解决方案。</p>
</blockquote>
<p>— Chris Hadfield, <em>一个宇航员的生活指南(An Astronaut's Guide to Life)</em>。</p>
</div>

    <div class="panel-body">
                    <a class="badge" href="/tag/Git">Git</a>

        <div class="pull-right">
            <div class="article-more-link">
                <a href="/2017/11/30/git-flight-rules">Read More</a>            </div>
        </div>
    </div>
</article>

</div>
<div data-key="32"><article class="panel panel-default">
    <div class="panel-body post-title">
        <h1>
            <a class="article-title" href="/2017/11/30/receive-payload-params-from-requery">Yii2.0中接收Payload参数值</a>        </h1>

        <div class="article-meta">
            <span class="post-date">
                <span class="glyphicon glyphicon-time"></span>
                <time>2017-11-30</time>
            </span>
                    </div>
    </div>

    <div class="panel-body markdown-body"><p>在一些前端Ajax请求中，发送到后端数据以Playload，这时在后端直接用<code>Yii::app()-&gt;getRequest()-&gt;get("xxx")</code>是无法直接获取到值的。需要在配置文件中加入相应的<code>parser</code>才可以实现直接取参数，修改配置文件<code>main.php</code>配置如下：</p>
<pre><code class="language-php hljs"><span class="hljs-string">'components'</span>=&gt;[
	<span class="hljs-comment">//...</span>
	<span class="hljs-string">'request'</span>      =&gt; [
		<span class="hljs-string">'parsers'</span> =&gt; [
			<span class="hljs-string">'application/json'</span> =&gt; <span class="hljs-string">'yii\web\JsonParser'</span>,
			<span class="hljs-string">'text/json'</span>        =&gt; <span class="hljs-string">'yii\web\JsonParser'</span>,
		],
	],
]
</code></pre>
<p>这时通过<code>Yii::app()-&gt;getRequest()-&gt;get("xxx")</code>方法就可以获取前端传递参数值了。</p>
</div>

    <div class="panel-body">
                    <a class="badge" href="/tag/php">php</a>                    <a class="badge" href="/tag/Yii2">Yii2</a>                    <a class="badge" href="/tag/payload">payload</a>                    <a class="badge" href="/tag/Ajax">Ajax</a>                    <a class="badge" href="/tag/JsonParser">JsonParser</a>

        <div class="pull-right">
            <div class="article-more-link">
                <a href="/2017/11/30/receive-payload-params-from-requery">Read More</a>            </div>
        </div>
    </div>
</article>

</div>
<div data-key="31"><article class="panel panel-default">
    <div class="panel-body post-title">
        <h1>
            <a class="article-title" href="/2017/11/29/replace-yii2-assets-files-to-cdn-resources">Yii2.0中静态资源生产环境替换为CDN资源</a>        </h1>

        <div class="article-meta">
            <span class="post-date">
                <span class="glyphicon glyphicon-time"></span>
                <time>2017-11-29</time>
            </span>
                    </div>
    </div>

    <div class="panel-body markdown-body"><p>Yii2.0 通过使用Assets的方式注册静态资源来管理资源依赖。有些比较大的文件（如：jQueryUI,bootstrap）在生产环境中替换为CDN资源可以加快页面访问速度。</p>
<p>配置方式如下：
修改 <code>environment/product/xxxx/config/main-local.php</code>文件，在<code>common</code>配置中添加如下规则,</p>
<pre><code class="language-php hljs"><span class="hljs-string">'components'</span>=&gt;[
		<span class="hljs-comment">// ...</span>
        <span class="hljs-string">'assetManager'</span> =&gt; [
            <span class="hljs-string">'bundles'</span> =&gt; [
                <span class="hljs-string">'yii\web\JqueryAsset'</span> =&gt; [
                    <span class="hljs-string">'jsOptions'</span> =&gt; [
                        <span class="hljs-string">'position'</span> =&gt; \yii\web\View::POS_HEAD,
                    ],
                    <span class="hljs-string">'js'</span>        =&gt; [
                        <span class="hljs-string">'//cdn.staticfile.org/jquery/3.2.1/jquery.min.js'</span>,
                    ],
                ],
            ],
        ],
		<span class="hljs-comment">// ...</span>
]
</code></pre>
</div>

    <div class="panel-body">
                    <a class="badge" href="/tag/Yii2">Yii2</a>                    <a class="badge" href="/tag/Assets">Assets</a>                    <a class="badge" href="/tag/CDN">CDN</a>                    <a class="badge" href="/tag/%E9%9D%99%E6%80%81%E8%B5%84%E6%BA%90">静态资源</a>

        <div class="pull-right">
            <div class="article-more-link">
                <a href="/2017/11/29/replace-yii2-assets-files-to-cdn-resources">Read More</a>            </div>
        </div>
    </div>
</article>

</div>
<div data-key="30"><article class="panel panel-default">
    <div class="panel-body post-title">
        <h1>
            <a class="article-title" href="/2017/11/29/xiaomi-mobile-after-can-not-write-system-directory">小米手机root后MIUI系统System目录不可写问题解决</a>        </h1>

        <div class="article-meta">
            <span class="post-date">
                <span class="glyphicon glyphicon-time"></span>
                <time>2017-11-29</time>
            </span>
                    </div>
    </div>

    <div class="panel-body markdown-body"><p>小米手机root后MIUI系统System目录直接写会提示没有权限，需要重新挂载目录进行权限切换，先把手机连接电脑，使用adb连接手机（adb 1.032以上版本），运行下面的命令：</p>
<pre><code class="hljs bash">adb root
adb <span class="hljs-built_in">disable</span>-verity
adb reboot
adb shell
mount -o rw,remount,rw /system
</code></pre>
<p>测试下效果，删除自带的搜狗输入法</p>
<pre><code class="hljs nginx"><span class="hljs-attribute">rm</span> -r /system/app/SogouInput
</code></pre>
</div>

    <div class="panel-body">
                    <a class="badge" href="/tag/root">root</a>                    <a class="badge" href="/tag/adb">adb</a>                    <a class="badge" href="/tag/MIUI">MIUI</a>                    <a class="badge" href="/tag/%E5%B0%8F%E7%B1%B3">小米</a>

        <div class="pull-right">
            <div class="article-more-link">
                <a href="/2017/11/29/xiaomi-mobile-after-can-not-write-system-directory">Read More</a>            </div>
        </div>
    </div>
</article>

</div>
<div data-key="28"><article class="panel panel-default">
    <div class="panel-body post-title">
        <h1>
            <a class="article-title" href="/2017/11/29/use-nginx-reverse-proxy-to-access-wechat-mp-images">使用Nginx反向代理实现微信公众号图片防盗链的破解</a>        </h1>

        <div class="article-meta">
            <span class="post-date">
                <span class="glyphicon glyphicon-time"></span>
                <time>2017-11-29</time>
            </span>
                    </div>
    </div>

    <div class="panel-body markdown-body"><p>微信公众号图片现在对非TX系的域名下直接引用加入了防盗链机制，我们可以通过使用Nginx反向代理来进行访问，不过这么实现会有下面的问题：</p>
<ul>
<li>流量走自己的服务器（建议网页配置LazyLoad，防止页面图片对服务器一次次加载造成浪费）</li>
<li>需要有自己的域名（直接使用服务器IP也可以）</li>
<li>国内要有主机（国外主机的速度让你有想死的冲动）</li>
</ul>
</div>

    <div class="panel-body">
                    <a class="badge" href="/tag/Nginx">Nginx</a>                    <a class="badge" href="/tag/%E5%8F%8D%E5%90%91%E4%BB%A3%E7%90%86">反向代理</a>                    <a class="badge" href="/tag/%E9%98%B2%E7%9B%97%E9%93%BE">防盗链</a>

        <div class="pull-right">
            <div class="article-more-link">
                <a href="/2017/11/29/use-nginx-reverse-proxy-to-access-wechat-mp-images">Read More</a>            </div>
        </div>
    </div>
</article>

</div>
<div data-key="27"><article class="panel panel-default">
    <div class="panel-body post-title">
        <h1>
            <a class="article-title" href="/2017/11/29/golang-connect-to-sim900a-send-receive-sms">Golang连接Sim900A模块进行短信收发</a>        </h1>

        <div class="article-meta">
            <span class="post-date">
                <span class="glyphicon glyphicon-time"></span>
                <time>2017-11-29</time>
            </span>
                    </div>
    </div>

    <div class="panel-body markdown-body"><pre><code class="language-go">package main

import(
    "fmt"
    "time"
    "github.com/xlab/at"
    //"github.com/xlab/at/sms"
)

const (
    CommandPort = "/dev/ttyUSB0"
    NotifyPort = "/dev/ttyUSB0"
)

var device *at.Device
func main(){
    fmt.Println("BEGIN...")
    device = &amp;at.Device{
        CommandPort: CommandPort,
        NotifyPort: NotifyPort,
    }

    if err := device.Open(); err!=nil{
        fmt.Println("ERR:", err);
        return
    }

    if err := device.Init(at.DeviceSIM900A()); err != nil {
        fmt.Println("ERR:", err);
        return
    }
    defer device.Close()

    CNUM, err := device.Send("AT+CNUM")
    if err != nil {
        fmt.Println("ERR:", err);
        return
    }

    fmt.Println("CNUM: ", CNUM)
    fmt.Println("IMEI: ", device.State.IMEI)
    fmt.Println("ModelName: ", device.State.ModelName)
    fmt.Println("OperatorName: ", device.State.OperatorName)
    fmt.Println("SignalStrength: ", device.State.SignalStrength)



	_,err = device.Send("AT+CMGF=0")
    if err != nil {
        fmt.Println("ERR:", err);
        return
    }

    if err != nil {
        fmt.Println("ERR:", err);
        return
    }

    fmt.Println("pending send sms..." )
    //phoneNumber := sms.PhoneNumber("18601013734")
    //phoneSMSTxt := "你好小张！"
    //err = device.SendSMS(phoneSMSTxt,phoneNumber)
    //if err != nil {
        //fmt.Println("ERR:", err);
        //return
    //}
    fmt.Println("DONE")

	go checkIncomeMsg()
	for {
		time.Sleep(time.Second)
	}
}

func checkIncomeMsg(){
	fmt.Println("begin check...")
	t := time.NewTicker(time.Second*5)
	defer t.Stop()

	//balanceUSSD := "*100#"
	go func(){
		device.Watch()
	}()

	for {
		select {
			case <- device.Closed():
				fmt.Println("device closed...")
				return
			case msg,ok:= &lt;-device.IncomingSms():
				fmt.Println("PONG: new message incomming...")
				if ok {
					//fmt.Printf("NEW MSG: %+v\n",msg)
					fmt.Println("=====================[ NEW MESSAGE ]==================")
					fmt.Println("From: ", msg.Address)
					fmt.Println("Time: ", time.Time(msg.ServiceCenterTime).Unix())
					fmt.Println("Text: ", msg.Text)
					fmt.Println("========================================================")
				}
			case &lt;-t.C:
				continue
		}
	}
}
</code></pre>
</div>

    <div class="panel-body">
                    <a class="badge" href="/tag/Sim900A">Sim900A</a>                    <a class="badge" href="/tag/Golang">Golang</a>                    <a class="badge" href="/tag/%E7%9F%AD%E4%BF%A1">短信</a>

        <div class="pull-right">
            <div class="article-more-link">
                <a href="/2017/11/29/golang-connect-to-sim900a-send-receive-sms">Read More</a>            </div>
        </div>
    </div>
</article>

</div>
<div data-key="26"><article class="panel panel-default">
    <div class="panel-body post-title">
        <h1>
            <a class="article-title" href="/2017/11/29/linux-query-network-status-scripts">Linux网络相关查询脚本</a>        </h1>

        <div class="article-meta">
            <span class="post-date">
                <span class="glyphicon glyphicon-time"></span>
                <time>2017-11-29</time>
            </span>
                    </div>
    </div>

    <div class="panel-body markdown-body"><ul>
<li>查看TCP连接状态<pre><code class="hljs ruby">netstat -nat <span class="hljs-params">|awk '{print $6}'|</span>sort<span class="hljs-params">|uniq -c|</span>sort -rn
netstat -n <span class="hljs-params">| awk '/^tcp/ {++S[$NF]};<span class="hljs-keyword">END</span> {<span class="hljs-keyword">for</span>(a <span class="hljs-keyword">in</span> S) print a, S[a]}' 或
netstat -n |</span> awk <span class="hljs-string">'/^tcp/ {++state[$NF]}; END {for(key in state) print key,"\t",state[key]}'</span>
netstat -n <span class="hljs-params">| awk '/^tcp/ {++arr[$NF]};<span class="hljs-keyword">END</span> {<span class="hljs-keyword">for</span>(k <span class="hljs-keyword">in</span> arr) print k,"\t",arr[k]}'
netstat -n |</span>awk <span class="hljs-string">'/^tcp/ {print $NF}'</span><span class="hljs-params">|sort|</span>uniq -c<span class="hljs-params">|sort -rn
netstat -ant |</span> awk <span class="hljs-string">'{print $NF}'</span> <span class="hljs-params">| grep -v '[a-z]' |</span> sort <span class="hljs-params">| uniq -c
</span></code></pre>
</li>
</ul>
<p>(以上每一行实现的效果基本相同，在此列出不同的写法，方便对脚本写法的更深理解）</p>
</div>

    <div class="panel-body">
                    <a class="badge" href="/tag/Linux">Linux</a>                    <a class="badge" href="/tag/%E7%BD%91%E7%BB%9C">网络</a>                    <a class="badge" href="/tag/TCP">TCP</a>

        <div class="pull-right">
            <div class="article-more-link">
                <a href="/2017/11/29/linux-query-network-status-scripts">Read More</a>            </div>
        </div>
    </div>
</article>

</div>
<div class="text-center"><ul class="pagination"><li class="prev disabled"><span>上一页</span></li>
<li class="active"><a href="/page/1" data-page="0">1</a></li>
<li><a href="/page/2" data-page="1">2</a></li>
<li><a href="/page/3" data-page="2">3</a></li>
<li><a href="/page/4" data-page="3">4</a></li>
<li class="next"><a href="/page/2" data-page="1">下一页</a></li></ul></div></div>
        </div>

        <div class="col-md-3">
            <div id="profile">
                <div class="panel panel-default text-center">
                    <div class="panel-body">
                        <img id="avatar" src="/images/avatar.jpg" alt="">                    </div>
                    <div class="panel-body">
                        <h2 id="name">Rogee</h2>
                        <h3 id="title">Web Developer &amp; Designer</h3>
                    </div>
                    <div class="panel-body" id="follow">
                        <a class="btn btn-info" href="https://github.com/rogeecn" style="border-radius: 34px" target="_blank">FOLLOW</a>                    </div>
                    <table class="table table-bordered" id="post-info">
                        <tbody><tr>
                            <td>
                                35                                <span>POSTS</span>
                            </td>
                            <td>
                                78                                <span>TAGS</span>
                            </td>
                        </tr>
                    </tbody></table>
                    <!--                    <table class="table" id="sns-info">-->
                    <!--                        <tr>-->
                    <!--                            <td><a href="#"><span class="glyphicon glyphicon-apple"></span></a></td>-->
                    <!--                            <td><a href="#"><span class="glyphicon glyphicon-apple"></span></a></td>-->
                    <!--                            <td><a href="#"><span class="glyphicon glyphicon-apple"></span></a></td>-->
                    <!--                            <td><a href="#"><span class="glyphicon glyphicon-apple"></span></a></td>-->
                    <!--                        </tr>-->
                    <!--                    </table>-->
                </div>
            </div>
        </div>
    </div>
</div>
<footer id="footer">
    <div class="container">
        <p>Copyright © 2017 <a href="/">711XD</a>. All Rights Reserved</p>
        <p><a href="http://www.miitbeian.gov.cn/" target="_blank">京ICP备17027412号</a></p>
    </div>
</footer>
<script src="//cdn.staticfile.org/highlight.js/9.12.0/highlight.min.js"></script>
<script src="//cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<script>
    $(function () {
        $('pre code').each(function (i, block) {
            hljs.highlightBlock(block);
        });
    })
</script>



<audio controls="controls" style="display: none;"></audio></body>`

	list, err := AllLinks(htmlBody, "href", FetchRule{
		StartWith: "/2017",
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("link cnt: %d", len(list))
	for _, item := range list {
		t.Log(item)
	}
	t.Log("done")
}
