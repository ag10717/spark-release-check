latest_tag=$(git describe --tags "$(git rev-list --tags --max-count=1)")
version_number="${latest_tag:1}"
new_version_number=$(version_number + 1)

# push build for package only

# 1. build make
# 2. clean directory
# 3. checkout release
# 4. commit 
# 5. tag
# 6. push
#
make build
git checkout -b release
find . -not -name "lib" -delete
git add "lib"
git commit -m "release version"
git tag -a "$new_version_number" -m "$new_version_number"
git push -u origin release
