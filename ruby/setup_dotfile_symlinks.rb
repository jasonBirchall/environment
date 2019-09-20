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

symlink(DOTFILES, "vimrc", "~/.vimrc")
symlink(DOTFILES, "vim/*", "~/.vim")
symlink(DOTFILES, "bashrc", " ~/.bashrc")
symlink(DOTFILES, "bash_profile", "~/.bash_profile")
symlink(DOTFILES, "bash_completion/*", "~/.completions")
symlink(DOTFILES, "terminator/config", "~/.config/terminator/config")
symlink(DOTFILES, "tmux", "~/.tmux.conf")
symlink(DOTFILES, "vim/plugins", "~/.vim/pack/plugins")
symlink(DOTFILES, "vim/colors", "~/.vim/colors")
symlink(DOTFILES, "ranger/*", "~/.config/ranger/")
symlink(DOTFILES, "git/gitconfig", "~/.gitconfig")

symlink(ENVIRONMENT, "bash/*", "~/bin")
symlink(ENVIRONMENT, "ruby/*", "~/bin")
