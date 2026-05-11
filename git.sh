echo "name for commit: "
read com
git add .
git commit -m "$com"
git remote add gitflic git@gitflic.ru:hhu67/proj.git
git remote add github git@github.com:hhu67/my.git
git remote add gitea git@gitek.duckdns.org:hhu67/my.git
git remote add berg ssh://git@codeberg.org/hhu67/sergay.git
git push gitflic main
git push github main
git push gitea main 
git push berg main
