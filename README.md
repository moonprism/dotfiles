<div align="center">
    <h1>.dotfiles</h1>
    <br />
</div>

系统： Arch / MacOS

终端： [Kitty](https://github.com/kovidgoyal/kitty)

字体： [Nerd Fonts Lilex](https://github.com/ryanoasis/nerd-fonts/tree/master/patched-fonts/Lilex)

配色： [Molokai](https://github.com/tomasr/molokai)

### Keymaps

- [xkeysnail](https://github.com/mooz/xkeysnail) (linux)
- [karabiner-elements](https://github.com/pqrs-org/Karabiner-Elements) (mac)

按键            | 映射               | 说明
--              | --                 | --
`<Caps>`        | \<Ctrl>             | 主要
`<Alt-h/j/k/l>` | Left/Down/Up/Right | 光标移动
`<Alt-c/v/x/f>` | <Ctrl-c/v/x/f>     | 复制/粘贴/剪切/搜索
`<Alt-ESC>`     | `                  | 反引号

> 通过 karabiner 可以配置轻击 Ctrl 为 ECS, 长按 Ctrl 为正常 Ctrl.
> linux 下也有大佬给出工具: <https://gitlab.com/interception/linux/plugins/caps2esc>

以后有机会上GH60.

### Shell

[Oh My Zsh](http://github.com/ohmyzsh/ohmyzsh)

- [zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting)
- [zsh-autosuggestions](https://github.com/zsh-users/zsh-autosuggestions)

命令行工具:

- [fzf](https://github.com/junegunn/fzf)

### IdeaVim

安装 ideavim 和 ideavim-easymotion 插件

### Tmux

安装 tmux 和 tmp 插件管理工具

复制、分屏选择都采用vi模式

按键    | 说明
--      | --
`<C-a>` | 指令前缀
`j`     | jump, 类似 vim-easymotion
`-`     | 水平分屏
`/`     | 垂直分屏
`r`     | 重载配置
`S`     | 保存session
`R`     | 载入session

### 其他

- HyperSwitch (mac 标签切换)
