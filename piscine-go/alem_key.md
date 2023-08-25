mkdir -p ~/.ssh

ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -N ''

cat ~/.ssh/id_ed25519.pub

git clone git@git.01.alem.school:dkurmang/piscine-go.git

git config --global user.email "danaeskazi@gmail.com"
git config --global user.name "Dana"
