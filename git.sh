#!/bin/bash

echo "name for commit: "
read com
git add .
git commit -m "$com"
declare -A remotes=(
    ["gitflic"]="git@gitflic.ru:hhu67/proj.git"
    ["github"]="git@github.com:hhu67/my.git"
    ["forgejo"]="git@git.vlv-s.site:hhu67/my.git"
    ["berg"]="ssh://git@codeberg.org/hhu67/my.git"
    ["gitea.org"]="git@gitea.com:hhu67/my.git"
)
manage_remote() {
    local name=$1
    local url=$2
    
    if git remote | grep -q "^${name}$"; then
        echo "Remote '$name' уже существует, пропускаю добавление."
    else
        echo "Добавляю remote: $name"
        git remote add "$name" "$url"
    fi
}
for name in "${!remotes[@]}"; do
    manage_remote "$name" "${remotes[$name]}"
    echo "Пушу в $name..."
    git push "$name" main
done
