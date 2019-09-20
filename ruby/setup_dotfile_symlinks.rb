#!/usr/bin/env ruby

# Create symlinks from dotfiles dir ONLY!
def symlink(source, destination)
  dir = "/home/json/Documents/workarea/dotfiles/"
  if (File.directory?(dir))
    `ln -sf #{dir}#{source} #{destination}`
    puts "created symlink #{destination}" 
  else
    puts "#{dir} doesn't exist"
  end
end

symlink("vimrc", "~/.vimrc")
symlink("vim/*", "~/.vim")
symlink("bashr", " ~/.bashrc")
symlink("bash_profile", "~/.bash_profile")
symlink("bash_completion", "/* ~/.completions")
symlink("terminator/config", "~/.config/terminator/config")
symlink("tmux", "~/.tmux.conf")
symlink("vim/plugins", "~/.vim/pack/plugins")
symlink("vim/colors", "~/.vim/colors")
symlink("ranger/*", "~/.config/ranger/")
symlink("git/gitconfig", "~/.gitconfig")
