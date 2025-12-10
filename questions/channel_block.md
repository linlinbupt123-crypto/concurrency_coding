select的逻辑
1,在执行 select 的那个瞬间，哪个 case ready，我才会选
    1.1,其他case当前不ready
    1.2, ch这个case允许阻塞等待
    1.3, select会随机或按调度策略选一个“可阻塞的”case并阻塞等待
    总结，select不是看ch有数据才选它，而是允许选它然后等数据
2, 一旦选了某个 case，它就进入该 case 的阻塞操作，不会再动态切换。这就是如果只有两个case，一个case是阻塞读/写channel,一个case是ctx.Done()的情况下，即便是ctx结束了，也不会被执行，因为阻塞在channel的读写case里了
⚠️关键理解：select 不是“持续探测”，它只执行一次
很多人误以为：
select 会不断找 ready 的 case
❗错误。
真正机制是：
调用 select → 做就绪判断
选定 1 个 case
进入 case 内执行
如果 case 是阻塞操作，那就阻塞到底，不会切换到其他 case
select 不会“反悔”
也不会说：“阻塞久了我换一个 case 吧”

answer1:
goroutine持续阻塞在读ch这一步。导致goroutine不能退出

answer2:
当select 选择一个case之后，就阻塞在这里不会退出

