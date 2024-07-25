# Novice 使用指北 `--整理中--`
*--- 不是什么技术力的项目，只是整理一些配置做个懒人包 ---*

## (一) 介绍

### $ Novice / 受众群体
- Novice面向即将或初步入门`archlinux`，不希望将时间过多消耗在折腾系统上，想要快速配置一个可用`archlinux` 环境 *(!!!非生产环境!!!)* ，需要提供以后 *自由折腾* 或 *调试空间* 的用户。
- 省流：适用于忙人或懒人的一键环境配置。
### $ Novice / 知悉
  1. Novice只负责各项软件环境的自动化安装与检查，安装已经被预设好的配置文件，节省用户入门使用archlinux的时间。也正因如此，Novice的元配置主要源自于开源社区和一些我的个人实装配置，在此向广大开源社区工作者致以敬意 ;
  2. 如果你已经配置好了hyprland环境或者只是想参考我的这份配置，只需提取仓库中/novice/configs的内容即可；如果需要安装，这将直接覆盖当前配置，请注意备份； 
  3. Novice所使用的配置文件就是我当前所使用的配置文件，并不是所有人都会喜欢她。如果可能，挪出时间自己进行修改可以使自己获得一个更加舒适的linux使用体验。初步的修改可以先从copy别人现成的配置开始，她们一般都是`.conf`,`config`文件，在Novice所提供的配置文件当中，我会尽可能的将官方文档对其描述和具体功能以代码注释的形式表示出来以方便各位理解，从而更好地自定义自己的archlinux。例如，我的第一套 `archlinux + hyprland` 配置就是从 `@Cascade` 大佬处复制而来。其个人博客的地址是 `cascade.moe`，欢迎各位拜谒 ;
  4. 在使用Novice之前，请事先安装好 archlinux 和 hyprland，这是入门的最低门槛。如果连装系统都不会......那就移步`B站@木子龙昊`处手把手教你怎么装系统，这是其B站个人主页链接： `space.bilibili.com/125070206` 。Novice保险推荐使用 `archinstall` 进行archlinux系统安装，可考虑采用妈妈级手段:`bilibili.com/video/BV15F4m1j7Qd`。在编辑完`/etc/pacman.d/mirrorlist`以后，请注意`archinstall`的部分信息：1.桌面环境应当选择`Hyprland`，这会自动化帮你安装`hyprland`桌面环境; 2.为了你的使用体验着想，请不要勾选`/home`的独立分区。因为系统默认的分区方案一般都是智障级别的 *(我的`/home`就被分配了`>400Gib`的空间)* ;


### $. Novice / 自动预装的软件列表和作用
- Novice安装 *配置* 之前设备上应当安装的程序如下，如果未安装，Novice会自动安装后再执行配置的部署工作。但仍然推荐自行检查程序的安装情况，你可以通过`pacman`包管理器来查看当前archlinux的包安装情况：`sudo pacman -Q`；或者，你可以通过直接在终端窗口内尝试打开应用来确定其安装情况。
  - 以下为基本环境类软件:
  1. alacritty ; `~/.config/alacritty/`；终端，你可以选择她来使用`fish`,`bash`,`zsh`等`shell`；请使用 `Alt + T` 或者 `Alt + A` 组合键启动她；
  2. fcitx5 ; `~/.config/fcitx5/`；在archlinux下的中文输入法软件，在Novice当中，将默认使用`双拼`输入方式，`自然码`字符映射方案；Novice默认其开机自启动；请使用 `Super + Space` 组合键打开或关闭输入法；
  3. fish ; `~/.config/fish/`；内置自动提示、语法高亮、自动补全、搜索历史等功能的现代化`shell`软件；她将是`alacritty`的默认shell；
  4. hyprland ; `~/.config/hypr/hyprland.conf`；在默认情况下，linux并不会提供一个可用的桌面环境，对于用户来说，他们能看到的只有默认的内置bash shell，安装桌面环境的目的就是提升其可用性并且提供现代化的GUI界面。*但是需要注意，一般情况下，linux在生产环境当中并不会安装桌面环境，安装桌面环境的linux在生产环境当中稳定性并不如Windows Server*。Hyprland基于Wayland开发，相比于成熟的KDE，在某些软件上兼容性不会很好，但使用更加现代化和美观。目前在一般生活场景下的软件均可正常使用。比如：`钉钉/Dingtalk`，`微信/WeChat Universal`, `QQ`, `火狐/FireFox`等；Novice默认其开机自启动；
  5. hyprlock ; `~/.config/hypr/hyprlock.conf`；锁屏工具，hyprland默认不会自动提供一个锁屏软件；请使用 `Alt + L` 组合键启动她；
  6. hyprpaper ; `~/.config/hypr/hyprpaper.conf`；壁纸软件，她将配置由Novice提供的新壁纸来替代hyprland默认的看板娘`hypr-chan`；你可以通过修改配置文件的方式将壁纸指向你希望显示的壁纸，也可以同名替换`~/Pictures`下的图像内容；Novice默认其开机自启动；
  7. kitty ; `~/.config/kitty/`；hyprland默认终端，由于Novice已经提供更加美观方面的alacritty，你可以选择不管她；你可以通过shell启动kitty；
  8. mako ;`~/.config/mako/`；轻量级的通知 Daemon，在修改配置文件之后，用`makoctl`来重载配置文件：`makoctl reload`； 你可以通过 notify-send 指令来发送一个通知，以此观察修改配置文件之后的效果：`notify-send 'EXAMPLE' 'this is a notification.' -u normal`；她将在有通知时自动唤醒；
  9. tofi ;`~/.config/tofi/`；应用启动器，用于绕过shell启动程序，如果你使用shell来启动，你应该会发现大部分应用会随着shell的关闭而关闭，一般情况下，shell启动的程序所使用的环境与你使用的shell启动时已经载入的环境相同，在这里，以proxy环境变量为例，倘若你在shell A当中export了一个代理，在shell B当中则没有设置这个代理，那么，shell A启动的程序也将跟随shell A使用这个代理。但有个别例外，比如Microsoft VSCode, JetBrains IDEs... 请使用 `Alt + R`组合键启动她(tofi)；
  10. vlc ;`~/.config/vlc/`；媒体播放器，用于播放视频和音频；你可以通过tofi启动她；也可以在打开视频时自动唤醒；
  11. waybar ;`~/.config/waybar/`；状态栏；你可以把她当作linux下的Windows任务栏；Novice将让hyprland默认其开机自启动；
  - 以下为生活与工作类软件:
  1. firefox ;`~/.mozilla/`；浏览器，在Novice当中她将自动安装AdGuard并启用 `EasyPrivacy` 等过滤器； 可以通过tofi启动；
  2. Telegram ;A globally accessible freemium, encrypted, cloud-based and centralized instant messaging (IM) service. The application also provides optional end-to-end encrypted chats (popularly known as "secret chats") and video calling, VoIP, file sharing and several other features. Recommend to use `IPv6` connections.
  3. Discord ; `/home/usr/.config/discord/` ; An instant messaging and VoIP social platform which allows communication through voice calls, video calls, text messaging, and media.
  4. ProxyChains ; A tool that forces any TCP connection made by any given application to go through proxies like TOR or any other SOCKS4, SOCKS5 or HTTP proxies.
  5. *(for CN usrs/国内)* Wechat ; `/usr/share/wechat-universal` ; 微信； tofi可以启动。请选择`WeChat Universal`；
  6. *(for CN usrs/国内)* QQ ; tofi选择`QQ`再`enter`即可；
  7. *(for CN usrs/国内)* Dingtalk ; 钉钉；
  8. OBS-Studio ; 录屏软件； 倘若不直播，我个人推荐 `I444 (8-bit, 4:4:4, 3 planes)` 搭配码率 `4000-6000kbps`， `60fps`, `(Audio) Stereo` 进行非游戏内容录制以获得更佳效果；
- Novice安装配置之前应安装的 *部分* `应用/applications`与 `依赖环境/dependecies` 排列如下 *(不分先后)* ：
`wqy-microhei` `wqy-microhei-lite` `wqy-zenhei` `ttf-arphic-ukai` `ttf-arphic-uming` `adobe-source-han-sans-cn-fonts` `adobe-source-han-serif-cn-fonts` `noto-fonts-cjk` `yay` `noto-fonts-cjk` `noto-fonts-emoji` `wget` `paru` `jdk21-openjdk` `vi` `nano` `go` `wechat-uos-qt` `linuxqq-appimage` `Dingtalk-bin` `nautilus` `hyprctl` `tofictl` `pseudo` `grim` `OBS-studio` `obs-cli` `togglefloating` `swayidle` `jdk17-openjdk` `dbus` `git` `tofi` `vlc` `alacritty` `fish` `hyprland` `fontconfig` `hyprlock` `hyprpaper` `kitty` `mako` `waybar` `dwindle`

### $. Novice / 效果展示

## (二) 使用
### 安装
使用`Alt+ T`唤起终端，此时应该会有黄色的警告标识遮住终端显示，多次`enter`将输入行顶下来即可。你可以使用以下方法检验网络的正常接入：

    ping microsoft.com
安装git:

    sudo pacman -S git
进入用户目录：

    cd ~
将`Novice`克隆到当前目录下：

    git clone https://github.com/Siracusant/Novice.git
*注意，此处可能会概率性出现卡在`cloning into ~`的情况，这是由于国内网络特性，github无法稳定连接导致。你可以选择等待片刻多次尝试，也可以选择先进行网络环境的配置。*
进入`Novice`目录：

    cd Novice/
运行`novice`：

    ./novice
此后等待程序检验和安装即可。

### 检验

### 快捷键
1. 唤起终端 `Alt + T`或`Alt + A`;
2. 窗口全屏 `Alt + F`或`F11`;
3. 窗口浮动 `Alt + V`;
4. 窗口微缩 `Alt + P`;
5. 打开tofi `Alt + R`;
6. 复制 `Ctrl + Shift + C`或`Ctrl + C`;
7. 粘贴 `Ctrl + Shift + V`或`Ctrl + V`;
8. 终止进程 `Alt + C`或`Ctrl + C`;
9. 资源管理 `Alt + E`;
10. 桌面锁屏 `Alt + L`;
11. 切换工作区 `Alt + 1`,`Alt + 2`...`Alt + 10`或`Alt + 滚轮滚动`;
12. 终止waybar `Alt + Ctrl + Super + R`;
13. 退出hyprland桌面环境 `Alt + Shift + M`;
14. 区域截屏 `Alt + Shift + S`或`PrintScreen`;
15. 全屏截屏 `Alt + Shift + W`;
16. OBS开始录制 `Alt + Shift + T`;
17. OBS结束录制 `Alt + Shift + P`;
18. 调整窗口大小 `Alt + 鼠标右键拖动`;
19. 移动窗口 `Alt + 鼠标左键拖动`;
20. 音量调大 `Fn + F12`或`F12` *(具体取决于你的键盘设计)*;
21. 音量调小 `Fn + F11`或`F11` *(具体取决于你的键盘设计)*;
22. 静音 `Fn + F10`或`F10` *(具体取决于你的键盘设计)*;
23. 启用输入法 `Super + Space`;
24. 内容自动填充 `Tab`;
25. 内容按组进行删除 `Alt + BackSpace`或`Ctrl + BackSpace`或`Shift + BackSpace`;
26. 内容按组进行光标移动 `Alt + ↑ ↓ ← →`或`Ctrl + ↑ ↓ ← →`或`Shift + ↑ ↓ ← →`;
27. 内容根据光标移动进行选中 `Shift + ↑ ↓ ← →`;
