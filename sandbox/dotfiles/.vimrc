
" install vplug:
" curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
" vimrc source and PlugInstall

syntax on
set number
set tabstop=4
set shiftwidth=4
set expandtab
set t_Co=256
set bg=dark

imap jj <Esc>
:nnoremap <Space> :


call plug#begin()
    Plug 'morhetz/gruvbox'
    Plug 'sbdchd/neoformat'
call plug#end()


colorscheme gruvbox
