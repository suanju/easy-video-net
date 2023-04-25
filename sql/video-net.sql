-- MySQL dump 10.13  Distrib 8.0.24, for Linux (x86_64)
--
-- Host: localhost    Database: video-net
-- ------------------------------------------------------
-- Server version	8.0.24

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `lv_article_classification`
--

DROP TABLE IF EXISTS `lv_article_classification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_article_classification` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `a_id` int NOT NULL DEFAULT '0' COMMENT '上级id 0表示顶级',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名称',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章分类';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_article_classification`
--

LOCK TABLES `lv_article_classification` WRITE;
/*!40000 ALTER TABLE `lv_article_classification` DISABLE KEYS */;
INSERT INTO `lv_article_classification` VALUES (1,0,'生活倒影','2023-02-05 00:00:00.000','2023-02-05 00:00:00.000'),(2,0,'学习人生','2023-02-05 00:00:00.000','2023-02-05 00:00:00.000'),(3,0,'游戏娱乐','2023-02-05 00:00:00.000','2023-02-05 00:00:00.000'),(4,0,'视听盛宴','2023-02-05 00:00:00.000','2023-02-05 00:00:00.000'),(5,1,'记录生活','2023-02-05 00:00:00.000','2023-02-05 00:00:00.000'),(6,1,'随手拍','2023-02-05 00:00:00.000','2023-02-05 00:00:00.000'),(7,1,'照片分享','2023-02-05 00:00:00.000','2023-02-05 00:00:00.000'),(8,6,'相机拍','2023-02-05 00:00:00.000','2023-02-05 00:00:00.000'),(9,6,'手机拍','2023-02-05 00:00:00.000','2023-02-05 00:00:00.000');
/*!40000 ALTER TABLE `lv_article_classification` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_article_contribution`
--

DROP TABLE IF EXISTS `lv_article_contribution`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_article_contribution` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户ID',
  `classification_id` int DEFAULT NULL COMMENT '分类ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章标题',
  `cover` json NOT NULL COMMENT '文章封面',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标签定义，分割	',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '富文本存储值',
  `content_storage_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '富文本存储类型',
  `is_comments` int NOT NULL DEFAULT '1' COMMENT '是否可以评论0否1是',
  `heat` int DEFAULT '0' COMMENT '文章热度',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_article_contribution`
--

LOCK TABLES `lv_article_contribution` WRITE;
/*!40000 ALTER TABLE `lv_article_contribution` DISABLE KEYS */;
INSERT INTO `lv_article_contribution` VALUES (19,30,2,'Go语言的错误error处理的方案','{\"src\": \"assets/static/img/articleContributionCover2c821364af4b3c683c6272e20bc1801dada4fe3d538934f6416e70c59c87ed44.jpg\", \"type\": \"aliyunOss\"}','golang ,学习','<p><span style=\"color: rgb(18, 18, 18);\">对于Go语言（golang）的错误设计，相信很多人已经体验过了，它是通过返回值的方式，来强迫调用者对错误进行处理，要么你忽略，要么你处理（处理也可以是继续返回给调用者），对于golang这种设计方式，我们会在代码中写大量的</span><code style=\"color: rgb(18, 18, 18); background-color: rgb(246, 246, 246);\">if</code><span style=\"color: rgb(18, 18, 18);\">判断，以便做出决定。</span></p><p><br></p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-function\"><span class=\"hljs-keyword\">func</span> <span class=\"hljs-title\">main</span><span class=\"hljs-params\">()</span></span> {\n    conent,err:=ioutil.ReadFile(<span class=\"hljs-string\">\"filepath\"</span>)\n    <span class=\"hljs-keyword\">if</span> err !=<span class=\"hljs-literal\">nil</span>{\n        <span class=\"hljs-comment\">//错误处理</span>\n    }<span class=\"hljs-keyword\">else</span> {\n        fmt.Println(<span class=\"hljs-type\">string</span>(conent))\n    }\n}\n</pre><p>这类代码，在我们编码中是非常的，大部分情况下<code style=\"background-color: rgb(246, 246, 246);\">error</code>都是<code style=\"background-color: rgb(246, 246, 246);\">nil</code>，也就是没有任何错误，但是非<code style=\"background-color: rgb(246, 246, 246);\">nil</code>的时候，意味着错误就出现了，我们需要对他进行处理。</p><h2>error 接口</h2><p><code style=\"background-color: rgb(246, 246, 246);\">error</code>其实一个接口，内置的，我们看下它的定义</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-comment\">// The error built-in interface type is the conventional interface for</span>\n<span class=\"hljs-comment\">// representing an error condition, with the nil value representing no error.</span>\n<span class=\"hljs-built_in\">type</span> <span class=\"hljs-built_in\">error</span> interface {\n    <span class=\"hljs-built_in\">Error</span>() <span class=\"hljs-keyword\">string</span>\n}\n</pre><p><br></p><p><span style=\"color: rgb(18, 18, 18);\">它只有一个方法&nbsp;</span><code style=\"color: rgb(18, 18, 18); background-color: rgb(246, 246, 246);\">Error</code><span style=\"color: rgb(18, 18, 18);\">，只要实现了这个方法，就是实现了</span><code style=\"color: rgb(18, 18, 18); background-color: rgb(246, 246, 246);\">error</code><span style=\"color: rgb(18, 18, 18);\">。现在我们自己定义一个错误试试。</span></p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">type</span> fileError <span class=\"hljs-keyword\">struct</span> {\n}\n\n<span class=\"hljs-function\"><span class=\"hljs-keyword\">func</span> <span class=\"hljs-params\">(fe *fileError)</span></span> Error() <span class=\"hljs-type\">string</span> {\n    <span class=\"hljs-keyword\">return</span> <span class=\"hljs-string\">\"文件错误\"</span>\n}\n</pre><h2>自定义 error</h2><p>自定义了一个<code style=\"background-color: rgb(246, 246, 246);\">fileError</code>类型，实现了<code style=\"background-color: rgb(246, 246, 246);\">error</code>接口。现在测试下看看效果。</p><p><br></p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-function\"><span class=\"hljs-keyword\">func</span> <span class=\"hljs-title\">main</span><span class=\"hljs-params\">()</span></span> {\n    conent, err := openFile()\n    <span class=\"hljs-keyword\">if</span> err != <span class=\"hljs-literal\">nil</span> {\n        fmt.Println(err)\n    } <span class=\"hljs-keyword\">else</span> {\n        fmt.Println(<span class=\"hljs-type\">string</span>(conent))\n    }\n}\n\n<span class=\"hljs-comment\">//只是模拟一个错误</span>\n<span class=\"hljs-function\"><span class=\"hljs-keyword\">func</span> <span class=\"hljs-title\">openFile</span><span class=\"hljs-params\">()</span></span> ([]<span class=\"hljs-type\">byte</span>, <span class=\"hljs-type\">error</span>) {\n    <span class=\"hljs-keyword\">return</span> <span class=\"hljs-literal\">nil</span>, &amp;fileError{}\n}\n</pre><p><br></p><p>我们运行模拟的代码，可以看到<code style=\"background-color: rgb(246, 246, 246);\">文件错误</code>的通知。</p><p>在实际的使用过程中，我们可能遇到很多错误，他们的区别是错误信息不一样，一种做法是每种错误都类似上面一样定义一个错误类型，但是这样太麻烦了。我们发现<code style=\"background-color: rgb(246, 246, 246);\">Error</code>返回的其实是个字符串，我们可以修改下，让这个字符串可以设置就可以了。</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-built_in\">type</span> fileError <span class=\"hljs-keyword\">struct</span> {\n    s <span class=\"hljs-keyword\">string</span>\n}\n\n<span class=\"hljs-built_in\">func</span> (fe *fileError) <span class=\"hljs-built_in\">Error</span>() <span class=\"hljs-keyword\">string</span> {\n    <span class=\"hljs-keyword\">return</span> fe.s\n}\n\n</pre>','aliyunOss',1,1,'2023-04-19 11:18:53.875','2023-04-19 11:20:21.106'),(20,30,2,'理解JavaScript Promise','{\"src\": \"assets/static/img/articleContributionCover8bf8e02da8af8c15bee0a5aeb40dd6c2fc1b3da8516f1cf28232feb5f69f6323.png\", \"type\": \"aliyunOss\"}','js','<h2>本文目标</h2><p>从一脸懵逼，到认识、掌握并使用Promise……</p><p>ES2015出来之后，什么箭头函数、类声明、解构赋值等新概念冒了出来，其中一个重要的概念就是Promise。不学会Promise，都不好意思说自己懂ES6了！</p><p>那么Promise到底是个什么鬼？！像我这样的JavaScript小菜，初次看到，着实一脸懵逼。</p><h2>本文结构</h2><p>本文从下面几个问题，逐步了解JavaScript Promise的概念和用法，重点是概念啦，用法什么的网上有很多资料参考</p><ul><li>Promise是什么？</li><li>有什么用？解决了什么问题？</li></ul><p>怎么使用Promise?</p><ul><li><span style=\"color: rgb(18, 18, 18);\">剩下部分整理了一些Promise习题和示例，以及参考资料</span>习题和示例</li></ul><p>参考资料</p><h2>是什么</h2><p>当我们谈到，或者看到别人提及Promise的时候，什么回调函数、异步编程、流程控制这样的术语冒了出来，还有，取了个Promise这样的名字，究竟是何含义。首先通通都别care，没那么复杂。</p><p>Promise说得通俗一点就是一种写代码的方式，并且是用来写JavaScript编程中的异步代码的。</p><p>基本用法</p><p>首先要认清最基本的用法。一般学习Promise看到的第一段代码是这样：</p><pre class=\"ql-syntax\" spellcheck=\"false\">let p = new Promise((resolve, reject) =&gt; {\n  // 做一些事情\n  // 然后在某些条件下resolve，或者reject\n  if (/* 条件随便写^_^ */) {\n    resolve()\n  } else {\n    reject()\n  }\n})\n\np.then(() =&gt; {\n    // 如果p的状态被resolve了，就进入这里\n}, () =&gt; {\n    // 如果p的状态被reject\n})\n</pre><p>解释一下</p><p>第一段调用了Promise构造函数，第二段是调用了promise实例的.then方法</p><p>1. 构造实例</p><ul><li>构造函数接受一个函数作为参数</li><li>调用构造函数得到实例p的同时，作为参数的函数会立即执行</li><li>参数函数接受两个回调函数参数resolve和reject</li><li>在参数函数被执行的过程中，如果在其内部调用resolve，会将p的状态变成fulfilled，或者调用reject，会将p的状态变成rejected</li></ul><p><br></p><p>2. 调用.then</p><ul><li>调用.then可以为实例p注册两种状态回调函数</li><li>当实例p的状态为fulfilled，会触发第一个函数执行</li></ul><p>当实例p的状态为rejected，则触发第二个函数执行</p><p><span style=\"color: rgb(18, 18, 18);\">总结</span></p><p>上面这样构造promise实例，然后调用.then.then.then的编写代码方式，就是promise。</p><p>其基本模式是：</p><p>将异步过程转化成promise对象对象有3种状态通过.then注册状态的回调已完成的状态能触发回调</p><p><span style=\"color: rgb(18, 18, 18);\">采用这种方式来处理编程中的异步任务，就是在使用promise了。</span></p><p>所以promise就是一种异步编程模式。</p><h2>有什么用</h2><p>从JavaScript中的异步任务说起</p><p>典型的异步任务，是一个setTimeout调用</p><pre class=\"ql-syntax\" spellcheck=\"false\">setTimeout(() =&gt; {\n  // 爱干啥干啥\n}, 1000)\n</pre><p>通常用JS写异步任务的时候，会分成两个部分：主过程和后续过程，在主过程执行成功后，触发后续过程执行。比如，在实际编程中经常需要使用AJAX向服务器请求数据，成功获取到数据后，才开始处理数据。于是代码分成获取数据部分和处理数据部分，像下面这样：</p><pre class=\"ql-syntax\" spellcheck=\"false\">getData((data) =&gt; {\n  // 处理data\n})\n</pre><p>上面这两个处理异步任务的编程方式都是采用的回调函数的形式。</p><p>那么既然有回调函数这种方式，为什么还需要promise？</p><p>当然是因为回调函数这种方式不好了，否则谁吃饱了没事干非要整个promise出来！</p><p>也许有人说，觉得回调函数的方式没什么不好啊，我也不想否认，就像有人用习惯了windows操作系统，你说windows垃圾，改用Linux吧，balabala……其实用者开心，用得舒服就好啊。然而系统是自己用的，代码却是写给别人看的。</p><p>回调哪里不好了？</p><p>现在假设有多个异步任务，且任务间有依赖关系（一个任务需要拿到另一个任务成功后的结果才能开始执行）的时候，回调的方式写出来的代码就会像下面这样：</p><pre class=\"ql-syntax\" spellcheck=\"false\">getData1(data1 =&gt; {\n  getData2(data1, data2 =&gt; {\n    getData3(data2, data3 =&gt; {\n      getData4(data3, data4 =&gt; {\n        getData5(data4, data5 =&gt; {\n          // 终于取到data5了\n        })\n      })\n    })\n  })\n})\n</pre><p>这种代码被称为回调地狱或者回调金字塔</p><p>假设上面的任务，想要换一下执行顺序，代码修改起来，就比较麻烦了。如果内容复杂，阅读代码的时候跳来跳去，也让人头大。</p><p>如果用promise改写一下：</p><pre class=\"ql-syntax\" spellcheck=\"false\">// 先把getData们都转成返回promise对象的函数\n\n// 然后\ngetData1()\n.then(getData2)\n.then(getData3)\n.then(getData4)\n.then(getData5)\n.then(data =&gt; {\n  // 取到最终data了\n})\n</pre><p>这样的代码，是线性的，符合人的阅读习惯，代码表示的流程清晰，便于阅读</p><p><br></p><p><span style=\"color: rgb(18, 18, 18);\">异步任务并行</span></p><p>有了多个异步任务后，下面假设想要多个异步任务并行执行，获取执行成功后，才处理结果。用回调方式来写，可采用下面的办法：</p><pre class=\"ql-syntax\" spellcheck=\"false\">let tasks = [getData1, getData2, getData3, getData4, getData5]\nlet datas = []\n\ntasks.forEach(task =&gt; {\n  task(data =&gt; {\n    datas.push(data)\n\n    if (datas.length == tasks.length) {\n      // datas里已经包含全部的数据了\n    }\n  })\n})\n</pre><p>上面结合数组，使用了forEach方法来为每个任务传入回调。如果改用promise来实现，代码会是这样：</p><pre class=\"ql-syntax\" spellcheck=\"false\">// 先要把getData们转成promise对象\n\n// 然后\nPromise.all([\n  getData1,\n  getData2,\n  getData3,\n  getData4,\n  getData5\n]).then(datas =&gt; {\n  // 已拿到全部的data，可以处理了\n})\n</pre><p>可以看到，下面的代码更清晰易读</p><h2>如何使用</h2><p>既然promise是一种更好的编程方式，那么现在就要好好了解下它的内容和各部分的作用了。其实这节重点是内容，怎么用还得靠看官自己去实践。不过，了解内容和作用是学会使用的前提和关键。</p><p>3种状态</p><p>首先，promise实例有三种状态：</p><ul><li>pending（待定）</li><li>fulfilled（已执行）</li><li>rejected（已拒绝）</li></ul><p><br></p><p><span style=\"color: rgb(18, 18, 18);\">fulfilled和rejected有可以说是已成功和已失败，这两种状态又归为已完成状态</span></p><p>resolve和reject</p><p>调用resolve和reject能将分别将promise实例的状态变成fulfilled和rejected，只有状态变成已完成（即fulfilled和rejected之一），才能触发状态的回调</p><p>Promise API</p><p>promise的内容分为构造函数、实例方法和静态方法</p><p>1个构造函数： new Promise2个实例方法：.then 和 .catch4个静态方法：Promise.all、Promise.race、Promise.resolve和Promise.reject</p><p><span style=\"color: rgb(18, 18, 18);\">其中Promise.race不常用，本文忽略</span></p><p>下面逐个讲下他们的作用</p><p>new Promise能将一个异步过程转化成promise对象。先有了promise对象，然后才有promise编程方式。.then用于为promise对象的状态注册回调函数。它会返回一个promise对象，所以可以进行链式调用，也就是.then后面可以继续.then。在注册的状态回调函数中，可以通过return语句改变.then返回的promise对象的状态，以及向后面.then注册的状态回调传递数据；也可以不使用return语句，那样默认就是将返回的promise对象resolve。.catch用于注册rejected状态的回调函数，同时该回调也是程序出错的回调，即如果前面的程序运行过程中出错，也会进入执行该回调函数。同.then一样，也会返回新的promise对象。调用Promise.resolve会返回一个状态为fulfilled状态的promise对象，参数会作为数据传递给后面的状态回调函数Promise.reject与Promise.resolve同理，区别在于返回的promise对象状态为rejected</p><p><span style=\"color: rgb(18, 18, 18);\">（有待补充...）</span></p>','aliyunOss',1,1,'2023-04-19 11:22:53.873','2023-04-19 11:23:09.330'),(21,30,2,'一文读懂 TypeScript 泛型','{\"src\": \"assets/static/img/articleContributionCover7020a1fb05a2fc66e069866b961a6f7453719940c4203d6d43f01ca9b4bf18b6.jpg\", \"type\": \"aliyunOss\"}','ts','<p>觉得 TypeScript 泛型有点难，想系统学习 TypeScript 泛型相关知识的小伙伴们看过来，本文从八个方面入手，全方位带你一步步学习 TypeScript 中泛型，详细的内容大纲请看下图：</p><p><span style=\"background-color: transparent;\"><img src=\"https://pic3.zhimg.com/80/v2-1048cf96ac2d875f77788a3e3bbb9f32_1440w.webp\" height=\"426\" width=\"1294\"></span></p><h3>一、泛型是什么</h3><p>软件工程中，我们不仅要创建一致的定义良好的 API，同时也要考虑可重用性。 组件不仅能够支持当前的数据类型，同时也能支持未来的数据类型，这在创建大型系统时为你提供了十分灵活的功能。</p><p>「在像 C# 和 Java 这样的语言中，可以使用泛型来创建可重用的组件，一个组件可以支持多种类型的数据。 这样用户就可以以自己的数据类型来使用组件。」</p><p>设计泛型的关键目的是在成员之间提供有意义的约束，这些成员可以是：类的实例成员、类的方法、函数参数和函数返回值。</p><p>为了便于大家更好地理解上述的内容，我们来举个例子，在这个例子中，我们将一步步揭示泛型的作用。首先我们来定义一个通用的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数，该函数接收一个参数并直接返回它：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title function_\">identity</span> (<span class=\"hljs-params\">value</span>) {\n  <span class=\"hljs-keyword\">return</span> value;\n}\n\n<span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(identity(<span class=\"hljs-number\">1</span>)) <span class=\"hljs-comment\">// 1</span>\n</pre><p>现在，我们将&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数做适当的调整，以支持 TypeScript 的 Number 类型的参数：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title\">identity</span> (value: Number) : <span class=\"hljs-type\">Number</span> {\n  <span class=\"hljs-keyword\">return</span> <span class=\"hljs-type\">value</span>;\n}\n\nconsole.log(identity(<span class=\"hljs-number\">1</span>)) // <span class=\"hljs-number\">1</span>\n</pre><p>这里&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;的问题是我们将&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Number</code>&nbsp;类型分配给参数和返回类型，使该函数仅可用于该原始类型。但该函数并不是可扩展或通用的，很明显这并不是我们所希望的。</p><p>我们确实可以把&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Number</code>&nbsp;换成&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">any</code>，我们失去了定义应该返回哪种类型的能力，并且在这个过程中使编译器失去了类型保护的作用。我们的目标是让&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数可以适用于任何特定的类型，为了实现这个目标，我们可以使用泛型来解决这个问题，具体实现方式如下：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title function_\">identity</span> &lt;<span class=\"hljs-title function_\">T</span>&gt;(<span class=\"hljs-params\">value: T</span>) : <span class=\"hljs-title function_\">T</span> {\n  <span class=\"hljs-keyword\">return</span> value;\n}\n\n<span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(identity&lt;<span class=\"hljs-built_in\">Number</span>&gt;(<span class=\"hljs-number\">1</span>)) <span class=\"hljs-comment\">// 1</span>\n</pre><p>对于刚接触 TypeScript 泛型的读者来说，首次看到&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">&lt;T&gt;</code>&nbsp;语法会感到陌生。但这没什么可担心的，就像传递参数一样，我们传递了我们想要用于特定函数调用的类型。</p><p><span style=\"background-color: transparent;\"><img src=\"https://pic1.zhimg.com/80/v2-14dcc2c0b314b13f9540c90816a9fcdc_1440w.webp\" height=\"376\" width=\"776\"></span></p><p>参考上面的图片，当我们调用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity&lt;Number&gt;(1)</code>&nbsp;，<code style=\"background-color: rgb(246, 246, 246);\">Number</code>&nbsp;类型就像参数&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">1</code>&nbsp;一样，它将在出现&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;的任何位置填充该类型。图中&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">&lt;T&gt;</code>&nbsp;内部的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;被称为类型变量，它是我们希望传递给 identity 函数的类型占位符，同时它被分配给&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">value</code>&nbsp;参数用来代替它的类型：此时&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;充当的是类型，而不是特定的 Number 类型。</p><p>其中&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;代表&nbsp;「Type」，在定义泛型时通常用作第一个类型变量名称。但实际上&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;可以用任何有效名称代替。除了&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;之外，以下是常见泛型变量代表的意思：</p><ul><li>K（Key）：表示对象中的键类型；V（Value）：表示对象中的值类型；E（Element）：表示元素类型。</li></ul><p>其实并不是只能定义一个类型变量，我们可以引入希望定义的任何数量的类型变量。比如我们引入一个新的类型变量&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">U</code>，用于扩展我们定义的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title function_\">identity</span> &lt;<span class=\"hljs-title function_\">T</span>, <span class=\"hljs-title function_\">U</span>&gt;(<span class=\"hljs-params\">value: T, message: U</span>) : <span class=\"hljs-title function_\">T</span> {\n  <span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(message);\n  <span class=\"hljs-keyword\">return</span> value;\n}\n\n<span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(identity&lt;<span class=\"hljs-built_in\">Number</span>, string&gt;(<span class=\"hljs-number\">68</span>, <span class=\"hljs-string\">\"Semlinker\"</span>));\n</pre><p><span style=\"background-color: transparent;\"><img src=\"https://pic3.zhimg.com/80/v2-0610d3da83188cb823c0d1b3c215d6d2_1440w.webp\" height=\"564\" width=\"1178\"></span></p><p>除了为类型变量显式设定值之外，一种更常见的做法是使编译器自动选择这些类型，从而使代码更简洁。我们可以完全省略尖括号，比如：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title function_\">identity</span> &lt;<span class=\"hljs-title function_\">T</span>, <span class=\"hljs-title function_\">U</span>&gt;(<span class=\"hljs-params\">value: T, message: U</span>) : <span class=\"hljs-title function_\">T</span> {\n  <span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(message);\n  <span class=\"hljs-keyword\">return</span> value;\n}\n\n<span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(identity(<span class=\"hljs-number\">68</span>, <span class=\"hljs-string\">\"Semlinker\"</span>));\n</pre><p>对于上述代码，编译器足够聪明，能够知道我们的参数类型，并将它们赋值给 T 和 U，而不需要开发人员显式指定它们。下面我们来看张动图，直观地感受一下类型传递的过程：</p><p class=\"ql-align-center\"><img src=\"https://pic2.zhimg.com/v2-21710f911723bb489952093c89fad6fd_b.jpg\" alt=\"动图封面\" width=\"780\"></p><p>（图片来源：<a href=\"https://link.zhihu.com/?target=https%3A//medium.com/better-programming/typescript-generics-90be93d8c292\" rel=\"noopener noreferrer\" target=\"_blank\" style=\"color: transparent; background-color: transparent;\">https://</a><a href=\"https://link.zhihu.com/?target=https%3A//medium.com/better-programming/typescript-generics-90be93d8c292\" rel=\"noopener noreferrer\" target=\"_blank\" style=\"color: inherit;\">medium.com/better-progr</a><a href=\"https://link.zhihu.com/?target=https%3A//medium.com/better-programming/typescript-generics-90be93d8c292\" rel=\"noopener noreferrer\" target=\"_blank\" style=\"color: transparent; background-color: transparent;\">amming/typescript-generics-90be93d8c292</a>）</p><blockquote>❝ 注意：「动图应该是 —— identity&lt;Number[]&gt;([1,2,3])」</blockquote><blockquote>❞</blockquote><p>如你所见，该函数接收你传递给它的任何类型，使得我们可以为不同类型创建可重用的组件。现在我们再来看一下&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title\">identity</span> &lt;T, U&gt;(value: T, message: U) : <span class=\"hljs-type\">T</span> {\n  console.log(message);\n  <span class=\"hljs-keyword\">return</span> value;\n}\n</pre><p>相比之前定义的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数，新的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数增加了一个类型变量&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">U</code>，但该函数的返回类型我们仍然使用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>。如果我们想要返回两种类型的对象该怎么办呢？针对这个问题，我们有多种方案，其中一种就是使用元组，即为元组设置通用的类型：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title\">identity</span> &lt;T, U&gt;(value: T, message: U) : [<span class=\"hljs-type\">T</span>, U] {\n  <span class=\"hljs-keyword\">return</span> <span class=\"hljs-type\">[value,</span> message];\n}\n</pre><p>虽然使用元组解决了上述的问题，但有没有其它更好的方案呢？答案是有的，你可以使用泛型接口。</p><h3>二、泛型接口</h3><p>为了解决上面提到的问题，首先让我们创建一个用于的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数通用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Identities</code>&nbsp;接口：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">interface</span> <span class=\"hljs-symbol\">Identities</span>&lt;<span class=\"hljs-symbol\">V</span>, <span class=\"hljs-symbol\">M</span>&gt; {\n  value: V,\n  message: M\n}\n</pre><p>在上述的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Identities</code>&nbsp;接口中，我们引入了类型变量&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">V</code>&nbsp;和&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">M</code>，来进一步说明有效的字母都可以用于表示类型变量，之后我们就可以将&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Identities</code>&nbsp;接口作为&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数的返回类型：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title function_\">identity</span>&lt;<span class=\"hljs-title function_\">T</span>, <span class=\"hljs-title function_\">U</span>&gt; (<span class=\"hljs-params\">value: T, message: U</span>): <span class=\"hljs-title function_\">Identities</span>&lt;<span class=\"hljs-title function_\">T</span>, <span class=\"hljs-title function_\">U</span>&gt; {\n  <span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(value + <span class=\"hljs-string\">\": \"</span> + <span class=\"hljs-built_in\">typeof</span> (value));\n  <span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(message + <span class=\"hljs-string\">\": \"</span> + <span class=\"hljs-built_in\">typeof</span> (message));\n  let identities: Identities&lt;T, U&gt; = {\n    value,\n    message\n  };\n  <span class=\"hljs-keyword\">return</span> identities;\n}\n\n<span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(identity(<span class=\"hljs-number\">68</span>, <span class=\"hljs-string\">\"Semlinker\"</span>));\n</pre><p>以上代码成功运行后，在控制台会输出以下结果：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-attr\">68:</span> <span class=\"hljs-string\">number</span>\n<span class=\"hljs-attr\">Semlinker:</span> <span class=\"hljs-string\">string</span>\n{<span class=\"hljs-attr\">value:</span> <span class=\"hljs-number\">68</span>, <span class=\"hljs-attr\">message:</span> <span class=\"hljs-string\">\"Semlinker\"</span>}\n</pre><p>泛型除了可以应用在函数和接口之外，它也可以应用在类中，下面我们就来看一下在类中如何使用泛型。</p><h3>三、泛型类</h3><p>在类中使用泛型也很简单，我们只需要在类名后面，使用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">&lt;T, ...&gt;</code>&nbsp;的语法定义任意多个类型变量，具体示例如下：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">interface</span> <span class=\"hljs-title class_\">GenericInterface</span>&lt;U&gt; {\n  <span class=\"hljs-attr\">value</span>: U\n  <span class=\"hljs-attr\">getIdentity</span>: <span class=\"hljs-function\">() =&gt;</span> U\n}\n\n<span class=\"hljs-keyword\">class</span> <span class=\"hljs-title class_\">IdentityClass</span>&lt;T&gt; <span class=\"hljs-keyword\">implements</span> <span class=\"hljs-title class_\">GenericInterface</span>&lt;T&gt; {\n  <span class=\"hljs-attr\">value</span>: T\n\n  <span class=\"hljs-title function_\">constructor</span>(<span class=\"hljs-params\">value: T</span>) {\n    <span class=\"hljs-variable language_\">this</span>.<span class=\"hljs-property\">value</span> = value\n  }\n\n  <span class=\"hljs-title function_\">getIdentity</span>(): T {\n    <span class=\"hljs-keyword\">return</span> <span class=\"hljs-variable language_\">this</span>.<span class=\"hljs-property\">value</span>\n  }\n\n}\n\n<span class=\"hljs-keyword\">const</span> myNumberClass = <span class=\"hljs-keyword\">new</span> <span class=\"hljs-title class_\">IdentityClass</span>&lt;<span class=\"hljs-title class_\">Number</span>&gt;(<span class=\"hljs-number\">68</span>);\n<span class=\"hljs-variable language_\">console</span>.<span class=\"hljs-title function_\">log</span>(myNumberClass.<span class=\"hljs-title function_\">getIdentity</span>()); <span class=\"hljs-comment\">// 68</span>\n\n<span class=\"hljs-keyword\">const</span> myStringClass = <span class=\"hljs-keyword\">new</span> <span class=\"hljs-title class_\">IdentityClass</span>&lt;<span class=\"hljs-built_in\">string</span>&gt;(<span class=\"hljs-string\">\"Semlinker!\"</span>);\n<span class=\"hljs-variable language_\">console</span>.<span class=\"hljs-title function_\">log</span>(myStringClass.<span class=\"hljs-title function_\">getIdentity</span>()); <span class=\"hljs-comment\">// Semlinker!</span>\n</pre><p>接下来我们以实例化&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">myNumberClass</code>&nbsp;为例，来分析一下其调用过程：</p><ul><li>在实例化&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">IdentityClass</code>&nbsp;对象时，我们传入&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Number</code>&nbsp;类型和构造函数参数值&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">68</code>；之后在&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">IdentityClass</code>&nbsp;类中，类型变量&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;的值变成&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Number</code>&nbsp;类型；<code style=\"background-color: rgb(246, 246, 246);\">IdentityClass</code>&nbsp;类实现了&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">GenericInterface&lt;T&gt;</code>，而此时&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;表示&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Number</code>&nbsp;类型，因此等价于该类实现了&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">GenericInterface&lt;Number&gt;</code>&nbsp;接口；而对于&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">GenericInterface&lt;U&gt;</code>&nbsp;接口来说，类型变量&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">U</code>&nbsp;也变成了&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Number</code>。这里我有意使用不同的变量名，以表明类型值沿链向上传播，且与变量名无关。</li></ul><p>泛型类可确保在整个类中一致地使用指定的数据类型。比如，你可能已经注意到在使用 Typescript 的 React 项目中使用了以下约定：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-class\"><span class=\"hljs-keyword\">type</span> <span class=\"hljs-title\">Props</span> </span>= {\n  className?: string\n   ...\n};\n\n<span class=\"hljs-class\"><span class=\"hljs-keyword\">type</span> <span class=\"hljs-title\">State</span> </span>= {\n  submitted?: bool\n   ...\n};\n\n<span class=\"hljs-class\"><span class=\"hljs-keyword\">class</span> <span class=\"hljs-title\">MyComponent</span> <span class=\"hljs-keyword\">extends</span> <span class=\"hljs-title\">React</span>.<span class=\"hljs-title\">Component&lt;Props</span>, <span class=\"hljs-title\">State&gt;</span> </span>{\n   ...\n}\n</pre><p>在以上代码中，我们将泛型与 React 组件一起使用，以确保组件的 props 和 state 是类型安全的。</p><p>相信看到这里一些读者会有疑问，我们在什么时候需要使用泛型呢？通常在决定是否使用泛型时，我们有以下两个参考标准：</p><ul><li>当你的函数、接口或类将处理多种数据类型时；当函数、接口或类在多个地方使用该数据类型时。</li></ul><p>很有可能你没有办法保证在项目早期就使用泛型的组件，但是随着项目的发展，组件的功能通常会被扩展。这种增加的可扩展性最终很可能会满足上述两个条件，在这种情况下，引入泛型将比复制组件来满足一系列数据类型更干净。</p><p>我们将在本文的后面探讨更多满足这两个条件的用例。不过在这样做之前，让我们先介绍一下 Typescript 泛型提供的其他功能。</p><h3>四、泛型约束</h3><p>有时我们可能希望限制每个类型变量接受的类型数量，这就是泛型约束的作用。下面我们来举几个例子，介绍一下如何使用泛型约束。</p><h3>4.1 确保属性存在</h3><p>有时候，我们希望类型变量对应的类型上存在某些属性。这时，除非我们显式地将特定属性定义为类型变量，否则编译器不会知道它们的存在。</p><p>一个很好的例子是在处理字符串或数组时，我们会假设&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">length</code>&nbsp;属性是可用的。让我们再次使用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数并尝试输出参数的长度：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title function_\">identity</span>&lt;<span class=\"hljs-title function_\">T</span>&gt;(<span class=\"hljs-params\">arg: T</span>): <span class=\"hljs-title function_\">T</span> {\n  <span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(arg.<span class=\"hljs-built_in\">length</span>); <span class=\"hljs-comment\">// Error</span>\n  <span class=\"hljs-keyword\">return</span> arg;\n}\n</pre><p>在这种情况下，编译器将不会知道&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;确实含有&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">length</code>&nbsp;属性，尤其是在可以将任何类型赋给类型变量&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;的情况下。我们需要做的就是让类型变量&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">extends</code>&nbsp;一个含有我们所需属性的接口，比如这样：</p><pre class=\"ql-syntax\" spellcheck=\"false\">interface <span class=\"hljs-built_in\">Length</span> {\n  <span class=\"hljs-attr\">length</span>: <span class=\"hljs-built_in\">number</span>;\n}\n\n<span class=\"hljs-keyword\">function</span> <span class=\"hljs-title function_\">identity</span>&lt;<span class=\"hljs-title function_\">T</span> <span class=\"hljs-title function_\">extends</span> <span class=\"hljs-title function_\">Length</span>&gt;(<span class=\"hljs-params\">arg: T</span>): <span class=\"hljs-title function_\">T</span> {\n  <span class=\"hljs-built_in\">console</span>.<span class=\"hljs-built_in\">log</span>(arg.<span class=\"hljs-built_in\">length</span>); <span class=\"hljs-comment\">// 可以获取length属性</span>\n  <span class=\"hljs-keyword\">return</span> arg;\n}\n</pre><p><code style=\"background-color: rgb(246, 246, 246);\">T extends Length</code>&nbsp;用于告诉编译器，我们支持已经实现&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Length</code>&nbsp;接口的任何类型。之后，当我们使用不含有&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">length</code>&nbsp;属性的对象作为参数调用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">identity</code>&nbsp;函数时，TypeScript 会提示相关的错误信息：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">identity</span>(<span class=\"hljs-number\">68</span>); // Error\n// Argument <span class=\"hljs-keyword\">of</span> <span class=\"hljs-keyword\">type</span> <span class=\"hljs-string\">\'68\'</span> <span class=\"hljs-keyword\">is</span> <span class=\"hljs-keyword\">not</span> assignable <span class=\"hljs-keyword\">to</span> parameter <span class=\"hljs-keyword\">of</span> <span class=\"hljs-keyword\">type</span> <span class=\"hljs-string\">\'Length\'</span>.(<span class=\"hljs-number\">2345</span>)\n</pre><p>此外，我们还可以使用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">,</code>&nbsp;号来分隔多种约束类型，比如：<code style=\"background-color: rgb(246, 246, 246);\">&lt;T extends Length, Type2, Type3&gt;</code>。而对于上述的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">length</code>&nbsp;属性问题来说，如果我们显式地将变量设置为数组类型，也可以解决该问题，具体方式如下：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-function\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title\">identity</span>&lt;<span class=\"hljs-title\">T</span>&gt;<span class=\"hljs-params\">(arg: T[])</span></span>: T[] {\n   console.<span class=\"hljs-built_in\">log</span>(<span class=\"hljs-built_in\">arg</span>.length);  \n   <span class=\"hljs-keyword\">return</span> <span class=\"hljs-built_in\">arg</span>; \n}\n\n// <span class=\"hljs-keyword\">or</span>\n<span class=\"hljs-function\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title\">identity</span>&lt;<span class=\"hljs-title\">T</span>&gt;<span class=\"hljs-params\">(arg: Array&lt;T&gt;)</span></span>: Array&lt;T&gt; {      \n  console.<span class=\"hljs-built_in\">log</span>(<span class=\"hljs-built_in\">arg</span>.length);\n  <span class=\"hljs-keyword\">return</span> <span class=\"hljs-built_in\">arg</span>; \n}\n</pre><h3>4.2 检查对象上的键是否存在</h3><p>泛型约束的另一个常见的使用场景就是检查对象上的键是否存在。不过在看具体示例之前，我们得来了解一下&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">keyof</code>&nbsp;操作符，「<code style=\"background-color: rgb(246, 246, 246);\">keyof</code>&nbsp;操作符是在 TypeScript 2.1 版本引入的，该操作符可以用于获取某种类型的所有键，其返回类型是联合类型。」&nbsp;\"耳听为虚，眼见为实\"，我们来举个&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">keyof</code>&nbsp;的使用示例：</p><pre class=\"ql-syntax\" spellcheck=\"false\">interface Person {\n  name: <span class=\"hljs-keyword\">string</span>;\n  age: <span class=\"hljs-keyword\">number</span>;\n  location: <span class=\"hljs-keyword\">string</span>;\n}\n\n<span class=\"hljs-keyword\">type</span> K1 = keyof Person; // <span class=\"hljs-string\">\"name\"</span> | <span class=\"hljs-string\">\"age\"</span> | <span class=\"hljs-string\">\"location\"</span>\n<span class=\"hljs-keyword\">type</span> K2 = keyof Person[];  // <span class=\"hljs-keyword\">number</span> | <span class=\"hljs-string\">\"length\"</span> | <span class=\"hljs-string\">\"push\"</span> | <span class=\"hljs-string\">\"concat\"</span> | ...\n<span class=\"hljs-keyword\">type</span> K3 = keyof { [x: <span class=\"hljs-keyword\">string</span>]: Person };  // <span class=\"hljs-keyword\">string</span> | <span class=\"hljs-keyword\">number</span>\n</pre><p>通过&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">keyof</code>&nbsp;操作符，我们就可以获取指定类型的所有键，之后我们就可以结合前面介绍的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">extends</code>&nbsp;约束，即限制输入的属性名包含在&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">keyof</code>&nbsp;返回的联合类型中。具体的使用方式如下：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">function</span> getProperty<span class=\"hljs-operator\">&lt;</span><span class=\"hljs-built_in\">T</span><span class=\"hljs-punctuation\">,</span> K extends keyof <span class=\"hljs-built_in\">T</span><span class=\"hljs-operator\">&gt;</span><span class=\"hljs-punctuation\">(</span>obj<span class=\"hljs-operator\">:</span> <span class=\"hljs-built_in\">T</span><span class=\"hljs-punctuation\">,</span> key<span class=\"hljs-operator\">:</span> K<span class=\"hljs-punctuation\">)</span><span class=\"hljs-operator\">:</span> <span class=\"hljs-built_in\">T</span><span class=\"hljs-punctuation\">[</span>K<span class=\"hljs-punctuation\">]</span> <span class=\"hljs-punctuation\">{</span>\n  <span class=\"hljs-built_in\">return</span> obj<span class=\"hljs-punctuation\">[</span>key<span class=\"hljs-punctuation\">]</span>;\n<span class=\"hljs-punctuation\">}</span>\n</pre><p>在以上的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">getProperty</code>&nbsp;函数中，我们通过&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">K extends keyof T</code>&nbsp;确保参数 key 一定是对象中含有的键，这样就不会发生运行时错误。这是一个类型安全的解决方案，与简单调用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">let value = obj[key];</code>&nbsp;不同。</p><p>下面我们来看一下如何使用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">getProperty</code>&nbsp;函数：</p><pre class=\"ql-syntax\" spellcheck=\"false\">enum Difficulty {\n  Easy,\n  Intermediate,\n  Hard\n}\n\n<span class=\"hljs-keyword\">function</span> getProperty&lt;T, K extends keyof T&gt;(obj: T, key: K): T<span class=\"hljs-literal\">[K]</span> {\n  return obj<span class=\"hljs-literal\">[<span class=\"hljs-identifier\">key</span>]</span>;\n}\n\n<span class=\"hljs-keyword\">let</span> tsInfo = {\n   name: <span class=\"hljs-string\">\"Typescript\"</span>,\n   supersetOf: <span class=\"hljs-string\">\"Javascript\"</span>,\n   difficulty: Difficulty.Intermediate\n}\n \n<span class=\"hljs-keyword\">let</span> difficulty: Difficulty = \n  get<span class=\"hljs-constructor\">Property(<span class=\"hljs-params\">tsInfo</span>, \'<span class=\"hljs-params\">difficulty</span>\')</span>; <span class=\"hljs-comment\">// OK</span>\n\n<span class=\"hljs-keyword\">let</span> supersetOf: <span class=\"hljs-built_in\">string</span> = \n  get<span class=\"hljs-constructor\">Property(<span class=\"hljs-params\">tsInfo</span>, \'<span class=\"hljs-params\">superset_of</span>\')</span>; <span class=\"hljs-comment\">// Error</span>\n</pre><p>在以上示例中，对于&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">getProperty(tsInfo, \'superset_of\')</code>&nbsp;这个表达式，TypeScript 编译器会提示以下错误信息：</p><pre class=\"ql-syntax\" spellcheck=\"false\">Argument <span class=\"hljs-keyword\">of</span> <span class=\"hljs-keyword\">type</span> \'<span class=\"hljs-string\">\"superset_of\"</span>\' <span class=\"hljs-keyword\">is</span> <span class=\"hljs-keyword\">not</span> assignable <span class=\"hljs-keyword\">to</span> <span class=\"hljs-keyword\">parameter</span> <span class=\"hljs-keyword\">of</span> <span class=\"hljs-keyword\">type</span> \n\'<span class=\"hljs-string\">\"difficulty\"</span> | <span class=\"hljs-string\">\"name\"</span> | <span class=\"hljs-string\">\"supersetOf\"</span>\'.(<span class=\"hljs-number\">2345</span>)\n</pre><p>很明显通过使用泛型约束，在编译阶段我们就可以提前发现错误，大大提高了程序的健壮性和稳定性。接下来，我们来介绍一下泛型参数默认类型。</p><h3>五、泛型参数默认类型</h3><p>在 TypeScript 2.3 以后，我们可以为泛型中的类型参数指定默认类型。当使用泛型时没有在代码中直接指定类型参数，从实际值参数中也无法推断出类型时，这个默认类型就会起作用。</p><p>泛型参数默认类型与普通函数默认值类似，对应的语法很简单，即&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">&lt;T=Default Type&gt;</code>，对应的使用示例如下：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">interface</span> <span class=\"hljs-symbol\">A</span>&lt;<span class=\"hljs-symbol\">T</span>=<span class=\"hljs-symbol\">string</span>&gt; {\n  name: T;\n}\n\n<span class=\"hljs-keyword\">const</span> strA: A = { name: <span class=\"hljs-string\">\"Semlinker\"</span> };\n<span class=\"hljs-keyword\">const</span> numB: A&lt;number&gt; = { name: <span class=\"hljs-number\">101</span> };\n</pre><p>泛型参数的默认类型遵循以下规则：</p><ul><li>有默认类型的类型参数被认为是可选的。必选的类型参数不能在可选的类型参数后。如果类型参数有约束，类型参数的默认类型必须满足这个约束。当指定类型实参时，你只需要指定必选类型参数的类型实参。 未指定的类型参数会被解析为它们的默认类型。如果指定了默认类型，且类型推断无法选择一个候选类型，那么将使用默认类型作为推断结果。一个被现有类或接口合并的类或者接口的声明可以为现有类型参数引入默认类型。一个被现有类或接口合并的类或者接口的声明可以引入新的类型参数，只要它指定了默认类型。</li></ul><h3>六、泛型条件类型</h3><p>在 TypeScript 2.8 中引入了条件类型，使得我们可以根据某些条件得到不同的类型，这里所说的条件是类型兼容性约束。尽管以上代码中使用了&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">extends</code>&nbsp;关键字，也不一定要强制满足继承关系，而是检查是否满足结构兼容性。</p><p>条件类型会以一个条件表达式进行类型关系检测，从而在两种类型中选择其一：</p><pre class=\"ql-syntax\" spellcheck=\"false\">T extends U ? <span class=\"hljs-keyword\">X</span> : <span class=\"hljs-keyword\">Y</span>\n</pre><p>以上表达式的意思是：若&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;能够赋值给&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">U</code>，那么类型是&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">X</code>，否则为&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Y</code>。在条件类型表达式中，我们通常还会结合&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">infer</code>&nbsp;关键字，实现类型抽取：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">interface</span> <span class=\"hljs-title class_\">Dictionary</span>&lt;T = <span class=\"hljs-built_in\">any</span>&gt; {\n  [<span class=\"hljs-attr\">key</span>: <span class=\"hljs-built_in\">string</span>]: T;\n}\n \n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-title class_\">StrDict</span> = <span class=\"hljs-title class_\">Dictionary</span>&lt;<span class=\"hljs-built_in\">string</span>&gt;\n\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-title class_\">DictMember</span>&lt;T&gt; = T <span class=\"hljs-keyword\">extends</span> <span class=\"hljs-title class_\">Dictionary</span>&lt;infer V&gt; ? V : <span class=\"hljs-built_in\">never</span>\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-title class_\">StrDictMember</span> = <span class=\"hljs-title class_\">DictMember</span>&lt;<span class=\"hljs-title class_\">StrDict</span>&gt; <span class=\"hljs-comment\">// string</span>\n</pre><p>在上面示例中，当类型 T 满足&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T extends Dictionary</code>&nbsp;约束时，我们会使用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">infer</code>&nbsp;关键字声明了一个类型变量 V，并返回该类型，否则返回&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">never</code>&nbsp;类型。</p><blockquote>❝ 在 TypeScript 中，<code style=\"background-color: rgb(246, 246, 246);\">never</code>&nbsp;类型表示的是那些永不存在的值的类型。 例如，&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">never</code>&nbsp;类型是那些总是会抛出异常或根本就不会有返回值的函数表达式或箭头函数表达式的返回值类型。</blockquote><blockquote>另外，需要注意的是，没有类型是&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">never</code>&nbsp;的子类型或可以赋值给&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">never</code>&nbsp;类型（除了&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">never</code>&nbsp;本身之外）。 即使&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">any</code>&nbsp;也不可以赋值给&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">never</code>。</blockquote><blockquote>❞</blockquote><p>除了上述的应用外，利用条件类型和&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">infer</code>&nbsp;关键字，我们还可以方便地实现获取 Promise 对象的返回值类型，比如：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">async</span> <span class=\"hljs-keyword\">function</span> <span class=\"hljs-title function_\">stringPromise</span>() {\n  <span class=\"hljs-keyword\">return</span> <span class=\"hljs-string\">\"Hello, Semlinker!\"</span>;\n}\n\n<span class=\"hljs-keyword\">interface</span> <span class=\"hljs-title class_\">Person</span> {\n  <span class=\"hljs-attr\">name</span>: <span class=\"hljs-built_in\">string</span>;\n  <span class=\"hljs-attr\">age</span>: <span class=\"hljs-built_in\">number</span>;\n}\n\n<span class=\"hljs-keyword\">async</span> <span class=\"hljs-keyword\">function</span> <span class=\"hljs-title function_\">personPromise</span>() {\n  <span class=\"hljs-keyword\">return</span> { <span class=\"hljs-attr\">name</span>: <span class=\"hljs-string\">\"Semlinker\"</span>, <span class=\"hljs-attr\">age</span>: <span class=\"hljs-number\">30</span> } <span class=\"hljs-keyword\">as</span> <span class=\"hljs-title class_\">Person</span>;\n}\n\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-title class_\">PromiseType</span>&lt;T&gt; = <span class=\"hljs-function\">(<span class=\"hljs-params\">args: <span class=\"hljs-built_in\">any</span>[]</span>) =&gt;</span> <span class=\"hljs-title class_\">Promise</span>&lt;T&gt;;\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-title class_\">UnPromisify</span>&lt;T&gt; = T <span class=\"hljs-keyword\">extends</span> <span class=\"hljs-title class_\">PromiseType</span>&lt;infer U&gt; ? U : <span class=\"hljs-built_in\">never</span>;\n\n<span class=\"hljs-keyword\">type</span> extractStringPromise = <span class=\"hljs-title class_\">UnPromisify</span>&lt;<span class=\"hljs-keyword\">typeof</span> stringPromise&gt;; <span class=\"hljs-comment\">// string</span>\n<span class=\"hljs-keyword\">type</span> extractPersonPromise = <span class=\"hljs-title class_\">UnPromisify</span>&lt;<span class=\"hljs-keyword\">typeof</span> personPromise&gt;; <span class=\"hljs-comment\">// Person</span>\n</pre><h3>七、泛型工具类型</h3><p>为了方便开发者 TypeScript 内置了一些常用的工具类型，比如 Partial、Required、Readonly、Record 和 ReturnType 等。出于篇幅考虑，这里我们只简单介绍其中几个常用的工具类型。</p><h3>7.1 Partial</h3><p><code style=\"background-color: rgb(246, 246, 246);\">Partial&lt;T&gt;</code>&nbsp;的作用就是将某个类型里的属性全部变为可选项&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">?</code>。</p><p>「定义：」</p><pre class=\"ql-syntax\" spellcheck=\"false\">/**\n * node_modules/typescript/lib/lib.es5.d.ts\n * <span class=\"hljs-type\">Make</span> all properties <span class=\"hljs-keyword\">in</span> <span class=\"hljs-type\">T</span> optional\n */\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-type\">Partial</span>&lt;<span class=\"hljs-type\">T</span>&gt; = {\n    [<span class=\"hljs-type\">P</span> in keyof <span class=\"hljs-type\">T</span>]?: <span class=\"hljs-type\">T</span>[<span class=\"hljs-type\">P</span>];\n};\n</pre><p>在以上代码中，首先通过&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">keyof T</code>&nbsp;拿到&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;的所有属性名，然后使用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">in</code>&nbsp;进行遍历，将值赋给&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">P</code>，最后通过&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T[P]</code>&nbsp;取得相应的属性值。中间的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">?</code>&nbsp;号，用于将所有属性变为可选。</p><p>「示例：」</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-class\"><span class=\"hljs-keyword\">interface</span> <span class=\"hljs-title\">Todo</span> </span>{\n  title: <span class=\"hljs-keyword\">string</span>;\n  description: <span class=\"hljs-keyword\">string</span>;\n}\n\n<span class=\"hljs-function\"><span class=\"hljs-keyword\">function</span> <span class=\"hljs-title\">updateTodo</span>(<span class=\"hljs-params\">todo: Todo, fieldsToUpdate: Partial&lt;Todo&gt;</span>) </span>{\n  <span class=\"hljs-keyword\">return</span> { ...todo, ...fieldsToUpdate };\n}\n\n<span class=\"hljs-keyword\">const</span> <span class=\"hljs-variable constant_\">todo1</span> = {\n  title: <span class=\"hljs-string\">\"organize desk\"</span>,\n  description: <span class=\"hljs-string\">\"clear clutter\"</span>\n};\n\n<span class=\"hljs-keyword\">const</span> <span class=\"hljs-variable constant_\">todo2</span> = <span class=\"hljs-title function_ invoke__\">updateTodo</span>(todo1, {\n  <span class=\"hljs-attr\">description</span>: <span class=\"hljs-string\">\"throw out trash\"</span>\n});\n</pre><p>在上面的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">updateTodo</code>&nbsp;方法中，我们利用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Partial&lt;T&gt;</code>&nbsp;工具类型，定义&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">fieldsToUpdate</code>&nbsp;的类型为&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Partial&lt;Todo&gt;</code>，即：</p><pre class=\"ql-syntax\" spellcheck=\"false\">{\n   title?: <span class=\"hljs-built_in\">string</span> | <span class=\"hljs-literal\">undefined</span>;\n   description?: <span class=\"hljs-built_in\">string</span> | <span class=\"hljs-literal\">undefined</span>;\n}\n</pre><h3>7.2 Record</h3><p><code style=\"background-color: rgb(246, 246, 246);\">Record&lt;K extends keyof any, T&gt;</code>&nbsp;的作用是将&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">K</code>&nbsp;中所有的属性的值转化为&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;类型。</p><p>「定义：」</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-comment\">/**\n * node_modules/typescript/lib/lib.es5.d.ts\n * Construct a type with a set of properties K of type T\n */</span>\n<span class=\"hljs-class\"><span class=\"hljs-keyword\">type</span> <span class=\"hljs-title\">Record&lt;K</span> <span class=\"hljs-keyword\">extends</span> <span class=\"hljs-title\">keyof</span> <span class=\"hljs-title\">any</span>, <span class=\"hljs-title\">T&gt;</span> </span>= {\n    [<span class=\"hljs-type\">P</span> in <span class=\"hljs-type\">K</span>]: <span class=\"hljs-type\">T</span>;\n};\n</pre><p>「示例：」</p><pre class=\"ql-syntax\" spellcheck=\"false\">interface PageInfo {\n  <span class=\"hljs-keyword\">title</span>: string;\n}\n\ntype <span class=\"hljs-keyword\">Page</span> = <span class=\"hljs-string\">\"home\"</span> | <span class=\"hljs-string\">\"about\"</span> | <span class=\"hljs-string\">\"contact\"</span>;\n\nconst <span class=\"hljs-keyword\">x</span>: Record&lt;<span class=\"hljs-keyword\">Page</span>, PageInfo&gt; = {\n  about: { <span class=\"hljs-keyword\">title</span>: <span class=\"hljs-string\">\"about\"</span> },\n  contact: { <span class=\"hljs-keyword\">title</span>: <span class=\"hljs-string\">\"contact\"</span> },\n  home: { <span class=\"hljs-keyword\">title</span>: <span class=\"hljs-string\">\"home\"</span> }\n};\n</pre><h3>7.3 Pick</h3><p><code style=\"background-color: rgb(246, 246, 246);\">Pick&lt;T, K extends keyof T&gt;</code>&nbsp;的作用是将某个类型中的子属性挑出来，变成包含这个类型部分属性的子类型。</p><p>「定义：」</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-comment\">// node_modules/typescript/lib/lib.es5.d.ts</span>\n\n<span class=\"hljs-comment\">/**\n * From T, pick a set of properties whose keys are in the union K\n */</span>\n<span class=\"hljs-class\"><span class=\"hljs-keyword\">type</span> <span class=\"hljs-title\">Pick&lt;T</span>, <span class=\"hljs-title\">K</span> <span class=\"hljs-keyword\">extends</span> <span class=\"hljs-title\">keyof</span> <span class=\"hljs-title\">T&gt;</span> </span>= {\n    [<span class=\"hljs-type\">P</span> in <span class=\"hljs-type\">K</span>]: <span class=\"hljs-type\">T</span>[<span class=\"hljs-type\">P</span>];\n};\n</pre><p>「示例：」</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">interface</span> <span class=\"hljs-title class_\">Todo</span> {\n  <span class=\"hljs-attr\">title</span>: <span class=\"hljs-built_in\">string</span>;\n  <span class=\"hljs-attr\">description</span>: <span class=\"hljs-built_in\">string</span>;\n  <span class=\"hljs-attr\">completed</span>: <span class=\"hljs-built_in\">boolean</span>;\n}\n\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-title class_\">TodoPreview</span> = <span class=\"hljs-title class_\">Pick</span>&lt;<span class=\"hljs-title class_\">Todo</span>, <span class=\"hljs-string\">\"title\"</span> | <span class=\"hljs-string\">\"completed\"</span>&gt;;\n\n<span class=\"hljs-keyword\">const</span> <span class=\"hljs-attr\">todo</span>: <span class=\"hljs-title class_\">TodoPreview</span> = {\n  <span class=\"hljs-attr\">title</span>: <span class=\"hljs-string\">\"Clean room\"</span>,\n  <span class=\"hljs-attr\">completed</span>: <span class=\"hljs-literal\">false</span>\n};\n</pre><h3>7.4 Exclude</h3><p><code style=\"background-color: rgb(246, 246, 246);\">Exclude&lt;T, U&gt;</code>&nbsp;的作用是将某个类型中属于另一个的类型移除掉。</p><p>「定义：」</p><pre class=\"ql-syntax\" spellcheck=\"false\">// node_modules/typescript/lib/lib.<span class=\"hljs-symbol\">es5</span>.d.ts\n\n/**\n * Exclude from <span class=\"hljs-built_in\">T</span> those types that are assignable to U\n */\n<span class=\"hljs-built_in\">type</span> Exclude&lt;<span class=\"hljs-built_in\">T</span>, U&gt; = <span class=\"hljs-built_in\">T</span> extends U ? never <span class=\"hljs-symbol\">:</span> <span class=\"hljs-built_in\">T</span>;\n</pre><p>如果&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;能赋值给&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">U</code>&nbsp;类型的话，那么就会返回&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">never</code>&nbsp;类型，否则返回&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;类型。最终实现的效果就是将&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;中某些属于&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">U</code>&nbsp;的类型移除掉。</p><p>「示例：」</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-built_in\">type</span> T0 = Exclude&lt;<span class=\"hljs-string\">\"a\"</span> | <span class=\"hljs-string\">\"b\"</span> | <span class=\"hljs-string\">\"c\"</span>, <span class=\"hljs-string\">\"a\"</span>&gt;; // <span class=\"hljs-string\">\"b\"</span> | <span class=\"hljs-string\">\"c\"</span>\n<span class=\"hljs-built_in\">type</span> T1 = Exclude&lt;<span class=\"hljs-string\">\"a\"</span> | <span class=\"hljs-string\">\"b\"</span> | <span class=\"hljs-string\">\"c\"</span>, <span class=\"hljs-string\">\"a\"</span> | <span class=\"hljs-string\">\"b\"</span>&gt;; // <span class=\"hljs-string\">\"c\"</span>\n<span class=\"hljs-built_in\">type</span> T2 = Exclude&lt;<span class=\"hljs-built_in\">string</span> | number | (<span class=\"hljs-function\"><span class=\"hljs-params\">()</span> =&gt;</span> void), Function&gt;; // <span class=\"hljs-built_in\">string</span> | number\n</pre><h3>7.5 ReturnType</h3><p><code style=\"background-color: rgb(246, 246, 246);\">ReturnType&lt;T&gt;</code>&nbsp;的作用是用于获取函数&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;的返回类型。</p><p>「定义：」</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-comment\">// node_modules/typescript/lib/lib.es5.d.ts</span>\n\n<span class=\"hljs-comment\">/**\n * Obtain the return type of a function type\n */</span>\n<span class=\"hljs-class\"><span class=\"hljs-keyword\">type</span> <span class=\"hljs-title\">ReturnType&lt;T</span> <span class=\"hljs-keyword\">extends</span> (<span class=\"hljs-params\">...args: any</span>) <span class=\"hljs-title\">=&gt;</span> <span class=\"hljs-title\">any&gt;</span> </span>= <span class=\"hljs-type\">T</span> <span class=\"hljs-keyword\">extends</span> (...args: any) =&gt; infer <span class=\"hljs-type\">R</span> ? <span class=\"hljs-type\">R</span> : any;\n</pre><p>「示例：」</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">type</span> <span class=\"hljs-variable constant_\">T0</span> = <span class=\"hljs-title class_\">ReturnType</span>&lt;<span class=\"hljs-function\">() =&gt;</span> <span class=\"hljs-built_in\">string</span>&gt;; <span class=\"hljs-comment\">// string</span>\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-variable constant_\">T1</span> = <span class=\"hljs-title class_\">ReturnType</span>&lt;<span class=\"hljs-function\">(<span class=\"hljs-params\">s: <span class=\"hljs-built_in\">string</span></span>) =&gt;</span> <span class=\"hljs-built_in\">void</span>&gt;; <span class=\"hljs-comment\">// void</span>\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-variable constant_\">T2</span> = <span class=\"hljs-title class_\">ReturnType</span>&lt;&lt;T&gt;<span class=\"hljs-function\">() =&gt;</span> T&gt;; <span class=\"hljs-comment\">// {}</span>\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-variable constant_\">T3</span> = <span class=\"hljs-title class_\">ReturnType</span>&lt;&lt;T <span class=\"hljs-keyword\">extends</span> U, U <span class=\"hljs-keyword\">extends</span> <span class=\"hljs-built_in\">number</span>[]&gt;<span class=\"hljs-function\">() =&gt;</span> T&gt;; <span class=\"hljs-comment\">// number[]</span>\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-variable constant_\">T4</span> = <span class=\"hljs-title class_\">ReturnType</span>&lt;<span class=\"hljs-built_in\">any</span>&gt;; <span class=\"hljs-comment\">// any</span>\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-variable constant_\">T5</span> = <span class=\"hljs-title class_\">ReturnType</span>&lt;<span class=\"hljs-built_in\">never</span>&gt;; <span class=\"hljs-comment\">// any</span>\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-variable constant_\">T6</span> = <span class=\"hljs-title class_\">ReturnType</span>&lt;<span class=\"hljs-built_in\">string</span>&gt;; <span class=\"hljs-comment\">// Error</span>\n<span class=\"hljs-keyword\">type</span> <span class=\"hljs-variable constant_\">T7</span> = <span class=\"hljs-title class_\">ReturnType</span>&lt;<span class=\"hljs-title class_\">Function</span>&gt;; <span class=\"hljs-comment\">// Error</span>\n</pre><p>简单介绍了泛型工具类型，最后我们来介绍如何使用泛型来创建对象。</p><h3>八、使用泛型创建对象</h3><h3>8.1 构造签名</h3><p>有时，泛型类可能需要基于传入的泛型 T 来创建其类型相关的对象。比如：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">class</span> <span class=\"hljs-title class_\">FirstClass</span> {\n  id: number | <span class=\"hljs-literal\">undefined</span>;\n}\n\n<span class=\"hljs-keyword\">class</span> <span class=\"hljs-title class_\">SecondClass</span> {\n  name: string | <span class=\"hljs-literal\">undefined</span>;\n}\n\n<span class=\"hljs-keyword\">class</span> <span class=\"hljs-title class_\">GenericCreator</span>&lt;T&gt; {\n  create(): T {\n    <span class=\"hljs-keyword\">return</span> <span class=\"hljs-keyword\">new</span> T();\n  }\n}\n\n<span class=\"hljs-keyword\">const</span> creator1 = <span class=\"hljs-keyword\">new</span> GenericCreator&lt;FirstClass&gt;();\n<span class=\"hljs-keyword\">const</span> firstClass: FirstClass = creator1.create();\n\n<span class=\"hljs-keyword\">const</span> creator2 = <span class=\"hljs-keyword\">new</span> GenericCreator&lt;SecondClass&gt;();\n<span class=\"hljs-keyword\">const</span> secondClass: SecondClass = creator2.create();\n</pre><p>在以上代码中，我们定义了两个普通类和一个泛型类&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">GenericCreator&lt;T&gt;</code>。在通用的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">GenericCreator</code>&nbsp;泛型类中，我们定义了一个名为&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">create</code>&nbsp;的成员方法，该方法会使用 new 关键字来调用传入的实际类型的构造函数，来创建对应的对象。但可惜的是，以上代码并不能正常运行，对于以上代码，在&nbsp;「TypeScript v3.9.2」&nbsp;编译器下会提示以下错误：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-string\">\'T\'</span> <span class=\"hljs-keyword\">only</span> refers <span class=\"hljs-keyword\">to</span> a <span class=\"hljs-keyword\">type</span>, but <span class=\"hljs-keyword\">is</span> being used <span class=\"hljs-keyword\">as</span> a <span class=\"hljs-keyword\">value</span> here.\n</pre><p>这个错误的意思是：<code style=\"background-color: rgb(246, 246, 246);\">T</code>&nbsp;类型仅指类型，但此处被用作值。那么如何解决这个问题呢？根据 TypeScript 文档，为了使通用类能够创建 T 类型的对象，我们需要通过其构造函数来引用 T 类型。对于上述问题，在介绍具体的解决方案前，我们先来介绍一下构造签名。</p><p>在 TypeScript 接口中，你可以使用&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">new</code>&nbsp;关键字来描述一个构造函数：</p><pre class=\"ql-syntax\" spellcheck=\"false\">interface <span class=\"hljs-built_in\">Point</span> {\n  <span class=\"hljs-keyword\">new</span> (x: <span class=\"hljs-built_in\">number</span>, <span class=\"hljs-attr\">y</span>: <span class=\"hljs-built_in\">number</span>): <span class=\"hljs-built_in\">Point</span>;\n}\n</pre><p>以上接口中的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">new (x: number, y: number)</code>&nbsp;我们称之为构造签名，其语法如下：</p><blockquote>❝&nbsp;<em>ConstructSignature:</em>&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">new</code>&nbsp;<em>TypeParametersopt</em>&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">(</code>&nbsp;<em>ParameterListopt</em>&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">)</code>&nbsp;<em>TypeAnnotationopt</em></blockquote><blockquote>❞</blockquote><p>在上述的构造签名中，<code style=\"background-color: rgb(246, 246, 246);\">TypeParametersopt</code>&nbsp;、<code style=\"background-color: rgb(246, 246, 246);\">ParameterListopt</code>&nbsp;和&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">TypeAnnotationopt</code>&nbsp;分别表示：可选的类型参数、可选的参数列表和可选的类型注解。与该语法相对应的几种常见的使用形式如下：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">new</span> <span class=\"hljs-type\">C</span>  \n<span class=\"hljs-keyword\">new</span> <span class=\"hljs-type\">C</span> ( ... )  \n<span class=\"hljs-keyword\">new</span> <span class=\"hljs-type\">C</span> &lt; ... &gt; ( ... )\n</pre><p>介绍完构造签名，我们再来介绍一个与之相关的概念，即构造函数类型。</p><h3>8.2 构造函数类型</h3><p>在 TypeScript 语言规范中这样定义构造函数类型：</p><blockquote>❝ An object type containing one or more construct signatures is said to be a&nbsp;<em>「constructor type」</em>. Constructor types may be written using constructor type literals or by including construct signatures in object type literals.</blockquote><blockquote>❞</blockquote><p>通过规范中的描述信息，我们可以得出以下结论：</p><ul><li>包含一个或多个构造签名的对象类型被称为构造函数类型；构造函数类型可以使用构造函数类型字面量或包含构造签名的对象类型字面量来编写。</li></ul><p>那么什么是构造函数类型字面量呢？构造函数类型字面量是包含单个构造函数签名的对象类型的简写。具体来说，构造函数类型字面量的形式如下：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">new</span> &lt; T1, T2, ... &gt; <span class=\"hljs-function\">(<span class=\"hljs-params\"> p1, p2, ... </span>) =&gt;</span> R\n</pre><p>该形式与以下对象类型字面量是等价的：</p><pre class=\"ql-syntax\" spellcheck=\"false\">{ <span class=\"hljs-literal\">new</span> &lt; T1, T2, <span class=\"hljs-params\">...</span> &gt; ( p1, p2, <span class=\"hljs-params\">...</span> ) : R }\n</pre><p>下面我们来举个实际的示例：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-comment\">// 构造函数类型字面量</span>\n<span class=\"hljs-keyword\">new</span><span class=\"hljs-function\"> (<span class=\"hljs-params\">x</span>: <span class=\"hljs-params\">number</span>, <span class=\"hljs-params\">y</span>: <span class=\"hljs-params\">number</span>) =&gt;</span> Point\n</pre><p>等价于以下对象类型字面量：</p><pre class=\"ql-syntax\" spellcheck=\"false\">{\n   <span class=\"hljs-keyword\">new</span> (x: <span class=\"hljs-built_in\">number</span>, <span class=\"hljs-attr\">y</span>: <span class=\"hljs-built_in\">number</span>): <span class=\"hljs-built_in\">Point</span>;\n}\n</pre><h3>8.3 构造函数类型的应用</h3><p>在介绍构造函数类型的应用前，我们先来看个例子：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">interface</span> <span class=\"hljs-title class_\">Point</span> {\n  <span class=\"hljs-keyword\">new</span> (<span class=\"hljs-attr\">x</span>: <span class=\"hljs-built_in\">number</span>, <span class=\"hljs-attr\">y</span>: <span class=\"hljs-built_in\">number</span>): <span class=\"hljs-title class_\">Point</span>;\n  <span class=\"hljs-attr\">x</span>: <span class=\"hljs-built_in\">number</span>;\n  <span class=\"hljs-attr\">y</span>: <span class=\"hljs-built_in\">number</span>;\n}\n\n<span class=\"hljs-keyword\">class</span> <span class=\"hljs-title class_\">Point2D</span> <span class=\"hljs-keyword\">implements</span> <span class=\"hljs-title class_\">Point</span> {\n  <span class=\"hljs-keyword\">readonly</span> <span class=\"hljs-attr\">x</span>: <span class=\"hljs-built_in\">number</span>;\n  <span class=\"hljs-keyword\">readonly</span> <span class=\"hljs-attr\">y</span>: <span class=\"hljs-built_in\">number</span>;\n\n  <span class=\"hljs-title function_\">constructor</span>(<span class=\"hljs-params\">x: <span class=\"hljs-built_in\">number</span>, y: <span class=\"hljs-built_in\">number</span></span>) {\n    <span class=\"hljs-variable language_\">this</span>.<span class=\"hljs-property\">x</span> = x;\n    <span class=\"hljs-variable language_\">this</span>.<span class=\"hljs-property\">y</span> = y;\n  }\n}\n\n<span class=\"hljs-keyword\">const</span> <span class=\"hljs-attr\">point</span>: <span class=\"hljs-title class_\">Point</span> = <span class=\"hljs-keyword\">new</span> <span class=\"hljs-title class_\">Point2D</span>(<span class=\"hljs-number\">1</span>, <span class=\"hljs-number\">2</span>);\n</pre><p>对于以上的代码，TypeScript 编译器会提示以下错误信息：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">Class</span> <span class=\"hljs-string\">\'Point2D\'</span> incorrectly <span class=\"hljs-keyword\">implements</span> <span class=\"hljs-keyword\">interface</span> <span class=\"hljs-string\">\'Point\'</span>.\n<span class=\"hljs-keyword\">Type</span> <span class=\"hljs-string\">\'Point2D\'</span> provides no match <span class=\"hljs-keyword\">for</span> the signature <span class=\"hljs-string\">\'new (x: number, y: number): Point\'</span>.\n</pre><p>相信很多刚接触 TypeScript 不久的小伙伴都会遇到上述的问题。要解决这个问题，我们就需要把对前面定义的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Point</code>&nbsp;接口进行分离，即把接口的属性和构造函数类型进行分离：</p><pre class=\"ql-syntax\" spellcheck=\"false\">interface <span class=\"hljs-built_in\">Point</span> {\n  <span class=\"hljs-attr\">x</span>: <span class=\"hljs-built_in\">number</span>;\n  y: <span class=\"hljs-built_in\">number</span>;\n}\n\ninterface PointConstructor {\n  <span class=\"hljs-keyword\">new</span> (x: <span class=\"hljs-built_in\">number</span>, <span class=\"hljs-attr\">y</span>: <span class=\"hljs-built_in\">number</span>): <span class=\"hljs-built_in\">Point</span>;\n}\n</pre><p>完成接口拆分之后，除了前面已经定义的&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">Point2D</code>&nbsp;类之外，我们又定义了一个&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">newPoint</code>&nbsp;工厂函数，该函数用于根据传入的 PointConstructor 类型的构造函数，来创建对应的 Point 对象。</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">class</span> <span class=\"hljs-title class_\">Point2D</span> <span class=\"hljs-keyword\">implements</span> <span class=\"hljs-title class_\">Point</span> {\n  <span class=\"hljs-keyword\">readonly</span> <span class=\"hljs-attr\">x</span>: <span class=\"hljs-built_in\">number</span>;\n  <span class=\"hljs-keyword\">readonly</span> <span class=\"hljs-attr\">y</span>: <span class=\"hljs-built_in\">number</span>;\n\n  <span class=\"hljs-title function_\">constructor</span>(<span class=\"hljs-params\">x: <span class=\"hljs-built_in\">number</span>, y: <span class=\"hljs-built_in\">number</span></span>) {\n    <span class=\"hljs-variable language_\">this</span>.<span class=\"hljs-property\">x</span> = x;\n    <span class=\"hljs-variable language_\">this</span>.<span class=\"hljs-property\">y</span> = y;\n  }\n}\n\n<span class=\"hljs-keyword\">function</span> <span class=\"hljs-title function_\">newPoint</span>(<span class=\"hljs-params\">\n  pointConstructor: PointConstructor,\n  x: <span class=\"hljs-built_in\">number</span>,\n  y: <span class=\"hljs-built_in\">number</span>\n</span>): <span class=\"hljs-title class_\">Point</span> {\n  <span class=\"hljs-keyword\">return</span> <span class=\"hljs-keyword\">new</span> <span class=\"hljs-title function_\">pointConstructor</span>(x, y);\n}\n\n<span class=\"hljs-keyword\">const</span> <span class=\"hljs-attr\">point</span>: <span class=\"hljs-title class_\">Point</span> = <span class=\"hljs-title function_\">newPoint</span>(<span class=\"hljs-title class_\">Point2D</span>, <span class=\"hljs-number\">1</span>, <span class=\"hljs-number\">2</span>);\n</pre><h3>8.4 使用泛型创建对象</h3><p>了解完构造签名和构造函数类型之后，下面我们来开始解决上面遇到的问题，首先我们需要重构一下&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">create</code>&nbsp;方法，具体如下所示：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-built_in\">class</span> GenericCreator<span class=\"hljs-operator\">&lt;</span><span class=\"hljs-built_in\">T</span><span class=\"hljs-operator\">&gt;</span> <span class=\"hljs-punctuation\">{</span>\n  create<span class=\"hljs-operator\">&lt;</span><span class=\"hljs-built_in\">T</span><span class=\"hljs-operator\">&gt;</span><span class=\"hljs-punctuation\">(</span><span class=\"hljs-built_in\">c</span><span class=\"hljs-operator\">:</span> <span class=\"hljs-punctuation\">{</span> new <span class=\"hljs-punctuation\">()</span><span class=\"hljs-operator\">:</span> <span class=\"hljs-built_in\">T</span> <span class=\"hljs-punctuation\">})</span><span class=\"hljs-operator\">:</span> <span class=\"hljs-built_in\">T</span> <span class=\"hljs-punctuation\">{</span>\n    <span class=\"hljs-built_in\">return</span> new <span class=\"hljs-built_in\">c</span><span class=\"hljs-punctuation\">()</span>;\n  <span class=\"hljs-punctuation\">}</span>\n<span class=\"hljs-punctuation\">}</span>\n</pre><p>在以上代码中，我们重新定义了&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">create</code>&nbsp;成员方法，根据该方法的签名，我们可以知道该方法接收一个参数，其类型是构造函数类型，且该构造函数不包含任何参数，调用该构造函数后，会返回类型 T 的实例。</p><p>如果构造函数含有参数的话，比如包含一个&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">number</code>&nbsp;类型的参数时，我们可以这样定义 create 方法：</p><pre class=\"ql-syntax\" spellcheck=\"false\">create<span class=\"hljs-operator\">&lt;</span><span class=\"hljs-built_in\">T</span><span class=\"hljs-operator\">&gt;</span><span class=\"hljs-punctuation\">(</span><span class=\"hljs-built_in\">c</span><span class=\"hljs-operator\">:</span> <span class=\"hljs-punctuation\">{</span> new<span class=\"hljs-punctuation\">(</span>a<span class=\"hljs-operator\">:</span> number<span class=\"hljs-punctuation\">)</span><span class=\"hljs-operator\">:</span> <span class=\"hljs-built_in\">T</span>; <span class=\"hljs-punctuation\">},</span> num<span class=\"hljs-operator\">:</span> number<span class=\"hljs-punctuation\">)</span><span class=\"hljs-operator\">:</span> <span class=\"hljs-built_in\">T</span> <span class=\"hljs-punctuation\">{</span>\n  <span class=\"hljs-built_in\">return</span> new <span class=\"hljs-built_in\">c</span><span class=\"hljs-punctuation\">(</span>num<span class=\"hljs-punctuation\">)</span>;\n<span class=\"hljs-punctuation\">}</span>\n</pre><p>更新完&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">GenericCreator</code>&nbsp;泛型类，我们就可以使用下面的方式来创建&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">FirstClass</code>&nbsp;和&nbsp;<code style=\"background-color: rgb(246, 246, 246);\">SecondClass</code>&nbsp;类的实例：</p><pre class=\"ql-syntax\" spellcheck=\"false\"><span class=\"hljs-keyword\">const</span> creator1 = <span class=\"hljs-keyword\">new</span> GenericCreator&lt;FirstClass&gt;();\n<span class=\"hljs-keyword\">const</span> firstClass: FirstClass = creator1.create(FirstClass);\n\n<span class=\"hljs-keyword\">const</span> creator2 = <span class=\"hljs-keyword\">new</span> GenericCreator&lt;SecondClass&gt;();\nconst secondClass: SecondClass = creator2.create(SecondClass);\n</pre>','aliyunOss',1,1,'2023-04-19 11:26:08.536','2023-04-19 11:26:25.229'),(22,30,2,'什么是Docker？','{\"src\": \"assets/static/img/articleContributionCover6d650d4d81f6e040ab5d7588f7a91fd977ee756d2c7f1bc5a8c82e5f35f4a6d0.png\", \"type\": \"aliyunOss\"}','docker','<p>作为程序员我们应怎样理解docker？</p><p><br></p><p><br></p><p>容器技术的起源</p><p>假设你们公司正在秘密研发下一个“今日头条”APP，我们姑且称为明日头条，程序员自己从头到尾搭建了一套环境开始写代码，写完代码后程序员要把代码交给测试同学测试，这时测试同学开始从头到尾搭建这套环境，测试过程中出现问题程序员也不用担心，大可以一脸无辜的撒娇，“明明在人家的环境上可以运行的”。</p><p>测试同学测完后终于可以上线了，这时运维同学又要重新从头到尾搭建这套环境，费了九牛二虎之力搭建好环境开始上线，糟糕，上线系统就崩溃了，这时心理素质好的程序员又可以施展演技了，“明明在人家的环境上可以运行的”。</p><p>从整个过程可以看到，不但我们重复搭建了三套环境还要迫使程序员转行演员浪费表演才华，典型的浪费时间和效率，聪明的程序员是永远不会满足现状的，因此又到了程序员改变世界的时候了，容器技术应运而生。</p><p>有的同学可能会说：“等等，先别改变世界，我们有虚拟机啊，VMware好用的飞起，先搭好一套虚拟机环境然后给测试和运维clone出来不就可以了吗？”</p><p>在没有容器技术之前，这确实是一个好办法，只不过这个办法还没有那么好。</p><p>先科普一下，现在云计算其底层的基石就是虚拟机技术，云计算厂商买回来一堆硬件搭建好数据中心后使用虚拟机技术就可以将硬件资源进行切分了，比如可以切分出100台虚拟机，这样就可以卖给很多用户了。</p><p>你可能会想这个办法为什么不好呢？</p><p><br></p><p><br></p><p>容器技术 vs 虚拟机</p><p>我们知道和一个单纯的应用程序相比，操作系统是一个很重而且很笨的程序，简称笨重，有多笨重呢？</p><p>我们知道操作系统运行起来是需要占用很多资源的，大家对此肯定深有体会，刚装好的系统还什么都没有部署，单纯的操作系统其磁盘占用至少几十G起步，内存要几个G起步。</p><p>假设我有一台机器，16G内存，需要部署三个应用，那么使用虚拟机技术可以这样划分：</p><p><span style=\"background-color: transparent;\"><img src=\"https://pic4.zhimg.com/80/v2-c20cb49c88034e73e09059668b8cecfb_1440w.webp\" height=\"402\" width=\"973\"></span></p><p>在这台机器上开启三个虚拟机，每个虚拟机上部署一个应用，其中VM1占用2G内存，VM2占用1G内存，VM3占用了4G内存。</p><p>我们可以看到虚拟本身就占据了总共7G内存，因此我们没有办法划分出更过虚拟机从而部署更多的应用程序，可是我们部署的是应用程序，要用的也是应用程序而不是操作系统。</p><p>如果有一种技术可以让我们避免把内存浪费在“无用”的操作系统上岂不是太香？这是问题一，主要原因在于操作系统太重了。</p><p>还有另一个问题，那就是启动时间问题，我们知道操作系统重启是非常慢的，因为操作系统要从头到尾把该检测的都检测了该加载的都加载上，这个过程非常缓慢，动辄数分钟，因此操作系统还是太笨了。</p><p>那么有没有一种技术可以让我们获得虚拟机的好处又能克服这些缺点从而一举实现鱼和熊掌的兼得呢？</p><p>答案是肯定的，这就是容器技术。</p><p><br></p><p><br></p><p>什么是容器</p><p>容器一词的英文是container，其实container还有集装箱的意思，集装箱绝对是商业史上了不起的一项发明，大大降低了海洋贸易运输成本。让我们来看看集装箱的好处：</p><ul><li>集装箱之间相互隔离长期反复使用快速装载和卸载规格标准，在港口和船上都可以摆放</li></ul><p><span style=\"background-color: transparent;\"><img src=\"https://pic3.zhimg.com/80/v2-f698870a2becd150a5376942be7368de_1440w.webp\" height=\"362\" width=\"560\"></span></p><p>回到软件中的容器，其实容器和集装箱在概念上是很相似的。</p><p>现代软件开发的一大目的就是隔离，应用程序在运行时相互独立互不干扰，这种隔离实现起来是很不容易的，其中一种解决方案就是上面提到的虚拟机技术，通过将应用程序部署在不同的虚拟机中从而实现隔离。</p><p><span style=\"background-color: transparent;\"><img src=\"https://pic1.zhimg.com/80/v2-0f6ede7f0b920b5d0d5571c937a04838_1440w.webp\" height=\"590\" width=\"543\"></span></p><p>但是虚拟机技术有上述提到的各种缺点，那么容器技术又怎么样呢？</p><p>与虚拟机通过操作系统实现隔离不同，容器技术只隔离应用程序的运行时环境但容器之间可以共享同一个操作系统，这里的运行时环境指的是程序运行依赖的各种库以及配置。</p><p><span style=\"background-color: transparent;\"><img src=\"https://pic2.zhimg.com/80/v2-907214eadd65987e84a0751c08143f91_1440w.webp\" height=\"684\" width=\"723\"></span></p><p>从图中我们可以看到容器更加的轻量级且占用的资源更少，与操作系统动辄几G的内存占用相比，容器技术只需数M空间，因此我们可以在同样规格的硬件上大量部署容器，这是虚拟机所不能比拟的，而且不同于操作系统数分钟的启动时间容器几乎瞬时启动，容器技术为打包服务栈提供了一种更加高效的方式，So cool。</p><p>那么我们该怎么使用容器呢？这就要讲到docker了。</p><p>注意，容器是一种通用技术，docker只是其中的一种实现。</p><p><br></p><p><br></p><p>什么是docker</p><p>docker是一个用Go语言实现的开源项目，可以让我们方便的创建和使用容器，docker将程序以及程序所有的依赖都打包到docker container，这样你的程序可以在任何环境都会有一致的表现，这里程序运行的依赖也就是容器就好比集装箱，容器所处的操作系统环境就好比货船或港口，程序的表现只和集装箱有关系(容器)，和集装箱放在哪个货船或者哪个港口(操作系统)没有关系。</p><p>因此我们可以看到docker可以屏蔽环境差异，也就是说，只要你的程序打包到了docker中，那么无论运行在什么环境下程序的行为都是一致的，程序员再也无法施展表演才华了，不会再有“在我的环境上可以运行”，真正实现“build once, run everywhere”。</p><p>此外docker的另一个好处就是快速部署，这是当前互联网公司最常见的一个应用场景，一个原因在于容器启动速度非常快，另一个原因在于只要确保一个容器中的程序正确运行，那么你就能确信无论在生产环境部署多少都能正确运行。</p><p><br></p><p><br></p><p>如何使用docker</p><p>docker中有这样几个概念：</p><ul><li>dockerfileimagecontainer</li></ul><p>实际上你可以简单的把image理解为可执行程序，container就是运行起来的进程。</p><p>那么写程序需要源代码，那么“写”image就需要dockerfile，dockerfile就是image的源代码，docker就是\"编译器\"。</p><p>因此我们只需要在dockerfile中指定需要哪些程序、依赖什么样的配置，之后把dockerfile交给“编译器”docker进行“编译”，也就是docker build命令，生成的可执行程序就是image，之后就可以运行这个image了，这就是docker run命令，image运行起来后就是docker container。</p><p>具体的使用方法就不再这里赘述了，大家可以参考docker的官方文档，那里有详细的讲解。</p><p><br></p><p><br></p><p>docker是如何工作的</p><p>实际上docker使用了常见的CS架构，也就是client-server模式，docker client负责处理用户输入的各种命令，比如docker build、docker run，真正工作的其实是server，也就是docker demon，值得注意的是，docker client和docker demon可以运行在同一台机器上。</p><p>接下来我们用几个命令来讲解一下docker的工作流程：</p><p><br></p><p><br></p><p>1，docker build</p><p>当我们写完dockerfile交给docker“编译”时使用这个命令，那么client在接收到请求后转发给docker daemon，接着docker daemon根据dockerfile创建出“可执行程序”image。</p><p><span style=\"background-color: transparent;\"><img src=\"https://pic3.zhimg.com/80/v2-f16577a98471b4c4b5b1af1036882caa_1440w.webp\" height=\"577\" width=\"1080\"></span></p><p><br></p><p><br></p><p>2，docker run</p><p>有了“可执行程序”image后就可以运行程序了，接下来使用命令docker run，docker daemon接收到该命令后找到具体的image，然后加载到内存开始执行，image执行起来就是所谓的container。</p><p><span style=\"background-color: transparent;\"><img src=\"https://pic4.zhimg.com/80/v2-672b29e2d53d2ab044269b026c6bc473_1440w.webp\" height=\"589\" width=\"1080\"></span></p><p><br></p><p><br></p><p>3，docker pull</p><p>其实docker build和docker run是两个最核心的命令，会用这两个命令基本上docker就可以用起来了，剩下的就是一些补充。</p><p>那么docker pull是什么意思呢？</p><p>我们之前说过，docker中image的概念就类似于“可执行程序”，我们可以从哪里下载到别人写好的应用程序呢？很简单，那就是APP Store，即应用商店。与之类似，既然image也是一种“可执行程序”，那么有没有\"Docker Image Store\"呢？答案是肯定的，这就是Docker Hub，docker官方的“应用商店”，你可以在这里下载到别人编写好的image，这样你就不用自己编写dockerfile了。</p><p>docker registry 可以用来存放各种image，公共的可以供任何人下载image的仓库就是docker Hub。那么该怎么从Docker Hub中下载image呢，就是这里的docker pull命令了。</p><p>因此，这个命令的实现也很简单，那就是用户通过docker client发送命令，docker daemon接收到命令后向docker registry发送image下载请求，下载后存放在本地，这样我们就可以使用image了。</p><p><span style=\"background-color: transparent;\"><img src=\"https://pic3.zhimg.com/80/v2-dac570abcf7e1776cc266a60c4b19e5e_1440w.webp\" height=\"305\" width=\"563\"></span></p><p>最后，让我们来看一下docker的底层实现。</p><p><br></p><p><br></p><p>docker的底层实现</p><p>docker基于Linux内核提供这样几项功能实现的：</p><ul><li>NameSpace</li><li>我们知道Linux中的PID、IPC、网络等资源是全局的，而NameSpace机制是一种资源隔离方案，在该机制下这些资源就不再是全局的了，而是属于某个特定的NameSpace，各个NameSpace下的资源互不干扰，这就使得每个NameSpace看上去就像一个独立的操作系统一样，但是只有NameSpace是不够。Control groups</li><li>虽然有了NameSpace技术可以实现资源隔离，但进程还是可以不受控的访问系统资源，比如CPU、内存、磁盘、网络等，为了控制容器中进程对资源的访问，Docker采用control groups技术(也就是cgroup)，有了cgroup就可以控制容器中进程对系统资源的消耗了，比如你可以限制某个容器使用内存的上限、可以在哪些CPU上运行等等。</li></ul><p>有了这两项技术，容器看起来就真的像是独立的操作系统了。</p><p><br></p><p><br></p><p>总结</p><p>docker是目前非常流行的技术，很多公司都在生产环境中使用，但是docker依赖的底层技术实际上很早就已经出现了，现在以docker的形式重新焕发活力，并且能很好的解决面临的问题，希望本文能对大家理解docker有所帮助，欢迎大家点赞哦~~</p>','aliyunOss',1,1,'2023-04-19 11:30:06.753','2023-04-19 11:30:37.946');
/*!40000 ALTER TABLE `lv_article_contribution` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_article_contribution_comments`
--

DROP TABLE IF EXISTS `lv_article_contribution_comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_article_contribution_comments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户id',
  `article_id` int NOT NULL COMMENT '文章id',
  `context` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论内容',
  `comment_id` int DEFAULT NULL COMMENT '评论上级id',
  `comment_user_id` int NOT NULL DEFAULT '0' COMMENT '评论上级的用户id (0为空)',
  `comment_first_id` int NOT NULL DEFAULT '0' COMMENT '一级评论(顶层评论)id',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=94 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_article_contribution_comments`
--

LOCK TABLES `lv_article_contribution_comments` WRITE;
/*!40000 ALTER TABLE `lv_article_contribution_comments` DISABLE KEYS */;
INSERT INTO `lv_article_contribution_comments` VALUES (93,30,22,'简单进行视频评论',0,0,0,'2023-04-19 15:05:05.809','2023-04-19 15:05:05.809');
/*!40000 ALTER TABLE `lv_article_contribution_comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_article_contribution_like`
--

DROP TABLE IF EXISTS `lv_article_contribution_like`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_article_contribution_like` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户id',
  `article_id` int NOT NULL COMMENT '文章id',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_article_contribution_like`
--

LOCK TABLES `lv_article_contribution_like` WRITE;
/*!40000 ALTER TABLE `lv_article_contribution_like` DISABLE KEYS */;
/*!40000 ALTER TABLE `lv_article_contribution_like` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_hone_rotograph`
--

DROP TABLE IF EXISTS `lv_hone_rotograph`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_hone_rotograph` (
  `id` int NOT NULL AUTO_INCREMENT,
  `cover` json NOT NULL COMMENT '封面',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '轮播图标题',
  `color` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '图片主题颜色',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '跳转类型 : video跳转视频 article跳转文章 live跳转直播',
  `to_id` int NOT NULL COMMENT '跳转id',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='首页轮播图';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_hone_rotograph`
--

LOCK TABLES `lv_hone_rotograph` WRITE;
/*!40000 ALTER TABLE `lv_hone_rotograph` DISABLE KEYS */;
INSERT INTO `lv_hone_rotograph` VALUES (1,'{\"src\": \"assets/static/img/videoContributionCover/4e78882cdb866a55dc20d52aaeea9ec592891a3fbf18985a4b70cde7680dd859.jpeg\", \"type\": \"local\"}','【囧菌翻唱】东京不太热','rgb(116,82,81)','video',3119,'2023-01-02 04:06:12.117','2023-01-10 03:08:09.185'),(2,'{\"src\": \"assets/static/img/videoContributionCover/ecb1f9671cd6451d7adb5e652d814b8adb96901396c0a57ffa0ed21867a4109c.jpeg\", \"type\": \"local\"}','所以生命啊，它苦涩如歌 | Destiny2来听！俏皮女声翻唱《情字最大》','rgb(16,46,57)','video',3108,'2023-01-02 04:06:12.117','2023-01-10 03:08:09.185'),(3,'{\"src\": \"assets/static/img/articleContributionCover6d650d4d81f6e040ab5d7588f7a91fd977ee756d2c7f1bc5a8c82e5f35f4a6d0.png\", \"type\": \"aliyunOss\"}','什么是Docker？','rgb(152,191,221)','live',30,'2023-01-02 04:06:12.117','2023-01-10 03:08:09.185');
/*!40000 ALTER TABLE `lv_hone_rotograph` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_live_info`
--

DROP TABLE IF EXISTS `lv_live_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_live_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` int unsigned DEFAULT NULL COMMENT '绑定的用户id',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '直播间标题',
  `img` json DEFAULT NULL COMMENT '直播间封面',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`),
  KEY `uid_2` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_live_info`
--

LOCK TABLES `lv_live_info` WRITE;
/*!40000 ALTER TABLE `lv_live_info` DISABLE KEYS */;
INSERT INTO `lv_live_info` VALUES (6,30,'这是一次简单的直播','{\"src\": \"assets/static/img/users/liveCovera1996f075e4eb8393cb8aff350c891bb7e833672c97778d13cc411b5fc796087.jpeg\", \"type\": \"aliyunOss\"}','2023-04-19 14:56:51.802','2023-04-19 14:56:51.802');
/*!40000 ALTER TABLE `lv_live_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_transcoding_task`
--

DROP TABLE IF EXISTS `lv_transcoding_task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_transcoding_task` (
  `id` int NOT NULL AUTO_INCREMENT,
  `task_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '任务id',
  `video_id` int NOT NULL COMMENT '视频id',
  `resolution` int NOT NULL COMMENT '分辨率类型如  1080p 720p',
  `dst` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '保存路径',
  `status` int NOT NULL COMMENT '0等待中1转码成功2转码失败',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '转码平台',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_transcoding_task`
--

LOCK TABLES `lv_transcoding_task` WRITE;
/*!40000 ALTER TABLE `lv_transcoding_task` DISABLE KEYS */;
/*!40000 ALTER TABLE `lv_transcoding_task` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_upload_method`
--

DROP TABLE IF EXISTS `lv_upload_method`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_upload_method` (
  `id` int NOT NULL AUTO_INCREMENT,
  `interface` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求接口名',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '上传方法',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '规定图片存储路径',
  `quality` decimal(2,1) NOT NULL DEFAULT '1.0' COMMENT '图片质量',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_upload_method`
--

LOCK TABLES `lv_upload_method` WRITE;
/*!40000 ALTER TABLE `lv_upload_method` DISABLE KEYS */;
INSERT INTO `lv_upload_method` VALUES (1,'liveCover','aliyunOss','assets/static/img/users/liveCover',0.2,NULL,NULL),(2,'userAvatar','aliyunOss','assets/static/img/users/headPortrait',0.2,NULL,NULL),(3,'videoContribution','local','assets/static/video/users/videoContribution',1.0,NULL,NULL),(4,'videoContributionCover','aliyunOss','assets/static/img/videoContributionCover',0.2,NULL,NULL),(5,'articleContribution','aliyunOss','assets/static/img/articleContribution',0.2,NULL,NULL),(6,'articleContributionCover','aliyunOss','assets/static/img/articleContributionCover',0.2,NULL,NULL),(7,'createFavoritesCover','aliyunOss','assets/static/img/createFavoritesCover',0.2,NULL,NULL);
/*!40000 ALTER TABLE `lv_upload_method` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_users`
--

DROP TABLE IF EXISTS `lv_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `openid` varchar(255) DEFAULT NULL COMMENT '微信小程序openid',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `salt` varchar(100) DEFAULT NULL COMMENT '密码盐',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `photo` json DEFAULT NULL COMMENT '头像',
  `gender` smallint DEFAULT NULL COMMENT '性别 0男 1女 2保密',
  `birth_date` datetime(3) DEFAULT NULL COMMENT '出生时间',
  `is_visible` smallint DEFAULT NULL COMMENT '是否可见 0 false 1 true',
  `signature` varchar(255) DEFAULT NULL COMMENT '个性签名',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Id` (`id`,`openid`),
  KEY `openid` (`openid`),
  KEY `idx_lv_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb3 COMMENT='用户信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_users`
--

LOCK TABLES `lv_users` WRITE;
/*!40000 ALTER TABLE `lv_users` DISABLE KEYS */;
INSERT INTO `lv_users` VALUES (30,'','2506152078@qq.com','橡皮擦','oJnNPG','37acdc0efac9a51e814c4c110843822d','{\"src\": \"assets/static/img/users/headPortraitc47ce475d3157bfdda6726a9a771e8ca220cff99b83e079c4310ae391969aa99.jpg\", \"type\": \"aliyunOss\"}',0,'2002-01-14 00:00:00.000',1,'6666666666','2022-11-03 00:13:55.751','2023-04-18 18:40:30.442',NULL),(36,'','2506152074@qq.com','酸橘子','lhxtxV','f3ec2441b60ca3e04e57bc911efe2661','{\"src\": \"assets/static/img/users/headPortrait56f900f93f2d3189dd319c58e933c2d5f2f18927d9b4a2763f95c8a7bcd0099f.jpg\", \"type\": \"aliyunOss\"}',0,'2023-04-19 15:09:06.564',0,'','2023-04-19 15:09:06.565','2023-04-19 15:10:52.834',NULL);
/*!40000 ALTER TABLE `lv_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_users_attention`
--

DROP TABLE IF EXISTS `lv_users_attention`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_users_attention` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户的id',
  `attention_id` int NOT NULL COMMENT '关注id',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户关注列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_users_attention`
--

LOCK TABLES `lv_users_attention` WRITE;
/*!40000 ALTER TABLE `lv_users_attention` DISABLE KEYS */;
INSERT INTO `lv_users_attention` VALUES (47,36,30,'2023-04-19 15:12:01.559','2023-04-19 15:12:01.559');
/*!40000 ALTER TABLE `lv_users_attention` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_users_chat_list`
--

DROP TABLE IF EXISTS `lv_users_chat_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_users_chat_list` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户id',
  `tid` int NOT NULL COMMENT '聊天对象id',
  `unread` int NOT NULL DEFAULT '0' COMMENT '未读消息数量',
  `last_message` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '最后一条消息',
  `created_at` datetime(3) NOT NULL COMMENT '初次聊天时间',
  `last_at` datetime(3) DEFAULT NULL COMMENT '最后聊天时间',
  `updated_at` datetime(3) NOT NULL COMMENT '最后进入时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_users_chat_list`
--

LOCK TABLES `lv_users_chat_list` WRITE;
/*!40000 ALTER TABLE `lv_users_chat_list` DISABLE KEYS */;
INSERT INTO `lv_users_chat_list` VALUES (56,36,30,0,'321','2023-04-19 15:12:36.286','2023-04-19 15:14:21.734','2023-04-19 15:14:21.735'),(57,30,36,0,'321','2023-04-19 15:12:42.164','2023-04-19 15:14:21.758','2023-04-19 15:18:24.555');
/*!40000 ALTER TABLE `lv_users_chat_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_users_chat_msg`
--

DROP TABLE IF EXISTS `lv_users_chat_msg`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_users_chat_msg` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '消息id',
  `uid` int NOT NULL COMMENT '发送者id',
  `tid` int NOT NULL COMMENT '接收者id',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'text' COMMENT '消息类型',
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '消息内容',
  `created_at` datetime(3) NOT NULL COMMENT '发送时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_users_chat_msg`
--

LOCK TABLES `lv_users_chat_msg` WRITE;
/*!40000 ALTER TABLE `lv_users_chat_msg` DISABLE KEYS */;
INSERT INTO `lv_users_chat_msg` VALUES (23,30,30,'sendChatMsgText','哈喽哈喽','2023-04-19 15:11:46.954','2023-04-19 15:11:46.954'),(24,36,30,'sendChatMsgText','你好啊','2023-04-19 15:12:42.149','2023-04-19 15:12:42.149'),(25,36,30,'sendChatMsgText','在不在','2023-04-19 15:12:55.388','2023-04-19 15:12:55.388'),(26,36,30,'sendChatMsgText','嘀嘀嘀','2023-04-19 15:13:15.464','2023-04-19 15:13:15.464'),(27,30,36,'sendChatMsgText','我在呢','2023-04-19 15:13:28.353','2023-04-19 15:13:28.353'),(28,30,36,'sendChatMsgText','123','2023-04-19 15:14:14.534','2023-04-19 15:14:14.534'),(29,36,30,'sendChatMsgText','321','2023-04-19 15:14:21.734','2023-04-19 15:14:21.734');
/*!40000 ALTER TABLE `lv_users_chat_msg` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_users_collect`
--

DROP TABLE IF EXISTS `lv_users_collect`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_users_collect` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户id',
  `favorites_id` int NOT NULL COMMENT '收藏夹id',
  `video_id` int NOT NULL COMMENT '视频iid',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户收藏表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_users_collect`
--

LOCK TABLES `lv_users_collect` WRITE;
/*!40000 ALTER TABLE `lv_users_collect` DISABLE KEYS */;
INSERT INTO `lv_users_collect` VALUES (71,30,33,3119,'2023-04-19 15:07:15.632','2023-04-19 15:07:15.632');
/*!40000 ALTER TABLE `lv_users_collect` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_users_favorites`
--

DROP TABLE IF EXISTS `lv_users_favorites`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_users_favorites` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '收藏夹创建者id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '收藏夹标题',
  `content` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '收藏夹内容介绍',
  `cover` json DEFAULT NULL COMMENT '收藏夹封面',
  `max` int NOT NULL DEFAULT '1000' COMMENT '最大存储量',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '收藏时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户收藏夹表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_users_favorites`
--

LOCK TABLES `lv_users_favorites` WRITE;
/*!40000 ALTER TABLE `lv_users_favorites` DISABLE KEYS */;
INSERT INTO `lv_users_favorites` VALUES (33,30,'这是一个收藏夹','这是一个收藏夹','{\"src\": \"assets/static/img/createFavoritesCover708aef1afe7f271c77ec28f8a4f48d8e802cf3727c1c76933bcd50354e6c7fce.jpeg\", \"type\": \"aliyunOss\"}',1000,'2023-04-19 15:07:08.097','2023-04-19 15:07:08.097');
/*!40000 ALTER TABLE `lv_users_favorites` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_users_notice`
--

DROP TABLE IF EXISTS `lv_users_notice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_users_notice` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '通知id',
  `uid` int NOT NULL COMMENT '通知用户id',
  `cid` int NOT NULL COMMENT '操作者id',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '视频评论videoComment  视频点赞 videoLike  文章评论articleComment  文章点赞 articleLike',
  `to_id` int NOT NULL COMMENT '对于类型id',
  `is_read` int NOT NULL DEFAULT '0' COMMENT '消息是否已读 0未读 1已读',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '通知内容',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户通知表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_users_notice`
--

LOCK TABLES `lv_users_notice` WRITE;
/*!40000 ALTER TABLE `lv_users_notice` DISABLE KEYS */;
INSERT INTO `lv_users_notice` VALUES (84,30,36,'videoComment',3119,1,'回复一下你啊','2023-04-19 15:11:05.274','2023-04-19 15:18:24.770'),(85,30,36,'videoComment',3119,1,'🤣🤣🤣','2023-04-19 15:11:11.637','2023-04-19 15:18:24.770'),(86,30,36,'videoLike',3119,1,'赞了您的作品','2023-04-19 15:11:14.066','2023-04-19 15:18:24.770'),(87,30,36,'videoLike',3115,1,'赞了您的作品','2023-04-19 15:11:27.254','2023-04-19 15:18:24.770'),(88,30,36,'videoLike',3118,1,'赞了您的作品','2023-04-19 15:11:58.049','2023-04-19 15:18:24.770');
/*!40000 ALTER TABLE `lv_users_notice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_users_record`
--

DROP TABLE IF EXISTS `lv_users_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_users_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户id',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '关联类型 : video跳转视频 article跳转文章 live跳转直播',
  `to_id` int DEFAULT NULL COMMENT '关联id',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=127 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_users_record`
--

LOCK TABLES `lv_users_record` WRITE;
/*!40000 ALTER TABLE `lv_users_record` DISABLE KEYS */;
INSERT INTO `lv_users_record` VALUES (109,30,'video',3108,'2023-04-18 21:29:43.591','2023-04-19 14:54:21.231'),(110,30,'video',3109,'2023-04-19 09:26:45.969','2023-04-19 09:50:52.455'),(111,30,'video',3110,'2023-04-19 09:50:54.703','2023-04-19 09:50:54.703'),(112,30,'video',3111,'2023-04-19 09:51:03.894','2023-04-19 10:09:33.288'),(113,30,'video',3112,'2023-04-19 09:51:10.923','2023-04-19 10:08:44.816'),(114,30,'video',3113,'2023-04-19 10:08:33.534','2023-04-19 10:08:33.534'),(115,30,'article',19,'2023-04-19 11:20:21.097','2023-04-19 11:20:21.097'),(116,30,'article',20,'2023-04-19 11:23:09.317','2023-04-19 11:23:22.037'),(117,30,'article',21,'2023-04-19 11:26:25.221','2023-04-19 11:26:25.221'),(118,30,'article',22,'2023-04-19 11:30:37.934','2023-04-19 15:19:46.842'),(119,30,'video',3119,'2023-04-19 14:52:21.540','2023-04-19 15:19:21.420'),(120,30,'video',3115,'2023-04-19 14:53:18.020','2023-04-19 14:53:18.020'),(121,30,'live',30,'2023-04-19 15:04:33.697','2023-04-19 15:22:40.970'),(122,30,'video',3118,'2023-04-19 15:07:22.162','2023-04-19 15:07:22.162'),(123,36,'video',3119,'2023-04-19 15:09:09.647','2023-04-19 15:10:55.433'),(124,36,'video',3115,'2023-04-19 15:11:25.980','2023-04-19 15:11:25.980'),(125,36,'video',3118,'2023-04-19 15:11:55.667','2023-04-19 15:11:55.667'),(126,36,'video',3116,'2023-04-19 15:12:35.249','2023-04-19 15:12:35.249');
/*!40000 ALTER TABLE `lv_users_record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_video_contribution`
--

DROP TABLE IF EXISTS `lv_video_contribution`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_video_contribution` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '视频标题',
  `video` json DEFAULT NULL COMMENT '视频(默认1080p)',
  `video_720p` json DEFAULT NULL COMMENT '720p视频',
  `video_480p` json DEFAULT NULL COMMENT '480p视频',
  `video_360p` json DEFAULT NULL COMMENT '360p视频',
  `media_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '媒资ID 阿里云oss 确定视频需要',
  `cover` json NOT NULL COMMENT '视频封面',
  `video_duration` int NOT NULL COMMENT '视频时长',
  `reprinted` int NOT NULL DEFAULT '0' COMMENT '是否转载0不是1是',
  `label` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标签定义，分割',
  `introduce` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '视频介绍',
  `heat` int NOT NULL DEFAULT '0' COMMENT '视频热度',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3120 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_video_contribution`
--

LOCK TABLES `lv_video_contribution` WRITE;
/*!40000 ALTER TABLE `lv_video_contribution` DISABLE KEYS */;
INSERT INTO `lv_video_contribution` VALUES (3108,30,'来听！俏皮女声翻唱《情字最大》','{\"src\": \"assets/static/video/users/videoContribution/8dd20333094d7c3531314e07ab14867d8feb46f7609c9def7c15730a05b9d7fb.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/8dd20333094d7c3531314e07ab14867d8feb46f7609c9def7c15730a05b9d7fb_output_720p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/8dd20333094d7c3531314e07ab14867d8feb46f7609c9def7c15730a05b9d7fb_output_480p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/8dd20333094d7c3531314e07ab14867d8feb46f7609c9def7c15730a05b9d7fb_output_360p.mp4\", \"type\": \"local\"}','','{\"src\": \"assets/static/img/videoContributionCover/ecb1f9671cd6451d7adb5e652d814b8adb96901396c0a57ffa0ed21867a4109c.jpeg\", \"type\": \"local\"}',156,0,'古风,古风翻唱,浅影啊','该版本已获得正版授权改编',1,'2023-04-18 21:29:34.894','2023-04-19 08:40:57.863'),(3109,30,'【洛少爷+囧菌】《国道217》Z新豪老师的歌真好听啊=3=！','{\"src\": \"assets/static/video/users/videoContribution/19c14b131c36c6ce9ec61ab0a2c56fc5fb17897196e745b8b49394bdda23c8d7.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/19c14b131c36c6ce9ec61ab0a2c56fc5fb17897196e745b8b49394bdda23c8d7_output_720p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/19c14b131c36c6ce9ec61ab0a2c56fc5fb17897196e745b8b49394bdda23c8d7_output_480p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/19c14b131c36c6ce9ec61ab0a2c56fc5fb17897196e745b8b49394bdda23c8d7_output_360p.mp4\", \"type\": \"local\"}','','{\"src\": \"assets/static/img/videoContributionCover/d1becc14f226bd6d14eca2e435782350661507e109869900fbd72498ea2639a8.jpeg\", \"type\": \"local\"}',232,0,'囧菌,洛少爷','本家：BV1EK411q78Q\n作词作曲编曲：Z新豪\n翻唱：封茗囧菌+洛少爷\n混音：防滑酱\n图绘：白秋\nPV：催归CUI\nPV素材：侑景-一罐',1,'2023-04-19 09:16:22.385','2023-04-19 09:26:45.959'),(3110,30,'演员失误造就的经典：群演入戏太深被揍，李云龙一脚踢出名场面','{\"src\": \"assets/static/video/users/videoContribution/9c7e60124636992b0e23d7e85b13451e981d2262073e7aa553d47ba26736174f.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/9c7e60124636992b0e23d7e85b13451e981d2262073e7aa553d47ba26736174f_output_720p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/9c7e60124636992b0e23d7e85b13451e981d2262073e7aa553d47ba26736174f_output_480p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/9c7e60124636992b0e23d7e85b13451e981d2262073e7aa553d47ba26736174f_output_360p.mp4\", \"type\": \"local\"}','','{\"src\": \"assets/static/img/videoContributionCover/ccc4f6c3c378796729181cba27ac7505b6e0eea447a7e41b1dc07c43ac6ca83e.jpeg\", \"type\": \"local\"}',95,0,'李云龙,影评','失误造就经典镜头',1,'2023-04-19 09:44:26.131','2023-04-19 09:50:54.694'),(3111,30,'大一野蛮学生几千块喜提8手烂仔桑塔纳2000后的土法整备',NULL,'{\"src\": \"assets/static/video/users/videoContribution/8dacb56866c34113f877e1df613f39656076ef5944747711073361e59763a2af.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/8dacb56866c34113f877e1df613f39656076ef5944747711073361e59763a2af_output_480p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/8dacb56866c34113f877e1df613f39656076ef5944747711073361e59763a2af_output_360p.mp4\", \"type\": \"local\"}','','{\"src\": \"assets/static/img/videoContributionCover/4f66fc4167584f4834ab7191fa977a8bcdd7decdb228b34bd0b08e1c96a572d5.jpg\", \"type\": \"local\"}',173,0,'汽车,汽车生活','记录一下买车两个月后的整备过程，视频有些无聊确实很抱歉 （以后有空还会再弄些贴纸',1,'2023-04-19 09:48:00.780','2023-04-19 09:51:03.883'),(3112,30,'【不良人|星雪】“天子的女人，又如何！”天暗星x姬如雪',NULL,'{\"src\": \"assets/static/video/users/videoContribution/04a84e5bdfd837096ae547774be3b860b6a9ea96d976d25af18cded7004de1ae.mp4\", \"type\": \"local\"}',NULL,NULL,'','{\"src\": \"assets/static/img/videoContributionCover/aab8ab2d6978c15566ea9c5412de959735cb2898967d9768b570ff0f4d4c0096.jpeg\", \"type\": \"local\"}',136,0,'不良人,画江湖','BGM：华之乱\n空镜感谢：阿琪的兴兴；Bluefox\n\n下架重传，大伙低调低调低调orz\n灵感来源是上周砍柴时冒出来的，看雪儿和小北篝火谈心那\n有句词正好能拼接到一块：“姬如雪，天下不良人皆知。”\n\n原本想等s6播完再剪，结果忍不过一星期。\n\n天子爱妻，天下不良人皆知。临终前，托不良人照料，新不良帅天暗星接下此任。殊不知，天子之死，另有隐情......',1,'2023-04-19 09:50:20.166','2023-04-19 09:51:10.914'),(3113,30,'快来听歌！浅影阿翻唱的《难却》闪亮回归！','{\"src\": \"assets/static/video/users/videoContribution/421ed5e75ad26a15b011895640cebd413b63a9e6e8ddfd0c5125354d43b64495.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/421ed5e75ad26a15b011895640cebd413b63a9e6e8ddfd0c5125354d43b64495_output_720p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/421ed5e75ad26a15b011895640cebd413b63a9e6e8ddfd0c5125354d43b64495_output_480p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/421ed5e75ad26a15b011895640cebd413b63a9e6e8ddfd0c5125354d43b64495_output_360p.mp4\", \"type\": \"local\"}','','{\"src\": \"assets/static/img/videoContributionCover/fb9d47f72d178669d138443c9749f5059300c2bab051db118d34c21c3f8cf306.jpeg\", \"type\": \"local\"}',182,0,'浅影啊','帅哥美女们早上好，相信大家都知道了，浅影姐姐把难却那一期视频删除了，许多小伙伴们表示非常可惜。作为浅影姐姐的头号粉丝，为了让大家看到本期视频，晓静花时间把这个视频找回来了，建议大家观看后把本期视频保存下来观，别忘了点个赞再走哦～\n《难却》',0,'2023-04-19 10:05:31.777','2023-04-19 10:09:16.615'),(3114,30,'梦幻联动！我把三位仙女up主翻唱的《七里香》','{\"src\": \"assets/static/video/users/videoContribution/02917b756c82ae32ce2add136cc4e7d4ec859f6469c60cadb96c70d67152fecf.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/02917b756c82ae32ce2add136cc4e7d4ec859f6469c60cadb96c70d67152fecf_output_720p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/02917b756c82ae32ce2add136cc4e7d4ec859f6469c60cadb96c70d67152fecf_output_480p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/02917b756c82ae32ce2add136cc4e7d4ec859f6469c60cadb96c70d67152fecf_output_360p.mp4\", \"type\": \"local\"}','','{\"src\": \"assets/static/img/videoContributionCover/24fa676f5cf3a9ea3427b4fc19b8261d66ae39b16abc5803c9935dd26a0e471c.jpeg\", \"type\": \"local\"}',115,0,'合作','热烈祝贺我们老大@浅影阿_ 获得100万粉丝！\n@九三的耳朵不是特别好 九三姐姐获得300万粉丝！\n@月梢Moon 月梢姐姐获得10万粉丝！\n七里香的名字很美，你们也是',0,'2023-04-19 11:31:53.205','2023-04-19 11:36:11.767'),(3115,30,'《玫瑰花的葬礼》当年听这首歌的都毕业了吧','{\"src\": \"assets/static/video/users/videoContribution/ba28308a2e9a89c08b561b7073061c1a804aeadfd5c25e559d2d6551cdea13cc.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/ba28308a2e9a89c08b561b7073061c1a804aeadfd5c25e559d2d6551cdea13cc_output_720p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/ba28308a2e9a89c08b561b7073061c1a804aeadfd5c25e559d2d6551cdea13cc_output_480p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/ba28308a2e9a89c08b561b7073061c1a804aeadfd5c25e559d2d6551cdea13cc_output_360p.mp4\", \"type\": \"local\"}','','{\"src\": \"assets/static/img/videoContributionCover/ef724dc87862ab20f28ec7b64d6dd4cc7a56feb1c551a8ccb92c23a8b9849215.jpg\", \"type\": \"local\"}',81,0,'玫瑰花的葬礼','好听爱听！！',2,'2023-04-19 11:33:23.006','2023-04-19 15:11:25.966'),(3116,30,'(ง •̀_•́)ง我要带上我的兄弟们在山顶上面摆造型儿！','{\"src\": \"assets/static/video/users/videoContribution/641cfc8d6c2fdf73ef5a31b2ec936acb405474326e59a619d6e59f27baefbc3e.mp4\", \"type\": \"local\"}',NULL,'{\"src\": \"assets/static/video/users/videoContribution/641cfc8d6c2fdf73ef5a31b2ec936acb405474326e59a619d6e59f27baefbc3e_output_480p.mp4\", \"type\": \"local\"}',NULL,'','{\"src\": \"assets/static/img/videoContributionCover/f5da00e2878759b3d9985caf90ccd9104ad8d5df8c1b80baed4d9f253c245491.jpeg\", \"type\": \"local\"}',31,0,'囧菌','(ง •̀_•́)ง我要带上我的兄弟们在山顶上面摆造型儿！',1,'2023-04-19 11:40:58.389','2023-04-19 15:12:35.232'),(3117,30,'【囧菌翻唱】KDA - More','{\"src\": \"assets/static/video/users/videoContribution/84c29d9bdf319e8d789e5dffb63f7e71ce0f5debe89b9fc7fa37ce9843a70160.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/84c29d9bdf319e8d789e5dffb63f7e71ce0f5debe89b9fc7fa37ce9843a70160_output_720p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/84c29d9bdf319e8d789e5dffb63f7e71ce0f5debe89b9fc7fa37ce9843a70160_output_480p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/84c29d9bdf319e8d789e5dffb63f7e71ce0f5debe89b9fc7fa37ce9843a70160_output_360p.mp4\", \"type\": \"local\"}','','{\"src\": \"assets/static/img/videoContributionCover/052344982f1061e1ecb331b0bde99f437785049caa9fe3694c6b5da430923812.jpeg\", \"type\": \"local\"}',171,0,'翻唱','KDA - MORE\n原唱：Madison Beer、Jaira Burns、刘柏辛、曺薇娟、田小娟\n翻唱：封茗囧菌\n混音：圈妹\n\n>3<这次的眼妆是有意化得夸张了些！毕竟这么酷的歌！！',0,'2023-04-19 11:44:41.705','2023-04-19 11:51:44.408'),(3118,30,'【囧菌翻唱】水星','{\"src\": \"assets/static/video/users/videoContribution/e9ea9cee56c66117229305b4a407515788a63174c6ef9e7325a212c3f3afed7f.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/e9ea9cee56c66117229305b4a407515788a63174c6ef9e7325a212c3f3afed7f_output_720p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/e9ea9cee56c66117229305b4a407515788a63174c6ef9e7325a212c3f3afed7f_output_480p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/e9ea9cee56c66117229305b4a407515788a63174c6ef9e7325a212c3f3afed7f_output_360p.mp4\", \"type\": \"local\"}','','{\"src\": \"assets/static/img/videoContributionCover/4cd758c3ced07af8c4f86f1b5543f96578362647a8d8d9cad65dd8b541bb4643.jpeg\", \"type\": \"local\"}',118,0,'水星','水星\n原唱：TOFUBEATS\n翻唱：封茗囧菌\n混音：我是百事糖',2,'2023-04-19 11:45:38.468','2023-04-19 15:11:55.657'),(3119,30,'【囧菌翻唱】东京不太热',NULL,'{\"src\": \"assets/static/video/users/videoContribution/37a3ce4a9e68bb2a2e5f0f7a747ad790e1a9d5a5484669b07c6cc58686f1d43d.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/37a3ce4a9e68bb2a2e5f0f7a747ad790e1a9d5a5484669b07c6cc58686f1d43d_output_480p.mp4\", \"type\": \"local\"}','{\"src\": \"assets/static/video/users/videoContribution/37a3ce4a9e68bb2a2e5f0f7a747ad790e1a9d5a5484669b07c6cc58686f1d43d_output_360p.mp4\", \"type\": \"local\"}','','{\"src\": \"assets/static/img/videoContributionCover/4e78882cdb866a55dc20d52aaeea9ec592891a3fbf18985a4b70cde7680dd859.jpeg\", \"type\": \"local\"}',225,0,'翻唱,东京不太热','自制 【本家：av2251501】【后期翻唱：囧菌】【MP3：5sing链接http://5sing.kugou.com/fc/14105458.html】【视频制作+压制：黑狸警长】【图片选自：pixvi 作者不知=口=百度搜索半天木有找到】',2,'2023-04-19 14:19:49.617','2023-04-19 15:09:09.632');
/*!40000 ALTER TABLE `lv_video_contribution` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_video_contribution_barrage`
--

DROP TABLE IF EXISTS `lv_video_contribution_barrage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_video_contribution_barrage` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户id',
  `video_id` int NOT NULL COMMENT '视频id',
  `time` float(10,6) NOT NULL DEFAULT '0.000000' COMMENT '出现时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `text` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '弹幕内容',
  `type` int DEFAULT '0' COMMENT '弹幕位置',
  `color` int NOT NULL COMMENT '弹幕颜色（十进制）',
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_video_contribution_barrage`
--

LOCK TABLES `lv_video_contribution_barrage` WRITE;
/*!40000 ALTER TABLE `lv_video_contribution_barrage` DISABLE KEYS */;
INSERT INTO `lv_video_contribution_barrage` VALUES (46,30,3119,9.055357,'橡皮擦','发个弹幕试试',0,16777215,'2023-04-19 15:05:20.807','2023-04-19 15:05:20.807'),(47,30,3119,18.571539,'橡皮擦','哈喽',0,16777215,'2023-04-19 15:05:30.201','2023-04-19 15:05:30.201'),(48,30,3119,24.546797,'橡皮擦','滴滴滴滴',0,16777215,'2023-04-19 15:05:36.151','2023-04-19 15:05:36.151'),(49,30,3119,2.022651,'橡皮擦','弹幕在这里',0,16777215,'2023-04-19 15:06:20.432','2023-04-19 15:06:20.432');
/*!40000 ALTER TABLE `lv_video_contribution_barrage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_video_contribution_comments`
--

DROP TABLE IF EXISTS `lv_video_contribution_comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_video_contribution_comments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户id',
  `video_id` int NOT NULL COMMENT '视频id',
  `context` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论内容',
  `comment_id` int DEFAULT NULL COMMENT '评论上级id',
  `comment_user_id` int NOT NULL DEFAULT '0' COMMENT '评论上级的用户id (0为空)',
  `comment_first_id` int NOT NULL DEFAULT '0' COMMENT '一级评论(顶层评论)id',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_video_contribution_comments`
--

LOCK TABLES `lv_video_contribution_comments` WRITE;
/*!40000 ALTER TABLE `lv_video_contribution_comments` DISABLE KEYS */;
INSERT INTO `lv_video_contribution_comments` VALUES (46,30,3119,'简单进行视频评论\n',0,0,0,'2023-04-19 15:06:36.927','2023-04-19 15:06:36.927'),(47,36,3119,'回复一下你啊',46,30,46,'2023-04-19 15:11:05.273','2023-04-19 15:11:05.273'),(48,36,3119,'🤣🤣🤣',0,0,0,'2023-04-19 15:11:11.637','2023-04-19 15:11:11.637');
/*!40000 ALTER TABLE `lv_video_contribution_comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lv_video_contribution_like`
--

DROP TABLE IF EXISTS `lv_video_contribution_like`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lv_video_contribution_like` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户id',
  `video_id` int NOT NULL COMMENT '视频id',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lv_video_contribution_like`
--

LOCK TABLES `lv_video_contribution_like` WRITE;
/*!40000 ALTER TABLE `lv_video_contribution_like` DISABLE KEYS */;
INSERT INTO `lv_video_contribution_like` VALUES (70,30,3119,'2023-04-19 15:07:18.081','2023-04-19 15:07:18.081'),(71,36,3119,'2023-04-19 15:11:14.056','2023-04-19 15:11:14.056'),(72,36,3115,'2023-04-19 15:11:27.246','2023-04-19 15:11:27.246'),(73,36,3118,'2023-04-19 15:11:58.041','2023-04-19 15:11:58.041');
/*!40000 ALTER TABLE `lv_video_contribution_like` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'video-net'
--

--
-- Dumping routines for database 'video-net'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-19 15:36:29
