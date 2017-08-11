
### Add subtree to module
git subtree add --prefix <prefix dir> <repo url> master --squash
ex.
git subtree add --prefix .vim/bundle/tpope-vim-surround https://bitbucket.org/vim-plugins-mirror/vim-surround.git master --squash

### Subtree pull
git subtree pull ...

### Reference
https://git-scm.com/book/en/v2/Git-Tools-Advanced-Merging#_subtree_merge
