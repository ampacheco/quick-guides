# Setup 
````
cat << EOF > ~/.vimrc 
syntax enable
set tabstop=2       " number of visual spaces per TAB
set softtabstop=2   " number of spaces in tab when editing
set expandtab       " tabs = spaces

set number		" show line numbers
set cursorline		" highlight the current line
set showcmd		" show command in bottom bar
filetype indent on	" load filetype specific indent files in ~/.vim/indent/*.vim
set wildmenu
set lazyredraw
set showmatch		" show matches of {} [] ()

set incsearch		" search as characters are entered
set hlsearch		" highlight matches
set ignorecase		" ignore upper/lower case when searching
let mapleader = "," 	" leader is comma
" turn off search highlight, mapping to ,<space>
nnoremap <leader><space> :nohlsearch<CR>

set backspace=indent,eol,start   " allows backspace to delete words
EOF

source <(kubectl completion bash) # setup autocomplete in bash into the current shell, bash-completion package should be installed first.
echo "source <(kubectl completion bash)" >> ~/.bashrc # add autocomplete permanently to your bash shell.

alias k=kubectl
complete -F __start_kubectl k
````
