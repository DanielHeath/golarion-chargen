# What is this

This setup is for a go project that produces a binary (it's not suitable for publishing a library).
The first person to set it up needs `gb` installed globalally (see `go get github.com/constabulary/gb/...` ).

```
# Install GB globally:
go get github.com/constabulary/gb/...

# Create output directories
mkdir -p src vendor/src pkg bin

# Ignore temporary and secret files:
echo 'pkg' >> .gitignore
echo 'bin' >> .gitignore
echo '.env.secret' >> .gitignore

# Use your global copy of gb to create a vendored copy of gb
gb vendor fetch github.com/constabulary/gb
gb vendor fetch github.com/constabulary/gb/vendor
make bin/gb

# Commit
git init .
git add .
git commit -m "Initial setup"
```

N.B. 

When 