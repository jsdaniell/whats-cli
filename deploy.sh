echo "Deploy new version to releases and NPM"
echo "Type the version: "

read -r version

git add -A

echo "Commit message: "
read -r commitDescription
git commit -am "${commitDescription}"
git push

git tag -a "v${version}" -am "Beta Release" && git push origin "v${version}"
npm publish

