#!/usr/bin/env ruby

# Dotfiles and environment are personal GitHub repositories,
# they must be cloned locally in the below directories
DOTFILES = "/home/json/Documents/workarea/dotfiles/"
ENVIRONMENT = "/home/json/Documents/workarea/environment/"

def symlink(dir, source, destination)
  check_dir_exist(DOTFILES)
  check_dir_exist(ENVIRONMENT)
  create_symlink(dir, source, destination)
end

def check_dir_exist(dir)
  raise "ERROR #{dir} doesn't exist" unless File.directory?(dir)
end

def create_symlink(dir, source, destination)
  `ln -sf #{dir}#{source} #{destination}`
  puts "created symlink #{destination}" 
end

# Files
symlink(DOTFILES, "bashrc", "~/.bashrc")
symlink(DOTFILES, "bash_profile", "~/.bash_profile")
symlink(DOTFILES, "git/gitconfig", "~/.gitconfig")
symlink(DOTFILES, "ranger", "~/.config/")
symlink(DOTFILES, "tmux", "~/.tmux.conf")
symlink(DOTFILES, "vimrc", "~/.vimrc")
symlink(DOTFILES, "i3/xinitrc", "~/.xinitrc")
symlink(DOTFILES, "i3/config", "~/.i3/config")
symlink(DOTFILES, "i3/i3status.conf", "~/.i3/i3status.conf")
symlink(DOTFILES, "vscode/keybindings.json", "~/.config/Code/User/keybindings.json")
symlink(DOTFILES, "vscode/settings.json", "~/.config/Code/User/settings.json")
symlink(DOTFILES, "apt/sources.list", "/etc/apt/sources.list")
symlink(DOTFILES, "ncspot/config.toml", "/home/json/snap/ncspot/17/.config/ncspot/config.toml")

# Directories
symlink(DOTFILES, "bash_completions/*", "~/.completions/")
symlink(DOTFILES, "terminator/config", "~/.config/terminator/config")
symlink(DOTFILES, "vim/plugins", "~/.vim/pack/")
symlink(DOTFILES, "vim/colors", "~/.vim/")
symlink(DOTFILES, "tmux-dir/plugins", "~/.tmux/")

symlink(ENVIRONMENT, "bash/*", "~/bin")
symlink(ENVIRONMENT, "ruby/*", "~/bin")
