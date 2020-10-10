# dotfiles

ideaVim, tmux, oh-my-zsh 配置

### install

```
source bootstrap.sh
```

### set up

phpstorm/idea 安装插件 ideaVim + ideaVim-Easymotion

tmux 安装tpm后 `<C-a>` + `I` install plugins

### keymap

#### zsh

- `<Esc>` 进入vi模式(插件)
- `./xx.zip` 直接解压各种压缩文件

#### tmux

复制、分屏选择都采用vi模式

- `<C-a>` 指令前缀
- `-` 水平分屏
- `\` 垂直分屏
- `r` 重载配置
- `S` 保存session
- `R` 载入session

#### ideaVim

[neovim配置](https://github.com/moonprism/nvim), 继承其中除插件外的大部分操作
